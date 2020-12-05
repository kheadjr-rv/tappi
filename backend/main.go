package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/terraform-exec/tfexec"
)

type executor struct {
	tf *tfexec.Terraform
}

func newExecutor(workingDir string, execPath string) (*executor, error) {
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		return nil, err
	}

	return &executor{
		tf: tf,
	}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	vars := mux.Vars(r)

	action, ok := vars["action"]
	if !ok {
		http.Error(w, "invalid action", http.StatusInternalServerError)
		return
	}

	exec, err := newExecutor("./terraform", "/usr/local/bin/terraform")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pr, pw := io.Pipe()
	defer pw.Close()

	exec.tf.SetStdout(io.MultiWriter(pw))
	exec.tf.SetStderr(io.MultiWriter(pw))

	w.WriteHeader(http.StatusOK)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	go func(res http.ResponseWriter, pipeReader *io.PipeReader) {
		buffer := make([]byte, 1024)
		for {
			n, err := pipeReader.Read(buffer)
			if err != nil {
				pipeReader.Close()
				break
			}

			data := buffer[0:n]
			res.Write(data)
			if f, ok := res.(http.Flusher); ok {
				f.Flush()
			}
			//reset buffer
			for i := 0; i < n; i++ {
				buffer[i] = 0
			}
		}
	}(w, pr)

	switch action {
	case "init":
		err = exec.tf.Init(r.Context(), tfexec.Upgrade(false), tfexec.LockTimeout("60s"))
		if err != nil {
			pw.Write([]byte(err.Error()))
			return
		}

	case "plan":
		_, err = exec.tf.Plan(r.Context())
		if err != nil {
			pw.Write([]byte(err.Error()))
			return
		}

	case "refresh":
		err = exec.tf.Refresh(r.Context())
		if err != nil {
			pw.Write([]byte(err.Error()))
			return
		}

	case "apply":
		err = exec.tf.Apply(r.Context())
		if err != nil {
			pw.Write([]byte(err.Error()))
			return
		}
	}

}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/{action}", handler).Methods(http.MethodGet, http.MethodOptions)

	r.Use(mux.CORSMethodMiddleware(r))

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 120,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// TODO move this

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

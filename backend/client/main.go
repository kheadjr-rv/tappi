package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080/init", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("%d - %s", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()

	for {
		buf := make([]byte, 1024)
		_, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

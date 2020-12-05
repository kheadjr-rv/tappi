package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type editor struct {
	app.Compo

	response string
}

func (e *editor) Render() app.UI {
	return app.Shell().
		Class("app-background").
		Menu(&menu{}).
		// Submenu(
		// 	newTableOfContents().
		// 		Links(p.links...),
		// ).
		OverlayMenu(&overlayMenu{}).
		Content(
			app.Main().
				Class("pane").
				Class("document").
				Body(
					app.H1().Text("Editor"),
					app.Section().Body(
						app.Div().
							ID("editor").
							Class("editor").
							Style("height", "500px").
							Style("width", "80%").
							Text(`# terraform code goes here`),
					),
					app.Section().Body(
						app.Flow().Content(
							app.Button().
								OnClick(e.OnClick).
								Text("Init").
								Value("init"),
							app.Button().
								OnClick(e.OnClick).
								Text("Plan").
								Value("plan"),
							app.Button().
								OnClick(e.OnClick).
								Text("Apply").
								Value("apply"),
							app.Button().
								OnClick(e.OnClick).
								Text("Refresh").
								Value("refresh"),
						).ItemsBaseWidth(50),
					),
					app.Section().Body(
						app.Textarea().
							Class("terminal").
							Cols(80).
							Rows(20).
							Text(e.response),
					),
					app.Script().Text(`
					var editor = ace.edit("editor");
					editor.setTheme("ace/theme/monokai");
					editor.session.setMode("ace/mode/terraform");
					editor.setOptions({
						autoScrollEditorIntoView: true,
						maxLines: 300,
						minLines: 10
					});
					`),
				),
		)
}

func (e *editor) OnClick(ctx app.Context, evt app.Event) {
	e.response = ""
	e.Update()

	action := ctx.JSSrc.Get("value").String()
	url := fmt.Sprintf("http://localhost:8080/%s", action)
	go e.doRequest(url) // Performs blocking operation on a new goroutine.
}

func (e *editor) doRequest(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		e.updateResponse(err.Error())
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		e.updateResponse(err.Error())
		return
	}

	defer resp.Body.Close()

	for {
		buf := make([]byte, 1024)
		_, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			e.updateResponse(err.Error())
			break
		}
		e.updateResponse(string(buf))
	}

}

func (e *editor) updateResponse(res string) {
	app.Dispatch(func() { // Ensures response field is updated on UI goroutine.
		e.response = e.response + res
		e.Update()
	})
}

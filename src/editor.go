package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type editor struct {
	app.Compo
}

func (s *editor) Render() app.UI {
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

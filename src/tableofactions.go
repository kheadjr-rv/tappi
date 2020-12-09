package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type tableOfActions struct {
	app.Compo

	Ilinks   []string
	selected string
}

func newTableOfActions() *tableOfActions {
	return &tableOfActions{}
}

func (t *tableOfActions) Links(v ...string) *tableOfActions {
	t.Ilinks = v
	return t
}

// func (t *tableOfActions) OnNav(ctx app.Context, u *url.URL) {
// 	t.selected = "#" + u.Fragment
// 	t.Update()
// }

func (t *tableOfActions) Render() app.UI {
	return app.Aside().
		Class("pane").
		Class("index").
		Body(
			app.H1().Text("Action"),
			app.Section().Body(
				app.Range(t.Ilinks).Slice(func(i int) app.UI {
					link := t.Ilinks[i]

					return &tableOfActionLink{
						Title: link,
						Focus: t.selected == githubIndex(link),
					}
				}),
			),
		)
}

type tableOfActionLink struct {
	app.Compo

	Title string
	Focus bool
}

func (l *tableOfActionLink) Render() app.UI {
	focus := ""
	if l.Focus {
		focus = "focus"
	}

	return app.A().
		Class(focus).
		Href(githubIndex(l.Title)).
		Text(l.Title)
}

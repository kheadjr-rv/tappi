package main

import (
	"path/filepath"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

const (
	githubURL = "https://github.com/kheadjr-rv/tappi"
)

func pages() map[string]func() app.UI {
	return map[string]func() app.UI{
		"":          newStart,
		"start":     newStart,
		"terraform": newTerraform,
		"editor":    newEditor,
	}
}

func newStart() app.UI {
	return newPage().
		Path("/web/documents/start.md").
		TableOfContents(
			"Getting started",
		)
}

func newTerraform() app.UI {
	return newPage().
		Path("/web/documents/terraform.md").
		TableOfContents(
			"Terraform",
			"    Handling Local Name Conflicts",
		)
}

func newEditor() app.UI {
	e := &editor{}
	return e.TableOfActions(
		"init",
		"plan",
		"apply",
		"refresh",
		"destroy",
	)
}

type page struct {
	app.Compo

	path  string
	links []string
}

func newPage() *page {
	return &page{}
}

func (p *page) Path(v string) *page {
	p.path = v
	return p
}

func (p *page) TableOfContents(v ...string) *page {
	p.links = v
	return p
}

func (p *page) Render() app.UI {
	return app.Shell().
		Class("app-background").
		Menu(&menu{}).
		Submenu(
			newTableOfContents().
				Links(p.links...),
		).
		OverlayMenu(&overlayMenu{}).
		Content(
			newDocument(p.path).
				Description(filepath.Base(p.path)),
		)
}

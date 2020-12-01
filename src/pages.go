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
		"":             newStart,
		"start":        newStart,
		"architecture": newArchitecture,
		"terraform":    newTerraform,
		// "components":       newCompo,
	}
}

func newStart() app.UI {
	return newPage().
		Path("/web/documents/start.md").
		TableOfContents(
			"Getting started",
			// "Prerequisite",
			// "Install",
			// "User interface",
			// "Server",
			// "Build and run",
			// "Tips",
			// "Next",
		)
}

func newArchitecture() app.UI {
	return newPage().
		Path("/web/documents/architecture.md").
		TableOfContents(
			"Architecture",
		)
}

func newTerraform() app.UI {
	return newPage().
		Path("/web/documents/terraform.md").
		TableOfContents(
			"Terraform",
			"Handling Local Name Conflicts",
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

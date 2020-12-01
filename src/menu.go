package main

import (
	"net/url"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type menu struct {
	app.Compo

	currentPath string
}

func (m *menu) OnNav(ctx app.Context, u *url.URL) {
	path := u.Path
	if path == "/" {
		path = "/start"
	}
	m.currentPath = path

	m.Update()
}

func (m *menu) Render() app.UI {
	return app.Nav().
		Class("menu").
		Body(
			app.Div().Body(
				app.A().
					Class("title").
					Href("/start").
					Text("TAPPI"),
			),
			app.Div().
				Class("content").
				Body(
					app.Section().Body(
						newMenuItem().
							Icon(`
							<svg style="width:24px;height:24px" viewBox="0 0 24 24">
    							<path fill="currentColor" d="M13.13 22.19L11.5 18.36C13.07 17.78 14.54 17 15.9 16.09L13.13 22.19M5.64 12.5L1.81 10.87L7.91 8.1C7 9.46 6.22 10.93 5.64 12.5M21.61 2.39C21.61 2.39 16.66 .269 11 5.93C8.81 8.12 7.5 10.53 6.65 12.64C6.37 13.39 6.56 14.21 7.11 14.77L9.24 16.89C9.79 17.45 10.61 17.63 11.36 17.35C13.5 16.53 15.88 15.19 18.07 13C23.73 7.34 21.61 2.39 21.61 2.39M14.54 9.46C13.76 8.68 13.76 7.41 14.54 6.63S16.59 5.85 17.37 6.63C18.14 7.41 18.15 8.68 17.37 9.46C16.59 10.24 15.32 10.24 14.54 9.46M8.88 16.53L7.47 15.12L8.88 16.53M6.24 22L9.88 18.36C9.54 18.27 9.21 18.12 8.91 17.91L4.83 22H6.24M2 22H3.41L8.18 17.24L6.76 15.83L2 20.59V22M2 19.17L6.09 15.09C5.88 14.79 5.73 14.47 5.64 14.12L2 17.76V19.17Z" />
							</svg>
							`).
							Text("Getting started").
							Selected(m.currentPath == "/start").
							Href("/start"),
						newMenuItem().
							Icon(`
							<svg style="width:24px;height:24px" viewBox="0 0 24 24">
    							<path fill="currentColor" d="M9,2V8H11V11H5C3.89,11 3,11.89 3,13V16H1V22H7V16H5V13H11V16H9V22H15V16H13V13H19V16H17V22H23V16H21V13C21,11.89 20.11,11 19,11H13V8H15V2H9Z" />
							</svg>
							`).
							Text("Architecture").
							Selected(m.currentPath == "/architecture").
							Href("/architecture"),
					),
					app.Section().Body(
						newMenuItem().
							Icon(`
							<svg style="width:24px;height:24px" viewBox="0 0 24 24">
								<path fill="currentColor" d="M10,5V11H21V5M16,18H21V12H16M4,18H9V5H4M10,18H15V12H10V18Z" />
							</svg>
							`).
							Text("Terraform").
							Selected(m.currentPath == "/terraform").
							Href("/terraform"),
					),
					app.Section().Body(
						newMenuItem().
							Icon(`
							<svg style="width:24px;height:24px" viewBox="0 0 24 24">
    							<path fill="currentColor" d="M12,2A10,10 0 0,0 2,12C2,16.42 4.87,20.17 8.84,21.5C9.34,21.58 9.5,21.27 9.5,21C9.5,20.77 9.5,20.14 9.5,19.31C6.73,19.91 6.14,17.97 6.14,17.97C5.68,16.81 5.03,16.5 5.03,16.5C4.12,15.88 5.1,15.9 5.1,15.9C6.1,15.97 6.63,16.93 6.63,16.93C7.5,18.45 8.97,18 9.54,17.76C9.63,17.11 9.89,16.67 10.17,16.42C7.95,16.17 5.62,15.31 5.62,11.5C5.62,10.39 6,9.5 6.65,8.79C6.55,8.54 6.2,7.5 6.75,6.15C6.75,6.15 7.59,5.88 9.5,7.17C10.29,6.95 11.15,6.84 12,6.84C12.85,6.84 13.71,6.95 14.5,7.17C16.41,5.88 17.25,6.15 17.25,6.15C17.8,7.5 17.45,8.54 17.35,8.79C18,9.5 18.38,10.39 18.38,11.5C18.38,15.32 16.04,16.16 13.81,16.41C14.17,16.72 14.5,17.33 14.5,18.26C14.5,19.6 14.5,20.68 14.5,21C14.5,21.27 14.66,21.59 15.17,21.5C19.14,20.16 22,16.42 22,12A10,10 0 0,0 12,2Z" />
							</svg>
							`).
							Text("GitHub").
							Href(githubURL).
							External(),
					),
				),
		)
}

type menuItem struct {
	app.Compo

	Iicon     string
	Ihref     string
	Iselected string
	Itext     string
	Itarget   string
	Irel      string
}

func newMenuItem() *menuItem {
	return &menuItem{}
}

func (i *menuItem) Icon(svg string) *menuItem {
	i.Iicon = svg
	return i
}

func (i *menuItem) Href(v string) *menuItem {
	i.Ihref = v
	return i
}

func (i *menuItem) Selected(v bool) *menuItem {
	if v {
		i.Iselected = "focus"
	}
	return i
}

func (i *menuItem) Text(v string) *menuItem {
	i.Itext = v
	return i
}

func (i *menuItem) External() *menuItem {
	i.Itarget = "_blank"
	i.Irel = "noopener"
	return i
}

func (i *menuItem) Render() app.UI {
	return app.A().
		Class("item").
		Class(i.Iselected).
		Href(i.Ihref).
		Target(i.Itarget).
		Rel(i.Irel).
		Body(
			app.Stack().
				Center().
				Content(
					app.Div().
						Class("icon").
						Body(
							app.Raw(i.Iicon),
						),
					app.Div().
						Class("label").
						Text(i.Itext),
				),
		)
}

type overlayMenu struct {
	app.Compo
}

func (m *overlayMenu) Render() app.UI {
	return app.Div().
		Class("overlay-menu").
		Body(
			&menu{},
		)
}

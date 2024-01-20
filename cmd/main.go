package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
	"twitch-htmx-server/templates"
)

func main() {
	err := templates.RegisterTheme("public/themes/default")
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("*", pageHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

type Site struct {
	Title             string
	Url               string
	MembersEnabled    bool `handlebars:"members_enabled"`
	Locale            string
	Logo              string
	MembersInviteOnly bool `handlebars:"members_invite_only"`
}

type Custom struct {
	ColorScheme      string `handlebars:"color_scheme"`
	NavigationLayout string `handlebars:"navigation_layout"`
}

type TplData struct {
	Site      Site   `handlebars:"@site"`
	Custom    Custom `handlebars:"@custom"`
	MetaTitle string `handlebars:"meta_title"`
	GhostHead string `handlebars:"ghost_head"`
	BodyClass string `handlebars:"body_class"`
}

// TODO: implement template helper function "asset",
// Example use case: <link rel="preload" as="script" href="{{asset "built/casper.js"}}" />

func pageHandler(c echo.Context) error {
	fmt.Println(c.Request().URL.Path)
	path := strings.TrimPrefix(c.Request().URL.Path, "/")

	tplData := TplData{
		Site: Site{
			Locale: "TR-tr",
		},
	}

	rendered, err := templates.RenderPage(path, tplData)
	if err != nil {
		return c.HTML(http.StatusInternalServerError,
			"<html><head><title>Internal Server Error</title></head><body>"+err.Error()+"</body></html>")
	}

	return c.HTML(http.StatusOK, rendered)
}

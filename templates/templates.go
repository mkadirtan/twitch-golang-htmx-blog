package templates

import (
	"errors"
	"github.com/mailgun/raymond/v2"
	"os"
	"twitch-htmx-server/state"
)

type PageDefinition struct {
	Template *raymond.Template
	Layout   string
}

type ThemeDefinition struct {
	pages    map[string]*PageDefinition
	partials map[string]*raymond.Template
}

var themes = make(map[string]*ThemeDefinition, 1)
var activeTheme string

func registerTheme(themesRoot string, themeName string) error {
	themeRoot := themesRoot + "/" + themeName
	pages, err := loadThemePages(themeRoot)
	if err != nil {
		return errors.Join(err, errors.New("theme pages load error"))
	}

	partials, err := loadThemePartials(themeRoot)
	if err != nil {
		return errors.Join(err, errors.New("theme partials load error"))
	}

	themes[themeName] = &ThemeDefinition{
		pages:    pages,
		partials: partials,
	}

	return nil
}

func RegisterThemes(themesRoot string, watchMode bool) error {
	err := registerThemes(themesRoot)
	if err != nil {
		return err
	}

	if watchMode {
		watchThemeChanges(themesRoot)
	}

	return nil
}

func registerThemes(themesRoot string) error {
	themes, err := os.ReadDir(themesRoot)
	if err != nil {
		return err
	}

	for _, theme := range themes {
		if !theme.IsDir() {
			continue
		}
		err = registerTheme(themesRoot, theme.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

func ActivateTheme(themeName string) error {
	theme := themes[themeName]
	if theme == nil {
		return errors.New("theme not found: " + themeName)
	}

	raymond.RemoveAllPartials()
	for partialName, partial := range theme.partials {
		raymond.RegisterPartialTemplate(partialName, partial)
	}

	// maybe add pages validation in the future here...

	activeTheme = themeName

	return nil
}

func RenderPage(pageName string, data state.TplData) (string, error) {
	if activeTheme == "" {
		return "", errors.New("no active theme")
	}
	theme := themes[activeTheme]
	if theme == nil {
		return "", errors.New("active theme not found: " + activeTheme)
	}

	pageDefinition := theme.pages[pageName]
	if pageDefinition == nil {
		return "", errors.New("page doesnt exist")
	}

	rendered, err := pageDefinition.Template.Exec(data)
	if err != nil {
		return "", err
	}

	if pageDefinition.Layout != "" {
		layout := theme.pages[pageDefinition.Layout]
		if layout == nil {
			return "", errors.New("missing layout: " + pageDefinition.Layout)
		}

		rendered, err = layout.Template.Exec(map[string]interface{}{
			"body": rendered,
		})

		if err != nil {
			return "", nil
		}
	}

	return rendered, nil
}

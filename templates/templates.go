package templates

import (
	"errors"
	"github.com/mailgun/raymond/v2"
	"twitch-htmx-server/state"
)

type PageDefinition struct {
	Template *raymond.Template
	Layout   string
}

var pages = map[string]*PageDefinition{}

// RegisterTheme expects theme root directory
func RegisterTheme(themeRoot string) error {
	err := registerThemePages(themeRoot)
	if err != nil {
		return errors.Join(err, errors.New("theme page register error"))
	}

	err = registerThemePartials(themeRoot)
	if err != nil {
		return errors.Join(err, errors.New("theme partials register error"))
	}

	return nil
}

func RenderPage(pageName string, data state.TplData) (string, error) {
	pageDefinition := pages[pageName]
	if pageDefinition == nil {
		return "", errors.New("page doesnt exist")
	}

	rendered, err := pageDefinition.Template.Exec(data)
	if err != nil {
		return "", err
	}

	if pageDefinition.Layout != "" {
		layout := pages[pageDefinition.Layout]
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

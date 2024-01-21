package templates

import (
	"fmt"
	"github.com/mailgun/raymond/v2"
	"path/filepath"
	"strings"
)

func loadThemePages(themeRoot string) (map[string]*PageDefinition, error) {
	pageGlob := themeRoot + "/*.hbs"
	fmt.Println("pageGlob: ", pageGlob)

	pageMatches, err := filepath.Glob(pageGlob)
	if err != nil {
		return nil, err
	}

	var pageDefinitions = make(map[string]*PageDefinition, 1)

	for _, pageFilepath := range pageMatches {
		pageName := strings.TrimSuffix(strings.TrimPrefix(pageFilepath, themeRoot+"/"), ".hbs")
		fmt.Println("pageName: ", pageName)

		template, parseErr := raymond.ParseFile(pageFilepath)
		if parseErr != nil {
			return nil, parseErr
		}

		layout, err := detectLayout(pageFilepath)
		if err != nil {
			return nil, err
		}

		pageDefinitions[pageName] = &PageDefinition{
			Template: template,
			Layout:   layout,
		}

		raymond.RemovePartial(pageName)
		raymond.RegisterPartialTemplate(pageName, template)
	}

	return pageDefinitions, nil
}

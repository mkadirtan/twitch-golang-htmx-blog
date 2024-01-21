package templates

import (
	"fmt"
	"github.com/mailgun/raymond/v2"
	"path/filepath"
	"slices"
	"strings"
)

func registerThemePages(themeRoot string) error {
	pageGlob := themeRoot + "/*.hbs"
	fmt.Println("pageGlob: ", pageGlob)

	pageMatches, err := filepath.Glob(pageGlob)
	if err != nil {
		return err
	}

	expectedPages := []string{
		"default",
		"index",
		"hello",
	}
	for _, pageFilepath := range pageMatches {
		pageName := strings.TrimSuffix(strings.TrimPrefix(pageFilepath, themeRoot+"/"), ".hbs")
		fmt.Println("pageName: ", pageName)

		var index int
		if index = slices.Index(expectedPages, pageName); index == -1 {
			fmt.Println("unknown page detected: ", pageFilepath, ". skipping registration.")
			continue
		}

		expectedPages = slices.Delete(expectedPages, index, index+1)

		template, parseErr := raymond.ParseFile(pageFilepath)
		if parseErr != nil {
			return parseErr
		}

		layout, err := detectLayout(pageFilepath)
		if err != nil {
			return err
		}

		pages[pageName] = &PageDefinition{
			Template: template,
			Layout:   layout,
		}

		raymond.RemovePartial(pageName)
		raymond.RegisterPartialTemplate(pageName, template)
	}

	// TODO: return error if expected pages not found

	return nil
}

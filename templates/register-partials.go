package templates

import (
	"fmt"
	"github.com/mailgun/raymond/v2"
	"path/filepath"
	"strings"
)

func loadThemePartials(themeRoot string) (map[string]*raymond.Template, error) {
	partialGlob := themeRoot + "/partials/**/*.hbs"
	partialMatches, err := filepath.Glob(partialGlob)

	if err != nil {
		return nil, err
	}

	var partials = make(map[string]*raymond.Template, 1)

	for _, pageFilepath := range partialMatches {
		partialName := strings.TrimSuffix(strings.TrimPrefix(pageFilepath, themeRoot+"/partials/"), ".hbs")

		fmt.Println("partialName: ", partialName)

		template, parseErr := raymond.ParseFile(pageFilepath)
		if parseErr != nil {
			return nil, parseErr
		}

		partials[partialName] = template
	}

	return partials, nil
}

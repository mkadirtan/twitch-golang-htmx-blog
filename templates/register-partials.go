package templates

import (
	"fmt"
	"github.com/mailgun/raymond/v2"
	"path/filepath"
	"strings"
)

func registerThemePartials(themeRoot string) error {
	partialGlob := themeRoot + "/partials/**/*.hbs"
	partialMatches, err := filepath.Glob(partialGlob)
	if err != nil {
		return err
	}

	for _, pageFilepath := range partialMatches {
		partialName := strings.TrimSuffix(strings.TrimPrefix(pageFilepath, themeRoot+"/partials/"), ".hbs")

		fmt.Println("partialName: ", partialName)

		template, parseErr := raymond.ParseFile(pageFilepath)
		if parseErr != nil {
			return parseErr
		}

		raymond.RemovePartial(partialName)
		raymond.RegisterPartialTemplate(partialName, template)
	}

	return nil
}

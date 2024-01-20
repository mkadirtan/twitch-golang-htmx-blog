package templates

import (
	"errors"
	"fmt"
	"github.com/mailgun/raymond/v2"
	"path/filepath"
	"slices"
	"strings"
)

var pages = map[string]*raymond.Template{}

// expects theme root directory
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

func RenderPage(pageName string, data interface{}) (string, error) {
	template := pages[pageName]
	if template == nil {
		return "", errors.New("page doesnt exist")
	}

	rendered, err := template.Exec(data)
	if err != nil {
		return "", err
	}

	return rendered, nil
}

func registerThemePartials(themeRoot string) error {
	partialGlob := filepath.Dir(themeRoot) + "**/*.hbs"
	partialMatches, err := filepath.Glob(partialGlob)
	if err != nil {
		return err
	}

	for _, pageFilepath := range partialMatches {
		partialName, found := strings.CutPrefix(pageFilepath, "public/views/partials")
		if !found {
			fmt.Println("unexpected error, couldn't obtain filePath from path, path: ", pageFilepath)
			continue
		}
		partialName, found = strings.CutSuffix(partialName, ".hbs")
		if !found {
			fmt.Println("unexpected error, couldn't obtain pageFilepath from path, path: ", pageFilepath)
			continue
		}

		template, parseErr := raymond.ParseFile(pageFilepath)
		if parseErr != nil {
			return parseErr
		}

		raymond.RegisterPartialTemplate(partialName, template)
	}

	return nil
}

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

		pages[pageName] = template
	}

	// TODO: return error if expected pages not found

	return nil
}

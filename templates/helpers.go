package templates

import (
	"bufio"
	"os"
	"strings"
)

func detectLayout(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "{{!<") && strings.HasSuffix(line, "}}") {
			content := strings.TrimPrefix(line, "{{!<")
			content = strings.TrimSuffix(content, "}}")
			content = strings.TrimSpace(content)
			return content, nil
		}
	}

	if err = scanner.Err(); err != nil {
		return "", err
	}

	return "", nil
}

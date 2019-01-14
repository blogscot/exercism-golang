// Package grep searches through files for user specified patterns
package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Search examines one or more files looking for a matching pattern.
// Possible flags are:
// `-i` Match lines using a case-insensitive comparison.
// `-x` Only match entire lines.
// `-n` Include line numbers.
// `-l` Print only the names of files that contain at least one match.
// `-v` Collect all lines that DO NOT match the pattern.
func Search(pattern string, flags, files []string) []string {
	var out = []string{}

	for _, file := range files {
		if found := searchFile(pattern, flags, file); len(found) != 0 {
			if len(files) > 1 && !useFilenames(flags) {
				for i := range found {
					found[i] = fmt.Sprintf("%s:%s", file, found[i])
				}
			}
			out = append(out, found...)
		}
	}
	return out
}

func searchFile(pattern string, flags []string, filename string) []string {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to open file:%s", file.Name())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	var lineNumber int
	var matches = []string{}
	invertDecorator := invertIf(contains(flags, "-v"))
	matchWholeLine := contains(flags, "-x")
	ignoreCase := contains(flags, "-i")

	var updateMatches = func() {
		if contains(flags, "-n") {
			line = fmt.Sprintf("%d:%s", lineNumber, line)
		}
		matches = append(matches, line)
	}

	for scanner.Scan() {
		line = scanner.Text()
		lineNumber++
		var newLine = line
		var newPattern = pattern

		if ignoreCase {
			newLine = strings.ToLower(line)
			newPattern = strings.ToLower(pattern)
		}

		if matchWholeLine {
			if invertDecorator(newLine == newPattern) {
				updateMatches()
			}
		} else if invertDecorator(strings.Contains(newLine, newPattern)) {
			updateMatches()
		}

		if len(matches) > 0 && useFilenames(flags) {
			return []string{filename}
		}
	}
	return matches
}

func invertIf(invertFlagSet bool) func(bool) bool {
	if invertFlagSet {
		return func(pred bool) bool {
			return !pred
		}
	}
	return func(pred bool) bool {
		return pred
	}
}

func useFilenames(flags []string) bool {
	return contains(flags, "-l")
}

func contains(flags []string, item string) bool {
	for _, flag := range flags {
		if flag == item {
			return true
		}
	}
	return false
}

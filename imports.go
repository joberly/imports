package imports

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

// Regular expressions for scanning lines of code.

// Single line comment
var comLine = regexp.MustCompile(`[\t ]*\/\/`)

// Single line C style comment
var comCLine = regexp.MustCompile(`[\t ]*\/\*.*\*\/`)

// Comment start
var comStart = regexp.MustCompile(`[\t ]*\/\*`)

// Comment end
var comEnd = regexp.MustCompile(`[\t ]*\*\/`)

// Single line import
// Match index 2 = short name
// Match index 3 = import path
var importSingle = regexp.MustCompile(`import[\t ]+(([\S]+)[\t ]+)?\"(\S+)\"`)

// Multiline import start
var importStart = regexp.MustCompile(`import[\t ]+\(`)

// Multiline import entry
// Match index 2 = short name
// Match index 3 = import path
var importEntry = regexp.MustCompile(`(([\S]+)[\t ]+)?\"(\S+)\"`)

// Multiline import end
var importEnd = regexp.MustCompile(`\)`)

// readImports reads the input r as a Golang source file, looking for imports
// and returning a list of those imports in the string slice.
func readImports(r io.Reader) ([]string, error) {
	if r == nil {
		return nil, fmt.Errorf("no input reader")
	}

	// Read each line looking for either a single line to process
	// or the first line of a multiline chunk to process. For multiline
	// chunks, process each line until the end is found.
	s := bufio.NewScanner(r)
	var importList []string
	for s.Scan() {
		// Check for things in the following order:
		// Single comment line.
		//   > Skip the line.
		// Multiline comment
		//   > Scan lines until the end is found.
		// Single import match.
		//   > Add it to the importList to return.
		// Multiline import start.
		//   > Add each import in the list to importList to return.

		line := s.Text()
		if comLine.MatchString(line) || comCLine.MatchString(line) {
			// Skip comments
			continue
		} else if comStart.MatchString(line) {
			// Skip all comment lines until comment end is found.
			for s.Scan() {
				if comEnd.MatchString(s.Text()) {
					break
				}
			}
			continue
		} else if matchesSingle := importSingle.FindStringSubmatch(line); matchesSingle != nil {
			// Found a single line import, add it to the list.
			importList = append(importList, matchesSingle[3])
		} else if importStart.MatchString(s.Text()) {
			// Found a multiline import start.
			// Keep scanning each line until the import end parenthesis appears.
			for s.Scan() {
				line := s.Text()
				// Check for an import entry
				if matchesEntry := importEntry.FindStringSubmatch(line); matchesEntry != nil {
					importList = append(importList, matchesEntry[3])
				} else if importEnd.MatchString(line) {
					break
				}
			}
		}
	}

	return importList, nil
}

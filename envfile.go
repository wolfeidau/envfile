package envfile

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	commentMarker = "#"
	envSeperator  = "="
)

// Read the contents of a file, parse the contents and return a map containing the values.
//
// Parsing will stop if there are any issues parsing the file and report the line number
// which caused the error a line number.
//
func ReadEnvFile(filename string, envMap map[string]string) (err error) {

	// does it exist
	f, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNo := 1
	for scanner.Scan() {
		line := scanner.Text()

		cleanLine := trimComment(line)

		if strings.TrimSpace(cleanLine) != "" {
			err, key, value := parseVariable(cleanLine)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to parse env var on line %d - %s", lineNo, err))
			}
			envMap[key] = value
		}

		lineNo++
	}

	return nil
}

func parseVariable(line string) (err error, key string, value string) {

	delimiterIndex := strings.Index(line, envSeperator)

	if delimiterIndex != -1 {
		tokens := strings.SplitN(line, envSeperator, 2)
		return nil, strings.TrimSpace(tokens[0]), strings.TrimFunc(tokens[1], unicode.IsSpace)
	}
	return errors.New(fmt.Sprintf("Cannot locate token %s", envSeperator)), "", ""

}

func trimComment(line string) string {

	commentIndex := strings.Index(line, commentMarker)

	if commentIndex != -1 {
		return line[0:commentIndex]
	}
	return line

}

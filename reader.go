package gonf

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"
)

const REGEX_COMMENT = "#(.)*$"
const REGEX_KEY_VALUE = "^[0-9A-Za-z_]*:(.)*$"
const REGEX_KEY = "^[0-9A-Za-z_]*:"

func fetchFromFile(filename string, only []string) []Set {
	scanner, err := fileAsScanner(filename)
	if err != nil {
		return make([]Set, 0)
	}
	r, _ := regexp.Compile(REGEX_KEY)
	res := make([]Set, 0)
	filters := buildFilters(only)
	for scanner.Scan() {
		l := scanner.Text()
		if !r.MatchString(l) {
			continue
		}
		l = cleanLine(l)
		if isWildcard(only) {
			res = append(res, lineToSetUsingRegexp(r, l))
			continue
		}
		k := fetchKeyUsingRegexp(r, l)
		if _, ok := filters[k]; ok {
			res = append(res, lineToSetUsingRegexp(r, l))
		}
	}
	return res
}

func fileAsScanner(filename string) (*bufio.Scanner, error) {
	block, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(block)
	return bufio.NewScanner(reader), nil
}

func cleanLine(line string) string {
	// trim line
	line = strings.Trim(line, " ")
	// remove comments
	r, _ := regexp.Compile(REGEX_COMMENT)
	return r.ReplaceAllString(line, "")
}

func lineToSetUsingRegexp(r *regexp.Regexp, line string) Set {
	k := fetchKeyUsingRegexp(r, line)
	v := fetchRawValueUsingRegexp(r, line)
	return Set{
		Key:      strings.Trim(k, " "),
		RawValue: strings.Trim(v, " "),
	}
}

func fetchKeyUsingRegexp(r *regexp.Regexp, line string) string {
	k := r.FindString(line)
	return strings.TrimSuffix(k, ":")
}

func fetchRawValueUsingRegexp(r *regexp.Regexp, line string) string {
	return r.ReplaceAllString(line, "")
}

func isWildcard(only []string) bool {
	return len(only) == 0
}

func buildFilters(only []string) map[string]bool {
	filters := make(map[string]bool)
	for _, v := range only {
		filters[v] = true
	}
	return filters
}

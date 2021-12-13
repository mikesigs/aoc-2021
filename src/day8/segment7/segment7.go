package segment7

import (
	"fmt"
	"regexp"
	"strings"
)

type IPattern interface {
	Subtract(except string) string
	Length() int
}

type Pattern struct {
	Signals string
	Digit   int
}

var regexCache = map[string]*regexp.Regexp{}

func (p *Pattern) NonMatchingSignalCount(except string) int {
	regex := fmt.Sprintf("[^%s]", except)
	if _, exists := regexCache[regex]; !exists {
		regexCache[regex] = regexp.MustCompile(regex)
	}

	return len(regexCache[regex].FindAllString(p.Signals, -1))

}

func (p1 *Pattern) Subtract(p2 *Pattern) *Pattern {
	regex := fmt.Sprintf("[^%s]", p2.Signals)
	if _, exists := regexCache[regex]; !exists {
		regexCache[regex] = regexp.MustCompile(regex)
	}

	remaining := strings.Join(regexCache[regex].FindAllString(p1.Signals, -1), "")
	return NewPattern(remaining)
}

func (p *Pattern) Length() int {
	return len(p.Signals)
}

func NewPattern(pattern string) *Pattern {
	return &Pattern{Signals: sortString(pattern), Digit: -1}
}

func sortString(input string) string {
	var output string
	for r := 'a'; r <= 'g'; r++ {
		if strings.ContainsRune(input, r) {
			output += string(r)
		}
	}
	return output
}

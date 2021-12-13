package segment7

import (
	"fmt"
	"regexp"
	"strings"
)

type IPattern interface {
	NonMatchingSignalCount(except string) string
	SetDigit(digit int)
}

type Pattern struct {
	Signals       string
}

var regexCache = map[string]*regexp.Regexp{}

func (p *Pattern) NonMatchingSignalCount(except string) int {
	regex := fmt.Sprintf("[^%s]", except)
	if _, exists := regexCache[regex]; !exists {
		regexCache[regex] = regexp.MustCompile(regex)
	}

	return len(regexCache[regex].FindAllString(p.Signals, -1))

}

func NewPattern(pattern string) Pattern {
	return Pattern{Signals: sortString(pattern)}
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

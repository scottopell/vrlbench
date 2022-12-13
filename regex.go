package main

import (
	"fmt"
	"regexp"
)

type RegexProcessor struct {
	re *regexp.Regexp
}

func NewRegexProcessor(pattern string) *RegexProcessor {
	return &RegexProcessor{
		regexp.MustCompile(pattern),
	}
}

func (r *RegexProcessor) Process(in string) string {
	return r.re.ReplaceAllString(in, "gogo")
}

func (r *RegexProcessor) Describe() string {
	return fmt.Sprintf("Go Regex, pattern=%q", r.re.String())
}

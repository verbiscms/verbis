package templates

import "regexp"

func (t *TemplateFunctions) regexMatch(regex string, str string) bool {
	match, _ := regexp.MatchString(regex, str)
	return match
}

func (t *TemplateFunctions) regexFindAll(regex string, str string, i int) []string {
	r := regexp.MustCompile(regex)
	return r.FindAllString(str, i)
}

func (t *TemplateFunctions) regexFind(regex string, str string) string {
	r := regexp.MustCompile(regex)
	return r.FindString(str)
}

func (t *TemplateFunctions) regexReplaceAll(regex string, str string, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(str, repl)
}

func (t *TemplateFunctions) regexReplaceAllLiteral(regex string, str string, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllLiteralString(str, repl)
}

func (t *TemplateFunctions) regexSplit(regex string, str string, n int) []string {
	r := regexp.MustCompile(regex)
	return r.Split(str, n)
}

func (t *TemplateFunctions) regexQuoteMeta(str string) string {
	return regexp.QuoteMeta(str)
}
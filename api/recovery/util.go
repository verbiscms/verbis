package recovery

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/tpl"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// tplLineNumber
//
// Returns the line number of the template that broke.
// If the line number could not be retrieved using
// a regex find, -1 will be returned.
func tplLineNumber(err *errors.Error) int {
	e := getError(err)
	reg := regexp.MustCompile(`:\d+:`)
	lc := string(reg.Find([]byte(e.Error())))
	line := strings.ReplaceAll(lc, ":", "")

	i, ok := strconv.Atoi(line)
	if ok != nil {
		return -1
	}
	return i
}

// tplFileContents
//
// Gets the file contents of the errored template..
// Logs errors.NOTFOUND if the path could not be found.
func tplFileContents(path string) string {
	const op = "Recovery.tplFileContents"

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.WithFields(log.Fields{
			"error": &errors.Error{Code: errors.NOTFOUND, Message: "Could not get the file contents with the path: " + path, Operation: op, Err: err},
		})
		return ""
	}

	return string(contents)
}

// resolver
//
//
func (r *Recover) resolver(custom bool) (string, tpl.TemplateExecutor, bool) {
	code := strconv.Itoa(r.code)
	if r.code == 0 || code == "0" {
		code = "500"
	}

	path := ""
	e := r.deps.Tmpl().Prepare(tpl.Config{
		Root:      r.deps.Paths.Theme + "/" + r.deps.Theme.TemplateDir,
		Extension: r.deps.Theme.FileExtension,
	})

	// Look for `errors-404.cms` for example
	path = Prefix + "-" + code
	if e.Exists(path) && custom {
		return path, e, true
	}

	// Look for `error.cms` for example
	path = Prefix
	if e.Exists(Prefix) && custom {
		return path, e, true
	}

	// Return native error page
	path, exec := r.verbisErrorResolver()
	return path, exec, false
}

func (r *Recover) verbisErrorResolver() (string, tpl.TemplateExecutor) {
	return "templates/error", r.deps.Tmpl().Prepare(tpl.Config{
		Root:      r.deps.Paths.Web,
		Extension: ".html",
		Master:    "layouts/main",
	})
}

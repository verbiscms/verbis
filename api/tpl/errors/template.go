package errors

//import (
//	"github.com/ainsleyclark/verbis/api/errors"
//	"github.com/ainsleyclark/verbis/api/tpl"
//	log "github.com/sirupsen/logrus"
//	"io"
//	"io/ioutil"
//	"regexp"
//	"strconv"
//	"strings"
//)
//
//// TemplateRecovery
////
////
//type TemplateRecovery struct {
//	*Recovery
//	File    string
//	Exec    tpl.TemplateConfig
//	Name    string
//	Writer  io.Writer
//}
//
//// getTemplate
////
////
//func (r *Recovery) Template(tr *TemplateRecovery) *TemplateRecovery {
//	tr.InternalServerError()
//	return &TemplateRecovery{
//		r,
//		tr.File,
//		tr.Exec,
//		tr.Name,
//		tr.Writer,
//	}
//}
//
//func (t *TemplateRecovery) getTemplate() *FileStack {
//	file := &FileStack{
//		File:     t.Exec.GetRoot() + "/" + t.File + t.Exec.GetExtension(),
//		Line:     t.lineNumber(),
//		Name:     t.File,
//		Contents: t.fileContents(),
//	}
//	return file
//}
//
//// tplLineNumber
////
////
//func (t *TemplateRecovery) lineNumber() int {
//	e := t.getError()
//	reg := regexp.MustCompile(`:\d+:`)
//	lc := string(reg.Find([]byte(e.Error())))
//	line := strings.ReplaceAll(lc, ":", "")
//
//	i, err := strconv.Atoi(line)
//	if err != nil {
//		return -1
//	}
//	return i
//}
//
//// tplFileContents gets the file contents of the errored file.
//// Returns INTERNAL if the path could not be found
//func (t *TemplateRecovery) fileContents() string {
//	const op = "Recovery.tplFileContents"
//
//	path := t.Exec.GetRoot() + "/" + t.File + t.Exec.GetExtension()
//	contents, err := ioutil.ReadFile(path)
//	if err != nil {
//		log.WithFields(log.Fields{
//			"error": &errors.Error{Code: errors.NOTFOUND, Message: "Could not get the file contents with the path: " + path, Operation: op, Err: err},
//		})
//		return ""
//	}
//
//	return string(contents)
//}

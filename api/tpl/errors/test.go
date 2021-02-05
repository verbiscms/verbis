package errors

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/spf13/cast"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

// If theres an error executing the template
// it will return the bytes of the error file
// with all the stack info etc
// Its up to the caller to log or do whatever it wants
// to do with the file contents


type Tester interface {

}

type Recovery struct {
	File    string
	Error   interface{}
	Deps    *deps.Deps
	Writer  io.Writer
	Name    string
	Context *gin.Context
	Post    *domain.PostData
	Exec    tpl.TemplateConfig
}


type tplData struct {
	Error      errors.Error
	Template  *FileStack
	Stack      []*FileStack
	SubMessage string
	File       []FileLine
	Highlight  int
	LineNumber int

}

// FileLine defines the error for templating it includes the
// line & content of the error file.
type FileLine struct {
	Line    int
	Content string
}

func New(r Recovery) *Recovery {
	return &r
}

func (r *Recovery) Recover() {

	t := r.Deps.Tmpl().Prepare(tpl.Config{
		Root:      r.Deps.Paths.Web,
		Extension: ".html",
		Master:    "layouts/main",
	})

	var b bytes.Buffer
	err := t.Execute(&b, "/templates/error", r.data())
	if err != nil {
		color.Green.Println(err)
		// We can't continue
		panic(err)
	}

	_, err = r.Writer.Write(b.Bytes())
	if err != nil {
		panic(err)
	}
}

func (r *Recovery) data() *tplData {

	//contents, err := r.fileContents()
	//if err != nil {
		// log
		//contents = ""
	//}

	return &tplData{
		Error:      r.getError(),
		Stack:      Stack(StackDepth, TraverseLength),
		SubMessage: "",
		//File:       fileLines(contents, 10),
		Highlight:  0,
		LineNumber: cast.ToInt(r.lineNumber()),
	}
}

// getFileContents gets the file contents of the errored file.
// Returns INTERNAL if the path could not be found
func (r *Recovery) fileContents() (string, error) {
	const op = "Recovery.getFileContents"

	path := r.Exec.GetRoot() + "/" + r.File + r.Exec.GetExtension()
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", &errors.Error{Code: errors.NOTFOUND, Message: "Could not get the file contents with the path: " + path, Operation: op, Err: err}
	}

	return string(contents), nil
}

func (r *Recovery) getError() errors.Error {
	var err errors.Error

	switch v := r.Error.(type) {
	case errors.Error:
		err = v
	case *errors.Error:
		err = *v
	case error:
		err = errors.Error{Code: errors.TEMPLATE, Message: "Unable to execute template", Operation: "TemplateEngine.Execute", Err: v}
	default:
		err = errors.Error{Code: errors.TEMPLATE, Message: "Unable to execute template", Operation: "TemplateEngine.Execute", Err: fmt.Errorf("templte engine: unable to execute template")}
	}

	return err
}

func (r *Recovery) lineNumber() string {
	e := r.getError()
	reg := regexp.MustCompile(`:\d+:`)
	lc := string(reg.Find([]byte(e.Error())))
	return strings.ReplaceAll(lc, ":", "")
}

// Lines gets the range of lines of a file in between a limit
// Returns an array of file lines
func (r *Recovery) getFileLines(file string, line int, limit int) []FileLine {
	split := strings.Split(file, "\n")

	var fileLines []FileLine
	counter := line - len(split)
	for i := 0; i < len(split) * 2; i++ {
		if counter >= 0 && counter < len(split) {
			fileLines = append(fileLines, FileLine{
				Line:    counter + 1,
				Content: html.UnescapeString(strings.Replace(split[counter], " ", "&nbsp;", -1)),
			})
		}
		counter++
	}

	return fileLines
}



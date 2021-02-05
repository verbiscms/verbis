package errors

import (
	"io/ioutil"
	"runtime"
	"strings"
)

const (
	// The amount of files in the stack to be retrieved
	StackDepth = 200
	// How many lines before and after the calling function
	// to retrieve
	LineLimit = 50
	// How many files to move up in the runtime.Caller
	// before obtaining the stack
	TraverseLength = 2
)

// FileStack defines the stack used for the error page
type FileStack struct {
	File string
	Line int
	Name string
	Contents string
}

// Stack
//
// Returns a slice of FileStack's by traversing the caller.
// using the depth and traverse arguments to loop over.
// If there was an error reading the file, or the
// runtime.Caller function failed, it will not
// be appended to the stack.
func Stack(depth int, traverse int) []*FileStack {
	var stack []*FileStack

	for c := traverse; c < depth; c++ {
		t, file, line, ok := runtime.Caller(c)

		if !ok {
			continue
		}

		contents, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}

		stack = append(stack, &FileStack{
			File: file,
			Line: line,
			Name: runtime.FuncForPC(t).Name(),
			Contents: string(contents),
		})
	}

	return stack
}

// Lines
//
// Splits the file into a array of lines by separating
// them by a new line.
func (f *FileStack) Lines() []*FileLine {
	lines := strings.Split(f.Contents, "\n")

	diff := LineLimit / 2

	var fileLines []*FileLine
	counter := 0
	for i := f.Line - diff; i < f.Line + diff; i ++ {
		if i > 0 && i < len(lines) {
			fileLines = append(fileLines, &FileLine{
				Line:    i + 1,
				Content: lines[i],
			})
		}
		counter++
	}

	return fileLines
}
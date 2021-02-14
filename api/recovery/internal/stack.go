package internal

//go:generate go run ../generator/main.go

// This needs to be cleaned up and look at the
// func namespace registry for templates

type stack struct {
	storage map[string][]byte
}

// Create new box for embed files
func newEmbedBox() *stack {
	return &stack{storage: make(map[string][]byte)}
}

// Add a file to box
func (e *stack) Add(file string, content []byte) {
	e.storage[file] = content
}

// Get file's content
func (e *stack) Get(file string) []byte {
	if f, ok := e.storage[file]; ok {
		return f
	}
	return nil
}

// Embed box expose
var stackRegistry = newEmbedBox()

// Add a file content to box
func Add(file string, content []byte) {
	stackRegistry.Add(file, content)
}

// Get a file from box
func Get(file string) []byte {
	return stackRegistry.Get(file)
}

package runtime

// File represents a file content.
type File struct {
	Path string
	Body string
}

// Entry represents a file that will be generated.
type Entry struct {
	File
	Template File
}

type ShouldRunFunc func(e *Entry) bool

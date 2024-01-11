package entity

type File struct {
	Name    string
	Path    string
	Content []byte
}

func NewFile(name string, path string, content []byte) *File {
	return &File{Name: name, Path: path, Content: content}
}

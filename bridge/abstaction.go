package main

// Abstraction
type File struct {
    fileOperation FileOperationImplementor
}

func (f *File) Open() {
    f.fileOperation.Open()
}

func (f *File) Read() {
    f.fileOperation.Read()
}

func (f *File) Write(data string) {
    f.fileOperation.Write(data)
}

func (f *File) Close() {
    f.fileOperation.Close()
}
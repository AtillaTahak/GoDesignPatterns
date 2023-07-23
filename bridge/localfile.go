package main
import "fmt"

// Concrete Implementor 1
type LocalFile struct {
    filePath string
}

func (l *LocalFile) Open() {
    fmt.Println("Opening local file:", l.filePath)
}

func (l *LocalFile) Read() {
    fmt.Println("Reading from local file")
}

func (l *LocalFile) Write(data string) {
    fmt.Println("Writing to local file:", data)
}

func (l *LocalFile) Close() {
    fmt.Println("Closing local file")
}


package main
import "fmt"
// Concrete Implementor 2
type RemoteFile struct {
    remoteURL string
}

func (r *RemoteFile) Open() {
    fmt.Println("Opening remote file:", r.remoteURL)
}

func (r *RemoteFile) Read() {
    fmt.Println("Reading from remote file")
}

func (r *RemoteFile) Write(data string) {
    fmt.Println("Writing to remote file:", data)
}

func (r *RemoteFile) Close() {
    fmt.Println("Closing remote file")
}
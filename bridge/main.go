package main

func main() {
    localFile := &LocalFile{filePath: "example.txt"}
    remoteFile := &RemoteFile{remoteURL: "http://example.com/data.txt"}

    localFileOperation := &File{fileOperation: localFile}
    remoteFileOperation := &File{fileOperation: remoteFile}

    localFileOperation.Open()
    localFileOperation.Read()
    localFileOperation.Write("Hello, Bridge Pattern!")
    localFileOperation.Close()

    remoteFileOperation.Open()
    remoteFileOperation.Read()
    remoteFileOperation.Write("Hello, Remote Bridge Pattern!")
    remoteFileOperation.Close()
}

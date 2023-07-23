package main

// Implementor interface
type FileOperationImplementor interface {
    Open()
    Read()
    Write(data string)
    Close()
}


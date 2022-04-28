package main

// in io package:

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// ReadWriter embedding an interface
type ReadWriter interface {
	Reader
	Writer

	//Read(p []byte) (n int, err error)
	//Write(p []byte) (n int, err error)

	//Read(p []byte) (n int, err error)
	//Writer
}

// ReadWriteCloser embedding an interface
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

func main() {

}

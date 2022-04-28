package main

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	//pStr := string(p)
	*c += WordCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	*c += LineCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {

}

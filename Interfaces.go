package main

import "fmt"

func main() {
	// -----> INTERFACES <-----

	var cw Writer = ConsoleWriter{}
	var fw Writer = FileWriter{}
	cw.Write()
	fw.Write()

	fmt.Println("\n--Compose--\n")
	var all WriterCloser = WriterCloserBuffer{}
	fmt.Println(all.Write())
	fmt.Println(all.Close())
}

// describe behaviors
// best practices - if we have one method, the intrface is called method+er, so Write+er(r) = Writer
type Writer interface {
	// creating methods
	Write() (int, error)
}

// implementing the interface --- START
type ConsoleWriter struct {

}
func (cw ConsoleWriter) Write() (int, error) {
	n, err := fmt.Println("Console")
	return n, err
}
// implementing the interface --- END

// second implementation
// implementing the interface  Writer--- START
type FileWriter struct {

}

func (fw FileWriter) Write() (int, error) {
	n, err := fmt.Println("FileWriter")
	return n, err
}
// implementing the interface  Writer--- END


type Closer interface {
	Close() (int, error)
}

type FileCloser struct {

}

func (fc FileCloser) Close() (int, error) {
	n, err := fmt.Println("FileCloser")
	return n, err
}

// COMPOSE INTERFACES TOGETHER
// only for example, not meaningful

// for this composition we need to implement another struct that is implementing
// WriterClose interface functions, so functions of the embedded interfaces
// Write from Writer and Close from Closer
type WriterCloser interface {
	Writer
	Closer
}
// implementing WriterCloser
type WriterCloserBuffer struct {

}

func (wcb WriterCloserBuffer) Write() (int, error) {
	n, err := fmt.Println("Writer Closer Interface - WRITE")
	return n, err
}

func (wcb WriterCloserBuffer) Close() (int, error) {
	n, err := fmt.Println("Writer Closer Interface - CLOSE")
	return n, err
}
// if we implement one interface function with pointer receiver func (wd *Writer)
// the interface is expecting reference object &struct{}

// if there is no pointer receiver functions interface expects just struct{}

// best practices
// 1 - use many, small interfaces (prefer single method interface)






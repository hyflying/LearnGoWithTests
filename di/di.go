package di

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func main() {
	Greet(os.Stdout, "Elodie")
	Countdown(os.Stdout)
}

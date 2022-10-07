package main

import (
	"flag"
	"fmt"
)

var (
	inputName    = flag.String("name", "CHENJIAN", "Input Your Name.")
	inputFlagvar int
)

func Init() {
	flag.IntVar(&inputFlagvar, "flagname", 1234, "Help")
}

func main() {
	Init()
	flag.Parse()

	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("name=", *inputName)
}

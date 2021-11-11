package main

import "fmt"
import "os"
import "strings"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("wordcount err: ")
		os.Exit(1)
	}
	if len(os.Args[1]) == 0 {
		fmt.Println("0")
	} else {
		arW := strings.Split(os.Args[1]," ")
		fmt.Println(len(arW))
	}
}

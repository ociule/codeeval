package main

import "fmt"
import "log"
import "os"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
    fi, _ := file.Stat()
    fmt.Println(fi.Size())
}

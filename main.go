package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("shadow.txt")
	checkErr(err)
	defer file.Close()
	finfo, err := file.Stat()
	checkErr(err)
	filesize := finfo.Size()
	fmt.Println(filesize)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

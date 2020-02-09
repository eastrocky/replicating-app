package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	exe, _ := os.Executable()
	fmt.Println(exe)
	input, err := ioutil.ReadFile(exe)
	if err != nil {
		fmt.Println(err)
		return
	}

	destinationFile := fmt.Sprintf("clones/%d", time.Now().Unix())
	err = ioutil.WriteFile(destinationFile, input, 0755)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func main() {
	exe, _ := os.Executable()
	input, _ := ioutil.ReadFile(exe)

	destinationFile := fmt.Sprintf("clones/%d", time.Now().Unix())
	ioutil.WriteFile(destinationFile, input, 0755)

	exec.Command(destinationFile).Start()
}

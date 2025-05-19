package main

import (
	"fmt"
	"os"

	"github.com/mmmmmmpc/uyuni-tools/uyunictl/cmd"
)

func Run() error {
	// Call your existing Run logic
	return cmd.NewUyunictlCommand().Execute()
}

func main() {
	err := Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mmmmmmpc/uyuni-tools/uyunictl/cmd"
)

func Run() error {
	// Command-line options
	user := flag.String("user", "", "Uyuni server username")
	password := flag.String("password", "", "Uyuni server password")
	apiURL := flag.String("apiURL", "", "Uyuni server API URL")
	flag.Parse()

	// Validate command-line options
	if *user == "" || *password == "" || *apiURL == "" {
		fmt.Println("Please provide the --user, --password, and --apiURL options")
		flag.PrintDefaults()
		os.Exit(1)
	}

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

package main

import (
	"os"
	"fmt"
	"flag"

	"github.com/uyuni-project/uyuni-tools/uyunictl/cmd"
)

// Run runs the `uyunictl` root command
func Run() error {
	return cmd.NewUyunictlCommand().Execute()
}

func main() {
	if err := Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	// Command-line options
	user := flag.String("user", "", "SUSE Manager username")
	password := flag.String("password", "", "SUSE Manager password")
	apiURL := flag.String("apiURL", "", "SUSE Manager API URL")
	flag.Parse()
	
	// Validate command-line options
	if *user == "" || *password == "" || *apiURL == "" {
		fmt.Println("Please provide the --user, --password, and --apiURL options")
		flag.PrintDefaults()
		os.Exit(1)
	}
	
	// Call the function to execute the request and print the data
	err := ExecuteRequest(*user, *password, *apiURL)	
}


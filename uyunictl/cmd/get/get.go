package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type System struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
}

func ExecuteRequest(user, password, apiURL string) error {
	// Create an HTTP client and make a request to authenticate
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.SetBasicAuth(user, password)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("authentication failed with status code: %d", resp.StatusCode)
	}

	// Make a request to retrieve the list of systems
	req, err = http.NewRequest("GET", apiURL+"system.listSystems", nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.SetBasicAuth(user, password)
	resp, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to retrieve systems: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to retrieve systems with status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the JSON response
	var systems []System
	err = json.Unmarshal(body, &systems)
	if err != nil {
		return fmt.Errorf("failed to parse response body: %v", err)
	}

	// Print the list of systems
	for _, sys := range systems {
		fmt.Printf("ID: %d, Name: %s, Hostname: %s\n", sys.ID, sys.Name, sys.Hostname)
	}

	return nil
}

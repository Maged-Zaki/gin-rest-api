package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PrettyPrintRequest(r *http.Request) {
	// Create a map to hold the request data
	reqData := map[string]interface{}{
		"Method": r.Method,
		"URL":    r.URL,
		"Header": r.Header,
		"Body":   r.Body,
	}

	// Marshal the request data to pretty JSON
	prettyJSON, err := json.MarshalIndent(reqData, "", "  ")
	if err != nil {
		fmt.Println("Failed to generate JSON:", err)
		return
	}

	// Print the pretty JSON with color
	fmt.Printf("\033[1;32m>>> %s\033[0m\n", string(prettyJSON))
}

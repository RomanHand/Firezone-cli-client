// firezone_client.go
package sender

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func Send(apiURL string, apiToken string, endpoint string, body io.Reader, method string) ([]byte) {
	client := &http.Client{}
	
	req, err := http.NewRequest(method, apiURL+"/v0/"+endpoint, body) 
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()
	if strings.HasPrefix(http.StatusText(resp.StatusCode), "2") {
		log.Fatalf("Error: received non-200 status code: %d\n", resp.StatusCode)
	}
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	fmt.Println(string(data))
	// data = []byte(strings.Replace(string(data), "%", "", -1))

	return data
	}
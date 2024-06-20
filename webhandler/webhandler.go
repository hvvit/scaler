/*
this package contains the functions to make get and post requests to the server
All the types of requests are implemented in this package
*/
package webhandler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

/*
this function accepts host address and path and returns the formatted url
host address should be in the format of "http://ip:port"
and path can be any string
for example if urlFormat("localhost:8123", "/app/status") is called it returns "http://localhost:8123/app/status"
*/
func urlFormat(host_addr string, path string) (string, error) {
	// format the url
	u, err := url.Parse(host_addr)
	if err != nil {
		return "", err

	}
	if u.Scheme == "" {
		return "", errors.New("scheme is missing in the host address")
	}
	address, err := url.JoinPath(host_addr, path)
	return address, err
}

/*
this function accepts host address and path
and returns the number of replicas and the current cpu utilization
here we return current replicas, current cpu utilization and error if any
logic to change the replicas is implemented in the watcher package
*/
func CheckGetRequest(host_addr string, path string) (int, float64, error) {
	// implement the get request to the server
	host_path, err := urlFormat(host_addr, path)
	if err != nil {
		return 0, 0.0, err
	}
	// make the get request to the server

	client := &http.Client{}
	defer client.CloseIdleConnections()
	req, err := http.NewRequest("GET", host_path, nil)
	if err != nil {
		return 0, 0.0, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0.0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		error_message := fmt.Sprintf("Error in getting the response from the server with status code: %d", resp.StatusCode)
		return 0, 0.0, errors.New(error_message)
	}
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return 0, 0.0, err
	}
	log.Println("current cpu utilization for host:", host_addr, "is", response.CPU.HighPriority)
	return response.Replicas, response.CPU.HighPriority, nil
}

/*
this function accepts host address, path and the number of replicas
and updates the number of replicas on the server
*/
func ChangeReplica(host_addr string, path string, replicas int) error {
	// implement the post request to the server
	host_path, err := urlFormat(host_addr, path)
	if err != nil {
		return err
	}
	// make the post request to the server
	client := &http.Client{}
	defer client.CloseIdleConnections()

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(PostResponse{Replicas: replicas})
	req, err := http.NewRequest("PUT", host_path, payloadBuf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// check if the status code is 204
	if resp.StatusCode != http.StatusNoContent {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			error_message := fmt.Sprintf("Error in getting the response from the server with status code: %d, with message: %s", resp.StatusCode, string(bodyBytes))
			return errors.New(error_message)
		}
		error_message := fmt.Sprintf("Error in getting the response from the server with status code: %d, with message: %s", resp.StatusCode, string(bodyBytes))
		return errors.New(error_message)
	}
	log.Println("replicas are updated to:", replicas)
	return nil
}

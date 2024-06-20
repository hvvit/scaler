/*
this file will handle all the structs required for handling web requests
*/
package webhandler

// Response struct to store the response from the server
type Response struct {
	CPU struct {
		HighPriority float64 `json:"highPriority"`
	} `json:"cpu"`
	Replicas int `json:"replicas"`
}

// PUT body requests struct
type PostResponse struct {
	Replicas int `json:"replicas"`
}

// Description: This file contains the logic to check the current replica and metrics value of the server
package watcher

import (
	"errors"
	"log"
	"math"

	"github.com/hvvit/scaler/webhandler"
)

// checkAndUpdateReplica function will check the current replica and metrics value of the server
// watcher will look for this function to implement the logic to check the current replica and metrics value of the server
func checkAndUpdateReplica(host_addr string, threshold_value float64) {
	// call the checkGetRequest function from webhandler package
	current_replicas, current_metrics_value, err := webhandler.CheckGetRequest(host_addr, "/app/status")
	if err != nil {
		log.Println("error in getting the response from the server", err)
	} else {
		log.Println("current replicas are:", current_replicas, ",current metrics value is:", current_metrics_value)
		if isChangeRequired(current_metrics_value, threshold_value) {
			err := updateReplica(host_addr, current_replicas, threshold_value, current_metrics_value)
			if err != nil {
				log.Println("error in updating the replica", err)
			}
		}
	}
}

// calculateDesiredReplica function will calculate the desired replica based on the current replica, threshold value and current metrics value
// we will use the algorithm mentioned in the kubernetes documentation
func calculateDesiredReplica(current_replicas int, threshold_value float64, current_metrics_value float64) (int, error) {
	// here i will try to implement kubentes autoscaler algorithm mentioned in the documentation
	// https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#algorithm-details

	if current_metrics_value == 0 {
		return 0, errors.New("current metrics value is zero")
	}
	if threshold_value == 0 {
		return 0, errors.New("threshold value is zero")
	}
	desiredReplicas := int(math.Ceil((current_metrics_value / threshold_value) * float64(current_replicas)))
	return desiredReplicas, nil
}

// updateReplica function will update the replica of the server
// if the desired replica is not equal to the current replica it will update the replica
func updateReplica(host_addr string, current_replicas int, threshold_value float64, current_metrics_value float64) error {
	desiredReplicas, err := calculateDesiredReplica(current_replicas, threshold_value, current_metrics_value)
	if err != nil {
		return err
	}
	log.Println("desired replicas are:", desiredReplicas)
	if desiredReplicas != current_replicas || desiredReplicas > 1 {
		err := webhandler.ChangeReplica(host_addr, "/app/replicas", desiredReplicas)
		if err != nil {
			return err
		}
	}
	return nil
}

// isChangeRequired function will check if the change is required or not
// logic for scale-in and scale-out can be implemented here.
func isChangeRequired(current_metrics_value float64, threshold_value float64) bool {
	return current_metrics_value != threshold_value
}

/*
Package watcher is responsible for watching the cpu utilization of the server and updating the replica count of the server.
*/
package watcher

import (
	"log"
	"time"
)

func Watch(threshold_value float64, scrape_time float64, host_addr string) {
	// implement watcher function, where it will sleep for scrape_time and then check the cpu utilization of the server
	// if the cpu utiliztion is not the value of threshold_value it will calculate the threshold value again.
	log.Println("Watcher is running")
	for {
		log.Println("will scrape url:", host_addr, "after", scrape_time, "seconds")
		time.Sleep(time.Duration(scrape_time) * time.Second)
		checkAndUpdateReplica(host_addr, threshold_value)
	}
}

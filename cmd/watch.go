/*
Copyright © 2024 HARSH VARDHAN <harsh.vardhan7896@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/hvvit/scaler/watcher"
	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "sub command to watch the server load and change replica if required",
	Long:  `takes parameters like threshold value, scrape time and url to watch the server load and change replica if required`,
	Run: func(cmd *cobra.Command, args []string) {
		threshold_value, _ := cmd.Flags().GetFloat64("threshold")
		scrape_time, _ := cmd.Flags().GetFloat64("scrape-time")
		url, _ := cmd.Flags().GetString("url")
		log.Println("watching the server load with threshold value: ", threshold_value, " scrape time: ", scrape_time, " url: ", url)
		// call the watcher function from watcher package
		watcher.Watch(threshold_value, scrape_time, url)
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
	watchCmd.Flags().Float64P("threshold", "t", 0.8, "Threshold value for the server load to monitor")
	watchCmd.Flags().Float64P("scrape-time", "s", 1.0, "Time to wait before scraping the server load")
	watchCmd.Flags().StringP("url", "u", "http://localhost:8123", "URL of the server to monitor")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// watchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

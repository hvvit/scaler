/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/hvvit/scaler/cmd"

// main function to execute the command
/*
A scaler is a tool to scale the upstream server based on the load of the server.
        The purpose of this exercise is to write an auto-scaler that manipulates the
number of replicas of a separate application based on CPU utilization metrics.
As your auto-scaler changes the replica count, CPU utilization will be impacted
accordingly.

Usage:
  scaler [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  watch       sub command to watch the server load and change replica if required

Flags:
  -h, --help      help for scaler
  -t, --toggle    Help message for toggle
  -v, --version   version for scaler

Use "scaler [command] --help" for more information about a command.
*/
func main() {
	cmd.Execute()
}

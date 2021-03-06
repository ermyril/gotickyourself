package main

import (
	"github.com/spf13/cobra"
)

func getUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:    "update",
		Short:  "Updates roles, projects and tasks list. This is done automatically on init, and before the first request of the day",
		Long:   ``,
		PreRun: initConfigFiles,
		Run:    runUpdateCmd,
	}
}

func runUpdateCmd(cmd *cobra.Command, args []string) {
	updateClients()
	updateProjects()
}

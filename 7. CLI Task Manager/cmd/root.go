package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Tasks is a cli tool to manage tasks",
}

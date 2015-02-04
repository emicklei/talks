package main

import "github.com/spf13/cobra"

var TorchCmd = &cobra.Command{
	Use:   "torch",
	Short: "torch a is a tool to inspect/debug/command a firespark server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

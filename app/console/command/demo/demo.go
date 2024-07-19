package demo

import (
	"fmt"
	"github.com/gohade/hade/framework/cobra"
	"log"
)

var Demo2Command = &cobra.Command{
	Use:   "demo2",
	Short: "demo for app",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		log.Println(container)
		fmt.Println("<---demo2--->")
		return nil
	},
}

package command

import "github.com/gohade/hade/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(initEnvCommand())
	//root.AddCommand(DemoCommand)
	root.AddCommand(initAppCommand())
	// cron
	root.AddCommand(initCronCommand())
}

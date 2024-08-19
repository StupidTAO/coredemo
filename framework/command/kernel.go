package command

import "github.com/gohade/hade/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// config 命令
	root.AddCommand(initConfigCommand())
	// env 命令
	root.AddCommand(initEnvCommand())
	// app 命令
	root.AddCommand(initAppCommand())
	// cron 命令
	root.AddCommand(initCronCommand())
	// build 命令
	root.AddCommand(initBuildCommand())
	// go build
	root.AddCommand(goCommand)
	// npm build
	root.AddCommand(npmCommand)
}

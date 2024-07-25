package console

import (
	"github.com/gohade/hade/app/console/command/demo"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/cobra"
	"github.com/gohade/hade/framework/command"
)

// RunCommand 初始化根Command并运行
func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		// 定义命令关键字
		Use: "hade",
		// 简短介绍
		Short: "hade 命令",
		// 根命令的详细介绍
		Long: "hade 框架命令行",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现cobra默认的completion字命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	rootCmd.SetContainer(container)
	// 绑定框架的命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)

	return rootCmd.Execute()
}

// 绑定业务的命令
func AddAppCommand(rootCmd *cobra.Command) {
	// demo 例子
	rootCmd.AddCommand(demo.InitFoo())
	rootCmd.AddCronCommand("* * * * *", demo.InitFoo())
}

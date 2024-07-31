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
	rootCmd.AddCommand(demo.FooCommand)

	// 每秒带调用一次Foo命令
	//rootCmd.AddCronCommand("* * * * * *", demo.FooCommand)

	// 启动一个分布式任务调度，调度的服务名称为init_func_for_test，每个节点每5s调用一次Foo命令，抢占到了调度任务的节点将抢占锁持续挂载2s才释放
	//rootCmd.AddDistributedCronCommand("foo_func_for_test", "*/5 * * * * *", demo.FooCommand, 2*time.Second)
}

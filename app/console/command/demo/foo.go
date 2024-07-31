package demo

import (
	"errors"
	"fmt"
	"github.com/gohade/hade/framework/cobra"
	"github.com/gohade/hade/framework/contract"
	"log"
)

// InitFoo 初始化Foo命令
func InitFoo() *cobra.Command {
	FooCommand.AddCommand(Foo1Command)
	return FooCommand
}

// FooCammand 代表Foo命令
var FooCommand = &cobra.Command{
	Use:     "foo",
	Short:   "foo的简要说明",
	Long:    "foo的长说明",
	Aliases: []string{"fo", "f"},
	Example: "foo命令的例子",
	RunE: func(cmd *cobra.Command, args []string) error {
		configService, ok := cmd.GetContainer().MustMake(contract.ConfigKey).(contract.Config)
		if !ok {
			return errors.New("configService contro is failed")
		}
		envService, ok := cmd.GetContainer().MustMake(contract.EnvKey).(contract.Env)
		if !ok {
			return errors.New("envService contro is failed")
		}
		fmt.Println("APP_ENV: ", envService.Get("APP_ENV"))
		fmt.Println("FOO_ENV: ", envService.Get("FOO_ENV"))
		fmt.Println("config url: ", configService.GetString("app.url"))

		return nil
	},
}

// Foo1Command 代表Foo命令的字命令Foo1
var Foo1Command = &cobra.Command{
	Use:     "foo1",
	Short:   "foo1的简要说明",
	Long:    "foo1的长说明",
	Aliases: []string{"fo1", "f1"},
	Example: "foo1命令的例子",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		log.Println("foo1 ", container)
		return nil
	},
}

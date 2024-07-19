package command

import (
	"fmt"
	"github.com/gohade/hade/framework/cobra"
	"github.com/gohade/hade/framework/contract"
)

// helpCommand show current envioment
var DemoCommand = &cobra.Command{
	Use:   "demo",
	Short: "demo for framework",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		appService, _ := container.MustMake(contract.AppKey).(contract.App)
		fmt.Println("app base folder: ", appService.BaseFolder())
	},
}

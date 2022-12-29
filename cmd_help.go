package goter

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

// 环境变量，是否打印当前版本号
var envVersionShowRuntime = false

// newVersionCmd 创建一个版本信息查看命令
func newVersionCmd(use, version string) *Command {
	return (&Command{Cmd: &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "打印当前版本号",
		Run: func(cmd *cobra.Command, args []string) {
			if envVersionShowRuntime {
				fmt.Println(fmt.Sprintf("%s version: v%s(%s/%s)", use, version, runtime.GOOS, runtime.GOARCH))
			} else {
				fmt.Println(fmt.Sprintf("%s version: v%s", use, version))
			}
		},
	}}).Bind(NewCmdFlagBool(false, "runtime", "R", "show runtime or not"))
}

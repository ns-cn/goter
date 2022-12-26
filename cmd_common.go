package goter

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

// newRootCmd 创建一个根命令
func newRootCmd(use, short string) *Command {
	return &Command{Cmd: &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}}
}

// newVersionCmd 创建一个版本信息查看命令
func newVersionCmd(use, version string) *Command {
	return &Command{Cmd: &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "打印当前版本号",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf("%s version: v%s(%s/%s)", use, version, runtime.GOOS, runtime.GOARCH))
		},
	}}
}

package goter

import (
	"github.com/spf13/cobra"
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

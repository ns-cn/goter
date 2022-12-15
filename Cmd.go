package goter

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

func Root(use, short, version string) *cobra.Command {
	root := &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	root.AddCommand(&cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "打印当前版本号",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf("%s version: v%s(%s/%s)", use, version, runtime.GOOS, runtime.GOARCH))
		},
	})
	return root
}

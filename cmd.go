package goter

import (
	"github.com/spf13/cobra"
)

// NewRootCmd 创建根命令，包含一个根命令以及根命令下的版本信息查看命令
func NewRootCmd(use, short, version string) *Command {
	root := newRootCmd(use, short)
	commandVersion := newVersionCmd(use, version)
	root.AddCommand(commandVersion)
	return root
}

// CmdStringFlag 命令的字符串参数
type CmdStringFlag struct {
	P               *string
	Name, Shorthand string
	Value           string
	Usage           string
}

// Command 命令的包装类
type Command struct {
	Cmd *cobra.Command
}

// Bind 绑定指定的参数，可指定具体的字符串参数
func (c *Command) Bind(flags ...CmdStringFlag) *Command {
	for _, flag := range flags {
		c.Cmd.Flags().StringVarP(flag.P, flag.Name, flag.Shorthand, flag.Value, flag.Usage)
	}
	return c
}

// AddCommand 为当前命令添加子命令
func (c *Command) AddCommand(subCommand *Command) *Command {
	c.Cmd.AddCommand(subCommand.Cmd)
	return c
}

// Execute 执行当前命令
func (c *Command) Execute() error {
	return c.Cmd.Execute()
}

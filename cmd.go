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

// Command 命令的包装类
type Command struct {
	Cmd *cobra.Command
}

// Bind 绑定指定的参数，可指定具体的字符串参数
func (c *Command) Bind(flags ...CmdFlagBinder) *Command {
	for _, flag := range flags {
		flag.Bind(c.Cmd)
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

// NewSubCommand 针对当前命令创建子命令：返回新创建命令
func (c *Command) NewSubCommand(sub *cobra.Command) *Command {
	subCommand := &Command{
		Cmd: sub,
	}
	c.AddCommand(subCommand)
	return subCommand
}

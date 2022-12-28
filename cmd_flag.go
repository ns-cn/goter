package goter

import "github.com/spf13/cobra"

type CmdFlag interface {
	Bind(c *cobra.Command)
}

// CmdFlagString 命令的字符串参数
type CmdFlagString struct {
	CmdFlag
	Name      string
	Shorthand string
	Usage     string

	P     *string
	Value string
}

func (f CmdFlagString) Bind(c *cobra.Command) {
	c.Flags().StringVarP(f.P, f.Name, f.Shorthand, f.Value, f.Usage)
}

// CmdFlagBool 命令的bool类型参数
type CmdFlagBool struct {
	CmdFlag
	Name      string
	Shorthand string
	Usage     string
	P         *bool
	Value     bool
}

func (f CmdFlagBool) Bind(c *cobra.Command) {
	c.Flags().BoolVarP(f.P, f.Name, f.Shorthand, f.Value, f.Usage)
}

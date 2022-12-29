package goter

import "github.com/spf13/cobra"

type CmdFlag interface {
	// Bind 命令参数的绑定方式声明
	Bind(c *cobra.Command)
}

// NewCmdFlagString 创建新的命令行字符串参数
func NewCmdFlagString(p *string, value string, name string, shorthand string, usage string) CmdFlagString {
	return CmdFlagString{P: p, Name: name, Shorthand: shorthand, Usage: usage, Value: value}
}

// NewCmdFlagStringSlice 创建新的命令行字符串数组参数
func NewCmdFlagStringSlice(p *[]string, value []string, name string, shorthand string, usage string) CmdFlagStringSlice {
	return CmdFlagStringSlice{P: p, Name: name, Shorthand: shorthand, Usage: usage, Value: value}
}

// NewCmdFlagBool 创建新的命令行bool参数
func NewCmdFlagBool(p *bool, value bool, name string, shorthand string, usage string) CmdFlagBool {
	return CmdFlagBool{P: p, Name: name, Shorthand: shorthand, Usage: usage, Value: value}
}

// NewCmdFlagInt 创建新的命令行int参数
func NewCmdFlagInt(p *int, defaultValue int, name string, shorthand string, usage string) CmdFlagInt {
	return CmdFlagInt{P: p, Name: name, Shorthand: shorthand, Usage: usage, Value: defaultValue}
}

// CmdFlagString 命令的字符串参数
type CmdFlagString struct {
	CmdFlag
	Name      string
	Shorthand string
	Usage     string
	P         *string
	Value     string
}

func (f CmdFlagString) Bind(c *cobra.Command) {
	c.Flags().StringVarP(f.P, f.Name, f.Shorthand, f.Value, f.Usage)
}

// CmdFlagStringSlice 命令的字符串数组参数
type CmdFlagStringSlice struct {
	CmdFlag
	Name      string
	Shorthand string
	Usage     string
	P         *[]string
	Value     []string
}

func (f CmdFlagStringSlice) Bind(c *cobra.Command) {
	c.Flags().StringSliceVarP(f.P, f.Name, f.Shorthand, f.Value, f.Usage)
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

// CmdFlagInt 命令的int类型参数
type CmdFlagInt struct {
	CmdFlag
	Name      string
	Shorthand string
	Usage     string
	P         *int
	Value     int
}

func (f CmdFlagInt) Bind(c *cobra.Command) {
	c.Flags().IntVarP(f.P, f.Name, f.Shorthand, f.Value, f.Usage)
}

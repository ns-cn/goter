package goter

import "github.com/spf13/cobra"

type CmdFlag interface {
	// Bind 命令参数的绑定方式声明
	Bind(c *cobra.Command)
}

// NewCmdFlagString 创建新的命令行字符串参数
func NewCmdFlagString(p *string, defaultValue string, name string, shorthand string, usage string) CmdFlagString {
	return CmdFlagString{P: p, Value: defaultValue, CmdFlagSign: CmdFlagSign{Name: name, Shorthand: shorthand, Usage: usage}}
}

// NewCmdFlagStringSlice 创建新的命令行字符串数组参数
func NewCmdFlagStringSlice(p *[]string, defaultValue []string, name string, shorthand string, usage string) CmdFlagStringSlice {
	return CmdFlagStringSlice{P: p, Value: defaultValue, CmdFlagSign: CmdFlagSign{Name: name, Shorthand: shorthand, Usage: usage}}
}

// NewCmdFlagBool 创建新的命令行bool参数
func NewCmdFlagBool(p *bool, defaultValue bool, name string, shorthand string, usage string) CmdFlagBool {
	return CmdFlagBool{P: p, Value: defaultValue, CmdFlagSign: CmdFlagSign{Name: name, Shorthand: shorthand, Usage: usage}}
}

// NewCmdFlagInt 创建新的命令行int参数
func NewCmdFlagInt(p *int, defaultValue int, name string, shorthand string, usage string) CmdFlagInt {
	return CmdFlagInt{P: p, Value: defaultValue, CmdFlagSign: CmdFlagSign{Name: name, Shorthand: shorthand, Usage: usage}}
}

type CmdFlagSign struct {
	Name      string
	Shorthand string
	Usage     string
}

// CmdFlagString 命令的字符串参数
type CmdFlagString struct {
	CmdFlag
	CmdFlagSign
	P     *string
	Value string
}

func (f CmdFlagString) Bind(c *cobra.Command) {
	c.Flags().StringVarP(f.P, f.Name, f.Shorthand, f.Value, f.Usage)
}

// CmdFlagStringSlice 命令的字符串数组参数
type CmdFlagStringSlice struct {
	CmdFlag
	CmdFlagSign
	P     *[]string
	Value []string
}

func (f CmdFlagStringSlice) Bind(c *cobra.Command) {
	c.Flags().StringSliceVarP(f.P, f.Name, f.Shorthand, f.Value, f.Usage)
}

// CmdFlagBool 命令的bool类型参数
type CmdFlagBool struct {
	CmdFlag
	CmdFlagSign
	P     *bool
	Value bool
}

func (f CmdFlagBool) Bind(c *cobra.Command) {
	c.Flags().BoolVarP(f.P, f.Name, f.Shorthand, f.Value, f.Usage)
}

// CmdFlagInt 命令的int类型参数
type CmdFlagInt struct {
	CmdFlag
	CmdFlagSign
	P     *int
	Value int
}

func (f CmdFlagInt) Bind(c *cobra.Command) {
	c.Flags().IntVarP(f.P, f.Name, f.Shorthand, f.Value, f.Usage)
}

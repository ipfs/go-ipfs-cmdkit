package cmdkit

import (
	"reflect"

	"github.com/ipfs/go-ipfs-cmds"
)

type (
	Argument     = cmds.Argument
	ArgumentType = cmds.ArgumentType
	Error        = cmds.Error
	ErrorType    = cmds.ErrorType
	HelpText     = cmds.HelpText
	OptMap       = cmds.OptMap
	Option       = cmds.Option
)

const (
	Invalid   = cmds.Invalid
	Bool      = cmds.Bool
	Int       = cmds.Int
	Uint      = cmds.Uint
	Int64     = cmds.Int64
	Uint64    = cmds.Uint64
	Float     = cmds.Float
	String    = cmds.String
	ArgString = cmds.ArgString
	ArgFile   = cmds.ArgFile
)

func BoolOption(names ...string) Option                   { return cmds.BoolOption(names...) }
func FloatOption(names ...string) Option                  { return cmds.FloatOption(names...) }
func Int64Option(names ...string) Option                  { return cmds.Int64Option(names...) }
func IntOption(names ...string) Option                    { return cmds.IntOption(names...) }
func NewOption(kind reflect.Kind, names ...string) Option { return cmds.NewOption(kind, names...) }
func StringOption(names ...string) Option                 { return cmds.StringOption(names...) }
func Uint64Option(names ...string) Option                 { return cmds.Uint64Option(names...) }
func UintOption(names ...string) Option                   { return cmds.UintOption(names...) }

func FileArg(name string, required, variadic bool, description string) Argument {
	return cmds.FileArg(name, required, variadic, description)
}

func StringArg(name string, required, variadic bool, description string) Argument {
	return cmds.StringArg(name, required, variadic, description)
}

func Errorf(code ErrorType, format string, args ...interface{}) Error {
	return cmds.Errorf(code, format, args...)
}

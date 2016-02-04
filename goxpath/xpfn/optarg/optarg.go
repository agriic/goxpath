package optarg

import (
	"fmt"

	"github.com/ChrisTrenkamp/goxpath/tree"
)

//Fn defines an XPath function that has one optional argument
type Fn func(arg []tree.Res) ([]tree.Res, error)

//Wrap is wraps the OptArgFn method with XPFn
type Wrap struct {
	Fn Fn
}

//Call satisfies the XPFn interface for optarg.Wrap
func (fn Wrap) Call(ctx []tree.Res, args ...[]tree.Res) ([]tree.Res, error) {
	if len(args) > 1 {
		return nil, fmt.Errorf("Too many arguments to function call.")
	}

	if len(args) == 1 {
		return fn.Fn(args[0])
	}

	return fn.Fn(ctx)
}
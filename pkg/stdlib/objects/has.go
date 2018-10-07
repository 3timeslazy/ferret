package objects

import (
	"context"

	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
)

/*
 * Returns the value stored by the given key.
 * @params (String) - The key name string.
 * @returns (Boolean) - True if the key exists else false.
 */
func Has(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 2, 2)

	if err != nil {
		return values.None, err
	}

	err = core.ValidateType(args[0], core.ObjectType)

	if err != nil {
		return values.None, err
	}

	err = core.ValidateType(args[1], core.StringType)

	if err != nil {
		return values.None, err
	}

	obj := args[0].(*values.Object)
	keyName := args[1].(values.String)

	_, has := obj.Get(keyName)

	return values.NewBoolean(has), nil
}

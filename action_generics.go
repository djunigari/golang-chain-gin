package chains

import (
	"reflect"

	chain "github.com/djunigari/golang-chain"
)

func ConvertJsonTo[T any](attName string) *chain.Action[Context] {
	var f chain.ActionFunc[Context] = func(ctx *chain.Context[Context]) {
		var value T
		if err := ctx.Extra.C.ShouldBindJSON(&value); err != nil {
			ctx.SetErr(err)
			return
		}

		ctx.Additional[attName] = &value
	}
	return chain.NewAction[Context]("chains.ConvertJsonTo").Function(f)
}

func Get[FROM any](from, attribute, attName string) *chain.Action[Context] {
	var f chain.ActionFunc[Context] = func(ctx *chain.Context[Context]) {
		obj, ok := ctx.Additional[from].(FROM)
		if !ok {
			ctx.SetErr(chain.ErrInvalidVariableType)
			return
		}

		value := reflect.ValueOf(obj)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		if value.Kind() != reflect.Struct {
			return
		}

		fieldValue := value.FieldByName(attribute)
		if !fieldValue.IsValid() {
			return
		}
		ctx.Additional[attName] = fieldValue.Interface()
	}
	return chain.NewAction[Context]("chains.Get").Function(f)
}

func RenameAdditionalKey(from, to string) *chain.Action[Context] {
	var f chain.ActionFunc[Context] = func(ctx *chain.Context[Context]) {
		ctx.RenameAdditionalKey(from, to)
	}
	return chain.NewAction[Context]("chains.RenameAdditionalKey").Function(f)
}

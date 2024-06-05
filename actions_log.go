package chains

import (
	"fmt"
	"reflect"

	chain "github.com/djunigari/golang-chain"
	logger "github.com/djunigari/golang-logger"
)

func logMethodName(errorMessage string, ctx *chain.Context[Context]) {
	if ctx.Err() != nil {
		var errorDetails string
		if ctx.ActionErr != nil {
			errorDetails = fmt.Sprintf("[%s] %s : %s", ctx.ActionErr.Name, ctx.Err(), ctx.ErrMsg())
		} else {
			errorDetails = fmt.Sprintf("%s : %s", ctx.Err(), ctx.ErrMsg())
		}
		logger.LogError(errorMessage, errorDetails)
	}
}

func LogGet(entityType string) *chain.Action[Context] {
	return chain.NewAction[Context]("chains.LogGet").
		IgnoreError(true).
		Function(func(ctx *chain.Context[Context]) {
			if ctx.Err() != nil {
				logMethodName("failed to get "+entityType, ctx)
			}
		})
}

func LogGetList(entityType string) *chain.Action[Context] {
	return chain.NewAction[Context]("chains.LogGetList").
		IgnoreError(true).
		Function(func(ctx *chain.Context[Context]) {
			if ctx.Err() != nil {
				logMethodName("failed to get list of "+entityType, ctx)
			}
		})
}

func LogCreate[T any](attName string) *chain.Action[Context] {
	structName := typeOf[T]()
	entityName := structName.Name()
	return chain.NewAction[Context]("chains.LogCreate").
		IgnoreError(true).
		Function(func(ctx *chain.Context[Context]) {
			if ctx.Err() == nil {
				if obj, ok := ctx.Additional[attName].(*T); ok {
					logger.LogCreatedSuccess(entityName, obj)
					return
				}
				ctx.SetErr(chain.ErrInvalidVariableType)
			}
			logMethodName("failed to create "+entityName, ctx)
		})
}

func LogUpdate[T any](attName string) *chain.Action[Context] {
	structName := typeOf[T]()
	entityName := structName.Name()
	return chain.NewAction[Context]("chains.LogUpdate").
		IgnoreError(true).
		Function(func(ctx *chain.Context[Context]) {
			if ctx.Err() == nil {
				if obj, ok := ctx.Additional[attName].(*T); ok {
					logger.LogUpdatedSuccess(entityName, obj)
					return
				}
				ctx.SetErr(chain.ErrInvalidVariableType)
			}

			logMethodName("failed to update "+entityName, ctx)

		})
}

func LogDelete[T, idType any](attName string) *chain.Action[Context] {
	structName := typeOf[T]()
	entityName := structName.Name()
	return chain.NewAction[Context]("chains.LogDelete").
		IgnoreError(true).
		Function(func(ctx *chain.Context[Context]) {
			if ctx.Err() == nil {
				if id, ok := ctx.Additional[attName].(idType); ok {
					logger.LogDeletedSuccess(entityName, fmt.Sprintf("id=%v", id))
					return
				}
				ctx.SetErr(chain.ErrInvalidVariableType)
			}

			logMethodName("failed to delete "+entityName, ctx)
		})
}

func typeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

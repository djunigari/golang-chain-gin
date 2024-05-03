package chains

import (
	"fmt"
	"reflect"

	chain "github.com/djunigari/golang-chain"
	logger "github.com/djunigari/golang-logger"
)

func logMethodName(errorMessage, errorDetails string, ctx *chain.Context[Context]) {
	if ctx.Err() != nil {
		if ctx.ActionErr != nil {
			logger.LogError(errorMessage, ctx.ActionErr.Name+":"+errorDetails)
		} else {
			logger.LogError(errorMessage, errorDetails)
		}
	}
}

func LogGet(entityType string) *chain.Action[Context] {
	var f chain.ActionFunc[Context] = func(ctx *chain.Context[Context]) {
		if ctx.Err() != nil {
			logMethodName("failed to get "+entityType, ctx.Err().Error(), ctx)
		}
	}
	return chain.NewAction[Context]("chains.LogGet").Function(f).IgnoreError(true)
}

func LogGetList(entityType string) *chain.Action[Context] {
	var f chain.ActionFunc[Context] = func(ctx *chain.Context[Context]) {
		if ctx.Err() != nil {
			logMethodName("failed to get list of "+entityType, ctx.Err().Error(), ctx)
		}
	}
	return chain.NewAction[Context]("chains.LogGetList").Function(f).IgnoreError(true)
}

func LogCreate[T any](attName string) *chain.Action[Context] {
	structName := typeOf[T]()
	entityName := structName.Name()

	var f chain.ActionFunc[Context] = func(ctx *chain.Context[Context]) {
		if ctx.Err() != nil {
			logMethodName("failed to create "+entityName, ctx.Err().Error(), ctx)
			return
		}
		obj, ok := ctx.Additional[attName].(*T)
		if !ok {
			ctx.SetErr(chain.ErrInvalidVariableType)
			return
		}

		logger.LogCreatedSuccess(entityName, obj)
	}
	return chain.NewAction[Context]("chains.LogCreate").Function(f).IgnoreError(true)
}

func LogUpdate[T any](attName string) *chain.Action[Context] {
	structName := typeOf[T]()
	entityName := structName.Name()

	var f chain.ActionFunc[Context] = func(ctx *chain.Context[Context]) {
		if ctx.Err() != nil {
			logMethodName("failed to update "+entityName, ctx.Err().Error(), ctx)
			return
		}

		obj, ok := ctx.Additional[attName].(*T)
		if !ok {
			ctx.SetErr(chain.ErrInvalidVariableType)
			return
		}
		logger.LogUpdatedSuccess(entityName, obj)
	}
	return chain.NewAction[Context]("chains.LogUpdate").Function(f).IgnoreError(true)
}

func LogDelete[T, idType any](attName string) *chain.Action[Context] {
	structName := typeOf[T]()
	entityName := structName.Name()

	var f chain.ActionFunc[Context] = func(ctx *chain.Context[Context]) {
		if ctx.Err() != nil {
			logMethodName("failed to delete "+entityName, ctx.Err().Error(), ctx)
			return
		}

		id, ok := ctx.Additional[attName].(idType)
		if !ok {
			ctx.SetErr(chain.ErrInvalidVariableType)
			return
		}
		logger.LogDeletedSuccess(entityName, fmt.Sprintf("id=%v", id))
	}
	return chain.NewAction[Context]("chains.LogDelete").Function(f).IgnoreError(true)
}

func typeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

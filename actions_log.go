package chains

import (
	"fmt"
	"reflect"

	chain "github.com/djunigari/golang-chain"
	logger "github.com/djunigari/golang-logger"
)

// func LogGet(entityType string, handlerName string) *chain.Action[Context] {
// 	return chain.NewAction[Context]("chains.LogGet").
// 		Function(func(ctx *chain.Context[Context]) {
// 			if ctx.Err() != nil {
// 				logMethodName("failed "+handlerName, ctx)
// 			}
// 		})
// }

// func LogGetList(entityType string, handlerName string) *chain.Action[Context] {
// 	return chain.NewAction[Context]("chains.LogGetList").
// 		Function(func(ctx *chain.Context[Context]) {
// 			if ctx.Err() != nil {
// 				logMethodName("failed "+handlerName, ctx)
// 			}
// 		})
// }

func LogCreate[T any](attName string) *chain.Action[Context] {
	structName := typeOf[T]()
	entityName := structName.Name()
	return chain.NewAction[Context]("chains.LogCreate").
		Function(func(ctx *chain.Context[Context]) {
			if obj, ok := ctx.Additional[attName].(*T); ok {
				logger.LogCreatedSuccess(entityName, obj)
				return
			}
			ctx.SetErr(chain.ErrInvalidVariableType)
		})
}

func LogUpdate[T any](attName string) *chain.Action[Context] {
	structName := typeOf[T]()
	entityName := structName.Name()
	return chain.NewAction[Context]("chains.LogUpdate").
		Function(func(ctx *chain.Context[Context]) {
			if obj, ok := ctx.Additional[attName].(*T); ok {
				logger.LogUpdatedSuccess(entityName, obj)
				return
			}
			ctx.SetErr(chain.ErrInvalidVariableType)
		})
}

func LogDelete[T, idType any](attName string) *chain.Action[Context] {
	structName := typeOf[T]()
	entityName := structName.Name()
	return chain.NewAction[Context]("chains.LogDelete").
		Function(func(ctx *chain.Context[Context]) {
			if id, ok := ctx.Additional[attName].(idType); ok {
				logger.LogDeletedSuccess(entityName, fmt.Sprintf("id=%v", id))
				return
			}
			ctx.SetErr(chain.ErrInvalidVariableType)
		})
}

func typeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

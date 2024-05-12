package chains

import (
	"reflect"
	"strconv"

	chain "github.com/djunigari/golang-chain"
	"github.com/google/uuid"
)

type ParamType interface {
	uint64 | string | uuid.UUID
}

type QueryParamType interface {
	int | uint64 | string | uuid.UUID
}

func GetParam[T ParamType](attName string) *chain.Action[Context] {
	var f chain.ActionFunc[Context] = func(ctx *chain.Context[Context]) {
		param := ctx.Extra.C.Param(attName)

		var x T
		switch t := reflect.TypeOf(x); t.Kind() {
		case reflect.Uint64:
			paramUint, err := strconv.ParseUint(param, 10, 64)
			if err != nil {
				ctx.SetErr(err)
				return
			}

			ctx.Additional[attName] = paramUint
		case reflect.String:
			if reflect.TypeOf(x) == reflect.TypeOf(uuid.UUID{}) {
				uid, err := uuid.Parse(param)
				if err != nil {
					ctx.SetErr(err)
					return
				}
				ctx.Additional[attName] = uid
			} else {
				ctx.Additional[attName] = param
			}
		default:
			ctx.Additional[attName] = param
		}
	}
	return chain.NewAction[Context]("chains.GetParam").Function(f)
}

func GetQueryParam[T QueryParamType](attName string, defaultValue T) *chain.Action[Context] {
	var f chain.ActionFunc[Context] = func(ctx *chain.Context[Context]) {
		param := ctx.Extra.C.Query(attName)
		if param == "" {
			ctx.Additional[attName] = defaultValue
			return
		}

		var x T
		switch t := reflect.TypeOf(x); t.Kind() {
		case reflect.Uint64:
			paramUint, err := strconv.ParseUint(param, 10, 64)
			if err != nil {
				ctx.SetErr(err)
				return
			}
			ctx.Additional[attName] = paramUint
		case reflect.Int:
			paramInt, err := strconv.Atoi(param)
			if err != nil {
				ctx.SetErr(err)
				return
			}
			ctx.Additional[attName] = paramInt
		case reflect.String:
			if reflect.TypeOf(defaultValue) == reflect.TypeOf(uuid.UUID{}) {
				uid, err := uuid.Parse(param)
				if err != nil {
					ctx.SetErr(err)
					return
				}
				ctx.Additional[attName] = uid
			} else {
				ctx.Additional[attName] = param
			}
		default:
			ctx.Additional[attName] = param
		}
	}
	return chain.NewAction[Context]("chains.GetQueryParam").Function(f)
}

func GetQueryParamFilters(queryParams ...string) *chain.Action[Context] {
	var f chain.ActionFunc[Context] = func(ctx *chain.Context[Context]) {

		filters := make(map[string]string)
		for _, filter := range queryParams {
			if value := ctx.Extra.C.Query(filter); value != "" {
				filters[filter] = value
			}
		}

		ctx.Additional["filters"] = filters
	}
	return chain.NewAction[Context]("chains.GetQueryParam").Function(f)
}

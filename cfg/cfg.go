package cfg

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/wayne011872/api-toolkit/mid"
	"github.com/wayne011872/microservice/di"
	"github.com/wayne011872/microservice/grpc_tool/interceptor"
)

type ModelCfg interface {
	Close() error
	Init(uuid string, di di.DI) error
	Copy() ModelCfg
}

type ModelCfgMgr interface {
	mid.GinMiddle
	interceptor.Interceptor
}

type ctxType string

const cfgKey = "model_config"

func GetFromGinCtx[T ModelCfg](ctx *gin.Context) (T, bool) {
	var result T
	val, ok := ctx.Get(cfgKey)
	if !ok {
		return result, false
	}
	return val.(T), true
}

func setToGinCtx[T ModelCfg](ctx *gin.Context, cfg T) {
	ctx.Set(cfgKey, cfg)
}

func GetFromCtx[T ModelCfg](ctx context.Context) (T, bool) {
	var result T
	val := ctx.Value(ctxType(cfgKey))
	if val == nil {
		return result, false
	}
	return val.(T), true
}
func setToCtx[T ModelCfg](ctx context.Context, cfg T) context.Context {
	return context.WithValue(ctx, ctxType(cfgKey), cfg)
}

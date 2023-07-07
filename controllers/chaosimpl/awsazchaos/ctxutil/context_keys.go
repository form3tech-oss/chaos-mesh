package ctxutil

import "context"

type ctxKey string

const (
	CtxKeySimulationId ctxKey = "simulationId"
	CtxKeyCurbFlag     ctxKey = "curbFlag"
)

func GetOptionalBool(ctx context.Context, key ctxKey) bool {
	val := ctx.Value(key)
	if b, _ := val.(bool); b {
		return true
	}
	return false
}

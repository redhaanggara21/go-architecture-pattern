package util

import (
	"context"
	"time"

	"red21.id/learn/bengkel/domain"
)

func ResponseInterceptor(ctx context.Context, response *domain.ApiResponse) {
	traceIdInf := ctx.Value("requestid")
	traceId := ""
	if traceIdInf != nil {
		traceId = traceIdInf.(string)
	}
	response.Timestamp = time.Now()
	response.TraceID = traceId
}

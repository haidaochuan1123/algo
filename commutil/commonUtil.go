package commutil

import "time"

// CalCostTime 返回耗时, 单位为毫秒
func CalCostTime(startTime time.Time) int64 {
	costTime := time.Since(startTime)
	return costTime.Milliseconds()
}

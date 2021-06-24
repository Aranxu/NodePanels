package util

import (
	"time"
)

func Now() string {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	return time.Now().In(cstZone).Format("2006-01-02 15:04:05")
}

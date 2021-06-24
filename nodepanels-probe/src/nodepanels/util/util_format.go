package util

import (
	"fmt"
	"math"
	"strconv"
)

func Int2string(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int642string(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int322string(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func String2int(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func String2int32(str string) int32 {
	i, _ := strconv.Atoi(str)
	return int32(i)
}

func Float642string(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Float322string(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}

func Uint642string(u uint64) string {
	return strconv.FormatUint(u, 10)
}

func Round(f float64) int {
	return int(math.Floor(f + 0/5))
}

func RoundFloat64(f float64, i int) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%."+Int2string(i)+"f", f), 64)
	return value
}

func RoundFloat32(f float32, i int) float32 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%."+Int2string(i)+"f", f), 64)
	return float32(value)
}

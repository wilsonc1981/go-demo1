package rand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // 初始化亂數種子
}

// GetRandomIntBetween [min, max)
func GetRandomIntBetween(min, max int) int {
	return rand.Intn(max-min) + min
}

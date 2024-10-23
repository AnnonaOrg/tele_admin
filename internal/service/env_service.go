package service

import (
	"os"
	"strings"
)

// 穿透屏蔽
func IsPenetrationShielding() bool {
	return strings.EqualFold("true", os.Getenv("PENETRATION_SHIELDING_ENABLE"))
}

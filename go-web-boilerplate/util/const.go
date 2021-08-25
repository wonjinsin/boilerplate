package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// CustomStr ...
type CustomStr string

// CustomStrs ...
const (
	TRID CustomStr = "trid"
)

// GetTRID ...
func GetTRID() string {
	t := time.Now()
	randInt := strconv.Itoa(rand.Intn(8999) + 1000)
	trid := strings.Replace(t.Format("20060102150405.00"), ".", "", -1) + randInt

	return trid
}

package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type customStr string

// CustomStrs ...
const (
	TRID customStr = "trid"
)

// GetTRID ...
func GetTRID() string {
	t := time.Now()
	randInt := strconv.Itoa(rand.Intn(8999) + 1000)
	trid := strings.Replace(t.Format("20060102150405.00"), ".", "", -1) + randInt

	return trid
}

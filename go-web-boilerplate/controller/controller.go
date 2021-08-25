package controller

import (
	"log"
	"os"
	"pikachu/util"
	"strconv"

	"github.com/labstack/echo"
)

var zlog *util.Logger

type pickachuStatus struct {
	TRID       string      `json:"trid"`
	ResultCode string      `json:"resultCode"`
	ResultMsg  string      `json:"resultMsg"`
	ResultData interface{} `json:"resultData,omitempty"`
}

func init() {
	var err error
	zlog, err = util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[controller] err[%s]", err.Error())
		os.Exit(1)
	}
}

func response(c echo.Context, code int, resultMsg string, resultData ...interface{}) error {
	strCode := strconv.Itoa(code)

	trid, ok := c.Request().Context().Value(util.TRID).(string)
	if !ok {
		trid = util.GetTRID()
	}

	res := pickachuStatus{
		TRID:       trid,
		ResultCode: strCode,
		ResultMsg:  resultMsg,
		ResultData: resultData,
	}

	return c.JSON(code, res)
}

package room_handler

import (
	util "action/utility"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRoomId(c echo.Context) error {
	n, err := util.RandomString(10)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, n)
}

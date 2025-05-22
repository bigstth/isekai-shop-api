package validation

import (
	_adminException "github.com/bigstth/isekai-shop-api/pkg/admin/exception"
	_playerException "github.com/bigstth/isekai-shop-api/pkg/player/exception"
	"github.com/labstack/echo/v4"
)

func AdminIDGetting(ctx echo.Context) (string, error) {
	if adminID, ok := ctx.Get("adminID").(string); !ok || adminID == "" {
		return "", &_adminException.AdminNotFound{AdminID: "Unknown"}
	} else {
		return adminID, nil
	}
}
func PlayerIDGetting(ctx echo.Context) (string, error) {
	if playerID, ok := ctx.Get("playerID").(string); !ok || playerID == "" {
		return "", &_playerException.PlayerNotFound{PlayerID: "Unknown"}
	} else {
		return playerID, nil
	}
}

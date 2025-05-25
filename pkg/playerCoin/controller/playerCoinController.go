package controller

import "github.com/labstack/echo/v4"

type PlayerCoinController interface {
	CoinAdding(ctx echo.Context) error
	Showing(ctx echo.Context) error
}

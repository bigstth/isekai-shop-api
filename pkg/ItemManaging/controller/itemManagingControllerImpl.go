package controller

import (
	"net/http"
	"strconv"

	"github.com/bigstth/isekai-shop-api/pkg/custom"
	_itemManagingModel "github.com/bigstth/isekai-shop-api/pkg/itemManaging/model"
	_itemManagingService "github.com/bigstth/isekai-shop-api/pkg/itemManaging/service"
	"github.com/bigstth/isekai-shop-api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type ItemManagingControllerImpl struct {
	itemManagingService _itemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(itemManagingService _itemManagingService.ItemManagingService) ItemManagingController {
	return &ItemManagingControllerImpl{itemManagingService}
}

func (c *ItemManagingControllerImpl) Creating(ctx echo.Context) error {
	adminID, err := validation.AdminIDGetting(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	itemCreatingRequest := new(_itemManagingModel.ItemCreatingReq)

	customEchoRequest := custom.NewCustomEchoRequest(ctx)

	if err := customEchoRequest.Bind(itemCreatingRequest); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	itemCreatingRequest.AdminID = adminID

	item, err := c.itemManagingService.Creating(itemCreatingRequest)

	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, item)
}

func (c *ItemManagingControllerImpl) Editing(ctx echo.Context) error {
	itemId, err := c.getItemID(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, "Invalid item ID")
	}

	itemEditingRequest := new(_itemManagingModel.ItemEditingReq)

	customEchoRequest := custom.NewCustomEchoRequest(ctx)

	if err := customEchoRequest.Bind(itemEditingRequest); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	item, err := c.itemManagingService.Editing(itemId, itemEditingRequest)

	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, item)
}

func (c *ItemManagingControllerImpl) getItemID(ctx echo.Context) (uint64, error) {
	itemID := ctx.Param("itemID")
	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return 0, err
	}
	return itemIDUint64, nil
}

func (c *ItemManagingControllerImpl) Archiving(ctx echo.Context) error {
	itemId, err := c.getItemID(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, "Invalid item ID")
	}

	err = c.itemManagingService.Archiving(itemId)

	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

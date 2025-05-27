package controller

import (
	"net/http"

	"github.com/ThakdanaiDL.git/shop-api/pkg/custom"
	_itemShopService "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/service"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService: itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {

	itemModelList, err := c.itemShopService.Listing()
	if err != nil {
		// return pctx.String(http.StatusInternalServerError, "Error fetching item list: "+err.Error())
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, itemModelList) // Return the list of items as JSON response

}

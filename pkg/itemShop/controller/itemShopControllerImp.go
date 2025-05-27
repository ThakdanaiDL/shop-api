package controller

import (
	"net/http"

	"github.com/ThakdanaiDL.git/shop-api/pkg/custom"
	_itemShopmodel "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/model"
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

	// itemfilter := new(_itemShopmodel.ItemShopFilter)
	// if err := pctx.Bind(itemfilter); err != nil {
	// 	return custom.Error(pctx, http.StatusBadRequest, err.Error())

	// }
	// validator := validator.New()
	// if err := validator.Struct(pctx); err != nil {
	// 	return custom.Error(pctx, http.StatusBadRequest, err.Error())
	// }

	//จากเดิมข้างบน เป็น เเบบ ข้างล่างสะอาดขึ้น

	itemfilter := new(_itemShopmodel.ItemShopFilter)
	customEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customEchoRequest.Bind(itemfilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	itemModelList, err := c.itemShopService.Listing(itemfilter)
	if err != nil {
		// return pctx.String(http.StatusInternalServerError, "Error fetching item list: "+err.Error())
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, itemModelList) // Return the list of items as JSON response

}

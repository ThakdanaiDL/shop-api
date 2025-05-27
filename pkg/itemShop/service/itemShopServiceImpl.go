package services

import (
	_itemshopModel "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/model"
	"github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/repository"
)

type itemShopServiceImp struct {
	itemShopRepository repository.ItemShopRepository
}

func NewItemShopRepositoryImpl(itemShopRepository repository.ItemShopRepository) ItemShopService {

	return &itemShopServiceImp{itemShopRepository: itemShopRepository}

}

func (s *itemShopServiceImp) Listing(itemFilter *_itemshopModel.ItemShopFilter) ([]*_itemshopModel.Item, error) {

	itemList, err := s.itemShopRepository.Listing(itemFilter) //ดึงรายการมาจาก repository
	if err != nil {
		return nil, err
	}

	itemModelList := make([]*_itemshopModel.Item, 0) // สร้าง sliceเปล่า เพื่อเก็บ item model

	for _, item := range itemList {

		itemModelList = append(itemModelList, item.ToItemShopModel()) // เเปลง item entity เป็น item model เเละเพิ่มเข้าไปใน slice

	}

	return itemModelList, nil

}

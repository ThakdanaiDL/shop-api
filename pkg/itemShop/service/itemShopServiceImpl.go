package services

import (
	"github.com/ThakdanaiDL.git/shop-api/entities"
	_itemshopModel "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/model"
	"github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/repository"
)

type itemShopServiceImp struct {
	itemShopRepository repository.ItemShopRepository
}

func NewItemShopRepositoryImpl(itemShopRepository repository.ItemShopRepository) ItemShopService {

	return &itemShopServiceImp{itemShopRepository: itemShopRepository}

}

func (s *itemShopServiceImp) Listing(itemFilter *_itemshopModel.ItemShopFilter) (*_itemshopModel.ItemResult, error) {

	itemList, err := s.itemShopRepository.Listing(itemFilter) //ดึงรายการมาจาก repository
	if err != nil {
		return nil, err
	}

	itemCounting, err := s.itemShopRepository.Counting(itemFilter) //ดึงจำนวนรายการทั้งหมดจาก repository
	if err != nil {
		return nil, err
	}

	size := itemFilter.Size
	page := itemFilter.Page
	totalPage := s.totalPageCalculation(itemCounting, size)

	result := s.toItemResultResponse(itemList, page, totalPage)

	return result, nil
}

func (r *itemShopServiceImp) totalPageCalculation(totalItem int64, size int64) int64 {
	totalPage := totalItem / size

	if totalItem%size != 0 {
		totalPage++
	}

	return totalPage

}

func (r *itemShopServiceImp) toItemResultResponse(itemEntitityList []*entities.Item, page int64, totalPage int64) *_itemshopModel.ItemResult {

	itemModelList := make([]*_itemshopModel.Item, 0) // สร้าง sliceเปล่า เพื่อเก็บ item model
	for _, item := range itemEntitityList {

		itemModelList = append(itemModelList, item.ToItemShopModel()) // เเปลง item entity เป็น item model เเละเพิ่มเข้าไปใน slice

	}

	return &_itemshopModel.ItemResult{
		Item: itemModelList,
		Paginate: _itemshopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}

// func (s *itemShopServiceImp) Count(itemFilter *_itemshopModel.ItemShopFilter) (int64, error) {
// 	count, err := s.Count(itemFilter)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return count, nil

// }

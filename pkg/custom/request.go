package custom

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// itemfilter := new(_itemShopmodel.ItemShopFilter)  				** ดึงข้อมูลที่เราจองไว้ใน struct itemfilter
// if err := pctx.Bind(itemfilter); err != nil {    				** เอาข้อมูลที่อยู่ใน echo.Context มา bind กับ itemfilter
// 	return custom.Error(pctx, http.StatusBadRequest, err.Error())

// }
// validator := validator.New()										** สร้าง validator ขึ้นมาใหม่
// if err := validator.Struct(pctx); err != nil {					** เชคข้อมูลที่ bind ว่าถูกต้องไหม
// 	return custom.Error(pctx, http.StatusBadRequest, err.Error())
// }

//ปกติ ต้องทำเเบบข้างบนทุกรอบ เเต่มาเขียนฟังก์ชันเเยก จะได้ทำให้โค้ดสะอาดขึ้น

// 1. ไอเดียคือ เวลาที่เรา req มา ข้อมูลจะอยู่ใน echo.Context
// 2. จากนั้นจะต้องเรียกใช้ .Bind() เพื่อ mapping ข้อมูลที่อยู่ใน echo.Context กับ struct ที่เราจองที่ไว้
// 3. ในที่นี้คือ itemfilter  จากนั้นจะสร้าง validator ขึ้นมา เเเละเอาข้อมูลที่ bind ไปเชค

type (
	EchoRequest interface {
		Bind(object any) error
	}

	customEchoRequest struct {
		ctx       echo.Context
		validator *validator.Validate
	}
)

var (
	once              sync.Once
	validatorInstance *validator.Validate
)

func NewCustomEchoRequest(echoRequest echo.Context) EchoRequest {

	once.Do(func() {
		validatorInstance = validator.New() // สร้าง instance ของ validator เพียงครั้งเดียว
	})

	return &customEchoRequest{ctx: echoRequest, validator: validatorInstance}
}

func (r *customEchoRequest) Bind(object any) error {
	if err := r.ctx.Bind(object); err != nil {
		return err
	}

	if err := r.validator.Struct(object); err != nil {
		return err
	}
	return nil

}

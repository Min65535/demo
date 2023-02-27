package handle

import "errors"

type HttpResp struct {
	Success bool        `json:"success"`
	ErrMsg  string      `json:"code"`
	Data    interface{} `json:"data"`
}

func NewHttpErrResp(errMsg string) *HttpResp {
	return &HttpResp{ErrMsg: errMsg}
}

func NewHttpSuccessResp(data interface{}) *HttpResp {
	return &HttpResp{
		Success: true,
		Data:    data,
	}
}


type AddOrderReq struct {
	OrderNo  string  `json:"order_no"`
	UserName string  `json:"user_name"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	FileUrl  string  `json:"file_url"`
}

type Order struct {
	// ID       uint
	OrderNo  string  `json:"order_no"`
	UserName string  `json:"user_name"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	FileUrl  string  `json:"file_url"`
}

func (r *AddOrderReq) IsValid() error {
	switch {
	case r.Amount == 0:
		return errors.New("Amount is not zero. ")
	case r.UserName == "":
		return errors.New("UserName is not found. ")
	default:
		return nil
	}
}

func (r *Order) IsValid() error {
	switch {
	case r.OrderNo == "":
		return errors.New("OrderNo is not found. ")
	case r.Amount == 0:
		return errors.New("Amount is not zero. ")
	default:
		return nil
	}
}

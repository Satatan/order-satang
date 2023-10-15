package dto

import "order_satang/util/enum"

type CoreResponse struct {
	Message enum.ApiMessageResponse `json:"message"`
	Result  interface{}             `json:"result"`
}

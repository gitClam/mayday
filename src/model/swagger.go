package model

import (
//"log"
//"github.com/kataras/iris/v12"
)

// Success response
// swagger:model Ok
type swaggScsResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - OK
		//required: true
		Code int `json:"code"`
	}
}

// Error Forbidden
// swagger:response forbidden
type swaggErrForbidden struct {
	// in:body
	Body struct {
		// HTTP status code 403 -  Forbidden
		//required: true
		Code int `json:"code"`
		// Detailed error message
		//required: true
		Message string `json:"message"`
	}
}

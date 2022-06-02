package types

type ResponseData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type AddOneRequest struct {
	A int `json:"a" form:"a" binding:"required"`
}

type AddUserRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
}

type DiffUserRequest struct {
	Id uint `json:"id"`
}

package dto

type UserReq struct {
	ID         uint   `json:"id" form:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Pwd        string `json:"pwd"`
	Age        uint   `json:"age"`
	RegisterIP string
	LoginIP    string
}
type UserResp struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Age   uint   `json:"age"`
	Pwd   string `json:"pwd"`
}

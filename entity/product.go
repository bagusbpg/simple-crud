package entity

type Product struct {
	Id     int    //`json:"id" form:"id"`
	UserID int    //`json:"userid" form:"userid"`
	Name   string `json:"name" form:"name"`
	Price  int    `json:"price" form:"price"`
}

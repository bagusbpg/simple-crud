package entity

type Book struct {
	Id        int    //`json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Language  string `json:"language" form:"language"`
	Pages     int    `json:"pages" form:"pages"`
	ISBN13    string `json:"isbn13" form:"isbn13"`
}

package mocks

import "go_modules/models"

var BookStore = []models.Book{
	{
		Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go",
	},
	{
		Id:     2,
		Title:  "Data Structures and Algorithms",
		Author: "Navneet Shukla",
		Desc:   "A book for Data Structures and Algorithms",
	},
}

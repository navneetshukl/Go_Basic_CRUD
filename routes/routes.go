package routes

import (
	"encoding/json"
	"fmt"
	"go_modules/mocks"
	"go_modules/models"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {

}

func GetBook(w http.ResponseWriter, r *http.Request) {

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Fatal(err)
	}
	req:=string(body)
	fmt.Println(req)

	var book models.Book
	json.Unmarshal(body,&book)

	//Append to the BookStore

	book.Id=10
	n:=len(mocks.BookStore)
	book.Id=mocks.BookStore[n-1].Id+1

	mocks.BookStore=append(mocks.BookStore, book)
	// Send a 201 created response

	w.Header().Add("COntent-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Added Successfully")
	

	for i := 0; i <= n; i++ {

		fmt.Println(mocks.BookStore[i])
		
	}

	
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
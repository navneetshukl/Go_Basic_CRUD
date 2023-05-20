package routes

import (
	"encoding/json"
	"fmt"
	"go_modules/mocks"
	"go_modules/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("COntent-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Fetched All Books Successfully")
	n:=len(mocks.BookStore)

	for i := 0; i < n; i++ {

		fmt.Println("{")

		fmt.Println("  \"id\": ",mocks.BookStore[i].Id)
		fmt.Println(" \"title\": ",mocks.BookStore[i].Title)
		fmt.Println(" \"author\": ",mocks.BookStore[i].Author)
		fmt.Println(" \"description\": ",mocks.BookStore[i].Desc)

		fmt.Println("}")
		
	}


}

func GetBook(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParamFromCtx(r.Context(), "id")
	fmt.Println(userID)
	id, _ := strconv.Atoi(userID)

	for _,book:= range mocks.BookStore{
		if book.Id == id {
			w.Header().Add("Content-Type","application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
		}
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParamFromCtx(r.Context(), "id")
	fmt.Println(userID)
	id, _ := strconv.Atoi(userID)
	defer r.Body.Close()

	body,err:=ioutil.ReadAll(r.Body)

	if err!=nil{
		log.Fatalln(err)
	}

	var updatedBook models.Book

	json.Unmarshal(body,&updatedBook)

	// Iterate over all the bookStore

	for index,book:=range mocks.BookStore{
		if book.Id == id {
			book.Title=updatedBook.Title
			book.Author=updatedBook.Author
			book.Desc=updatedBook.Desc
			mocks.BookStore[index]=book

			w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)

            json.NewEncoder(w).Encode("Updated")
            break

		}
	}

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

	userID := chi.URLParamFromCtx(r.Context(), "id")
	fmt.Println(userID)
	id, _ := strconv.Atoi(userID)

	for index, book := range mocks.BookStore {
        if book.Id == id {
            // Delete book and send a response if the book Id matches dynamic Id
            mocks.BookStore = append(mocks.BookStore[:index], mocks.BookStore[index+1:]...)

            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode("Deleted")
            break
        }
    }
}
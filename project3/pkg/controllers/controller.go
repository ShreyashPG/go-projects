package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/ShreyashPG/project3/pkg/utils"
	"github.com/ShreyashPG/project3/pkg/model"
)

var NewBook model.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:= model.GetAllBooks()
	res,_:= json.Marshal(newBooks)
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook (w http.ResponseWriter, r *http.Request){
	CreateBook:= &model.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _:=json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _:= model.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r * http.Request){
	vars:= mux.Vars(r)

	bookId:= vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err!=nil {
		fmt.Println("error while parsing")
	}
	book :=model.DeleteBook(ID)
	res,_:=json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook =&model.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId:= vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err!=nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db :=model.GetBookById(Id)
	if updateBook.Name!= ""{
		bookDetails.Name = updateBook.Name
	}
	if(updateBook.Author !=""){
		bookDetails.Author = updateBook.Author
	}
	if(updateBook.Publication !=""){
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}





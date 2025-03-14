package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){

	if err:= r.ParseForm(); err!=nil {
		fmt.Fprintf(w, "ParseForm() err : %v",err)
	}
	fmt.Fprintf(w, "POST request Successful\n")

	name:=r.FormValue("name")
	email:= r.FormValue("email")
	age:= r.FormValue("age")

	fmt.Fprintf(w, "Name = %s\n",name)
	fmt.Fprintf(w, "Address = %s\n",email)
	fmt.Fprintf(w, "Age = %s\n",age)


}


func helloHandler(w http.ResponseWriter, r *http.Request){
	if(r.URL.Path !="/hello"){
		http.Error(w, "404 not found",http.StatusNotFound);
		return
	}

	if(r.Method!="GET"){
			http.Error(w, "Method not allowed",http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello\n")
}


func main (){
	fileServer:=http.FileServer(http.Dir("./static"))
		http.Handle("/",fileServer)

	http.HandleFunc("/form",formHandler);
	http.HandleFunc("/hello",helloHandler);


	fmt.Printf("Starting server at port 8080\n");

	if err:= http.ListenAndServe(":8080",nil) ;err!=nil{
		log.Fatal(err);
	}


}
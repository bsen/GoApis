package main

import (
	"fmt"
	userController "goapi/controller"
	"log"
	"net/http"
)

func main(){
	var port string = "8002"
	http.HandleFunc("/users",DummyUsers)
	fmt.Printf("http://localhost:%v\n",port)
	log.Fatal(http.ListenAndServe(":" + port,nil))	
}

func DummyUsers(w http.ResponseWriter, r * http.Request){
	var email string = r.URL.Query().Get("email")
    var pass string = r.URL.Query().Get("pass")
    userController.HandleSignUp(email,pass)
}



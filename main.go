package main

import (
	"encoding/json"
	"fmt"
	userController "free-apis/controller"
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
	response := userController.GetUsersJson()
	w.Header().Set("Content-Type","application/json")
	jsonData,err := json.Marshal(response)
    if err != nil{
		http.Error(w,err.Error(), http.StatusBadRequest)
	}
	w.Write(jsonData)
}



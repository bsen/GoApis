package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
	var port int = 8000
	http.HandleFunc("/users",DummyUsers)
	http.HandleFunc("/",OperateSum)
	fmt.Printf("http://localhost:%v\n",port)
	log.Fatal(http.ListenAndServe(":8000",nil))
	
}

var usersData = `
{
    "users": [
        {
            "ID": "001",
            "Name": "John Wick",
            "City": "New York",
            "Address": "1 Wall Street, Apt 5",
            "Contacts": {
                "Phone number": "111-222-3333",
                "E-mail": "johnwick@gmail.com",
                "Skype": "johnwick123"
            }
        },
        {
            "ID": "002",
            "Name": "Peter Parker",
            "City": "Queens, NY",
            "Address": "Chelsea Street, 410",
            "Contacts": {
                "Phone number": "123-456-7890",
                "E-mail": "spiderman@gmail.com",
                "Skype": "spiderman4ever"
            }
        },
        {
            "ID": "003",
            "Name": "Rustin Cohle",
            "City": "West Monroe, Louisiana",
            "Address": "14 Oklakhoma Street, Apt 13",
            "Contacts": {
                "Phone number (office)": "131-313-1313",
                "Phone number (home)": "666-666-6666",
                "E-mail": "truedetective@gmail.com",
                "Skype": "rustcohle888"
            }
        },
        {
            "ID": "004",
            "Name": "Tommy Shelby",
            "City": "Birmingham",
            "Address": "Mount Steet, 19/18",
            "Contacts": {
                "Phone number": "777-777-777",
                "E-mail": "peakyblinders@gmail.com",
                "Skype": "shelby1918"
            }
        }
    ]
}
`
func DummyUsers(w http.ResponseWriter, r * http.Request){
	fmt.Fprint(w, usersData)
}



func OperateSum (w http.ResponseWriter, r * http.Request){
	fmt.Println("Queries: ",r.URL.Query().Get("name"))
	fmt.Fprint(w,r.URL.Query())
}
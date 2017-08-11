package main

import(
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Many boy")
	})

	http.HandleFunc("/product",product)
	http.HandleFunc("/user", user)
  
	http.ListenAndServe(":8080",nil)
}

func product(w http.ResponseWriter, t *http.Request){
	fmt.Fprintf(w, "Product request")
}
/*function for request activity on http*/
func user(hr http.ResponseWriter, t *http.Request){
	fmt.Fprintf(hr, "user request")
}

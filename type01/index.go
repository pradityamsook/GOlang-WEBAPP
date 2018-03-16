package main

import(
	"html/template"
	"fmt"
	"net/http"
)

func editHandler(w http.ResponeWriter, r * http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err: =loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFIles("edit.html")
	t.Execute(w, p)
}

func main(){
	
	http.HandlerFunc("/view/", viewHandler)
	http.HandlerFunc("/edit/", editHandler)
	http.HandlerFunc("/save/", saveHandler)

	userDB := map[string] int{
		"pradit" : 21,
		"golang" : 25,
		"java"   : 10,
		"python" : 100,
	}

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Many boy")
	})

	http.HandleFunc("/user/",func(w http.ResponseWriter, r *http.Request){
		name := r.URL.Path[len("/user/"):]
		age := userDB[name]
		fmt.Fprintf(w, "I am %s %d years old",name, age)
	})

	
  
	http.ListenAndServe(":8080",nil)

}

func product(w http.ResponseWriter, t *http.Request){
	fmt.Fprintf(w, "Product request")
}

/*function for request activity on http*/
func user(hr http.ResponseWriter, t *http.Request){
	fmt.Fprintf(hr, "user request")
}

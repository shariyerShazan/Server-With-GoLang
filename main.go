package main

import(
	"fmt" 
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter , r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w , "404 not found" , http.StatusNotFound)
		return
	}
	if r.Method != "GET" {	
		http.Error(w , "Method not supported" , http.StatusNotFound)
		return
	}
	fmt.Fprintf(w , "Hello")
}

func formHanlder(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w , "parseForm() err: %v" , err)
		return
	}
	fmt.Fprintf(w , "Post request succesSfully")

	name := r.FormValue("name")
	adress := r.FormValue("adress")

	fmt.Fprintf(w , "Name : %s\n", name)
	fmt.Fprintf(w , "Adress : %s\n", adress)
}

func main(){
    fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/" , fileServer)
	http.HandleFunc("/hello" , helloHandler)
	http.HandleFunc("/form" , formHanlder)

	fmt.Printf("Starting server at port :3333\n")
	if err := http.ListenAndServe(":3333" , nil); err != nil {
		log.Fatal(err)
	}
}
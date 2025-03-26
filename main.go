package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func openBrowser(url string){
	switch runtime.GOOS{ //this is not goose this is GO OS lol
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "linux":
		exec.Command("xdg-open", url).Start()
	}
}

func formHandler(w http.ResponseWriter, r *http.Request){

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "ParseForm() POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name=%s\n", name)
	fmt.Fprintf(w, "address=%s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 hello path not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello world")
}



func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	url := "https://localhost:8080/"
	fmt.Printf("Port starting at: %s", url)

	openBrowser(url)

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}
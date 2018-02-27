package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/upload", upload)
	log.Println("Running on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}

func upload(w http.ResponseWriter, r *http.Request){
	
	if r.Method == http.MethodPost {
		file, handle, err := r.FormFile("myFile")
		if err != nil {
			log.Printf("Error loading the file %v", err)
			fmt.Fprint(w, "Error loading the file %v", err)
			return 
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("Error reading the file %v", err)
			fmt.Fprint(w, "Error reading the file %v", err)
			return 
		}

		err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
		if err != nil {
			log.Printf("Error writing the file %v", err)
			fmt.Fprint(w, "Error writing the file %v", err)
			return 
		}

		fmt.Fprint(w, "Succesful at loading the file")
	}//end if	
}//end upload

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	port := ":3001"
	fmt.Println("Starting server at port", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got a request to home page")
		w.Write([]byte("Hello Homepage"))
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			log.Println("Got a request to students using method " + r.Method)
			w.Write([]byte("Hello students from " + r.Method))
		case http.MethodPost:

			// processing x-www-form-urlencoded
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing data", http.StatusBadRequest)
			}
			fmt.Println(r.Form)
			processedData := make(map[string]interface{})

			for k, v := range r.Form {
				processedData[k] = v[0]
			}

			fmt.Println(processedData)

			// processing raw json
			byteBody, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error in processing raw json", http.StatusBadRequest)
			}

			fmt.Println(string(byteBody))

			var user User
			err = json.Unmarshal(byteBody, &user)
			if err != nil {
				http.Error(w, "Error converting json body to go struct", http.StatusBadRequest)
				return
			}
			fmt.Println(user)

			defer r.Body.Close()

			log.Println("Got a request to students using method " + r.Method)
			w.Write([]byte("Hello students from " + r.Method))
		case http.MethodPut:
			log.Println("Got a request to students using method " + r.Method)
			w.Write([]byte("Hello students from " + r.Method))
		case http.MethodPatch:
			log.Println("Got a request to students using method " + r.Method)
			w.Write([]byte("Hello students from " + r.Method))
		case http.MethodDelete:
			log.Println("Got a request to students using method " + r.Method)
			w.Write([]byte("Hello students from " + r.Method))
		default:
			log.Println("Got a request to students from unsupported method:", r.Method)
			w.Write([]byte("Hello students from " + r.Method))
		}

	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got a request to teachers")
		w.Write([]byte("Hello teachers"))
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got a request to execs")
		w.Write([]byte("Hello execs"))
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}

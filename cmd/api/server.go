package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Got a request to home page")
	w.Write([]byte("Hello Homepage"))
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {

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

}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		// Path params
		// teachers/{teacherId}
		// teachers/3

		fmt.Println("Path from url:", r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		teacherId := strings.TrimSuffix(path, "/")
		fmt.Println("Extracted teacherId as a path param:", teacherId)

		//Query params
		// teachers/?key=value&sortby=email&sortorder=asc

		queryParams := r.URL.Query()

		fmt.Println(queryParams.Get("key"))
		fmt.Println(queryParams.Get("sortby"))
		sortOrder := queryParams.Get("sortorder")

		// default value
		if sortOrder == "" {
			sortOrder = "desc"
		}
		fmt.Println()

		log.Println("Got a request to teachers using method " + r.Method)
		w.Write([]byte("Hello teachers from " + r.Method))
	case http.MethodPost:
		log.Println("Got a request to teachers using method " + r.Method)
		w.Write([]byte("Hello teachers from " + r.Method))
	case http.MethodPut:
		log.Println("Got a request to teachers using method " + r.Method)
		w.Write([]byte("Hello teachers from " + r.Method))
	case http.MethodPatch:
		log.Println("Got a request to teachers using method " + r.Method)
		w.Write([]byte("Hello teachers from " + r.Method))
	case http.MethodDelete:
		log.Println("Got a request to teachers using method " + r.Method)
		w.Write([]byte("Hello teachers from " + r.Method))
	default:
		log.Println("Got a request to teachers from unsupported method:", r.Method)
		w.Write([]byte("Hello teachers from " + r.Method))
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		log.Println("Got a request to execs using method " + r.Method)
		w.Write([]byte("Hello execs from " + r.Method))
	case http.MethodPost:
		log.Println("Got a request to execs using method " + r.Method)
		w.Write([]byte("Hello execs from " + r.Method))
	case http.MethodPut:
		log.Println("Got a request to execs using method " + r.Method)
		w.Write([]byte("Hello execs from " + r.Method))
	case http.MethodPatch:
		log.Println("Got a request to execs using method " + r.Method)
		w.Write([]byte("Hello execs from " + r.Method))
	case http.MethodDelete:
		log.Println("Got a request to execs using method " + r.Method)
		w.Write([]byte("Hello execs from " + r.Method))
	default:
		log.Println("Got a request to execs from unsupported method:", r.Method)
		w.Write([]byte("Hello execs from " + r.Method))
	}
}

func main() {
	port := ":3001"
	fmt.Println("Starting server at port", port)

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/students/", studentsHandler)

	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/execs/", execsHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}

package main2

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
	fmt.Println("END GET")


}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
	fmt.Println("END POST")

}

func put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
	fmt.Println("END PUT")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
	fmt.Println("END DELETE")

}

func params(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Begin: params...")

	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
		fmt.Println("user=", userID)
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
		fmt.Println("comment=", commentID)
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
	fmt.Println("End: params...")

}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)

	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"os"
)

func getPort() string {
	godotenv.Load()
	p := os.Getenv("PORT")
	fmt.Println("ENV"+p)
	if p != "" {
		return ":" + p
	}
	return ":4000"
}	
func testRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("On test route")
	json.NewEncoder(w).Encode("on test route only for testing purpose")
}

func main() {
	PORT := getPort()
	fmt.Println(PORT)
	router := mux.NewRouter()
	router.HandleFunc("/test", testRoute).Methods("GET")
	log.Fatal(http.ListenAndServe(PORT,router))
}
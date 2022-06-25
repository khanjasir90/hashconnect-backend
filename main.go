package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func goDotEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error while reading Env File")
		return ""
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Println("Error while getting key related to key")
		return ""
	}
	return value
}

func testRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("On test route")
	json.NewEncoder(w).Encode("on test route only for testing purpose")
}

func main() {
	PORT := goDotEnvVariable("PORT")
	if PORT == "" {
		PORT = ":3000"
	}
	fmt.Println(PORT)
	router := mux.NewRouter()
	router.HandleFunc("/test", testRoute).Methods("GET")
	log.Fatal(http.ListenAndServe(PORT,router))
}
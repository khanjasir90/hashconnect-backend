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
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Println("Error while getting key related to key")
	}
	return value
}

func testRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On test route")
	json.NewEncoder(w).Encode("On test route")
}

func main() {
	PORT := goDotEnvVariable("PORT")
	fmt.Println(PORT)
	router := mux.NewRouter()
	router.HandleFunc("/test", testRoute)
	log.Fatal(http.ListenAndServe(PORT,router))
}
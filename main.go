package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
	// "github.com/spf13/viper"
)

func main() {
	fmt.Println("your server run in port")
	http.ListenAndServe("127.0.0.1:3000", nil)
}

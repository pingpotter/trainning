package main
import (
	"fmt"
	"log"
	"net/http"
	"time"
	"./wallet"
)

type WalletX struct {
	Id    	  int16   `json:"id"`
	Citizen   int16   `json:"citizen"`
	Fullname  string `json:"name"`
	Opendate   time.Time   `json:"time"`
	Balance   int16	  `json:"id"`
}

var appName = "wallet service"

func main() {

	fmt.Printf("Starting %v\n", appName)
	router := wallet.NewRouter()
	log.Fatal(http.ListenAndServe(":6677", router))

}
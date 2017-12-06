package main
import (
	"log"
	"net/http"
	"routes"
)
func main() {
	log.Fatal(http.ListenAndServe(":8088", routes.Engine()))
}
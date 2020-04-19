package main
import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting up staticserve")

	log.Println("Configuring static file handler")
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	log.Println("Handler configured")

	log.Println("Starting up listener and serving on :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

package main

import (

	//"CanCommerce/middleware"
	//models "CanCommerce/models"

	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	dbUtil "CanCommerce/db"
)

func init() {
	dbUtil.Connect()
}

func main() {

	handler := http.NewServeMux()

	handler.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		name := strings.Replace(r.URL.Path, "/product/", "", 1)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, fmt.Sprintf("Hello %s\n", name))
	})

	handler.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, "Hello world\n")
	})

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)

		io.WriteString(w, "Hello Root\n")
	})

	err := http.ListenAndServe(":9000", handler)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

	//router.GET("/cart", middleware.Middleware(index))
	//router.GET("/login", middleware.Middleware(index))
	//router.GET("/register", middleware.Middleware(index))

}

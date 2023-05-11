package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rinpr/crud-api-golang/controllers"
)

var RegisterImageDataRoutes = func(router *mux.Router) {
	router.HandleFunc("/image-data", controllers.CreateImageData).Methods("POST")
	router.HandleFunc("/image-data", controllers.GetImagesData).Methods("GET")
	router.HandleFunc("/image-data/{imageId}", controllers.GetImageData).Methods("GET")
	router.HandleFunc("/image-data/{imageId}", controllers.UpdateImageData).Methods("PUT")
	// router.HandleFunc("/image-data/{imageId}", controllers.DeleteImageData).Methods("DELETE")
}

func ListenAndServe() {
	fmt.Println("Listen and serve on port 8080!")
	r := mux.NewRouter()
	RegisterImageDataRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

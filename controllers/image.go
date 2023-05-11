package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/rinpr/crud-api-golang/database"
	"github.com/rinpr/crud-api-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateImageData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var image models.ImageData

	json.NewDecoder(r.Body).Decode(&image)
	image.Time = time.Now().Format(time.RFC3339Nano)
	fmt.Println(image.Time)

	collection := database.ImageData()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, image)

	json.NewEncoder(w).Encode(result)
}

func GetImagesData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var images []models.ImageData
	collection := database.ImageData()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var image models.ImageData
		cursor.Decode(&image)
		images = append(images, image)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(w).Encode(images)
}

func GetImageData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["imageId"])
	var image models.ImageData
	collection := database.ImageData()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, models.ImageData{Id: id}).Decode(&image)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(w).Encode(image)
}

func UpdateImageData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["imageId"])
    var image models.ImageData
    collection := database.ImageData()
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err := collection.FindOne(ctx, models.ImageData{Id: id}).Decode(&image)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }

    // Decode the request body into an ImageData struct
    err = json.NewDecoder(r.Body).Decode(&image)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"message": "Invalid request payload"}`))
        return
    }
    // Update the image data
    update := bson.M{
        "$set": bson.M{
            "time":       image.Time,
            "path":       image.Path,
            "sizebefore": image.SizeBefore,
            "sizeafter":  image.SizeAfter,
            "issuccess":  image.IsSuccess,
        },
    }
    err = collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, update).Decode(&image)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }

    json.NewEncoder(w).Encode(image)
}



func DeleteImageData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["imageId"])
    collection := database.ImageData()
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    _, err := collection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "Image data deleted successfully"}`))
}

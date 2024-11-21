package control

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/joko345/golangmyskill/mongo-go/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserControl mengelola operasi CRUD untuk user
type UserControl struct {
	client *mongo.Client
}

// NewUserControl membuat instance baru dari UserControl
func NewUserControl(client *mongo.Client) *UserControl {
	return &UserControl{client}
}

// GetUser mendapatkan user berdasarkan ID
func (uc UserControl) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // Mendapatkan ID dari parameter URL

	// Validasi ID apakah berbentuk ObjsectId
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Koneksi ke database dan koleksi
	collection := uc.client.Database("mongo-golang").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cari user berdasarkan ID
	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Konversi data user ke JSON
	uj, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Kirim data sebagai respon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser membuat user baru
func (uc UserControl) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User

	// Decode data JSON dari body request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Buat ID baru untuk user
	user.Id = primitive.NewObjectID()

	// Koneksi ke database dan koleksi
	collection := uc.client.Database("mongo-golang").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Simpan data ke MongoDB
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Konversi data user ke JSON dan kirimkan respon
	uj, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser menghapus user berdasarkan ID
func (uc UserControl) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Validasi ID apakah berbentuk ObjectId
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Koneksi ke database dan koleksi
	collection := uc.client.Database("mongo-golang").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Hapus user berdasarkan ID
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Kirimkan respon sukses
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted User: %s\n", id)
}

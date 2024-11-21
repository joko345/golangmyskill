package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/joko345/golangmyskill/mongo-go/control"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := httprouter.New()
	client := getSession()               // Mendapatkan MongoDB client
	uc := control.NewUserControl(client) // Pass client ke kontroler

	// Konfigurasi route
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	// Jalankan server
	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe("localhost:8080", r)
}

// getSession: Membuat koneksi ke MongoDB menggunakan MongoDB Go Driver
func getSession() *mongo.Client {
	// URI MongoDB
	mongoURI := "mongodb://localhost:27017/"

	// Konfigurasi client MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(fmt.Sprintf("Gagal membuat MongoDB client: %s", err))
	}

	// Mengatur timeout untuk koneksi
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Menghubungkan ke server MongoDB
	err = client.Connect(ctx)
	if err != nil {
		panic(fmt.Sprintf("Gagal terhubung ke MongoDB: %s", err))
	}

	// Menguji koneksi
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("Tidak dapat terhubung ke MongoDB: %s", err))
	}

	fmt.Println("Berhasil terhubung ke MongoDB")
	return client
}

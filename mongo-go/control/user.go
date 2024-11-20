package control

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joko345/golangmyskill/mongo-go/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserControl struct {
	session *mgo.Session
}

func NewUserControl(s *mgo.Session) *UserControl {
	return &UserControl{s}
}

func (uc UserControl) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // deklarasi id untuk id pada mongodb

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // respon
		return
	}

	oid := bson.ObjectIdHex(id) // cek apakah id hex/numerik

	u := models.User{}                                                                   // struct user
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil { // users adalah nama koleksi/tabel
		w.WriteHeader(http.StatusNotFound) // cari id dengan value oid di mongodb yang diwakili variable u
		return
	}

	uj, err := json.Marshal(u) // simpan data ke json uj
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj) // print value ke body sebagai respon
}

func (uc UserControl) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u) // decode data json ke format mongodb

	u.Id = bson.NewObjectId() // buat id otomatis

	uc.session.DB("mongo-golang").C("users").Insert(u)
	// setelah input ubah data ke json lagi dan kirim ke user
	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserControl) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted User: %s\n", oid)
}

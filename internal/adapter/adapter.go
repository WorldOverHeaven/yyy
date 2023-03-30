package adapter

import (
	"YYY/internal/model"
	"YYY/internal/service"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

type Adapter struct {
	db *service.Service
}

func NewAdapter(db *service.Service) *Adapter {
	return &Adapter{db: db}
}

func (a *Adapter) PostGroup(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var group model.Group

	err := decoder.Decode(&group)
	if err != nil {
		log.Println("unmarshal error ", err)
		w.WriteHeader(500)
		return
	}

	if group.Name == "" {
		log.Println("error, name is null")
		w.WriteHeader(500)
		return
	}

	err = a.db.AddGroup(&group)
	if err != nil {
		log.Println("db error ", err)
		w.WriteHeader(500)
		return
	}

	log.Println(group)
	_, err = fmt.Fprintf(w, strconv.FormatInt(group.ID, 10))
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func (a *Adapter) GetGroups(w http.ResponseWriter, r *http.Request) {

	group, err := a.db.GetGroups()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	data, err := json.Marshal(group)

	_, err = fmt.Fprintf(w, string(data))
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func (a *Adapter) GetGroupsByID(w http.ResponseWriter, r *http.Request) {

	idString := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	group, err := a.db.GetGroupByID(id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	data, err := json.Marshal(group)

	_, err = fmt.Fprintf(w, string(data))
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// protHandler é responsável por gerenciar todos os recursos da struct protocolo
func protHandler(w http.ResponseWriter, request *http.Request) {
	identifier := strings.TrimPrefix(request.URL.Path, "/protocolos/")
	switch {
	case request.Method == "GET" && len(identifier) > 0:
		findByID(w, request, identifier)
	case request.Method == "GET":
		findAll(w, request)
	case request.Method == "POST":
		create(w, request)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Desculpa.... Nada encontrado! :(")
	}
}

func create(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	p := Protocolo{}

	p.Nome = request.FormValue("nome")
	p.Checado, _ = strconv.ParseBool(request.FormValue("checado"))
	p.Criacao = time.Now()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	c := session.DB("test").C("protocolos")
	err = c.Insert(&p)

	if err != nil {
		panic(err)
	}

	resp, _ := json.Marshal(p)
	fmt.Fprint(w, string(resp))
}

func findAll(w http.ResponseWriter, req *http.Request) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	db := session.DB("test").C("protocolos")
	var protocolos []Protocolo

	err = db.Find(bson.M{}).All(&protocolos)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Ocorreu um erro :(")
	}

	resp, _ := json.Marshal(protocolos)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, string(resp))
}

func findByID(w http.ResponseWriter, r *http.Request, id string) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	db := session.DB("test").C("protocolos")

	// Busca pelo id
	result := Protocolo{}
	err = db.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpe ocorreu um erro!! :-(")
	} else {
		resp, _ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(resp))
	}

	defer session.Close()
}

func main() {
	http.HandleFunc("/protocolos/", protHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

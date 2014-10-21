package server

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"gitlab.mitre.org/fhir/models"
	"github.com/gorilla/mux"
	"os"
)

func OperationOutcomeIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.OperationOutcome
	c := Database.C("operationoutcomes")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(result)
}

func OperationOutcomeShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	}	else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("operationoutcomes")

	result := models.OperationOutcome{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(result)
}

func OperationOutcomeCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	operationoutcome := &models.OperationOutcome{}
	err := decoder.Decode(operationoutcome)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("operationoutcomes")
	i := bson.NewObjectId()
	operationoutcome.Id = i.Hex()
	err = c.Insert(operationoutcome)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://" + host + "/operationoutcome/" + i.Hex())
}

func OperationOutcomeUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	}	else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	operationoutcome := &models.OperationOutcome{}
	err := decoder.Decode(operationoutcome)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("operationoutcomes")
	operationoutcome.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, operationoutcome)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func OperationOutcomeDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	}	else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("operationoutcomes")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
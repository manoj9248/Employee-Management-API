package endpoints

import (
	"Employee_crud_mux/db"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type employee struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Age       int                `json:"age,omitempty" bson:"age,omitempty"`
	Position  string             `json:"position,omitempty" bson:"position,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
type Handler struct {
	DB *mongo.Collection
}

func EndPoints(s *mux.Router) {
	urlHandler := &Handler{
		DB: db.DBinit().Database("testdb").Collection("employee"),
	}
	s.HandleFunc("/createEmployee", urlHandler.createEmployee).Methods("POST")
	s.HandleFunc("/getAllEmployees", urlHandler.getAllEmployees).Methods("GET")
	s.HandleFunc("/getAllEmployeeById/{id}", urlHandler.getAllEmployeeById).Methods("GET")
	s.HandleFunc("/patchEmployee/{id}", urlHandler.patchEmployee).Methods("PATCH")
	s.HandleFunc("/updateEmployee/{id}", urlHandler.updateEmployee).Methods("PUT")
	s.HandleFunc("/deleteEmployee/{id}", urlHandler.deleteEmployee).Methods("DELETE")
}
func (h *Handler) createEmployee(w http.ResponseWriter, r *http.Request) {
	var emp employee
	emp.CreatedAt = time.Now()
	json.NewDecoder(r.Body).Decode(&emp)
	result, err := h.DB.InsertOne(context.Background(), emp)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.InsertedID)
}
func (h *Handler) getAllEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []employee
	cursor, err := h.DB.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var emp employee
		if err := cursor.Decode(&emp); err != nil {
			log.Println(err)
			json.NewEncoder(w).Encode(err)
			return
		}
		employees = append(employees, emp)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
func (h *Handler) getAllEmployeeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var emp employee

	// MongoDB findOne() method to get employee by ID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = h.DB.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&emp)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}
func (h *Handler) patchEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var empUpdates map[string]interface{}
	json.NewDecoder(r.Body).Decode(&empUpdates)
	updateFields := bson.M{}
	for key, value := range empUpdates {
		updateFields[key] = value
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	_, err = h.DB.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": updateFields})
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Employee patched successfully")
}
func (h *Handler) updateEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var emp employee
	json.NewDecoder(r.Body).Decode(&emp)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = h.DB.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": emp})
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Employee updated successfully")

}
func (h *Handler) deleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	_, err = h.DB.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Employee deleted successfully")
}

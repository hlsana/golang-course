package handler

import (
	"encoding/json"
	"main/hw09/models"
	"net/http"

	"github.com/gorilla/mux"
)

var class = models.Class{
	Name: "English",
	Students: map[string]models.Student{
		"1": {ID: "1", Name: "James", Scores: map[string]int{"English": 90}},
		"2": {ID: "2", Name: "Jennie", Scores: map[string]int{"English": 85}},
		"3": {ID: "3", Name: "Jacob", Scores: map[string]int{"English": 100}},
		"4": {ID: "4", Name: "Jessica", Scores: map[string]int{"English": 73}},
		"5": {ID: "5", Name: "Janine", Scores: map[string]int{"English": 82}},
	},
}

func GetClassInfo(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Name     string            `json:"name"`
		Students map[string]string `json:"students"`
	}{
		Name:     class.Name,
		Students: make(map[string]string),
	}
	for id, student := range class.Students {
		response.Students[id] = student.Name
	}
	json.NewEncoder(w).Encode(response)
}

// GetStudentInfo handles the request for student information
func GetStudentInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	student, exists := class.Students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	response := struct {
		Name   string         `json:"name"`
		Scores map[string]int `json:"scores"`
	}{
		Name:   student.Name,
		Scores: student.Scores,
	}
	json.NewEncoder(w).Encode(response)
}

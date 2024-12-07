package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/vishalvivekm/pyqqserver/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"errors"
)

type Handler struct {
	db *mongo.Client
}
func NewHandler(db *mongo.Client) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetSubjects(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server is hit, with request params: ", mux.Vars(r))
	vars := mux.Vars(r)

	semester := vars["semester"]
	branch := vars["branch"]

	semesterLabel,ok := model.SemListNew[semester]
	if !ok {
		log.Printf("semester not found")
	}
	branchLabel := findBranchLabel(branch)

	subjects, err := h.FindSubjectsByBranchAndSemFromMongo(branchLabel, semesterLabel)
	if err != nil {
		log.Printf("Error fetching subjects: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := `{"status": "error", "message": "Error fetching subjects: ` + err.Error() + `"}`
		w.Write([]byte(msg))
		return
	}
	if len(subjects) == 0 {
		log.Printf("No subjects found")
		w.WriteHeader(http.StatusNotFound)
		msg := fmt.Sprintf(`{"status": "success", "message": "no subjects found for branch: %s and semester: %s"}`, branch, semester)
		fmt.Println()
		w.Write([]byte(msg))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subjects)
}

// getSubjectDetails handles fetching details for a specific subject
func (h *Handler) GetSubjectDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server is hit, with request params: ", mux.Vars(r))
	vars := mux.Vars(r)
	// semester := vars["semester"]
	// branch := vars["branch"]
	subject := vars["subject"]

	// semesterLabel := findSemesterLabel(semester)
	// branchLabel := findBranchLabel(branch)

	subjectDetails, err := h.FindSubjectDetailsByBranchAndSemFromMongo(subject)
	if err != nil {
		log.Printf("Error fetching subject details: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(subjectDetails) == 0 {
		log.Printf("No subject details found")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		msg := map[string]string {
			"status": "success",
			"message": "no subject details found for subject: " + subject,
		}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subjectDetails)
}
func (h *Handler) GetResources(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("server is hit, with request params: ", params)
	resourceType := params["type"]
	subject := params["subject"]

	switch(resourceType) {
	case "notes":
		notes, err := h.GetNotes(subject)
		if err != nil {
			log.Println("Error fetching notes: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	case "pyq":
		pyqs, err := h.GetPYQs(subject)
		if err != nil {
			log.Println("Error frtching pyqs: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pyqs)
	case "books":
		books, err := h.GetBooks(subject)
		if err != nil {
			log.Println("Error fetching books: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	default:
		log.Println("Invalid resource type", resourceType)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"status": "error", "message": "Invalid resource type: ` + resourceType +`"}`))
		fmt.Println()
		return
	}

}

func (h *Handler) GetNotes(subject string) (model.Notes, error) {
	collection := h.db.Database("PYQHUb").Collection("notes")

	filter := bson.M{
		"subjectId":  subject,
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results model.Notes
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	fmt.Println("results", results)
	if len(results) == 0 {
		return nil, errors.New("no notes found")
	}

	return results, nil
}

func (h *Handler) GetPYQs(subject string) (model.PYQs, error) {
	collection := h.db.Database("PYQHUb").Collection("pyq")

	filter := bson.M{
		"subjectId":  subject,
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results model.PYQs
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	fmt.Println("results", results)
	if len(results) == 0 {
		return nil, fmt.Errorf("No notes found")
	}

	return results, nil
}
func (h *Handler) GetBooks(subject string) (model.Books, error) {
	collection := h.db.Database("PYQHUb").Collection("books")

	filter := bson.M{
		"subjectId":  subject,
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results model.Books
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	fmt.Println("results", results)
	if len(results) == 0 {
		return nil, fmt.Errorf("No notes found")
	}

	return results, nil
}

// Utility functions to find original labels
func findSemesterLabel(encodedSemester string) string {
	for label, value := range model.SemesterList {
		if value == encodedSemester {
			return label
		}
	}
	return ""
}

func findBranchLabel(encodedBranch string) string {
	for label, value := range model.BranchList {
		if value == encodedBranch {
			return label
		}
	}
	return ""
}


func (h *Handler) FindSubjectsByBranchAndSemFromMongo(branch string, semester int) ([]string, error) {
	collection := h.db.Database("PYQHUb").Collection("subjects")

	filter := bson.M{
		"branches":  branch,
		"semester":  semester,
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []model.Subject
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	fmt.Println("results", results)
	// Extract subject IDs
	var subjectIDs []string
	for _, subject := range results {
		subjectIDs = append(subjectIDs, subject.SubjectID)
	}

	return subjectIDs, nil
}

func (h *Handler) FindSubjectDetailsByBranchAndSemFromMongo(subjectId string) ([]model.SubjectDetail, error) {

	collection := h.db.Database("PYQHUb").Collection("subjectdetails")

	filter := bson.M{
		"subjectID":  subjectId,
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []model.SubjectDetail
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	fmt.Println("results", results)
	
	if len(results) == 0 {
		return nil,nil // errror is nil
	}

	return results, nil
}
package handlers

import (
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/requests"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetPetReport(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewPetReport(r)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if req.StartTime == 0 || req.EndTime == 0 {
		http.Error(w, "Missing start_time or end_time parameters", http.StatusBadRequest)
		return
	}

	startTime := primitive.Timestamp{T: uint32(req.StartTime), I: 0}
	endTime := primitive.Timestamp{T: uint32(req.EndTime), I: 0}

	petsDB := MongoDB(r).Pets()
	healthDB := MongoDB(r).HealthData()

	pet, err := petsDB.Get(req.PetID)
	if err != nil {
		http.Error(w, "Failed to retrieve pet information", http.StatusInternalServerError)
		return
	}

	filter := bson.M{
		"pet_id": pet.ID,
		"time": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
	}

	healthData, err := healthDB.GetByFilter(filter)
	if err != nil {
		http.Error(w, "Failed to retrieve health data", http.StatusInternalServerError)
		return
	}

	ownerName := requests.GetUserName(MongoDB(r), pet.OwnerID)

	pdf, err := GeneratePetReportPDF(pet, healthData, ownerName)
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `attachment; filename="pet_report.pdf"`)
	w.WriteHeader(http.StatusOK)

	err = pdf.Output(w)
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}
}

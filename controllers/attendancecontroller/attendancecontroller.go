package attendancecontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"simple-elearning-backend/config"
	"simple-elearning-backend/helper"
	"simple-elearning-backend/models"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

// Index retrieves a list of all attendance records
func Index(w http.ResponseWriter, r *http.Request) {
	var attendances []models.Attendance

	if err := config.DB.Find(&attendances).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List Attendance Records", attendances)
}

// Create adds a new attendance record
func Create(w http.ResponseWriter, r *http.Request) {
	var attendance models.Attendance

	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&attendance).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Create Attendance Record", nil)
}

// Detail retrieves details of a specific attendance record
func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var attendance models.Attendance

	if err := config.DB.First(&attendance, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Attendance Record Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail Attendance Record", attendance)
}

// Update modifies an existing attendance record
func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var attendance models.Attendance

	if err := config.DB.First(&attendance, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Attendance Record Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("id = ?", id).Updates(&attendance).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Update Attendance Record", nil)
}

// Delete removes a specific attendance record
func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var attendance models.Attendance

	res := config.DB.Delete(&attendance, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Attendance Record Not Found", nil)
		return
	}

	helper.Response(w, 200, "Success Delete Attendance Record", nil)
}

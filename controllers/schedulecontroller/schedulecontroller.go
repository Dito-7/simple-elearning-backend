package schedulecontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"simple-elearning-backend/config"
	"simple-elearning-backend/helper"
	"simple-elearning-backend/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Index retrieves a list of all schedules
func Index(w http.ResponseWriter, r *http.Request) {
	var schedules []models.Schedule

	if err := config.DB.Find(&schedules).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List of Schedules", schedules)
}

// Create adds a new schedule
func Create(w http.ResponseWriter, r *http.Request) {
	var schedule models.Schedule

	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	schedule.CreatedAt = time.Now()
	schedule.UpdatedAt = time.Now()

	if err := config.DB.Create(&schedule).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Schedule created successfully", nil)
}

// Detail retrieves the details of a specific schedule by ID
func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var schedule models.Schedule

	if err := config.DB.First(&schedule, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Schedule not found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Schedule details", schedule)
}

// Update modifies an existing schedule
func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var schedule models.Schedule

	if err := config.DB.First(&schedule, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Schedule not found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	schedule.UpdatedAt = time.Now()

	if err := config.DB.Save(&schedule).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Schedule updated successfully", nil)
}

// Delete removes a specific schedule by ID
func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var schedule models.Schedule

	res := config.DB.Delete(&schedule, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Schedule not found", nil)
		return
	}

	helper.Response(w, 200, "Schedule deleted successfully", nil)
}

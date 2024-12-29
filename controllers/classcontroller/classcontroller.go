package classcontroller

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

// Index retrieves a list of all classes
func Index(w http.ResponseWriter, r *http.Request) {
	var classes []models.Class
	var classResponses []models.ClassResponse

	if err := config.DB.Find(&classes).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	// Map classes to ClassResponse
	for _, class := range classes {
		classResponses = append(classResponses, models.ClassResponse{
			ID:   class.ID,
			Name: class.Name,
		})
	}

	helper.Response(w, 200, "List of Classes", classResponses)
}

// Create adds a new class
func Create(w http.ResponseWriter, r *http.Request) {
	var class models.Class

	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	class.CreatedAt = time.Now()
	class.UpdatedAt = time.Now()

	if err := config.DB.Create(&class).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Class created successfully", nil)
}

// Detail retrieves the details of a specific class by ID
func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var class models.Class

	if err := config.DB.First(&class, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Class not found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Class details", class)
}

// Update modifies an existing class
func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var class models.Class

	if err := config.DB.First(&class, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Class not found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	class.UpdatedAt = time.Now()

	if err := config.DB.Save(&class).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Class updated successfully", nil)
}

// Delete removes a specific class by ID
func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var class models.Class

	res := config.DB.Delete(&class, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Class not found", nil)
		return
	}

	helper.Response(w, 200, "Class deleted successfully", nil)
}

package example

import (
	"encoding/json"
	"net/http"

	exampleModel "github.com/faisd405/go-restapi-chi/src/app/example/model"
	database "github.com/faisd405/go-restapi-chi/src/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func errorHandler(w http.ResponseWriter, r *http.Request, err error, httpStatus, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if message == "" {
		message = "Internal server error"
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "error",
		"message": message,
		"error":   err.Error(),
	})
}

func Index(w http.ResponseWriter, r *http.Request) {
	var examples []exampleModel.Example

	database.DB.Find(&examples)

	w.Header().Set("Content-Type", "application/json")

	var response = map[string]interface{}{
		"status":  "success",
		"message": "Data found",
		"data":    examples,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Show(w http.ResponseWriter, r *http.Request) {
	var example exampleModel.Example
	id := chi.URLParam(r, "id")

	if err := database.DB.First(&example, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  "error",
				"message": "Data not found",
			})
			return
			// errorHandler(w, r, err, "Example not found")
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  "error",
				"message": "Internal server error",
			})
			return
			// errorHandler(w, r, err, "")
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Data found",
		"data":    example,
	})
}

func Create(w http.ResponseWriter, r *http.Request) {
	var example exampleModel.Example

	err := json.NewDecoder(r.Body).Decode(&example)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	validate = validator.New()
	err = validate.Struct(example)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	database.DB.Create(&example)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Data created",
		"data":    example,
	})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var example exampleModel.Example
	id := chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&example)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	validate = validator.New()
	err = validate.Struct(example)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if database.DB.Model(&example).Where("id = ?", id).Updates(&example).RowsAffected == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Example not updated",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Data Updated",
	})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var example exampleModel.Example
	id := chi.URLParam(r, "id")

	if database.DB.First(&example, id).RowsAffected == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Data not found",
		})
		return
	}

	database.DB.Delete(&example)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Data deleted",
	})
}

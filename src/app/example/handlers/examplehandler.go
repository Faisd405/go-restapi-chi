package handlers

import (
	"net/http"
	"strconv"

	exampleInterface "github.com/faisd405/go-restapi-chi/src/app/example/interfaces"
	exampleModel "github.com/faisd405/go-restapi-chi/src/app/example/models"
	exampleRequest "github.com/faisd405/go-restapi-chi/src/app/example/request"
	"github.com/faisd405/go-restapi-chi/src/helper/database"
	"github.com/faisd405/go-restapi-chi/src/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type exampleHandler struct {
	ExampleService exampleInterface.ExampleService
}

func NewExampleHandler(exampleService exampleInterface.ExampleService) *exampleHandler {
	return &exampleHandler{
		ExampleService: exampleService,
	}
}

// use a single instance of Validate, it caches struct info
// var validate *validator.Validate

func (handler *exampleHandler) Index(w http.ResponseWriter, r *http.Request) {

	params := database.BuildParams(r)

	example, err := handler.ExampleService.FindAll(r.Context(), params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := map[string]interface{}{
		"status":      "success",
		"data":        example,
		"limit":       params["limit"],
		"currentPage": params["currentPage"],
	}

	utils.WriteJson(w, result)
}

func (handler *exampleHandler) Show(w http.ResponseWriter, r *http.Request) {
	IDArg := chi.URLParam(r, "id")
	ID, err := strconv.ParseInt(IDArg, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	example, err := handler.ExampleService.FindById(r.Context(), uint(ID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		result := map[string]interface{}{
			"status":  "error",
			"message": "Example not found",
		}

		utils.WriteJson(w, result)
		return
	}

	result := map[string]interface{}{
		"status": "success",
		"data":   example,
	}

	utils.WriteJson(w, result)
}

func (handler *exampleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input exampleRequest.ExampleRequest
	err := utils.ReadJson(r, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	example := exampleModel.Example{
		Example1: input.Example1,
		Example2: input.Example2,
	}

	err = handler.ExampleService.Create(r.Context(), &example)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := map[string]interface{}{
		"status":  "success",
		"message": "Example created successfully",
		"data":    example,
	}

	utils.WriteJson(w, result)
}

func (handler *exampleHandler) Update(w http.ResponseWriter, r *http.Request) {
	IDArg := chi.URLParam(r, "id")
	ID, err := strconv.ParseUint(IDArg, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var input exampleRequest.ExampleRequest
	err = utils.ReadJson(r, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	example := exampleModel.Example{
		Example1: input.Example1,
		Example2: input.Example2,
	}

	err = handler.ExampleService.Update(r.Context(), uint(ID), &example)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := map[string]interface{}{
		"status":  "success",
		"message": "Example updated successfully",
		"data":    example,
	}

	utils.WriteJson(w, result)

}

func (handler *exampleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	IDArg := chi.URLParam(r, "id")
	ID, err := strconv.ParseUint(IDArg, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = handler.ExampleService.Delete(r.Context(), uint(ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := map[string]interface{}{
		"status":  "success",
		"message": "Example deleted successfully",
	}

	utils.WriteJson(w, result)
}

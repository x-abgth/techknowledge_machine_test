package handlers

import (
	"errors"
	"net/http"

	"example.com/machine_test/task_4/internal/usecase"
	"example.com/machine_test/task_4/pkg/models"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var validate *validator.Validate

func NewValidator() {
	validate = validator.New()
}

func (app *JsonResponse) NoPageHandler(w http.ResponseWriter, r *http.Request) {
	app.ErrorJSON(w, errors.New("the page you are trying to access does not exist"))
}

func (app *JsonResponse) GetAllItems(w http.ResponseWriter, r *http.Request) {
	data, err := usecase.GetAllItemsUsecase()
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "Successfully fetched all items details",
		Data:    data,
	}

	app.WriteJSON(w, http.StatusFound, resp)
}

func (app *JsonResponse) GetOneItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := usecase.GetOneItemUsecase(id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "Successfully fetched the item details",
		Data:    data,
	}

	app.WriteJSON(w, http.StatusFound, resp)
}

func (app *JsonResponse) CreateOneItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	err := app.ReadJSON(w, r, &item)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	err = validate.Struct(item)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	err = usecase.InsertItemUsecase(item)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "Successfully inserted the item",
	}

	app.WriteJSON(w, http.StatusOK, resp)
}

func (app *JsonResponse) UpdateOneItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var item models.Item
	err := app.ReadJSON(w, r, &item)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	err = validate.Struct(item)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	err = usecase.UpdateItemUsecase(id, item)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "Successfully updated the item",
	}

	app.WriteJSON(w, http.StatusOK, resp)
}

func (app *JsonResponse) DeleteOneItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := usecase.DeleteItemUsecase(id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "Successfully deleted the item",
	}

	app.WriteJSON(w, http.StatusOK, resp)
}

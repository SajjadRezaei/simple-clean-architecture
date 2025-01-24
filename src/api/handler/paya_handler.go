package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"simpleBank/src/api/dto"
	"simpleBank/src/helper"
	"simpleBank/src/usecase"
)

type PayaHandler struct {
	service *usecase.PayaUseCase
}

func NewPayaHandler(useCase *usecase.PayaUseCase) *PayaHandler {
	return &PayaHandler{service: useCase}
}

func (h *PayaHandler) CreatePayaRequest(w http.ResponseWriter, r *http.Request) {
	var req dto.PayaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.SendErrorResponse(w, err)
		return
	}

	err := req.Validate()
	if err != nil {
		helper.SendErrorResponse(w, err)
		return
	}

	createdPaya, err := h.service.CreatePayaRequest(dto.ToCreatePaya(req))
	if err != nil {
		helper.SendErrorResponse(w, err)
		return
	}

	helper.SendSuccessResponse(w, map[string]any{
		"request": createdPaya,
	}, "")
}

func (h *PayaHandler) UpdatePayaRequest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var req dto.UpdatePayaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.SendErrorResponse(w, err)
		return
	}

	updatedPaya, err := h.service.UpdatePayaRequest(id, req.Status, req.Note)
	if err != nil {
		helper.SendErrorResponse(w, err)
		return
	}

	helper.SendSuccessResponse(w, map[string]any{"request": updatedPaya}, "")
}

func (h *PayaHandler) ListPayaRequests(w http.ResponseWriter, r *http.Request) {
	helper.SendSuccessResponse(w, map[string]any{
		"requests": h.service.GetPayaRequests(),
	}, "")
}

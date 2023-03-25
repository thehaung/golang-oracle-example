package route

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/thehaung/golang-oracle-example/domain"
	"github.com/thehaung/golang-oracle-example/internal/util/httputil"
	"github.com/thehaung/golang-oracle-example/internal/util/typeutil"
	"net/http"
	"strconv"
)

type staffRoutes struct {
	staffUseCase domain.StaffUseCase
}

func newStaffRoute(router *chi.Mux, useCase domain.StaffUseCase) {
	route := &staffRoutes{
		staffUseCase: useCase,
	}

	router.Route("/staff", func(r chi.Router) {
		r.Post("/", route.createStaff)
		r.Get("/", route.listStaff)
		r.Get("/{id}", route.findStaffById)
		r.Put("/{id}", route.updateStaff)
		r.Delete("/{id}", route.deleteStaffById)
	})
}

func (s *staffRoutes) createStaff(w http.ResponseWriter, r *http.Request) {
	httpUtil := httputil.NewHttpUtil(w, r)

	var req domain.CreateStaffRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpUtil.WriteError(http.StatusBadRequest, "Error unmarshalling the request", err)
		return
	}

	if err := typeutil.Validator().Struct(req); err != nil {
		httpUtil.WriteError(http.StatusBadRequest, "Error unmarshalling the request", err)
		return
	}

	staff := domain.Staff{
		ID:           req.ID,
		Name:         req.Name,
		TeamName:     req.TeamName,
		Organization: req.Organization,
		Title:        req.Title,
		OnboardDate:  req.OnboardDate,
	}

	create, err := s.staffUseCase.Create(context.Background(), staff)
	if err != nil {
		httpUtil.WriteError(http.StatusInternalServerError, "Can't create Staff. Try again!", err)
		return
	}

	httpUtil.WriteJson(http.StatusCreated, create)
}

func (s *staffRoutes) findStaffById(w http.ResponseWriter, r *http.Request) {
	httpUtil := httputil.NewHttpUtil(w, r)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id == 0 {
		httpUtil.WriteError(http.StatusBadRequest, "id is invalid type", err)
		return
	}

	staff, err := s.staffUseCase.FindById(context.Background(), int64(id))
	if err != nil {
		httpUtil.WriteError(http.StatusInternalServerError, "Can't find Staff. Try again!", err)
		return
	}

	httpUtil.WriteJson(http.StatusOK, staff)
}

func (s *staffRoutes) listStaff(w http.ResponseWriter, r *http.Request) {
	httpUtil := httputil.NewHttpUtil(w, r)

	staffs, err := s.staffUseCase.List(context.Background())
	if err != nil {
		httpUtil.WriteError(http.StatusInternalServerError, "Can't find Staffs. Try again!", err)
		return
	}

	httpUtil.WriteJson(http.StatusOK, staffs)
}

func (s *staffRoutes) updateStaff(w http.ResponseWriter, r *http.Request) {
	httpUtil := httputil.NewHttpUtil(w, r)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id == 0 {
		httpUtil.WriteError(http.StatusBadRequest, "id is invalid type", err)
		return
	}

	var req domain.UpdateStaffRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpUtil.WriteError(http.StatusBadRequest, "Error unmarshalling the request", err)
		return
	}

	if err = typeutil.Validator().Struct(req); err != nil {
		httpUtil.WriteError(http.StatusBadRequest, "Request Body is invalid. Try again!", err)
		return
	}

	staff := domain.Staff{
		ID:           int64(id),
		Name:         req.Name,
		TeamName:     req.TeamName,
		Organization: req.Organization,
		Title:        req.Title,
		OnboardDate:  req.OnboardDate,
	}
	updated, err := s.staffUseCase.Update(context.Background(), staff)
	if err != nil {
		httpUtil.WriteError(http.StatusInternalServerError, "Can't update Staff. Try again!", err)
		return
	}

	httpUtil.WriteJson(http.StatusOK, updated)
}

func (s *staffRoutes) deleteStaffById(w http.ResponseWriter, r *http.Request) {
	httpUtil := httputil.NewHttpUtil(w, r)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id == 0 {
		httpUtil.WriteError(http.StatusBadRequest, "id is invalid type", err)
		return
	}

	err = s.staffUseCase.Delete(context.Background(), int64(id))
	if err != nil {
		httpUtil.WriteError(http.StatusInternalServerError, "Can't delete Staff. Try again!", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package route

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/thehaung/golang-oracle-example/domain"
	"github.com/thehaung/golang-oracle-example/internal/util/httputil"
	"github.com/thehaung/golang-oracle-example/internal/util/typeutil"
	"net/http"
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
		httpUtil.WriteError(http.StatusInternalServerError, "Error unmarshalling the request", err)
		return
	}

	httpUtil.WriteJson(http.StatusCreated, create)
}

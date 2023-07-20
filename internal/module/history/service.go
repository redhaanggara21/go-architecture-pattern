package history

import (
	"context"

	"red21.id/learn/bengkel/domain"
)

type service struct {
	vehicleRepository domain.VehicleRepository
	historyRepository domain.HistoryRepository
}

func NewService(vehicle domain.VehicleRepository, historyReposirepository domain.HistoryRepository) domain.HistoryService {
	return &service{vehicleRepository, historyReposirepository}
}

func (s service) FindHistorical(ctx context.Context, vin string) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVIN(ctx, vin)
	if err != nil {
		return domain.ApiResponse{
			Code:    "99",
			Message: err.Error(),
		}
	}
	if vehicle == (domain.Vehicle{}) {
		return domain.ApiResponse{
			Code:    "01",
			Message: "VEHICLE NOT FOUND",
		}
	}
}

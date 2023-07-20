package vehicle

import (
	"context"

	"red21.id/learn/bengkel/domain"
	"red21.id/learn/bengkel/internal/module/customer"
	"red21.id/learn/bengkel/internal/module/history"
)

type service struct {
	vehicleRepository domain.VehicleRepository
	historyRepository domain.HistoryRepository
}

func NewService(vehicleRepository domain.VehicleRepository, historyRepository domain.HistoryRepository) domain.VehicleService {
	return &service{vehicleRepository,historyRepository}
}

func (s service) FindHistorical(ctx context.Context, vin string) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVIN(ctx, vin)
	if err != nil {
		return domain.ApiResponse{
			Code: "99",
			Message: err.Error(),
		}
	}

	if vehicle == (domain.Vehicle{}){
		return domain.ApiResponse{
			Code: "01",
			Message: "VEHICLE NOT FOUND",
		}
	}

	histories, err := s.historyRepository.FindByVehicle(ctx, vehicle.ID)
	if err != nil {
		return domain.ApiResponse{
			Code: "99",
			Message: err.Error(),
		}
	}

	var historiesData []domain.HistoryData
	for _, v:= range histories{
		historiesData = append(historiesData, domain.HistoryData{
			VehicleID: v.Vehicle.ID,
			CustomerID: v.CustomerID,
			PIC: v.PIC,
			PlateNumber: v.PlateNumber,
			Notes: v.Notes,
		})
	}

	result := domain.VehicleHistorical{
		ID: vehicle.ID,
		VIN: vehicle.VIN,
		Brand: vehicle.Brand,
		Histories: historiesData, 
	}

	return domain.ApiResponse{
		Code: "00",
		Message: "APPROVED",
		Data: result,
	}
}

func (s service) StoreHistorical(ctx context.Context, req domain.VehicleHistoricalRequest) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVIN(ctx, req.VIN)
	
	if err != {
		return domain.ApiResponse{
			Code: "99",
			Message: err.Error()
		}
	}

	if vehicle == (domain.Vehicle{}){
		err = s.vehicleRepository.Insert(ctx, &vehicle)
		if err != nil {
			return domain.ApiResponse{
				Code: "99",
				Message: err.Error(),
			}
		}

		history := domain.HistoryDetail{
			VehicleID: vehicle.ID,
			CustomerID: customer.ID,
			PIC: req.PIC,
			PlateNumber: req.PlateNumber,
			Notes: req.Notes,
		}
		
		err = s.historyRepository.Insert(ctx, &history)
		if err != nil {
			return domain.ApiResponse{
				Code: "99",
				Message: err.Error(),
			}
		} 

		return domain.ApiResponse{
			Code: : "00",
			Message: "APPROVED"
		}
	}
}
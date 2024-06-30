package vehicle

import (
	"bengkel_app/domain"
	"context"
	"time"
)

type service struct {
	vehicleRepository domain.VehicleRepository
	historyRepository domain.HistoryRepository
}

func NewService(vehicleRepository domain.VehicleRepository, histroryRepository domain.HistoryRepository) domain.VehicleService {
	return &service{
		vehicleRepository, histroryRepository,
	}
}

func (s service) FindHistorical(ctx context.Context, vin string) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVIN(ctx, vin)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	if vehicle == (domain.Vehicle{}) {
		return domain.ApiResponse{
			Code:    "404",
			Message: "Vehicle Not Found",
		}
	}

	histories, err := s.historyRepository.FindByVehicle(ctx, vehicle.ID)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	var historiesData []domain.HistoryData
	for _, v := range histories {
		historiesData = append(historiesData, domain.HistoryData{
			VehicleID:   v.ID,
			PIC:         v.PIC,
			Notes:       v.Notes,
			CustomerID:  v.CustomerID,
			PlateNumber: v.PlateNumber,
			ComeAt:      v.CreatedAt.Format(time.RFC822Z),
		})
	}

	result := domain.VehicleHistorical{
		ID:        vehicle.ID,
		VIN:       vehicle.VIN,
		Brand:     vehicle.Brand,
		Histories: historiesData,
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "Success",
		Data:    result,
	}
}

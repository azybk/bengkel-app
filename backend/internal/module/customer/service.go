package customer

import (
	"bengkel_app/domain"
	"context"
	"time"
)

type service struct {
	customerRepository domain.CustomerRepository
}

func NewService(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &service{customerRepository}
}

func (s service) All(ctx context.Context) domain.ApiResponse {
	customers, err := s.customerRepository.FindAll(ctx)
	if err != nil {
		return domain.ApiResponse{
			Code: "500",
			Message: "Internal Server Error",

		}
	}

	var customersData []domain.CustomerData

	for _, v := range customers {
		customersData = append(customersData, domain.CustomerData{
			ID: v.ID,
			Name: v.Name,
			Phone: v.Phone,
		})
	}

	return domain.ApiResponse{
		Code: "200",
		Message: "Success",
		Data: customersData,
	}
}

func (s service) Save(ctx context.Context, customerData domain.CustomerData) domain.ApiResponse {
	customer := domain.Customer {
		Name: customerData.Name,
		Phone: customerData.Phone,
		CreatedAt: time.Now(),
	}

	err := s.customerRepository.Insert(ctx, &customer)
	if err != nil {
		return domain.ApiResponse{
			Code: "500",
			Message: "Internal Server Error",
		}
	}

	return domain.ApiResponse{
		Code: "200",
		Message: "Success",
	}
}
package customer

import (
	"context"
	"time"

	"red21.id/learn/bengkel/domain"
)

type service struct {
	customerRepository domain.CustomerRepository
}

func NewService(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &service{
		customerRepository,
	}
}

func (s service) All(ctx context.Context) domain.ApiResponse {
	customers, err := s.customerRepository.FindAll(ctx)
	if err != nil {
		return domain.ApiResponse{
			Code:    "991",
			Message: "SYSTEM MALFUNCTION",
		}
	}

	var customerData []domain.CustomerData
	for _, v := range customers {
		customerData = append(customerData, domain.CustomerData{
			ID:    v.ID,
			Name:  v.Name,
			Phone: v.Phone,
		})
	}
	return domain.ApiResponse{
		Code:    "00",
		Message: "APPROVE",
		Data:    customerData,
	}
}

func (s service) Save(ctx context.Context, customerData domain.CustomerData) domain.ApiResponse {
	customer := domain.Customer{
		Name:      customerData.Name,
		Phone:     customerData.Phone,
		CreatedAt: time.Now(),
	}

	err := s.customerRepository.Insert(ctx, &customer)
	if err != nil {
		return domain.ApiResponse{
			Code:    "911",
			Message: "SYSTEM MALFUNCTION",
		}
	}

	return domain.ApiResponse{
		Code:    "00",
		Message: "APPROVE",
	}
}

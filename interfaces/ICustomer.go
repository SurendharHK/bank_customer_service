package interfaces

import "github.com/SurendharHK/bank_customer_service/models"

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.DBResponse, error)
}

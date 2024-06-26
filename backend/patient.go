package patients

import "context"

// Patient is domain object
type Patient struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order Order  `json:"order"`
}

// Order is domain object
type Order struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type PatientRepository interface {
	GetAllPatients(ctx context.Context) (*[]Patient, error)
	GetOrderByID(ctx context.Context, id int32) (*Order, error)
	UpdateOrder(ctx context.Context, order *Order) error
}

package ports

import "payment-system-three/internal/models"

type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	FindAdminByEmail(email string) (*models.Admin, error)
	TokenInBlacklist(token *string) bool
	CreateUser(user *models.User) error
	CreateAdmin(admin *models.Admin) error
	UpdateUser(user *models.User) error
	UpdateAdmin(user *models.Admin) error
}

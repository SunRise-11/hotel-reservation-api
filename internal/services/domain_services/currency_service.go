package domain_services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type CurrencyService struct {
	Repository *repositories.CurrencyRepository
}

// NewCurrencyService returns new CurrencyService
func NewCurrencyService(r *repositories.CurrencyRepository) *CurrencyService {
	return &CurrencyService{Repository: r}
}

// Create creates new currency.
func (s *CurrencyService) Create(currency *models.Currency, tenantID uint64) (*models.Currency, error) {

	return s.Repository.Create(currency, tenantID)
}

// Update updates currency.
func (s *CurrencyService) Update(currency *models.Currency, tenantID uint64) (*models.Currency, error) {

	return s.Repository.Update(currency, tenantID)
}

// Find returns currency and if it does not find the currency, it returns nil.
func (s *CurrencyService) Find(id uint64, tenantID uint64) (*models.Currency, error) {

	return s.Repository.Find(id, tenantID)
}

// FindAll returns paginates list of currencies
func (s *CurrencyService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)
}

// FindBySymbol returns currency by alias name.
func (s *CurrencyService) FindBySymbol(symbol string, tenantID uint64) (*models.Currency, error) {

	return s.Repository.FindBySymbol(symbol, tenantID)
}

// Seed
func (s *CurrencyService) Seed(jsonFilePath string, tenantID uint64) error {

	return s.Repository.Seed(jsonFilePath, tenantID)
}

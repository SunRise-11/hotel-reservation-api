package services

import (
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type GuestService struct {
	Repository *repositories.GuestRepository
}

// NewGuestService returns new GuestService
func NewGuestService() *GuestService {
	return &GuestService{}
}

// Create creates new Guest.
func (s *GuestService) Create(guest *models.Guest) (*models.Guest, error) {

	return s.Repository.Create(guest)
}

// Update updates Guest.
func (s *GuestService) Update(guest *models.Guest) (*models.Guest, error) {

	return s.Repository.Update(guest)
}

// Find returns Guest and if it does not find the Guest, it returns nil.
func (s *GuestService) Find(id uint64) (*models.Guest, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of cities.
func (s *GuestService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// ReservationsCount returns guest reserves count
func (s *GuestService) ReservationsCount(guestId uint64) (error, uint64) {

	return s.Repository.ReservationsCount(guestId)
}

// CheckIn checkin guest
func (s *GuestService) CheckIn(guestId uint64) error {

	return s.Repository.CheckIn(guestId)
}

// CheckOut checkout guest
func (s *GuestService) CheckOut(guestId uint64) error {

	return s.Repository.CheckOut(guestId)
}
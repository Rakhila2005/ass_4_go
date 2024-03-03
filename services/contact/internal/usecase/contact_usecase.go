package useCase

import "day2/services/contact/internal/domain"

// ContactUseCase
type ContactUseCase struct {
}

// NewContactUseCase  ContactUseCase
func NewContactUseCase() *ContactUseCase {
    return &ContactUseCase{}
}

// Create Contact
func (uc *ContactUseCase) CreateContact(fullName, phoneNumber string) (*domain.Contact, error) {
}

// Update Contact
func (uc *ContactUseCase) UpdateContact(contactID int, fullName, phoneNumber string) (*domain.Contact, error) {
}

// Delete Contact
func (uc *ContactUseCase) DeleteContact(contactID int) error {
}

// Get Contact by ID
func (uc *ContactUseCase) GetContact(contactID int) (*domain.Contact, error) {
}

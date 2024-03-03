package useCase

import "day2/services/contact/internal/domain"

// GroupUseCase use case 
type GroupUseCase struct {
}

// NewGroupUseCase GroupUseCase
func NewGroupUseCase() *GroupUseCase {
    return &GroupUseCase{}
}

// Create Group
func (uc *GroupUseCase) CreateGroup(name string) (*domain.Group, error) {
}

// Read Group
func (uc *GroupUseCase) ReadGroup(groupID int) (*domain.Group, error) {
}

// Add Contact To Group 
func (uc *GroupUseCase) AddContactToGroup(contactID, groupID int) error {
}

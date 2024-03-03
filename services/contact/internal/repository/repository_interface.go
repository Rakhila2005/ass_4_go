package repository

import "day2/services/contact/internal/domain"

type ContactRepository interface {
    Create(contact *domain.Contact) (*domain.Contact, error)
    Update(contact *domain.Contact) (*domain.Contact, error)
    Delete(contactID int) error
    Get(contactID int) (*domain.Contact, error)
}

type GroupRepository interface {
    Create(group *domain.Group) (*domain.Group, error)
    Read(groupID int) (*domain.Group, error)
    AddContactToGroup(contactID, groupID int) error
}

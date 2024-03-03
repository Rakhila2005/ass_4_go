package usecase

import "day2/services/contact/internal/domain"

type ContactUseCase interface {
    CreateContact(fullName, phoneNumber string) (*domain.Contact, error)
    UpdateContact(contactID int, fullName, phoneNumber string) (*domain.Contact, error)
    DeleteContact(contactID int) error
    GetContact(contactID int) (*domain.Contact, error)
}

type GroupUseCase interface {
    CreateGroup(name string) (*domain.Group, error)
    ReadGroup(groupID int) (*domain.Group, error)
    AddContactToGroup(contactID, groupID int) error
}

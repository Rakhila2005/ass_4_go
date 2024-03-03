package internal

import (
    "day2/services/internal/repository"
    "day2/services/internal/useCase"
    "day2/services/internal/delivery"
)

func NewRepository() repository.ContactRepository {
    return nil 
}

func NewUseCase(repo repository.ContactRepository) useCase.ContactUseCase {
    return nil 
}
func NewDelivery(useCase useCase.ContactUseCase) delivery.ContactDelivery {
    return nil 
}

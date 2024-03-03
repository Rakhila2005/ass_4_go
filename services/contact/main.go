package main

import (
    "day2/services/internal"
    "day2/services/internal/repository"
    "day2/services/internal/usecase"
    "day2/services/internal/delivery"
    "net/http"
)

func main() {

    repo := internal.NewRepository()

    uc := internal.NewUseCase(repo)

    delivery := internal.NewDelivery(uc)

    http.HandleFunc("/create_contact", delivery.CreateContactHandler)
    http.HandleFunc("/update_contact", delivery.UpdateContactHandler)
    http.HandleFunc("/delete_contact", delivery.DeleteContactHandler)
    http.HandleFunc("/get_contact", delivery.GetContactHandler)
    http.HandleFunc("/create_group", delivery.CreateGroupHandler)
    http.HandleFunc("/read_group", delivery.ReadGroupHandler)
    http.HandleFunc("/add_contact_to_group", delivery.AddContactToGroupHandler)

    http.ListenAndServe(":8080", nil)
}

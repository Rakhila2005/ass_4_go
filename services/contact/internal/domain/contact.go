package domain
type Contact struct {
    ID           int
    FullName     string
    PhoneNumber  string
}

func NewContact(id int, fullName, phoneNumber string) *Contact {
    return &Contact{
        ID:           id,
        FullName:     fullName,
        PhoneNumber:  phoneNumber,
    }
}

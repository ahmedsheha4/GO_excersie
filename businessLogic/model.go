package businesslogic

//PhoneNumber model
type PhoneNumber struct {
	Country string `json:"country"`
	IsValid bool   `json:"isValid"`
	Code    string `json:"code"`
	Number  string `json:"number"`
}

//Response (HTTP response) struct
type Response struct {
	Phones    []PhoneNumber
	TotalRows int
}

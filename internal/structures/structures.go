package structures

import "time"

type GetUsersStruct struct {
	Users []struct {
		DisabledAt any    `json:"disabled_at"`
		Email      string `json:"email"`
		ID         string `json:"id"`
		InsertedAt         time.Time `json:"inserted_at"`
		LastSignedInAt     any    `json:"last_signed_in_at"`
		LastSignedInMethod any    `json:"last_signed_in_method"`
		Role               string `json:"role"`
		UpdatedAt          time.Time `json:"updated_at"`
	} `json:"data"`
}
type CreateUserBody struct {
	User struct {
		Email                 string `json:"email"`
		Role                  string `json:"role"`
		Password              string `json:"password"`
		Password_confirmation string `json:"password_confirmation"`
	} `json:"user"`
	
}
type SoloUserStruct struct {
	User struct {
		DisabledAt         any       `json:"disabled_at"`
		Email              string    `json:"email"`
		ID                 string    `json:"id"`
		InsertedAt         time.Time `json:"inserted_at"`
		LastSignedInAt     any       `json:"last_signed_in_at"`
		LastSignedInMethod any       `json:"last_signed_in_method"`
		Role               string    `json:"role"`
		UpdatedAt          time.Time `json:"updated_at"`
	} `json:"data"`
}
type DelUserStruct struct {
	Error string `json:"error"`
}
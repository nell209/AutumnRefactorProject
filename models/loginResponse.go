package models

type LoginResponse struct {
	Response
	AccessToken string `json:"accessToken"`
}

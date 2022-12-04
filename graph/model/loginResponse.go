package model

type LoginResponse struct {
	Response
	AccessToken string `json:"accessToken"`
}

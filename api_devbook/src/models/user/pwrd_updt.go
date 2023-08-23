package models_user

type PswrdUpdate struct {
	New     string `json:"new,omitempty"`
	Current string `json:"current,omitempty"`
}

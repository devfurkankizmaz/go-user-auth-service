package model

import "context"

type Info struct {
	Name string `json:"name"`
}

type InfoService interface {
	GetInfoByID(c context.Context, userID string) (*Info, error)
}

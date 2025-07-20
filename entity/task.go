package entity

import "time"

type Task struct {
	Id        int       `json:"id"`
	Desc      string    `json:"desc"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

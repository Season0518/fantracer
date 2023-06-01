package models

type Group struct {
    ID         int64    `json:"id" xorm:"unique"`
    Name       string `json:"name"`
    Permission string `json:"permission"`
}
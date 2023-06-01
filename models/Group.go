package models

type Group struct {
    ID         int    `json:"id" xorm:"unique"`
    Name       string `json:"name"`
    Permission string `json:"permission"`
}
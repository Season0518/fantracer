package models

type Group struct {
    ID         int    `json:"id"`
    Name       string `json:"name"`
    Permission string `json:"permission"`
}
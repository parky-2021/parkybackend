package models

import "github.com/zebresel-com/mongodm"

type User struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`
	UserName             string `json:"username" bson:"username"`
}

package entity

import "time"

type User struct {
	ID   int    `json:"id" bson:"id" xml:"id" yaml:"id"`
	Code string `json:"code" bson:"code" xml:"code" yaml:"code"`
	Name string `json:"name" bson:"name" xml:"name" yaml:"name"`

	DateCreated *time.Time `json:"date_created" bson:"date_created" xml:"dateCreated" yaml:"dateCreated"`
	DateUpdated *time.Time `json:"date_updated" bson:"date_updated" xml:"dateUpdated" yaml:"dateUpdated"`
	DateDeleted *time.Time `json:"date_deleted" bson:"date_deleted" xml:"dateDeleted" yaml:"dateDeleted"`
}

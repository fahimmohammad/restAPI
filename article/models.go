package article

import "time"

// Article export
type Article struct {
	ID          string    `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`
	Count       int       `json:"count" bson:"count"`
	CreatedAt   time.Time `json:"createat" bson:"createat"`
	Rating      float64   `json:"rating" bson:"rating"`
	IsPublished bool      `json:"ispublished" bson:"ispublished"`
}

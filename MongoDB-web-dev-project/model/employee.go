package model

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Employee struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}

func (e *Employee) BeforeCreate() error {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return nil
}

func (e *Employee) BeforeUpdate() error {
	e.UpdatedAt = time.Now()
	return nil
} 


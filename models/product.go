package models

import (
	"concurrency-experimentals/utils"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Quantity     int       `json:"quantity"`
	Price        float64   `json:"price"`
	Available    int8      `json:"available"`
	PhotoURL     string    `json:"photo_url"`
	Ratings      int       `json:"ratings"`
	Category     string    `json:"category"`
	Manufacturer string    `json:"manufacturer"`
	IsDeleted    int8      `json:"is_deleted"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

func (p *Product) GenerateID() string {
	p.ID = bson.NewObjectId().Hex()
	return p.ID
}

func (p *Product) GenerateQuantity(r *rand.Rand) int {
	p.Quantity = r.Intn(1000)
	p.SetAvailable(false)
	if p.Quantity > 0 {
		p.SetAvailable(true)
	}
	return p.Quantity
}

func (p *Product) GenerateCreatedAt(r *rand.Rand) time.Time {
	p.CreatedAt = time.Now().AddDate(-r.Intn(5), -r.Intn(12), -r.Intn(28))
	return p.CreatedAt
}

func (p *Product) GenerateUpdatedAt(r *rand.Rand) time.Time {
	p.UpdatedAt = p.CreatedAt.AddDate(0, r.Intn(12), r.Intn(28))
	return p.UpdatedAt
}

func (p *Product) GenerateDeletedAt(r *rand.Rand) time.Time {
	p.SetIsDeleted(false)
	p.DeletedAt = time.Time{}
	if r.Int()%1000 == 0 {
		p.SetIsDeleted(true)
		p.DeletedAt = p.UpdatedAt
	}
	return p.DeletedAt
}

func (p *Product) SetAvailable(available bool) {
	p.Available = utils.BoolToInt(available)
}

func (p *Product) SetIsDeleted(isDeleted bool) {
	p.IsDeleted = utils.BoolToInt(isDeleted)
}

func (p *Product) ToMongoProductModel() *MongoProductModel {
	return &MongoProductModel{
		ID:          bson.ObjectIdHex(p.ID),
		Name:        p.Name,
		Description: p.Description,
		Quantity:    p.Quantity,
		Price:       p.Price,
		Available:   utils.IntToBool(p.Available),
		PhotoURL:    p.PhotoURL,
		Ratings:     p.Ratings,
		Category:    p.Category,
		IsDeleted:   utils.IntToBool(p.IsDeleted),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   p.DeletedAt,
	}
}

type MongoProductModel struct {
	ID           bson.ObjectId `bson:"id"`
	Name         string        `bson:"name"`
	Description  string        `bson:"description"`
	Quantity     int           `bson:"quantity"`
	Price        float64       `bson:"price"`
	Available    bool          `bson:"available"`
	PhotoURL     string        `bson:"photo_url"`
	Ratings      int           `bson:"rating"`
	Category     string        `bson:"category"`
	Manufacturer string        `bson:"manufacturer"`
	IsDeleted    bool          `bson:"is_deleted"`
	CreatedAt    time.Time     `bson:"created_at"`
	UpdatedAt    time.Time     `bson:"updated_at"`
	DeletedAt    time.Time     `bson:"deleted_at"`
}

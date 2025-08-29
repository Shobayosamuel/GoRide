package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	ID primitive.ObjectID
	UserID string
	PackageSlug string // for example; van, luxury, sedan
	TotalPriceInCents float64
}
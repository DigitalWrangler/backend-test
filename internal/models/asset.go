package models

type Asset struct {
    ID          string  `bson:"_id,omitempty" json:"id,omitempty"`
    Name        string  `bson:"name" json:"name"`
    Description string  `bson:"description" json:"description"`
    Value       float64 `bson:"value" json:"value"`
} 
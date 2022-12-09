package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LokasiTemuan struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Lokasi    string             `json:"lokasi,omitempty" bson:"lokasi,omitempty"`
	Kelurahan string             `json:"kelurahan,omitempty" bson:"kelurahan,omitempty"`
	Kecamatan string             `json:"kecamatan,omitempty" bson:"kecamatan,omitempty"`
	Kota      string             `json:"kota,omitempty" bson:"kota,omitempty"`
	Provinsi  string             `json:"provinsi,omitempty" bson:"provinsi,omitempty"`
}

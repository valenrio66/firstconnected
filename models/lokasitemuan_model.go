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

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}

type Villages struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Province     string             `bson:"province,omitempty"`
	District     string             `bson:"district,omitempty"`
	Sub_district string             `bson:"sub_district,omitempty"`
	Village      string             `bson:"village,omitempty"`
	Border       Geometry           `bson:"border,omitempty"`
}

type Koordinat struct {
	Longitude int64 `json:"longitude,omitempty" bson:"longitude,omitempty"`
	Latitude  int64 `json:"latitude,omitempty" bson:"latitude,omitempty"`
}

type MongoGeometry struct {
	MongoString    string
	DBName         string
	CollectionName string
	LocationField  string
}

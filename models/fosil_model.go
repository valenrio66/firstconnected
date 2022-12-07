package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Fosil struct {
	Id             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	NoRegister     string             `json:"no_register" bson:"No Register" validate:"required"`
	NoInventaris   string             `json:"no_inventaris" bson:"No Inventaris" validate:"required"`
	NamaKoleksi    string             `json:"nama_koleksi" bson:"Nama Koleksi" validate:"required"`
	LokasiTemuan   string             `json:"lokasi_temuan" bson:"Lokasi Temuan" validate:"required"`
	TahunPerolehan string             `json:"tahun_perolehan" bson:"Tahun Perolehan" validate:"required"`
	Determinator   string             `json:"determinator,omitempty" validate:"required"`
	Keterangan     string             `json:"keterangan,omitempty" validate:"required"`
}

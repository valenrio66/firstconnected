package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SumberDayaGeologi struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	NoRegister       string             `json:"no_register" bson:"no_reg" validate:"required"`
	NoInventaris     string             `json:"no_inventaris" bson:"no_invent" validate:"required"`
	KodeBmn          string             `json:"kode_bmn" bson:"kode_bmn" validate:"required"`
	NupBmn           string             `json:"nup_bmn" bson:"nup_bmn" validate:"required"`
	MerkBmn          string             `json:"merk_bmn" bson:"merk_bmn" validate:"required"`
	KelompokKoleksi  string             `json:"kelompok_koleksi" bson:"kelompok_koleksi" validate:"required"`
	JenisKoleksi     string             `json:"jenis_koleksi" bson:"jenis_koleksi" validate:"required"`
	SubJenisKoleksi  string             `json:"sub_jenis_koleksi" bson:"sub_jenis_koleksi" validate:"required"`
	KodeJenisKoleksi string             `json:"kode_jenis_koleksi" bson:"kode_jenis_koleksi" validate:"required"`
	RuangSimpan      string             `json:"ruang_simpan" bson:"ruang_simpan" validate:"required"`
	LokasiSimpan     string             `json:"lokasi_simpan" bson:"lokasi_simpan" validate:"required"`
	Kondisi          string             `json:"kondisi" bson:"kondisi" validate:"required"`
	NamaKoleksi      string             `json:"nama_koleksi" bson:"nama_koleksi" validate:"required"`
	Keterangan       string             `json:"keterangan" bson:"keterangan" validate:"required"`
	LokasiTemuan     string             `json:"lokasi_temuan" bson:"lokasi_temuan" validate:"required"`
	Pulau            string             `json:"pulau" bson:"pulau" validate:"required"`
	CaraPerolehan    string             `json:"cara_perolehan" bson:"cara_perolehan" validate:"required"`
	TahunPerolehan   string             `json:"tahun_perolehan" bson:"tahun_perolehan" validate:"required"`
	Kolektor         string             `json:"kolektor" bson:"kolektor" validate:"required"`
	Kepemilikan      string             `json:"kepemilikan" bson:"kepemilikan" validate:"required"`
	Operator         string             `json:"operator" bson:"operator" validate:"required"`
	TanggalDicatat   string             `json:"tanggal_dicatat" bson:"tanggal_dicatat" validate:"required"`
	NilaiPerolehan   string             `json:"nilai_perolehan" bson:"nilai_perolehan" validate:"required"`
	NilaiBuku        string             `json:"nilai_buku" bson:"nilai_buku" validate:"required"`
}

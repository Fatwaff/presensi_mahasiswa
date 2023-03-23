package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama          string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Npm           int             	 `bson:"npm,omitempty" json:"npm,omitempty"`
	Nama_kelas    Kelas           	 `bson:"nama_kelas,omitempty" json:"nama_kelas,omitempty"`
	Jurusan       Prodi              `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
}

type Kelas struct {
	ID           	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
    Kode_kelas  	string 				`bson:"kode_kelas,omitempty" json:"kode_kelas,omitempty"`
    Nama_kelas  	string 				`bson:"nama_kelas,omitempty" json:"nama_kelas,omitempty"`
    Kapasitas  		int 				`bson:"kapasitas,omitempty" json:"kapasitas,omitempty"`
}

type Prodi struct {
	ID           	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
    Nama_prodi    	string 				`bson:"nama_prodi,omitempty" json:"nama_prodi,omitempty"`
    Ketua_prodi     string 				`bson:"ketua_prodi,omitempty" json:"ketua_prodi,omitempty"`
    Fakultas        string 				`bson:"fakultas,omitempty" json:"fakultas,omitempty"`
    Kapasitas  		int 				`bson:"kapasitas,omitempty" json:"kapasitas,omitempty"`
}

type MataKuliah struct {
	ID           	primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_matkul  	int                `bson:"kode_matkul,omitempty" json:"kode_matkul,omitempty"`
	Nama_matkul  	string             `bson:"nama_matkul,omitempty" json:"nama_matkul,omitempty"`
	Sks          	int         	   `bson:"sks,omitempty" json:"sks,omitempty"`
	Dosen_pengajar  Dosen 			   `bson:"dosen_pengajar,omitempty" json:"dosen_pengajar,omitempty"`
	Jadwal_kuliah   JadwalKuliah       `bson:"jadwal_kuliah,omitempty" json:"jadwal_kuliah,omitempty"`
	Ruang_kuliah    RuangKuliah        `bson:"ruang_kuliah,omitempty" json:"ruang_kuliah,omitempty"`
}

type Dosen struct {
	ID           	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
    Nama_dosen      string 				`bson:"nama_dosen,omitempty" json:"nama_dosen,omitempty"`
    Nik          	string 				`bson:"nik,omitempty" json:"nik,omitempty"`
    Bidang_studi  	string 				`bson:"bidang_studi,omitempty" json:"bidang_studi,omitempty"`
}

type JadwalKuliah struct {
	Jam_masuk     string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar    string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Hari          string   `bson:"hari,omitempty" json:"hari,omitempty"`
}

type RuangKuliah struct {
	ID           primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
    Nama_ruang   string 				`bson:"nama_ruang,omitempty" json:"nama_ruang,omitempty"`
    Kapasitas    int 					`bson:"kapasitas,omitempty" json:"kapasitas,omitempty"`
    Gedung       string 				`bson:"gedung,omitempty" json:"gedung,omitempty"`
    Lantai       int 					`bson:"lantai,omitempty" json:"lantai,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Kehadiran    string           	`bson:"kehadiran,omitempty" json:"kehadiran,omitempty"`
	Biodata      Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Mata_kuliah  MataKuliah         `bson:"mata_kuliah,omitempty" json:"mata_kuliah,omitempty"`
}
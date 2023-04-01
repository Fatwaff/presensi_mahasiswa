package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama          string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Npm           int             	 `bson:"npm,omitempty" json:"npm,omitempty"`
	Nama_kelas    Kelas           	 `bson:"namakelas,omitempty" json:"namakelas,omitempty"`
	Jurusan       Prodi              `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
}

type Kelas struct {
	ID           	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
    Kode_kelas  	string 				`bson:"kodekelas,omitempty" json:"kodekelas,omitempty"`
    Nama_kelas  	string 				`bson:"namakelas,omitempty" json:"namakelas,omitempty"`
    Kapasitas  		int 				`bson:"kapasitas,omitempty" json:"kapasitas,omitempty"`
}

type Prodi struct {
	ID           	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
    Nama_prodi    	string 				`bson:"namaprodi,omitempty" json:"namaprodi,omitempty"`
    Ketua_prodi     string 				`bson:"ketuaprodi,omitempty" json:"ketuaprodi,omitempty"`
    Fakultas        string 				`bson:"fakultas,omitempty" json:"fakultas,omitempty"`
    Kapasitas  		int 				`bson:"kapasitas,omitempty" json:"kapasitas,omitempty"`
}

type MataKuliah struct {
	ID           	primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_matkul  	int                `bson:"kodematkul,omitempty" json:"kodematkul,omitempty"`
	Nama_matkul  	string             `bson:"namamatkul,omitempty" json:"namamatkul,omitempty"`
	Sks          	int         	   `bson:"sks,omitempty" json:"sks,omitempty"`
	Dosen_pengajar  Dosen 			   `bson:"dosenpengajar,omitempty" json:"dosenpengajar,omitempty"`
	Jadwal_kuliah   JadwalKuliah       `bson:"jadwalkuliah,omitempty" json:"jadwalkuliah,omitempty"`
	Ruang_kuliah    RuangKuliah        `bson:"ruangkuliah,omitempty" json:"ruangkuliah,omitempty"`
}

type Dosen struct {
	ID           	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
    Nama_dosen      string 				`bson:"namadosen,omitempty" json:"namadosen,omitempty"`
    Nik          	string 				`bson:"nik,omitempty" json:"nik,omitempty"`
    Bidang_studi  	string 				`bson:"bidangstudi,omitempty" json:"bidangstudi,omitempty"`
}

type JadwalKuliah struct {
	Jam_masuk     string   `bson:"jammasuk,omitempty" json:"jammasuk,omitempty"`
	Jam_keluar    string   `bson:"jamkeluar,omitempty" json:"jamkeluar,omitempty"`
	Hari          string   `bson:"hari,omitempty" json:"hari,omitempty"`
}

type RuangKuliah struct {
	ID           primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
    Nama_ruang   string 				`bson:"namaruang,omitempty" json:"namaruang,omitempty"`
    Kapasitas    int 					`bson:"kapasitas,omitempty" json:"kapasitas,omitempty"`
    Gedung       string 				`bson:"gedung,omitempty" json:"gedung,omitempty"`
    Lantai       int 					`bson:"lantai,omitempty" json:"lantai,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Kehadiran    string           	`bson:"kehadiran,omitempty" json:"kehadiran,omitempty"`
	Biodata      Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Mata_kuliah  MataKuliah         `bson:"matakuliah,omitempty" json:"matakuliah,omitempty"`
}
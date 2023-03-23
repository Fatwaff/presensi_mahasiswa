package Fatwa_1214038

import (
	"fmt"
	"testing"

	"github.com/Fatwaff/presensi_mahasiswa/model"
	"github.com/Fatwaff/presensi_mahasiswa/module"
)

func TestInsertMahasiswa(t *testing.T) {
	nama := "Dirga Febrian"
	npm := 1214039
	kelas := model.Kelas{
		Nama_kelas :  "Dirga Febrian",
	}
	jurusan := model.Prodi{
		Nama_prodi :   "D4 Teknik Informatika",
		Fakultas : "Sekolah Vokasi",
	}
	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, npm, kelas, jurusan)
	fmt.Println(hasil)
}
func TestInsertKelas(t *testing.T) {
	kode_kelas := "L3-14"
	nama_kelas := "Lab 14"
	kapasitas := 40
	hasil := module.InsertKelas(module.MongoConn, "kelas", kode_kelas, nama_kelas, kapasitas)
	fmt.Println(hasil)
}
func TestInsertProdi(t *testing.T) {
	nama_prodi := "D4 Teknik Informatika"
	ketua_prodi := "Roni Andarsyah"
	fakultas := "Sekolah Vokasi"
	kapasitas := 500
	hasil := module.InsertProdi(module.MongoConn, "prodi", nama_prodi, ketua_prodi, fakultas, kapasitas)
	fmt.Println(hasil)
}
func TestInserMatkul(t *testing.T) {
	kode_matkul := 21711
	nama_matkul := "Pemrograman 3"
	sks := 3
	dosen_pengajar := model.Dosen{
		Nama_dosen :  "Indra Riksa Herlambang",
	}
	jadwal_kuliah := model.JadwalKuliah{
		Jam_masuk :   "13:30",
		Jam_keluar :  "17:00",
		Hari : "Jum'at",
	}
	ruang_kuliah := model.RuangKuliah{
		Nama_ruang :   "Lab 314",
	}
	hasil := module.InsertMatkul(module.MongoConn, "matkul", kode_matkul, nama_matkul, sks, dosen_pengajar, jadwal_kuliah, ruang_kuliah)
	fmt.Println(hasil)
}
func TestInsertDosen(t *testing.T) {
	nama_dosen := "Indra Riksa Herlambang"
	nik := "123.45.678"
	bidang_studi := "Teknik Informatika"
	hasil := module.InsertDosen(module.MongoConn, "dosen", nama_dosen, nik, bidang_studi)
	fmt.Println(hasil)
}
func TestInsertRuang(t *testing.T) {
	nama_ruang := "Lab 314"
	kapasitas := 40
	gedung := "Pendidikan"
	lantai := 3
	hasil := module.InsertRuang(module.MongoConn, "ruang", nama_ruang, kapasitas, gedung, lantai)
	fmt.Println(hasil)
}
func TestInsertPresensi(t *testing.T) {
	kehadiran := "masuk"
	biodata := model.Mahasiswa{
		Nama :  "Dirga Febrian",
		Npm : 1214038,
		Nama_kelas :  model.Kelas{
			Nama_kelas : "2B",
		},
		Jurusan : model.Prodi{
			Nama_prodi : "D4 Teknik Informatika",
		},
	}
	matkul := model.MataKuliah{
		Kode_matkul :   21711,
		Nama_matkul :   "Pemrograman 3",
		Sks : 3,
		Dosen_pengajar :   model.Dosen{
			Nama_dosen: "Indra Riksa Herlambang",
		},
		Jadwal_kuliah :   model.JadwalKuliah{
			Jam_masuk : "13:30",
			Jam_keluar : "17:00",
			Hari : "Jum'at",
		},
		Ruang_kuliah :  model.RuangKuliah{
			Nama_ruang : "Lab 314",
		},
	}
	hasil := module.InsertPresensi(module.MongoConn, "presensi", kehadiran, biodata, matkul)
	fmt.Println(hasil)
}

func TestGetMahasiwaFromNpm(t *testing.T) {
	npm := 1214039
	biodata := module.GetMahasiswaFromNpm(npm, module.MongoConn, "mahasiswa")
	fmt.Println(biodata)
}
func TestGetKelasFromKodeKelas(t *testing.T) {
	kode_kelas := "L3-14"
	data := module.GetKelasFromKodeKelas(kode_kelas, module.MongoConn, "kelas")
	fmt.Println(data)
}
func TestGetProdiFromNamaProdi(t *testing.T) {
	nama_prodi := "D4 Teknik Informatika"
	data := module.GetProdiFromNamaProdi(nama_prodi, module.MongoConn, "prodi")
	fmt.Println(data)
}
func TestGetMatkulFromKodeMatkul(t *testing.T) {
	kode_matkul := 21711
	data := module.GetMatkulFromKodeMatkul(kode_matkul, module.MongoConn, "matkul")
	fmt.Println(data)
}
func TestGetDosenFromNamaDosen(t *testing.T) {
	nama_dosen := "Indra Riksa Herlambang"
	data := module.GetDosenFromNamaDosen(nama_dosen, module.MongoConn, "dosen")
	fmt.Println(data)
}
func TestGetRuangFromNamaRuang(t *testing.T) {
	nama_ruang := "Lab 314"
	data := module.GetRuangFromNamaRuang(nama_ruang, module.MongoConn, "ruang")
	fmt.Println(data)
}
func TestGetPresensiFromNamaMahasiswa(t *testing.T) {
	nama_mahasiswa := "Dirga Febrian"
	data := module.GetPresensiFromNamaMahasiswa(nama_mahasiswa, module.MongoConn, "presensi")
	fmt.Println(data)
}
func TestGetAllPresensiFromKehadiran(t *testing.T) {
	kehadiran := "masuk"
	data := module.GetAllPresensiFromKehadiran(kehadiran, module.MongoConn, "presensi")
	fmt.Println(data)
}
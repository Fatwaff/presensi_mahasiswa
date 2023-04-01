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
		Nama_kelas :  "2B",
	}
	jurusan := model.Prodi{
		Nama_prodi :   "D4 Teknik Informatika",
		Fakultas : "Sekolah Vokasi",
	}
	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, npm, kelas, jurusan)
	fmt.Println(hasil)
}
func TestInsertKelas(t *testing.T) {
	kodekelas := "L3-14"
	namakelas := "Lab 14"
	kapasitas := 40
	hasil := module.InsertKelas(module.MongoConn, "kelas", kodekelas, namakelas, kapasitas)
	fmt.Println(hasil)
}
func TestInsertProdi(t *testing.T) {
	namaprodi := "D4 Teknik Informatika"
	ketuaprodi := "Roni Andarsyah"
	fakultas := "Sekolah Vokasi"
	kapasitas := 500
	hasil := module.InsertProdi(module.MongoConn, "prodi", namaprodi, ketuaprodi, fakultas, kapasitas)
	fmt.Println(hasil)
}
func TestInserMatkul(t *testing.T) {
	kodematkul := 21711
	namamatkul := "Pemrograman 3"
	sks := 3
	dosenpengajar := model.Dosen{
		Nama_dosen :  "Indra Riksa Herlambang",
	}
	jadwalkuliah := model.JadwalKuliah{
		Jam_masuk :   "13:30",
		Jam_keluar :  "17:00",
		Hari : "Jum'at",
	}
	ruangkuliah := model.RuangKuliah{
		Nama_ruang :   "Lab 314",
	}
	hasil := module.InsertMatkul(module.MongoConn, "matkul", kodematkul, namamatkul, sks, dosenpengajar, jadwalkuliah, ruangkuliah)
	fmt.Println(hasil)
}
func TestInsertDosen(t *testing.T) {
	namadosen := "Indra Riksa Herlambang"
	nik := "123.45.678"
	bidangstudi := "Teknik Informatika"
	hasil := module.InsertDosen(module.MongoConn, "dosen", namadosen, nik, bidangstudi)
	fmt.Println(hasil)
}
func TestInsertRuang(t *testing.T) {
	namaruang := "Lab 314"
	kapasitas := 40
	gedung := "Pendidikan"
	lantai := 3
	hasil := module.InsertRuang(module.MongoConn, "ruang", namaruang, kapasitas, gedung, lantai)
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
	kodekelas := "L3-14"
	data := module.GetKelasFromKodeKelas(kodekelas, module.MongoConn, "kelas")
	fmt.Println(data)
}
func TestGetProdiFromNamaProdi(t *testing.T) {
	namaprodi := "D4 Teknik Informatika"
	data := module.GetProdiFromNamaProdi(namaprodi, module.MongoConn, "prodi")
	fmt.Println(data)
}
func TestGetMatkulFromKodeMatkul(t *testing.T) {
	kodematkul := 21711
	data := module.GetMatkulFromKodeMatkul(kodematkul, module.MongoConn, "matkul")
	fmt.Println(data)
}
func TestGetDosenFromNamaDosen(t *testing.T) {
	namadosen := "Indra Riksa Herlambang"
	data := module.GetDosenFromNamaDosen(namadosen, module.MongoConn, "dosen")
	fmt.Println(data)
}
func TestGetRuangFromNamaRuang(t *testing.T) {
	namaruang := "Lab 314"
	data := module.GetRuangFromNamaRuang(namaruang, module.MongoConn, "ruang")
	fmt.Println(data)
}
func TestGetPresensiFromNamaMahasiswa(t *testing.T) {
	namamahasiswa := "Dirga Febrian"
	data := module.GetPresensiFromNamaMahasiswa(namamahasiswa, module.MongoConn, "presensi")
	fmt.Println(data)
}
func TestGetAllPresensiFromKehadiran(t *testing.T) {
	kehadiran := "masuk"
	data := module.GetAllPresensiFromKehadiran(kehadiran, module.MongoConn, "presensi")
	fmt.Println(data)
}
func TestGetAllPresensi(t *testing.T) {
	data := module.GetAllPresensi(module.MongoConn, "presensi")
	fmt.Println(data)
}
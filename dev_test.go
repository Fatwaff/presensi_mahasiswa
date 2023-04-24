package Fatwa_1214038

import (
	"fmt"
	"testing"

	"github.com/Fatwaff/presensi_mahasiswa/model"
	"github.com/Fatwaff/presensi_mahasiswa/module"
)

func TestInsertMahasiswa(t *testing.T) {
	nama := "test"
	npm := 12140301
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
	kodekelas := "TI-B2"
	namakelas := "2B"
	kapasitas := 40
	hasil := module.InsertKelas(module.MongoConn, "kelas", kodekelas, namakelas, kapasitas)
	fmt.Println(hasil)
}
func TestInsertProdi(t *testing.T) {
	namaprodi := "Manajemen Bisnis"
	ketuaprodi := "Peter Parker"
	fakultas := "Vokasi"
	kapasitas := 1500
	hasil := module.InsertProdi(module.MongoConn, "prodi", namaprodi, ketuaprodi, fakultas, kapasitas)
	fmt.Println(hasil)
}
func TestInserMatkul(t *testing.T) {
	kodematkul := 21682
	namamatkul := "Kewirausahaan"
	sks := 3
	dosenpengajar := model.Dosen{
		Nama_dosen :  "Bheben oscar,. S.MB.MM",
	}
	jadwalkuliah := model.JadwalKuliah{
		Jam_masuk :   "10:20",
		Jam_keluar :  "12:00",
		Hari : "Selasa",
	}
	ruangkuliah := model.RuangKuliah{
		Nama_ruang :   "Ruang 102",
	}
	hasil := module.InsertMatkul(module.MongoConn, "matkul", kodematkul, namamatkul, sks, dosenpengajar, jadwalkuliah, ruangkuliah)
	fmt.Println(hasil)
}
func TestInsertDosen(t *testing.T) {
	namadosen := "test"
	nik := "111.22.333"
	bidangstudi := "Teknik Informatika"
	hasil := module.InsertDosen(module.MongoConn, "dosen", namadosen, nik, bidangstudi)
	fmt.Println(hasil)
}
func TestInsertRuang(t *testing.T) {
	namaruang := "Lab 312"
	kapasitas := 40
	gedung := "Pendidikan"
	lantai := 3
	hasil := module.InsertRuang(module.MongoConn, "ruang", namaruang, kapasitas, gedung, lantai)
	fmt.Println(hasil)
}
func TestInsertPresensi(t *testing.T) {
	kehadiran := "masuk"
	biodata := model.Mahasiswa{
		Nama :  "test",
		Npm : 12140301,
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
	namamahasiswa := "test"
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
func TestGetAllMahasiswa(t *testing.T) {
	data := module.GetAllMahasiswa(module.MongoConn, "mahasiswa")
	fmt.Println(data)
}
func TestGetAllKelas(t *testing.T) {
	data := module.GetAllKelas(module.MongoConn, "kelas")
	fmt.Println(data)
}
func TestGetAllProdi(t *testing.T) {
	data := module.GetAllProdi(module.MongoConn, "prodi")
	fmt.Println(data)
}
func TestGetAllMataKuliah(t *testing.T) {
	data := module.GetAllMataKuliah(module.MongoConn, "matkul")
	fmt.Println(data)
}
func TestGetAllDosen(t *testing.T) {
	data := module.GetAllDosen(module.MongoConn, "dosen")
	fmt.Println(data)
}
func TestGetAllRuangKuliah(t *testing.T) {
	data := module.GetAllRuangKuliah(module.MongoConn, "ruang")
	fmt.Println(data)
}
func TestGetAllDataRuangKuliah(t *testing.T) {
	var ruang []model.RuangKuliah
	module.GetAllData(module.MongoConn, "ruang", ruang)
}
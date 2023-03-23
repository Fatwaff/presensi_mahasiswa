package module

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Fatwaff/presensi_mahasiswa/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "presensiMahasiswa",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertMahasiswa(db *mongo.Database, col string, nama string, npm int, kelas  model.Kelas, jurusan model.Prodi) (InsertedID interface{}) {
	var mahasiswa model.Mahasiswa
	mahasiswa.Nama = nama
	mahasiswa.Npm = npm
	mahasiswa.Nama_kelas = kelas
	mahasiswa.Jurusan = jurusan
	return InsertOneDoc(db, col, mahasiswa)
}

func InsertKelas(db *mongo.Database, col string, kode_kelas string, nama_kelas string, kapasitas int) (InsertedID interface{}) {
	var kelas model.Kelas
	kelas.Kode_kelas = kode_kelas
	kelas.Nama_kelas =  nama_kelas
	kelas.Kapasitas = kapasitas
	return InsertOneDoc(db, col, kelas)
}

func InsertProdi(db *mongo.Database, col string, nama_prodi string, ketua_prodi string, fakultas string, kapasitas int) (InsertedID interface{}) {
	var prodi model.Prodi
	prodi.Nama_prodi = nama_prodi
	prodi.Ketua_prodi = ketua_prodi
	prodi.Fakultas = fakultas
	prodi.Kapasitas = kapasitas
	return InsertOneDoc(db, col, prodi)
}

func InsertMatkul(db *mongo.Database, col string, kode_matkul int, nama_matkul string, sks int, dosen_pengajar model.Dosen, jadwal_kuliah model.JadwalKuliah, ruang_kuliah model.RuangKuliah) (InsertedID interface{}) {
	var matkul model.MataKuliah
	matkul.Kode_matkul = kode_matkul
	matkul.Nama_matkul = nama_matkul
	matkul.Sks = sks
	matkul.Dosen_pengajar = dosen_pengajar
	matkul.Jadwal_kuliah = jadwal_kuliah
	matkul.Ruang_kuliah = ruang_kuliah
	return InsertOneDoc(db, col, matkul)
}

func InsertDosen(db *mongo.Database, col string, nama_dosen string, nik string, bidang_studi string) (InsertedID interface{}) {
	var dosen model.Dosen
	dosen.Nama_dosen = nama_dosen
	dosen.Nik = nik
	dosen.Bidang_studi = bidang_studi
	return InsertOneDoc(db, col, dosen)
}

func InsertRuang(db *mongo.Database, col string, nama_ruang string, kapasitas int, gedung string, lantai int) (InsertedID interface{}) {
	var ruang model.RuangKuliah
	ruang.Nama_ruang = nama_ruang
	ruang.Kapasitas = kapasitas
	ruang.Gedung = gedung
	ruang.Lantai = lantai
	return InsertOneDoc(db, col, ruang)
}

func InsertPresensi(db *mongo.Database, col string, kehadiran string, biodata model.Mahasiswa, matkul model.MataKuliah) (InsertedID interface{}) {
	var presensi model.Presensi
	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Kehadiran = kehadiran
	presensi.Biodata = biodata
	presensi.Mata_kuliah = matkul
	return InsertOneDoc(db, col, presensi)
}

func GetMahasiswaFromNpm(npm int, db *mongo.Database, col string) (mhs model.Mahasiswa) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"npm": npm}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("getMahasiswaFromNpm: %v\n", err)
	}
	return mhs
}
func GetKelasFromKodeKelas(kode_kelas string, db *mongo.Database, col string) (kls model.Kelas) {
	kelas := db.Collection(col)
	filter := bson.M{"kode_kelas": kode_kelas}
	err := kelas.FindOne(context.TODO(), filter).Decode(&kls)
	if err != nil {
		fmt.Printf("GetKelasFromKodeKelas: %v\n", err)
	}
	return kls
}
func GetProdiFromNamaProdi(nama_prodi string, db *mongo.Database, col string) (prd model.Prodi) {
	prodi := db.Collection(col)
	filter := bson.M{"nama_prodi": nama_prodi}
	err := prodi.FindOne(context.TODO(), filter).Decode(&prd)
	if err != nil {
		fmt.Printf("GetProdiFromNamaProdi: %v\n", err)
	}
	return prd
}
func GetMatkulFromKodeMatkul(kode_matkul int, db *mongo.Database, col string) (mk model.MataKuliah) {
	matkul := db.Collection(col)
	filter := bson.M{"kode_matkul": kode_matkul}
	err := matkul.FindOne(context.TODO(), filter).Decode(&mk)
	if err != nil {
		fmt.Printf("GetMatkulFromKodeMatkul: %v\n", err)
	}
	return mk
}
func GetDosenFromNamaDosen(nama_dosen string, db *mongo.Database, col string) (dsn model.Dosen) {
	dosen := db.Collection(col)
	filter := bson.M{"nama_dosen": nama_dosen}
	err := dosen.FindOne(context.TODO(), filter).Decode(&dsn)
	if err != nil {
		fmt.Printf("GetDosenFromNamaDosen: %v\n", err)
	}
	return dsn
}
func GetRuangFromNamaRuang(nama_ruang string, db *mongo.Database, col string) (rgn model.RuangKuliah) {
	ruang := db.Collection(col)
	filter := bson.M{"nama_ruang": nama_ruang}
	err := ruang.FindOne(context.TODO(), filter).Decode(&rgn)
	if err != nil {
		fmt.Printf("GetRuangFromNamaRuang: %v\n", err)
	}
	return rgn
}
func GetPresensiFromNamaMahasiswa(nama_mahasiswa string, db *mongo.Database, col string) (prs model.Presensi) {
	presensi := db.Collection(col)
	filter := bson.M{"biodata.nama": nama_mahasiswa}
	err := presensi.FindOne(context.TODO(), filter).Decode(&prs)
	if err != nil {
		fmt.Printf("GetPresensiFromNamaMahasiswa: %v\n", err)
	}
	return prs
}

func GetAllPresensiFromKehadiran(kehadiran string, db *mongo.Database, col string) (aprs []model.Presensi) {
	presensi := db.Collection(col)
	filter := bson.M{"kehadiran": kehadiran}
	cursor, err := presensi.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllPresensiFromKehadiran: %v\n", err)
	}
	err = cursor.All(context.TODO(), &aprs)
	if err != nil{
		fmt.Println(err)
	}
	return 
}
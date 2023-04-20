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

func InsertKelas(db *mongo.Database, col string, kodekelas string, namakelas string, kapasitas int) (InsertedID interface{}) {
	var kelas model.Kelas
	kelas.Kode_kelas = kodekelas
	kelas.Nama_kelas =  namakelas
	kelas.Kapasitas = kapasitas
	return InsertOneDoc(db, col, kelas)
}

func InsertProdi(db *mongo.Database, col string, namaprodi string, ketuaprodi string, fakultas string, kapasitas int) (InsertedID interface{}) {
	var prodi model.Prodi
	prodi.Nama_prodi = namaprodi
	prodi.Ketua_prodi = ketuaprodi
	prodi.Fakultas = fakultas
	prodi.Kapasitas = kapasitas
	return InsertOneDoc(db, col, prodi)
}

func InsertMatkul(db *mongo.Database, col string, kodematkul int, namamatkul string, sks int, dosenpengajar model.Dosen, jadwalkuliah model.JadwalKuliah, ruangkuliah model.RuangKuliah) (InsertedID interface{}) {
	var matkul model.MataKuliah
	matkul.Kode_matkul = kodematkul
	matkul.Nama_matkul = namamatkul
	matkul.Sks = sks
	matkul.Dosen_pengajar = dosenpengajar
	matkul.Jadwal_kuliah = jadwalkuliah
	matkul.Ruang_kuliah = ruangkuliah
	return InsertOneDoc(db, col, matkul)
}

func InsertDosen(db *mongo.Database, col string, namadosen string, nik string, bidangstudi string) (InsertedID interface{}) {
	var dosen model.Dosen
	dosen.Nama_dosen = namadosen
	dosen.Nik = nik
	dosen.Bidang_studi = bidangstudi
	return InsertOneDoc(db, col, dosen)
}

func InsertRuang(db *mongo.Database, col string, namaruang string, kapasitas int, gedung string, lantai int) (InsertedID interface{}) {
	var ruang model.RuangKuliah
	ruang.Nama_ruang = namaruang
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
func GetKelasFromKodeKelas(kodekelas string, db *mongo.Database, col string) (kls model.Kelas) {
	kelas := db.Collection(col)
	filter := bson.M{"kodekelas": kodekelas}
	err := kelas.FindOne(context.TODO(), filter).Decode(&kls)
	if err != nil {
		fmt.Printf("GetKelasFromKodeKelas: %v\n", err)
	}
	return kls
}
func GetProdiFromNamaProdi(namaprodi string, db *mongo.Database, col string) (prd model.Prodi) {
	prodi := db.Collection(col)
	filter := bson.M{"namaprodi": namaprodi}
	err := prodi.FindOne(context.TODO(), filter).Decode(&prd)
	if err != nil {
		fmt.Printf("GetProdiFromNamaProdi: %v\n", err)
	}
	return prd
}
func GetMatkulFromKodeMatkul(kodematkul int, db *mongo.Database, col string) (mk model.MataKuliah) {
	matkul := db.Collection(col)
	filter := bson.M{"kodematkul": kodematkul}
	err := matkul.FindOne(context.TODO(), filter).Decode(&mk)
	if err != nil {
		fmt.Printf("GetMatkulFromKodeMatkul: %v\n", err)
	}
	return mk
}
func GetDosenFromNamaDosen(namadosen string, db *mongo.Database, col string) (dsn model.Dosen) {
	dosen := db.Collection(col)
	filter := bson.M{"namadosen": namadosen}
	err := dosen.FindOne(context.TODO(), filter).Decode(&dsn)
	if err != nil {
		fmt.Printf("GetDosenFromNamaDosen: %v\n", err)
	}
	return dsn
}
func GetRuangFromNamaRuang(namaruang string, db *mongo.Database, col string) (rgn model.RuangKuliah) {
	ruang := db.Collection(col)
	filter := bson.M{"namaruang": namaruang}
	err := ruang.FindOne(context.TODO(), filter).Decode(&rgn)
	if err != nil {
		fmt.Printf("GetRuangFromNamaRuang: %v\n", err)
	}
	return rgn
}
func GetPresensiFromNamaMahasiswa(namamahasiswa string, db *mongo.Database, col string) (prs model.Presensi) {
	presensi := db.Collection(col)
	filter := bson.M{"biodata.nama": namamahasiswa}
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

func GetAllPresensi(db *mongo.Database, col string) (data []model.Presensi) {
	dataPresensi := db.Collection(col)
	filter := bson.M{}
	cursor, err := dataPresensi.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllPresensi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllMahasiswa(db *mongo.Database, col string) (data []model.Mahasiswa) {
	dataMahasiswa := db.Collection(col)
	filter := bson.M{}
	cursor, err := dataMahasiswa.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllMahasiswa :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllKelas(db *mongo.Database, col string) (data []model.Kelas) {
	dataKelas := db.Collection(col)
	filter := bson.M{}
	cursor, err := dataKelas.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllKelas :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllProdi(db *mongo.Database, col string) (data []model.Prodi) {
	dataProdi := db.Collection(col)
	filter := bson.M{}
	cursor, err := dataProdi.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllProdi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllMataKuliah(db *mongo.Database, col string) (data []model.MataKuliah) {
	dataMataKuliah := db.Collection(col)
	filter := bson.M{}
	cursor, err := dataMataKuliah.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllMataKuliah :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllDosen(db *mongo.Database, col string) (data []model.Dosen) {
	dataDosen := db.Collection(col)
	filter := bson.M{}
	cursor, err := dataDosen.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllDosen :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllRuangKuliah(db *mongo.Database, col string) (data []model.RuangKuliah) {
	dataRuangKuliah := db.Collection(col)
	filter := bson.M{}
	cursor, err := dataRuangKuliah.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllRuangKuliah :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllRuangKuliahTest(db *mongo.Database, col string, data interface{}) interface{} {
	dataRuangKuliah := db.Collection(col)
	filter := bson.M{}
	cursor, err := dataRuangKuliah.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllRuangKuliah :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
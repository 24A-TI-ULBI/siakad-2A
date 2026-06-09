package helper

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"backend/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnect connects to MongoDB and returns a *mongo.Database.
// Mengikuti pola boilerplate gocroot — ada fallback SRV lookup kalau koneksi Atlas gagal.
func MongoConnect(mconn model.DBInfo) (db *mongo.Database, err error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
	if err != nil {
		mconn.DBString = SRVLookup(mconn.DBString)
		client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
		if err != nil {
			return
		}
	}
	db = client.Database(mconn.DBName)
	return
}

// SRVLookup converts mongodb+srv:// URI to standard mongodb:// URI via DNS SRV lookup.
func SRVLookup(srvuri string) (mongouri string) {
	atsplits := strings.Split(srvuri, "@")
	userpass := strings.Split(atsplits[0], "//")[1]
	mongouri = "mongodb://" + userpass + "@"
	slashsplits := strings.Split(atsplits[1], "/")
	domain := slashsplits[0]
	dbname := slashsplits[1]

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}
	_, srvs, err := r.LookupSRV(context.Background(), "mongodb", "tcp", domain)
	if err != nil {
		return
	}
	var srvlist string
	for _, srv := range srvs {
		srvlist += strings.TrimSuffix(srv.Target, ".") + ":" + strconv.FormatUint(uint64(srv.Port), 10) + ","
	}
	txtrecords, _ := r.LookupTXT(context.Background(), domain)
	var txtlist string
	for _, txt := range txtrecords {
		txtlist += txt
	}
	mongouri = mongouri + strings.TrimSuffix(srvlist, ",") + "/" + dbname + "?ssl=true&" + txtlist
	return
}

// GetCollection returns a MongoDB collection dari database utama.
// Semua modul wajib pakai fungsi ini.
func GetCollection(collectionName string) *mongo.Collection {
	return GetDB().Collection(collectionName)
}

// Singleton: satu koneksi MongoDB untuk seluruh lifetime server
var dbInstance *mongo.Database

// InitDB menginisialisasi koneksi MongoDB sekali saat startup.
// Wajib dipanggil di main.go sebelum server listen.
func InitDB() {
	dbName := os.Getenv("MONGODB_NAME")
	if dbName == "" {
		dbName = "kampus"
	}
	connStr := os.Getenv("MONGOSTRING")
	if connStr == "" {
		log.Fatal("MONGOSTRING env variable kosong! Isi di file .env")
	}

	mconn := model.DBInfo{
		DBString: connStr,
		DBName:   dbName,
	}

	db, err := MongoConnect(mconn)
	if err != nil {
		log.Fatalf("Gagal koneksi MongoDB saat init: %v", err)
	}

	// Ping untuk memastikan koneksi benar-benar aktif
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := db.Client().Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB Ping gagal: %v", err)
	}

	dbInstance = db
	log.Println("✅ MongoDB berhasil terkoneksi ke database:", dbName)
}

// GetDB returns the main *mongo.Database instance dari config.
func GetDB() *mongo.Database {
	if dbInstance == nil {
		InitDB()
	}
	return dbInstance
}

// GetContext returns a context with timeout for MongoDB operations
func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

// GetOneDoc fetches a single document from a collection
func GetOneDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T, err error) {
	err = db.Collection(collection).FindOne(context.TODO(), filter).Decode(&doc)
	return
}

// GetAllDoc fetches all documents from a collection
func GetAllDoc[T any](db *mongo.Database, collection string, filter bson.M) (docs []T, err error) {
	ctx := context.Background()
	cursor, err := db.Collection(collection).Find(ctx, filter)
	if err != nil {
		return
	}
	err = cursor.All(ctx, &docs)
	return
}

// InsertOneDoc inserts a single document into a collection
func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}, err error) {
	result, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return
	}
	insertedID = result.InsertedID
	return
}

// UpdateDoc updates a single document in a collection
func UpdateDoc(db *mongo.Database, collection string, filter bson.M, updatefield bson.M) (updateresult *mongo.UpdateResult, err error) {
	updateresult, err = db.Collection(collection).UpdateOne(context.TODO(), filter, updatefield)
	return
}

// DeleteDoc deletes a single document from a collection
func DeleteDoc(db *mongo.Database, collection string, filter bson.M) (deleteresult *mongo.DeleteResult, err error) {
	deleteresult, err = db.Collection(collection).DeleteOne(context.TODO(), filter)
	return
}

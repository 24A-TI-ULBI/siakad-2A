package model

type Absensi struct {
	ID         string `json:"id" bson:"_id"`
	NPM        string `json:"npm" bson:"npm"`
	MatkulCode string `json:"matkulCode" bson:"matkulCode"`
	Tanggal    string `json:"tanggal" bson:"tanggal"`
	Status     string `json:"status" bson:"status"`
	Timestamp  string `json:"timestamp" bson:"timestamp"`
}

type RekapMatkul struct {
	MatkulCode string  `json:"matkulCode"`
	Hadir      int     `json:"hadir"`
	Izin       int     `json:"izin"`
	Alpha      int     `json:"alpha"`
	TotalSesi  int     `json:"totalSesi"`
	Persentase float64 `json:"persentase"`
}

type StudentRekap struct {
	NPM                 string        `json:"npm"`
	TotalSesi           int           `json:"totalSesi"`
	Hadir               int           `json:"hadir"`
	Izin                int           `json:"izin"`
	Alpha               int           `json:"alpha"`
	PersentaseKehadiran float64       `json:"persentaseKehadiran"`
	RekapPerMatkul      []RekapMatkul `json:"rekapPerMatkul"`
}

type StudentMatkulRekap struct {
	NPM        string  `json:"npm"`
	Hadir      int     `json:"hadir"`
	Izin       int     `json:"izin"`
	Alpha      int     `json:"alpha"`
	TotalSesi  int     `json:"totalSesi"`
	Persentase float64 `json:"persentase"`
}

type CourseRekap struct {
	MatkulCode string               `json:"matkulCode"`
	Pendaftar  []StudentMatkulRekap `json:"pendaftar"`
}

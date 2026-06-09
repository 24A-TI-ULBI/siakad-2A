package model

type Matkul struct {
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	Dosen    string `json:"dosen"`
	Hari     string `json:"hari"`
	Jam      string `json:"jam"`
	Ruangan  string `json:"ruangan"`
	SKS      int    `json:"sks"`
	Semester int    `json:"semester"`
}

package data

type Student struct {
	ID      int    `json:"id"`
	Nama    string `json:"nama"`
	NIM     string `json:"nim"`
	Jurusan string `json"jurusan"`
}

var Students = []Student{
	{ID: 1, Nama: "Andi", NIM: "22/111111", Jurusan: "Informatika"},
	{ID: 2, Nama: "Budi", NIM: "22/222222", Jurusan: "Sistem Informasi"},
}
package models

type DaftarJurusan struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Jurusan string `json:"jurusan"`
}

type DaftarMahasiswa struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Nama      string `json:"nama"`
	Nim       string `json:"nim"`
	Alamat    string `json:"alamat"`
	Tanggal   string `json:"created_at"  gorm:"column:created_at"`
	IdJurusan int    `json:"id_jurusan"`
}

type MahasiswaJurusan struct {
	Nama    string `json:"nama"`
	Nim     string `json:"nim"`
	Jurusan string `json:"jurusan"`
}

func (DaftarJurusan) TableName() string {
	return "jurusan"
}

func (DaftarMahasiswa) TableName() string {
	return "mahasiswa"
}

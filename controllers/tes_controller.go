package controllers

import (
	"log"
	"net/http"
	"tesbackend/config"
	"tesbackend/models"

	"github.com/gin-gonic/gin"
)

// Get daftar Jurusan
func GetDaftarJurusan(c *gin.Context) {
	var jurusan []models.DaftarJurusan
	config.DB.Find(&jurusan)
	c.JSON(http.StatusOK, jurusan)
}

// GET Daftar Mahasiswa berdasarkan ID Jurusan
func GetDaftrMahasiswa(c *gin.Context) {
	id := c.Param("id")
	var detail []models.DaftarMahasiswa

	result := config.DB.Where("id_jurusan = ?", id).Find(&detail)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak di temukan"})
		return
	}

	c.JSON(http.StatusOK, detail)
}

// Create Jurusan baru
func CreateJurusan(c *gin.Context) {
	var newJurusan models.DaftarJurusan

	if err := c.ShouldBindJSON(&newJurusan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}

	config.DB.Create(&newJurusan)

	c.JSON(http.StatusCreated, gin.H{"message": "Mahasiswa berhasil ditambahkan", "data": newJurusan})

}

// Create Mahasiswa baru
func CreateMahasiswa(c *gin.Context) {
	var newMahasiswa models.DaftarMahasiswa

	if err := c.ShouldBindJSON(&newMahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}

	var jurusan models.DaftarJurusan
	if err := config.DB.First(&jurusan, newMahasiswa.IdJurusan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Jurusan tidak ditemukan"})
		return
	}

	// config.DB.Create(&newDetail)
	if err := config.DB.Create(&newMahasiswa).Error; err != nil {
		log.Println("Gagal menyimpan ke database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	log.Println("Data berhasil disimpan:", newMahasiswa)
	c.JSON(http.StatusCreated, gin.H{"message": "Mahasiswa berhasil di tambahkan", "data": newMahasiswa})
}

// Update data Mahasiswa berdasarkan ID
func UpdateMahasiswa(c *gin.Context) {
	id := c.Param("id")
	var mahasiswa models.DaftarMahasiswa

	// Cek apakah mahasiswa dengan ID tsb ada
	if err := config.DB.First(&mahasiswa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
		return
	}

	// Bind JSON input ke variabel baru
	var input models.DaftarMahasiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}

	// Update nilai
	mahasiswa.Nama = input.Nama
	mahasiswa.Nim = input.Nim
	mahasiswa.Alamat = input.Alamat
	mahasiswa.IdJurusan = input.IdJurusan
	mahasiswa.Tanggal = input.Tanggal

	if err := config.DB.Save(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa berhasil diperbarui", "data": mahasiswa})
}

// Delete Mahasiswa berdasarkan ID
func DeleteMahasiswa(c *gin.Context) {
	id := c.Param("id")
	var mahasiswa models.DaftarMahasiswa

	// Cek apakah mahasiswa dengan ID tsb ada
	if err := config.DB.First(&mahasiswa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
		return
	}

	// Hapus dari DB
	if err := config.DB.Delete(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa berhasil dihapus"})
}

// Get semua Mahasiswa beserta nama Jurusan (JOIN)
func GetMahasiswaDenganJurusan(c *gin.Context) {
	var hasil []models.MahasiswaJurusan

	query := `
		SELECT m.nama, m.nim, j.jurusan
		FROM mahasiswa m
		JOIN jurusan j ON m.id_jurusan = j.id
	`

	if err := config.DB.Raw(query).Scan(&hasil).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, hasil)
}

// Get semua data Mahasiswa
func GetSemuaMahasiswa(c *gin.Context) {
	var mahasiswa []models.DaftarMahasiswa

	if err := config.DB.Find(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, mahasiswa)
}

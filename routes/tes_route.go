package routes

import (
	"tesbackend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//Endpoint usaha
	r.GET("/jurusan", controllers.GetDaftarJurusan)
	r.GET("/jurusan/:id", controllers.GetDaftrMahasiswa)
	r.POST("/jurusan", controllers.CreateJurusan)
	r.POST("/jurusan/mahasiswa", controllers.CreateMahasiswa)

	r.PUT("/mahasiswa/:id", controllers.UpdateMahasiswa)
	r.DELETE("/mahasiswa/:id", controllers.DeleteMahasiswa)

	r.GET("/mahasiswa/jurusan", controllers.GetMahasiswaDenganJurusan)

	r.GET("/mahasiswa", controllers.GetSemuaMahasiswa)

	return r
}

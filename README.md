# TesBE_KelasPintar

Untuk menjalankan aplikasi cukup tulis di terminal : go run main.go

Untuk melakukan Kegiatan CRUD bisa menggunakan info link di bawah ini

METHOD GET 
1. data Mahasiswa + Jurusan gunakan route : http://localhost:8080/mahasiswa/jurusan
2. data Mahasiswa sesuai Id jurusan       : http://localhost:8080/jurusan/:id
                                  Contoh  : http://localhost:8080/jurusan/2
3. data semua jurusan                     : http://localhost:8080/jurusan
4. data semua mahasiswa                   : http://localhost:8080/mahasiswa

METHOD POST
1. menambahkan data jurusan               : http://localhost:8080/jurusan
   gunakan body seperti ini :
    {
      "id": 3,
      "jurusan": "Informatika"
    }
2. menambahkan data mahasiswa             : http://localhost:8080/jurusan/mahasiswa
   gunakan body seperti ini :
    {
      "id": 3,
      "id_jurusan": 2,
      "nama": "Trisna 3",
      "nim": "654321",
      "created_at": "2025-04-09 08:16:21.606854",
      "alamat": "Depok"
    }

    //id_jurusan sudah autu konek dengan id pada table jurusan

METHOD PUT
1. Update data pada tabel Mahasiswa        : http://localhost:8080/mahasiswa/:id
    Menggunakan id pada mahasiswa, Contoh  : http://localhost:8080/mahasiswa/1
    yang akan di update adalah data di tabel Mahasiswa yang memiliki id 1
   Gunakan Body seperti ini :
   {
    "nama": "sholah",
    "nim": "32602000002",
    "alamat": "Cileungsi",
    "created_at": "2025-04-09 08:16:21.606854",
    "id_jurusan": 3
   }

METHOD DEL
1. Menghapus Data pada tabel Mahasiswa      : http://localhost:8080/mahasiswa/:id
   Menggunakan id pada mahasiswa, Contoh  : http://localhost:8080/mahasiswa/4
   yang akan di update adalah data di tabel Mahasiswa yang memiliki id 4

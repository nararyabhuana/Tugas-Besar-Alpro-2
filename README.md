# TaskMate

TaskMate adalah aplikasi manajemen tugas berbasis **Command Line Interface (CLI)** yang dibuat menggunakan bahasa **Go (Golang)** sebagai tugas besar mata kuliah **Algoritma dan Pemrograman 2**.

Program ini memungkinkan pengguna untuk mengelola daftar tugas, mulai dari menambahkan, melihat, mengubah, menghapus, mencari, mengurutkan, hingga melihat statistik tugas.

---

## Fitur

* Menambahkan data tugas
* Menampilkan seluruh daftar tugas
* Mengubah data tugas
* Menghapus data tugas
* Mencari tugas menggunakan:

  * Sequential Search
  * Binary Search
* Mengurutkan data menggunakan:

  * Selection Sort (berdasarkan tingkat kesulitan)
  * Insertion Sort (berdasarkan durasi)
* Menampilkan statistik tugas:

  * Jumlah tugas yang telah selesai
  * Rata-rata durasi pengerjaan

---

## Struktur Data

Setiap tugas memiliki informasi sebagai berikut:

| Field     | Tipe Data | Keterangan                               |
| --------- | --------- | ---------------------------------------- |
| Nama      | String    | Nama tugas                               |
| Ruangan   | String    | Lokasi atau ruangan pelaksanaan tugas    |
| Kesulitan | Integer   | Tingkat kesulitan tugas                  |
| Durasi    | Integer   | Estimasi durasi pengerjaan               |
| Status    | String    | Status tugas (misalnya: selesai / belum) |

Data disimpan menggunakan array statis dengan kapasitas maksimal 100 data.

---

## Algoritma yang Digunakan

### Searching

* Sequential Search
* Binary Search

### Sorting

* Selection Sort berdasarkan tingkat kesulitan
* Insertion Sort berdasarkan durasi
* Selection Sort internal berdasarkan nama untuk kebutuhan Binary Search

---

## Menu Program

```
1. Tambah Tugas
2. Lihat Tugas
3. Edit Tugas
4. Hapus Tugas
5. Cari Tugas
6. Urutkan Tugas
7. Statistik
0. Keluar
```

---

## Cara Menjalankan Program

### Clone Repository

```bash
git clone https://github.com/nararyabhuana/Tugas-Besar-Alpro-2.git
```

### Masuk ke Folder Project

```bash
cd Tugas-Besar-Alpro-2
```

### Jalankan Program

```bash
go run .
```

atau

```bash
go run main.go
```

---

## Teknologi

* Go (Golang)
* Command Line Interface (CLI)

---

## Tujuan Proyek

Project ini dibuat sebagai implementasi konsep dasar algoritma dan struktur data, meliputi:

* Array
* Struct
* Fungsi
* CRUD
* Searching
* Sorting
* Pengolahan data sederhana
* Statistik data

---

## Pengembang

Disusun sebagai Tugas Besar Algoritma dan Pemrograman 2 oleh kelompok pengembang TaskMate.

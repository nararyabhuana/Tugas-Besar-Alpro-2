package main

import (
	"fmt"
)

type Tugas struct {
	nama      string
	ruangan   string
	kesulitan int
	durasi    int
	status    string
}

var tugas [100]Tugas
var jumlah int

func refresh() {
	fmt.Print("\033[H\033[2J")
}

func lanjut() {
	var y string

	fmt.Print("\nLanjut? (y) ")
	fmt.Scan(&y)

	refresh()
}

func main() {
	var pilih int

	for {

		refresh()

		fmt.Println("=================================")
		fmt.Println("          TASKMATE")
		fmt.Println("=================================")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Lihat Tugas")
		fmt.Println("3. Edit Tugas")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Cari Tugas")
		fmt.Println("6. Urutkan Tugas")
		fmt.Println("7. Statistik")
		fmt.Println("0. Keluar")
		fmt.Println("=================================")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambah()
		case 2:
			lihat()
		case 3:
			edit()
		case 4:
			hapus()
		case 5:
			cari()
		case 6:
			urut()
		case 7:
			statistik()
		case 0:
			fmt.Println("\nProgram selesai")
			return
		}
	}
}

func tambah() {

	if jumlah >= 100 {
		fmt.Println("\nData penuh")
		lanjut()
		return
	}

	refresh()

	fmt.Print("Nama tugas : ")
	fmt.Scan(&tugas[jumlah].nama)

	fmt.Print("Ruangan : ")
	fmt.Scan(&tugas[jumlah].ruangan)

	fmt.Print("Kesulitan : ")
	fmt.Scan(&tugas[jumlah].kesulitan)

	fmt.Print("Durasi : ")
	fmt.Scan(&tugas[jumlah].durasi)

	fmt.Print("Status : ")
	fmt.Scan(&tugas[jumlah].status)

	jumlah++

	fmt.Println("\nTugas berhasil ditambah")

	lanjut()
}

func lihat() {

	if jumlah == 0 {
		fmt.Println("\nBelum ada data")
		lanjut()
		return
	}

	refresh()

	for i := 0; i < jumlah; i++ {

		fmt.Println("=================================")
		fmt.Println("Tugas", i+1)
		fmt.Println("=================================")
		fmt.Println("Nama :", tugas[i].nama)
		fmt.Println("Ruangan :", tugas[i].ruangan)
		fmt.Println("Kesulitan :", tugas[i].kesulitan)
		fmt.Println("Durasi :", tugas[i].durasi)
		fmt.Println("Status :", tugas[i].status)
		fmt.Println()
	}

	lanjut()
}

func edit() {
	var nama string
	var ketemu bool

	refresh()

	fmt.Print("Masukkan nama tugas : ")
	fmt.Scan(&nama)

	for i := 0; i < jumlah; i++ {

		if tugas[i].nama == nama {

			fmt.Print("Nama baru : ")
			fmt.Scan(&tugas[i].nama)

			fmt.Print("Ruangan baru : ")
			fmt.Scan(&tugas[i].ruangan)

			fmt.Print("Kesulitan baru : ")
			fmt.Scan(&tugas[i].kesulitan)

			fmt.Print("Durasi baru : ")
			fmt.Scan(&tugas[i].durasi)

			fmt.Print("Status baru : ")
			fmt.Scan(&tugas[i].status)

			fmt.Println("\nData berhasil diubah")

			ketemu = true
			break
		}
	}

	if ketemu == false {
		fmt.Println("\nData tidak ditemukan")
	}

	lanjut()
}

func hapus() {
	var nama string
	var ketemu bool

	refresh()

	fmt.Print("Masukkan nama tugas : ")
	fmt.Scan(&nama)

	for i := 0; i < jumlah; i++ {

		if tugas[i].nama == nama {

			for j := i; j < jumlah-1; j++ {
				tugas[j] = tugas[j+1]
			}

			jumlah--

			fmt.Println("\nData berhasil dihapus")

			ketemu = true
			break
		}
	}

	if ketemu == false {
		fmt.Println("\nData tidak ditemukan")
	}

	lanjut()
}

func cari() {
	var pilih int

	refresh()

	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		sequential()
	} else if pilih == 2 {
		binary()
	}
}

func sequential() {
	var nama string
	var ketemu bool

	fmt.Print("\nCari nama tugas : ")
	fmt.Scan(&nama)

	for i := 0; i < jumlah; i++ {

		if tugas[i].nama == nama {

			fmt.Println("\nData ditemukan")
			fmt.Println("Nama :", tugas[i].nama)
			fmt.Println("Ruangan :", tugas[i].ruangan)
			fmt.Println("Kesulitan :", tugas[i].kesulitan)
			fmt.Println("Durasi :", tugas[i].durasi)
			fmt.Println("Status :", tugas[i].status)

			ketemu = true
			break
		}
	}

	if ketemu == false {
		fmt.Println("\nData tidak ditemukan")
	}

	lanjut()
}

func binary() {
	var nama string
	var kiri, kanan, tengah int
	var ketemu bool

	sortNama()

	fmt.Print("\nCari nama tugas : ")
	fmt.Scan(&nama)

	kiri = 0
	kanan = jumlah - 1

	for kiri <= kanan {

		tengah = (kiri + kanan) / 2

		if tugas[tengah].nama == nama {

			fmt.Println("\nData ditemukan")
			fmt.Println("Nama :", tugas[tengah].nama)
			fmt.Println("Ruangan :", tugas[tengah].ruangan)
			fmt.Println("Kesulitan :", tugas[tengah].kesulitan)
			fmt.Println("Durasi :", tugas[tengah].durasi)
			fmt.Println("Status :", tugas[tengah].status)

			ketemu = true
			break

		} else if nama < tugas[tengah].nama {

			kanan = tengah - 1

		} else {

			kiri = tengah + 1
		}
	}

	if ketemu == false {
		fmt.Println("\nData tidak ditemukan")
	}

	lanjut()
}

func urut() {
	var pilih int

	refresh()

	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selection()
	} else if pilih == 2 {
		insertion()
	}
}

func selection() {
	var temp Tugas

	for i := 0; i < jumlah-1; i++ {

		min := i

		for j := i + 1; j < jumlah; j++ {

			if tugas[j].kesulitan < tugas[min].kesulitan {
				min = j
			}
		}

		temp = tugas[i]
		tugas[i] = tugas[min]
		tugas[min] = temp
	}

	fmt.Println("\nData berhasil diurutkan")

	lanjut()
}

func insertion() {
	var temp Tugas
	var j int

	for i := 1; i < jumlah; i++ {

		temp = tugas[i]
		j = i - 1

		for j >= 0 && tugas[j].durasi > temp.durasi {

			tugas[j+1] = tugas[j]
			j--
		}

		tugas[j+1] = temp
	}

	fmt.Println("\nData berhasil diurutkan")

	lanjut()
}

func sortNama() {
	var temp Tugas

	for i := 0; i < jumlah-1; i++ {

		min := i

		for j := i + 1; j < jumlah; j++ {

			if tugas[j].nama < tugas[min].nama {
				min = j
			}
		}

		temp = tugas[i]
		tugas[i] = tugas[min]
		tugas[min] = temp
	}
}

func statistik() {
	var selesai int
	var total int

	refresh()

	for i := 0; i < jumlah; i++ {

		if tugas[i].status == "selesai" {
			selesai++
		}

		total += tugas[i].durasi
	}

	fmt.Println("Jumlah tugas selesai :", selesai)

	if jumlah > 0 {
		fmt.Println("Rata-rata durasi :", total/jumlah)
	}

	lanjut()
}

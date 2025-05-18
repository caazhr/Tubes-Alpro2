package main

import (
	"fmt"
	"strings"
)

const maxInvestasi = 100

type Investasi struct {
	NamaAset      string
	JenisAset     string
	DanaAwal      float64
	NilaiSekarang float64
}

var dataInvestasi [maxInvestasi]Investasi
var jumlahData int

func main() {
	var pilihan int
	for {
		fmt.Println("\n===--- Daftar Menu ---===")
		fmt.Println("1. Tambah Data Investasi")
		fmt.Println("2. Ubah Data Investasi")
		fmt.Println("3. Hapus Data Investasi")
		fmt.Println("4. Cari Investasi")
		fmt.Println("5. Urutkan Investasi")
		fmt.Println("6. Hitung dan tampilkan keuntungan/kerugian")
		fmt.Println("7. Tampilkan Laporan")
		fmt.Println("8. Keluar")
		fmt.Println("===-------------------===")
		fmt.Print("Pilih menu (1-8): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahInvestasi()
		case 2:
			ubahInvestasi()
		case 3:
			hapusInvestasi()
		case 4:
			menuCariInvestasi()
		case 5:
			menuUrutInvestasi()
		case 6:
			hitungKeuntungan()
		case 7:
			tampilkanLaporan()
		case 8:
			fmt.Println("Program Selesai.")
			fmt.Println("===--- Terima Kasih ---===")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahInvestasi() {
	if jumlahData >= maxInvestasi {
		fmt.Println("Kapasitas penyimpanan penuh!")
		return
	}
	fmt.Print("Masukkan nama aset: ")
	fmt.Scan(&dataInvestasi[jumlahData].NamaAset)
	fmt.Print("Masukkan jenis aset: ")
	fmt.Scan(&dataInvestasi[jumlahData].JenisAset)
	fmt.Print("Masukkan jumlah dana awal: ")
	fmt.Scan(&dataInvestasi[jumlahData].DanaAwal)
	fmt.Print("Masukkan nilai sekarang: ")
	fmt.Scan(&dataInvestasi[jumlahData].NilaiSekarang)
	jumlahData++
	fmt.Println("Data berhasil ditambahkan.")
}

func ubahInvestasi() {
	fmt.Print("Masukkan nama aset yang akan diubah: ")
	var nama string
	fmt.Scan(&nama)
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	fmt.Print("Masukkan jenis aset baru: ")
	fmt.Scan(&dataInvestasi[idx].JenisAset)
	fmt.Print("Masukkan dana awal baru: ")
	fmt.Scan(&dataInvestasi[idx].DanaAwal)
	fmt.Print("Masukkan nilai sekarang baru: ")
	fmt.Scan(&dataInvestasi[idx].NilaiSekarang)
	fmt.Println("Data berhasil diubah.")
}

func hapusInvestasi() {
	fmt.Print("Masukkan nama aset yang akan dihapus: ")
	var nama string
	fmt.Scan(&nama)
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	for i := idx; i < jumlahData-1; i++ {
		dataInvestasi[i] = dataInvestasi[i+1]
	}
	jumlahData--
	fmt.Println("Data berhasil dihapus.")
}

func sequentialSearch(nama string) int {
	for i := 0; i < jumlahData; i++ {
		if strings.EqualFold(dataInvestasi[i].NamaAset, nama) {
			return i
		}
	}
	return -1
}

func binarySearch(nama string) int {
	low, high := 0, jumlahData-1
	for low <= high {
		mid := (low + high) / 2
		if strings.EqualFold(dataInvestasi[mid].NamaAset, nama) {
			return mid
		} else if strings.ToLower(dataInvestasi[mid].NamaAset) < strings.ToLower(nama) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func menuCariInvestasi() {
	var nama string
	fmt.Print("Masukkan nama aset yang dicari: ")
	fmt.Scan(&nama)
	fmt.Println("Pilih metode pencarian:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (array harus sudah diurutkan)")
	var metode int
	fmt.Scan(&metode)
	var idx int
	if metode == 1 {
		idx = sequentialSearch(nama)
	} else {
		idx = binarySearch(nama)
	}
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
	} else {
		fmt.Printf("Ditemukan: %+v\n", dataInvestasi[idx])
	}
}

func selectionSortByKeuntungan(ascending bool) {
	for i := 0; i < jumlahData-1; i++ {
		idx := i
		for j := i + 1; j < jumlahData; j++ {
			if ascending {
				if keuntungan(j) < keuntungan(idx) {
					idx = j
				}
			} else {
				if keuntungan(j) > keuntungan(idx) {
					idx = j
				}
			}
		}
		dataInvestasi[i], dataInvestasi[idx] = dataInvestasi[idx], dataInvestasi[i]
	}
}

func insertionSortByNama(ascending bool) {
	for i := 1; i < jumlahData; i++ {
		temp := dataInvestasi[i]
		j := i - 1
		for j >= 0 && ((ascending && strings.ToLower(dataInvestasi[j].NamaAset) > strings.ToLower(temp.NamaAset)) || (!ascending && strings.ToLower(dataInvestasi[j].NamaAset) < strings.ToLower(temp.NamaAset))) {
			dataInvestasi[j+1] = dataInvestasi[j]
			j--
		}
		dataInvestasi[j+1] = temp
	}
}

func menuUrutInvestasi() {
	fmt.Println("Urut berdasarkan:")
	fmt.Println("1. Nama Aset (Insertion Sort)")
	fmt.Println("2. Keuntungan (Selection Sort)")
	var pilihan, arah int
	fmt.Scan(&pilihan)
	fmt.Println("Urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Scan(&arah)
	ascending := arah == 1
	if pilihan == 1 {
		insertionSortByNama(ascending)
	} else {
		selectionSortByKeuntungan(ascending)
	}
	fmt.Println("Data berhasil diurutkan.")
}

func keuntungan(i int) float64 {
	return dataInvestasi[i].NilaiSekarang - dataInvestasi[i].DanaAwal
}

func hitungKeuntungan() {
	fmt.Println("\n=== Keuntungan/Kerugian Investasi ===")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%s: %.2f\n", dataInvestasi[i].NamaAset, keuntungan(i))
	}
}

func tampilkanLaporan() {
	fmt.Println("\n=== Laporan Portofolio Investasi ===")
	fmt.Printf("%-20s %-10s %-15s %-15s %-10s\n", "Nama Aset", "Jenis", "Dana Awal", "Nilai Sekarang", "Untung")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%-20s %-10s %-15.2f %-15.2f %-10.2f\n",
			dataInvestasi[i].NamaAset,
			dataInvestasi[i].JenisAset,
			dataInvestasi[i].DanaAwal,
			dataInvestasi[i].NilaiSekarang,
			keuntungan(i))

	}
}

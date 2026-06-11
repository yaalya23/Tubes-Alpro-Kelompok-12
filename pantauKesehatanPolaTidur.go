/*
Nama Anggota :
 1. Nurul Aulia Zainal (103052530012)
 2. Fathiya Alya Ruasnadita (103052500023)

Nomor Topik : 12
Judul Topik : Aplikasi Pemantauan Kesehatan dan Pola Tidur Sederhana

Deskripsi Program :
Program berbasis konsol ini dirancang untuk melakukan manajemen data riwayat tidur pengguna,
menghitung durasi tidur secara otomatis (termasuk penanganan pergantian hari), memberikan
klasifikasi kualitas tidur, serta menyajikan laporan analisis statistik harian.

Tantangan yang dihadapi :
 1. Memproses kalkulasi durasi waktu desimal secara akurat dari parameter jam dan menit input manual.
 2. Mengendalikan seluruh alur pencarian dan penggeseran indeks array statis murni dengan logika kondisi.
 3. Memastikan fungsi Binary Search hanya tereksekusi secara valid setelah kondisi data array terurut.

Masalah yang dihadapi:
 Pembatasan penggunaan keyword kontrol seperti 'break' dan 'continue' menuntut efisiensi akurasi 
 pada kondisi terminasi di setiap struktur perulangan (for-loop).
*/
package main

import "fmt"

const nMAX int = 1000

type tWaktu struct {
	Jam   int
	Menit int
}

type polaTidur struct {
	Tanggal					string
	WaktuTidur, WaktuBangun tWaktu
	Durasi      			float64
	Kualitas, Saran			string
}

var tabTidur [nMAX]polaTidur
var nData int = 0

func hitungDurasi(tidur, bangun tWaktu) float64 {
	var totMenitTidur, totMenitBangun, selisihMenit int

	totMenitTidur = (tidur.Jam * 60) + tidur.Menit
	totMenitBangun = (bangun.Jam * 60) + bangun.Menit

	if totMenitBangun < totMenitTidur {
		selisihMenit = ((24 * 60) - totMenitTidur) + totMenitBangun
	} else {
		selisihMenit = totMenitBangun - totMenitTidur
	}

	return float64(selisihMenit) / 60.0
}

func kualitasDanSaran(durasi float64, kualitas *string, saran *string) {
	if durasi >= 7.0 && durasi <= 9.0 {
		*kualitas = "Baik"
		*saran = "Pola tidur ideal dan memenuhi standar kesehatan harian."
	} else if durasi > 9.0 {
		*kualitas = "Cukup"
		*saran = "Durasi tidur berlebih, kurangi durasi agar tubuh tidak lemas."
	} else {
		*kualitas = "Kurang"
		*saran = "Durasi tidur kurang dari batas minimal kesehatan 7 jam harian."
	}
}

func tambahRiwayat(tanggal string, tidur, bangun tWaktu) {
	if nData < nMAX {
		tabTidur[nData].Tanggal = tanggal
		tabTidur[nData].WaktuTidur = tidur
		tabTidur[nData].WaktuBangun = bangun
		tabTidur[nData].Durasi = hitungDurasi(tidur, bangun)
		kualitasDanSaran(tabTidur[nData].Durasi, &tabTidur[nData].Kualitas, &tabTidur[nData].Saran)
		nData++
		fmt.Println("Data riwayat tidur berhasil disimpan.")
	} else {
		fmt.Println("Gagal menambahkan data: kapasitas penyimpanan penuh.")
	}
}

func cariIdxTanggal(tanggal string) int {
	var i int = 0
	var ketemu int = -1
	
	for i < nData && ketemu == -1 {
		if tabTidur[i].Tanggal == tanggal {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func ubahRiwayat(tanggalLama, tanggalBaru string, tidurBaru, bangunBaru tWaktu) {
	var idx int
	idx = cariIdxTanggal(tanggalLama)

	if idx != -1 {
		tabTidur[idx].Tanggal = tanggalBaru
		tabTidur[idx].WaktuTidur = tidurBaru
		tabTidur[idx].WaktuBangun = bangunBaru
		tabTidur[idx].Durasi = hitungDurasi(tidurBaru, bangunBaru)
		kualitasDanSaran(tabTidur[idx].Durasi, &tabTidur[idx].Kualitas, &tabTidur[idx].Saran)
		fmt.Println("Data riwayat tidur berhasil diperbarui.")
	} else {
		fmt.Println("Kesalahan: Data tanggal tersebut tidak ditemukan.")
	}
}

func hapusRiwayat(tanggal string) {
	var idx, i int
	idx = cariIdxTanggal(tanggal)

	if idx != -1 {
		i = idx
		for i < nData-1 {
			tabTidur[i] = tabTidur[i+1]
			i++
		}
		nData--
		fmt.Println("Data riwayat tidur berhasil dihapus.")
	} else {
		fmt.Println("Kesalahan: Data tanggal tersebut tidak ditemukan.")
	}
}

func cetakSemuaData() {
	var i int = 0
	if nData == 0 {
		fmt.Println("Database kosong, belum ada data pola tidur yang tercatat.")
	}
	for i < nData {
		fmt.Printf("Tanggal: %s | Tidur: %02d:%02d | Bangun: %02d:%02d | Durasi: %.2f Jam | Kualitas: %s\n",
			tabTidur[i].Tanggal, tabTidur[i].WaktuTidur.Jam, tabTidur[i].WaktuTidur.Menit,
			tabTidur[i].WaktuBangun.Jam, tabTidur[i].WaktuBangun.Menit, tabTidur[i].Durasi, tabTidur[i].Kualitas)
		fmt.Printf("Catatan Medis: %s\n--------------------------------------------------\n", tabTidur[i].Saran)
		i++
	}
}

func selectionSortDurasi(urutMenaik bool) {
	var i, j, idxE int
	var temp polaTidur

	i = 0
	for i < nData-1 {
		idxE = i
		j = i + 1
		for j < nData {
			if urutMenaik {
				if tabTidur[j].Durasi < tabTidur[idxE].Durasi {
					idxE = j
				}
			} else {
				if tabTidur[j].Durasi > tabTidur[idxE].Durasi {
					idxE = j
				}
			}
			j++
		}
		temp = tabTidur[i]
		tabTidur[i] = tabTidur[idxE]
		tabTidur[idxE] = temp
		i++
	}
	fmt.Println("Pengurutan berdasarkan durasi selesai.")
}

func insertionSortTanggal(urutMenaik bool) {
	var i, j int
	var temp polaTidur

	i = 1
	for i < nData {
		temp = tabTidur[i]
		j = i - 1

		if urutMenaik {
			for j >= 0 && tabTidur[j].Tanggal > temp.Tanggal {
				tabTidur[j+1] = tabTidur[j]
				j--
			}
		} else {
			for j >= 0 && tabTidur[j].Tanggal < temp.Tanggal {
				tabTidur[j+1] = tabTidur[j]
				j--
			}
		}
		tabTidur[j+1] = temp
		i++
	}
	fmt.Println("Pengurutan berdasarkan tanggal selesai.")
}

func sequentialSearchTanggal(targetTanggal string) int {
	var i int = 0
	var ketemu int = -1

	for i < nData && ketemu == -1 {
		if tabTidur[i].Tanggal == targetTanggal {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func binarySearchTanggal(targetTanggal string) int {
	var kiri int = 0
	var kanan int = nData - 1
	var tengah int
	var ketemu int = -1

	for kiri <= kanan && ketemu == -1 {
		tengah = (kiri + kanan) / 2

		if tabTidur[tengah].Tanggal == targetTanggal {
			ketemu = tengah
		} else if tabTidur[tengah].Tanggal < targetTanggal {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return ketemu
}

func tampilkanHasilCari(idx int) {
	if idx != -1 {
		fmt.Println("\nData Ditemukan:")
		fmt.Printf("Tanggal      : %s\n", tabTidur[idx].Tanggal)
		fmt.Printf("Waktu Tidur  : %02d:%02d\n", tabTidur[idx].WaktuTidur.Jam, tabTidur[idx].WaktuTidur.Menit)
		fmt.Printf("Waktu Bangun : %02d:%02d\n", tabTidur[idx].WaktuBangun.Jam, tabTidur[idx].WaktuBangun.Menit)
		fmt.Printf("Durasi Tidur : %.2f Jam\n", tabTidur[idx].Durasi)
		fmt.Printf("Kualitas     : %s\n", tabTidur[idx].Kualitas)
		fmt.Printf("Saran        : %s\n", tabTidur[idx].Saran)
	} else {
		fmt.Println("\nData tidak ditemukan.")
	}
}

func tampilkanLaporanMingguan() {
	var i, startIdx, hitungHari int
	var totDurasi, rataRata float64

	if nData == 0 {
		fmt.Println("\nData kosong, rekapitulasi mingguan tidak dapat diproses.")
		return
	}

	fmt.Println("\n==================================================")
	fmt.Println("LAPORAN KESEHATAN DAN POLA TIDUR MINGGUAN")
	fmt.Println("==================================================")

	if nData > 7 {
		startIdx = nData - 7
	} else {
		startIdx = 0
	}

	fmt.Println("Rekapitulasi 7 data terakhir:")
	i = startIdx
	totDurasi = 0.0
	hitungHari = 0
	
	for i < nData {
		fmt.Printf("- %s: %.2f Jam (%s)\n", tabTidur[i].Tanggal, tabTidur[i].Durasi, tabTidur[i].Kualitas)
		totDurasi += tabTidur[i].Durasi
		hitungHari++
		i++
	}

	rataRata = totDurasi / float64(hitungHari)

	fmt.Println("--------------------------------------------------")
	fmt.Printf("Total Durasi Tidur : %.2f Jam\n", totDurasi)
	fmt.Printf("Rata-rata Harian    : %.2f Jam/hari\n", rataRata)
	fmt.Println("--------------------------------------------------")

	fmt.Print("Status Mingguan: ")
	if rataRata >= 7.0 && rataRata <= 9.0 {
		fmt.Println("Kondisi optimal. Rata-rata durasi tidur harian memenuhi standar kesehatan.")
	} else if rataRata > 9.0 {
		fmt.Println("Kondisi kurang optimal akibat durasi tidur berlebih (hipersomnia).")
	} else {
		fmt.Println("Kondisi buruk. Rata-rata durasi tidur harian berada di bawah ambang batas minimal.")
	}
	fmt.Println("==================================================")
}

func Welcome() {
	fmt.Println("==================================================")
	fmt.Println("       SELAMAT DATANG DI APLIKASI KESEHATAN       ")
	fmt.Println("             DAN PEMANTAUAN POLA TIDUR            ")
	fmt.Println("==================================================")
	fmt.Println("Program ini digunakan untuk mencatat dan menganalisis")
	fmt.Println("kualitas tidur harian serta laporan mingguan Anda.")
	fmt.Println("==================================================")
}

func CetakMenu() {
	fmt.Println("\n================= MENU PROGRAM =================")
	fmt.Println("1. Input Riwayat Tidur baru")
	fmt.Println("2. Cetak dan Urutkan Riwayat Tidur")
	fmt.Println("3. Cari Riwayat Tidur Berdasarkan Tanggal")
	fmt.Println("4. Edit Riwayat Tidur")
	fmt.Println("5. Hapus Riwayat Tidur")
	fmt.Println("6. Cetak Laporan Evaluasi Mingguan")
	fmt.Println("7. Selesai")
	fmt.Println("==================================================")
	fmt.Println("Pilih menu (1-7):")
}

func main() {
	var pilihan int = 0
	var subPilihan, jamT, menitT, jamB, menitB int
	var tgl, tglBaru, targetTgl string
	var urutAsc string
	var urutMenaik bool
	var idxHasil int
	var wTidur, wBangun tWaktu
	
	// Penanda untuk memastikan data sudah di-sorting sebelum bisa di-Binary Search
	var dataTerurut bool = false

	// Fungsi sambutan dipanggil satu kali sebelum masuk ke dalam loop menu
	Welcome()

	for pilihan != 7 {
		CetakMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Tanggal (DD-MM-YYYY):")
			fmt.Scan(&tgl)
			fmt.Println("Jam Tidur (Jam Menit):")
			fmt.Scan(&jamT, &menitT)
			fmt.Println("Jam Bangun (Jam Menit):")
			fmt.Scan(&jamB, &menitB)

			wTidur.Jam = jamT
			wTidur.Menit = menitT
			wBangun.Jam = jamB
			wBangun.Menit = menitB

			tambahRiwayat(tgl, wTidur, wBangun)
			
			// Jika ada data baru masuk, urutan array berubah sehingga status terurut di-reset
			dataTerurut = false

		case 2:
			fmt.Println("\nOpsi Tampilan:")
			fmt.Println("1. Tanpa Pengurutan")
			fmt.Println("2. Urutkan Durasi (Selection Sort)")
			fmt.Println("3. Urutkan Tanggal (Insertion Sort)")
			fmt.Println("Pilihan:")
			fmt.Scan(&subPilihan)

			if subPilihan == 2 || subPilihan == 3 {
				fmt.Println("Urutkan menaik / ascending? (y/n):")
				fmt.Scan(&urutAsc)
				urutMenaik = (urutAsc == "y" || urutAsc == "Y")
				
				if subPilihan == 2 {
					selectionSortDurasi(urutMenaik)
					dataTerurut = false
				} else {
					insertionSortTanggal(urutMenaik)
					dataTerurut = true
				}
			}
			fmt.Println("\n--- DAFTAR RIWAYAT TIDUR ---")
			cetakSemuaData()

		case 3:
			fmt.Println("\nMetode Pencarian:")
			fmt.Println("1. Sequential Search")
			fmt.Println("2. Binary Search (Data harus sudah terurut)")
			fmt.Println("Pilihan:")
			fmt.Scan(&subPilihan)

			// Validasi pengaman: Jika memilih Binary Search tetapi data belum diurutkan
			if subPilihan == 2 && !dataTerurut {
				fmt.Println("\nPeringatan: Anda harus mengurutkan data (Menu 2) sebelum menggunakan Binary Search!")
			} else {
				fmt.Println("Tanggal yang dicari (DD-MM-YYYY):")
				fmt.Scan(&targetTgl)

				if subPilihan == 1 {
					idxHasil = sequentialSearchTanggal(targetTgl)
				} else {
					idxHasil = binarySearchTanggal(targetTgl)
				}
				tampilkanHasilCari(idxHasil)
			}

		case 4:
			fmt.Println("Masukkan tanggal data lama (DD-MM-YYYY):")
			fmt.Scan(&tgl)
			fmt.Println("Masukkan tanggal baru (DD-MM-YYYY):")
			fmt.Scan(&tglBaru)
			fmt.Println("Masukkan Jam Tidur baru (Jam Menit):")
			fmt.Scan(&jamT, &menitT)
			fmt.Println("Masukkan Jam Bangun baru (Jam Menit):")
			fmt.Scan(&jamB, &menitB)

			wTidur.Jam = jamT
			wTidur.Menit = menitT
			wBangun.Jam = jamB
			wBangun.Menit = menitB

			ubahRiwayat(tgl, tglBaru, wTidur, wBangun)
			
			// Reset status karena isi data di dalam array mengalami perubahan nilai
			dataTerurut = false

		case 5:
			fmt.Println("Masukkan tanggal data yang akan dihapus:")
			fmt.Scan(&tgl)
			hapusRiwayat(tgl)
			
			// Reset status karena struktur indeks array bergeser setelah penghapusan
			dataTerurut = false

		case 6:
			tampilkanLaporanMingguan()

		case 7:
			fmt.Println("Program dihentikan.")
			
		default:
			fmt.Println("Kesalahan: Opsi tidak tersedia.")
		}
	}
}
package main
import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type Data struct {
	noreservasi int
	nip string
	nama string
	usia int
	kursi string
}

type Array1[9][4] Data
type Array2[12][4] Data
type Array3[14][4] Data
type Array4[150] Data
var Prioritas Array1
var Eksekutif Array2
var Ekonomi Array3
var Bayi Array4
var Bio Data
var totalpri, totaleks, totaleko int

func main() {
	var menu int
	var quit bool
	quit = false
	clear()

	for !quit {

		fmt.Println()
		fmt.Println("      APLIKASI PENJUALAN TIKET KERETA API")
		fmt.Println("=================================================")
		fmt.Println("1. RESERVASI")
		fmt.Println("2. TEMPAT DUDUK KOSONG")
		fmt.Println("3. CARI ORANG TUA BAYI")
		fmt.Println("4. DATA BAYI PER GERBONG")
		fmt.Println("5. SEMUA PENUMPANG BERDASARKAN USIA")
		fmt.Println("6. STATISTIK HARIAN")
		fmt.Println("7. CETAK DATA PENUMPANG")
		fmt.Println("8. EXIT")
		fmt.Println("=================================================")
		fmt.Println()
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&menu)
			
		switch menu {
			case 1 : reservasi()
			case 2 : hitungkursi()
			case 3 : cariortu()
			case 4 : tampilbayi()
			case 5 : sorting()
			case 6 : okupansi()
			case 7 : cetakdatapenumpang() 
			case 8 : quit = true
			default : main()
		}

	}

}

func reservasi() {
	var kelas,jumlahtiket,norev,baris,col int
	var enter,pindah,kolom string
	var dewasa,anakanak,bayi int

	norev = rand.Intn(7000)
	norev = norev

	fmt.Println("\n            RESERVASI TIKET \n========================================")
	fmt.Println("1. Prioritas | 2. Eksekutif | 3. Ekonomi ")
	fmt.Println()
	fmt.Print("Pilih Kelas Tiket : ")
	fmt.Scan(&kelas)
	for kelas < 1 || kelas > 3 {
		fmt.Print("Pilih Kelas Tiket : ")
		fmt.Scan(&kelas)
	}
	fmt.Println()
	fmt.Print("Input Jumlah Tiket (Maks 10): ")
	fmt.Scan(&jumlahtiket)
	for jumlahtiket < 1 || jumlahtiket > 10 {
		fmt.Print("Input Jumlah Tiket (Maks 10) : ")
		fmt.Scan(&jumlahtiket)
	}
	fmt.Println()
	fmt.Println("Input Data Penumpang : ")
	fmt.Println()

	for i:=0 ; i < jumlahtiket ; i++ {

		fmt.Println("DATA PENUMPANG ",i+1)
		fmt.Print("NIP  : ")
		fmt.Scan(&Bio.nip)
		fmt.Print("NAMA : ")
		fmt.Scan(&Bio.nama)
		fmt.Print("USIA : ")
		fmt.Scan(&Bio.usia)

		if Bio.usia > 2 && Bio.usia < 18 {
			anakanak++
		} else if Bio.usia >= 18 {
			dewasa++
		} else {
			bayi++
		}

		Bio.noreservasi = norev
		// jika kelas 1 maka masuk ke array tabel prioritas
		if kelas == 1 {
			if bayi > 0 && Bio.usia <= 2 {
				e:=0
				for d:=0; d < 150; d++ {
					if Bayi[d].noreservasi == 0 {
						for e < 1 {
							Bayi[d] = Bio	
							e++
						}
					}	
				}
			} else {
				c:=0
				for a:=0; a < 9; a++ {
					for b:=0; b < 4; b++ { 
						if Prioritas[a][b].noreservasi == 0 {
							for c < 1 {
								Prioritas[a][b] = Bio
								totalpri++
								c++
							}
						}
					} 
				}
			}
		} else if kelas == 2 {
			if bayi > 0 && Bio.usia <= 2 {
				e:=0
				for d:=0; d < 150; d++ {
					if Bayi[d].noreservasi == 0 {
						for e < 1 {
							Bayi[d] = Bio	
							e++
						}
					}	
				}
			} else {
				c:=0
				for a:=0; a < 12; a++ {
					for b:=0; b < 4; b++ { 
						if Eksekutif[a][b].noreservasi == 0 {
							for c < 1 {
								Eksekutif[a][b] = Bio	
								totaleks++
								c++
							}
						}
					} 
				}
			}
		} else {
			if bayi > 0 && Bio.usia <= 2 {
				e:=0
				for d:=0; d < 150; d++ {
					if Bayi[d].noreservasi == 0 {
						for e < 1 {
							Bayi[d] = Bio	
							e++
						}
					}	
				}
			} else {
				c:=0
				for a:=0; a < 14; a++ {
					for b:=0; b < 4; b++ { 
						if Ekonomi[a][b].noreservasi == 0 {
							for c < 1 {
								Ekonomi[a][b] = Bio	
								totaleko++
								c++
							}
						}
					} 
				}
			}
		}

		fmt.Println()
	}

	if kelas == 1{
		tiket:=0
		for k:=0; k < 9; k++ {
			for l:=0; l < 4; l++ {
				if Prioritas[k][l].noreservasi == norev && tiket < anakanak+dewasa {
					fmt.Print("NAMA : ",Prioritas[k][l].nama, "  |  SEAT : ", kursi(k,l), "     Mau Dipindah ? (y/n) ")
					fmt.Scan(&pindah)
					if pindah=="y" || pindah=="Y"{
						fmt.Print("Pindah ke kursi mana? (Cth. 5 A) : ")
						fmt.Scan(&baris, &kolom)

						if baris > 0 && baris < 10{
							for !cekkosong(baris, kolom, kelas){
								fmt.Print("Kursi Terisi, Pilih Kursi Lain : ")
								fmt.Scan(&baris, &kolom)
							} 
							baris -= 1
							if kolom =="A" || kolom == "a" {
								col = 0
							} else if kolom =="B" || kolom == "b" {
								col = 1
							} else if kolom =="C" || kolom == "c" {
								col = 2
							} else if kolom =="D" || kolom == "d" {
								col = 3
							}
							Prioritas[baris][col] = Prioritas[k][l]

							Prioritas[k][l].noreservasi = 0
							Prioritas[k][l].nip = ""
							Prioritas[k][l].nama = ""
							Prioritas[k][l].usia = 0
							tiket++
						} 
					}else{
						tiket++
					}
				}
			} 
		}	
	} else if kelas == 2{
		tiket:=0
		for k:=0; k < 12; k++ {
			for l:=0; l < 4; l++ {
				if Eksekutif[k][l].noreservasi == norev && tiket < anakanak+dewasa {
					fmt.Print("NAMA : ",Eksekutif[k][l].nama, "  |  SEAT : ", kursi(k,l), "     Mau Dipindah ? (y/n) ")
					fmt.Scan(&pindah)
					if pindah=="y" || pindah=="Y"{
						fmt.Print("Pindah ke kursi mana? (Cth. 5 A) : ")
						fmt.Scan(&baris, &kolom)

						if baris > 0 && baris < 10{
							for !cekkosong(baris, kolom, kelas){
								fmt.Print("Kursi Terisi, Pilih Kursi Lain : ")
								fmt.Scan(&baris, &kolom)
							} 
							baris -= 1
							if kolom =="A" || kolom == "a" {
								col = 0
							} else if kolom =="B" || kolom == "b" {
								col = 1
							} else if kolom =="C" || kolom == "c" {
								col = 2
							} else if kolom =="D" || kolom == "d" {
								col = 3
							}
							Eksekutif[baris][col] = Eksekutif[k][l]

							Eksekutif[k][l].noreservasi = 0
							Eksekutif[k][l].nip = ""
							Eksekutif[k][l].nama = ""
							Eksekutif[k][l].usia = 0
							tiket++
						}
					}else{
						tiket++
					}
				}
			} 
		}	
	} else {
		tiket:=0
		for k:=0; k < 14; k++ {
			for l:=0; l < 4; l++ {
				if Ekonomi[k][l].noreservasi == norev && tiket < anakanak+dewasa {
					fmt.Print("NAMA : ",Ekonomi[k][l].nama, "  |  SEAT : ", kursi(k,l), "     Mau Dipindah ? (y/n) ")
					fmt.Scan(&pindah)
					if pindah=="y" || pindah=="Y"{
						fmt.Print("Pindah ke kursi mana? (Cth. 5 A) : ")
						fmt.Scan(&baris, &kolom)

						if baris > 0 && baris < 10{
							for !cekkosong(baris, kolom, kelas){
								fmt.Print("Kursi Terisi, Pilih Kursi Lain : ")
								fmt.Scan(&baris, &kolom)
							} 
							baris -= 1
							if kolom =="A" || kolom == "a" {
								col = 0
							} else if kolom =="B" || kolom == "b" {
								col = 1
							} else if kolom =="C" || kolom == "c" {
								col = 2
							} else if kolom =="D" || kolom == "d" {
								col = 3
							}
							Ekonomi[baris][col] = Ekonomi[k][l]

							Ekonomi[k][l].noreservasi = 0
							Ekonomi[k][l].nip = ""
							Ekonomi[k][l].nama = ""
							Ekonomi[k][l].usia = 0
							tiket++
						} 
					}else{
						tiket++
					}
				}
			} 
		}	
	}

	fmt.Println()
	fmt.Println("===========================================================")
	fmt.Println("                           TIKET                           ")
	fmt.Println("===========================================================")
	fmt.Println()
	fmt.Println()
	if kelas == 1 {
		tampiltiket(norev, kelas)
	} else if kelas == 2 {
		tampiltiket(norev, kelas)
	} else {
		tampiltiket(norev, kelas)
	}
	fmt.Print("PRESS ANY KEY AND ENTER TO CONTINUE...")
    fmt.Scan(&enter) // wait for Enter Key

	fmt.Println()
	fmt.Println("===========================================================")
	fmt.Println("                           HARGA                           ")
	fmt.Println("===========================================================")
	fmt.Println("DEWASA :", dewasa," ORANG" ,"  ANAK ANAK :", anakanak," ORANG" ,"  BAYI :", bayi," ORANG"  )
	fmt.Println("===========================================================")
	fmt.Println("                TOTAL HARGA : Rp.",harga(dewasa, anakanak, kelas))
	fmt.Println("===========================================================")
	fmt.Println()
	fmt.Print("PRESS ANY KEY AND ENTER TO CONTINUE...")
	fmt.Scan(&enter) // wait for Enter Key
	clear()
}

func tampiltiket(nores int, kelas int) {
	if kelas == 1 {
		for k:=0; k < 9; k++ {
			for l:=0; l < 4; l++ {
				if Prioritas[k][l].noreservasi == nores {
					fmt.Println("===========================================================")
					fmt.Println("                     KA. SRILELAWANGSA                     ")
					fmt.Println("PRIORITAS                                            #",Prioritas[k][l].noreservasi)
					fmt.Println()
					fmt.Println("NAMA : ",Prioritas[k][l].nama)
					fmt.Println("SEAT : ",kursi(k,l))
					fmt.Println("===========================================================")
					fmt.Println()
				}
			} 
		}
		for d:=0; d < 150; d++ {
			if Bayi[d].noreservasi == nores {
				fmt.Println("===========================================================")
				fmt.Println("                     KA. SRILELAWANGSA                     ")
				fmt.Println("PRIORITAS                                            #",Bayi[d].noreservasi)
				fmt.Println()
				fmt.Println("NAMA : ",Bayi[d].nama)
				fmt.Println("SEAT : ",cariseatortu(Bayi[d].nip))
				fmt.Println("===========================================================")
				fmt.Println()
			}	
		}
	}else if kelas == 2 {
		for k:=0; k < 12; k++ {
			for l:=0; l < 4; l++ {
				if Eksekutif[k][l].noreservasi == nores {
					fmt.Println("===========================================================")
					fmt.Println("                     KA. SRILELAWANGSA                     ")
					fmt.Println("EKSEKUTIF                                            #",Eksekutif[k][l].noreservasi)
					fmt.Println()
					fmt.Println("NAMA : ",Eksekutif[k][l].nama)
					fmt.Println("SEAT : ",kursi(k,l))
					fmt.Println("===========================================================")
					fmt.Println()
				}
			} 
		}
		for d:=0; d < 150; d++ {
			if Bayi[d].noreservasi == nores {
				fmt.Println("===========================================================")
				fmt.Println("                     KA. SRILELAWANGSA                     ")
				fmt.Println("EKSEKUTIF                                            #",Bayi[d].noreservasi)
				fmt.Println()
				fmt.Println("NAMA : ",Bayi[d].nama)
				fmt.Println("SEAT : ",cariseatortu(Bayi[d].nip))
				fmt.Println("===========================================================")
				fmt.Println()
			}	
		}
	}else {
		for k:=0; k < 14; k++ {
			for l:=0; l < 4; l++ {
				if Ekonomi[k][l].noreservasi == nores {
					fmt.Println("===========================================================")
					fmt.Println("                     KA. SRILELAWANGSA                     ")
					fmt.Println("EKONOMI                                            #",Ekonomi[k][l].noreservasi)
					fmt.Println()
					fmt.Println("NAMA : ",Ekonomi[k][l].nama)
					fmt.Println("SEAT : ",kursi(k,l))
					fmt.Println("===========================================================")
					fmt.Println()
				}
			} 
		}
		for d:=0; d < 150; d++ {
			if Bayi[d].noreservasi == nores {
				fmt.Println("===========================================================")
				fmt.Println("                     KA. SRILELAWANGSA                     ")
				fmt.Println("EKONOMI                                            #",Bayi[d].noreservasi)
				fmt.Println()
				fmt.Println("NAMA : ",Bayi[d].nama)
				fmt.Println("SEAT : ",cariseatortu(Bayi[d].nip))
				fmt.Println("===========================================================")
				fmt.Println()
			}	
		}
	}
	
}

func cekkosong(baris int,kolom string, kelas int) bool{
	var col int

	baris -= 1
	if kolom =="A" || kolom == "a" {
		col = 0
	} else if kolom =="B" || kolom == "b" {
		col = 1
	} else if kolom =="C" || kolom == "c" {
		col = 2
	} else if kolom =="D" || kolom == "d" {
		col = 3
	}

	if kelas==1 {
		if Prioritas[baris][col].noreservasi != 0 {
			return false
		}else{
			return true
		}
	}else if kelas==2 {
		if Eksekutif[baris][col].noreservasi != 0 {
			return false
		}else{
			return true
		}
	}else if kelas==3 {
		if Ekonomi[baris][col].noreservasi != 0 {
			return false
		}else{
			return true
		}
	}else{
		return true
	}

}

func harga(dewasa int, anak int, kelas int) int {
	var hargadewasa, hargaanak, cost int
	
	if kelas == 1 {
		cost = 100000
	} else if kelas == 2 {
		cost = 80000
	} else if kelas == 3 {
		cost = 50000
	}
	var costanak = cost / 2

	if dewasa%2 == 0 {
		if dewasa > 1 && anak == 1{
			anak--
			dewasa++
		} 
	} else 	if dewasa%4 > 0 {
		sel := 4 - dewasa % 4
		if anak >= sel {
			anak -= sel
			dewasa += sel
		}
	}

	for dewasa > 0 {
		if dewasa < 3 {
			hargadewasa += dewasa * cost
			dewasa -= 3
		} else {
			hargadewasa += 2 * cost
			dewasa -= 4
		}
	}
	hargaanak = anak * costanak
	total := hargadewasa + hargaanak
	return total
}

func kursi(baris,kolom int) string {
	kursi:=""
	num := baris+1

	if kolom==0 {
		seat := "A"	
		kursi = fmt.Sprintf("%d%s", num, seat)	
	}else if kolom==1 {
		seat := "B"	
		kursi = fmt.Sprintf("%d%s", num, seat)
	}else if kolom==2 {
		seat := "C"	
		kursi = fmt.Sprintf("%d%s", num, seat)
	}else if kolom==3 {
		seat := "D"	
		kursi = fmt.Sprintf("%d%s", num, seat)
	}
	return kursi
}

func cariortu(){
	var namabayi,nipbayi,namaortu,enter string
	
	fmt.Println("\n              CARI ORANGTUA\n========================================")
	fmt.Println()
	fmt.Print("Nama Bayi : ")
	fmt.Scan(&namabayi)
	for d:=0; d < 150; d++ {
		if Bayi[d].nama == namabayi {
			nipbayi = Bayi[d].nip
		}	
	}

	for i:=0; i < 9; i++ {
		for j:=0; j < 4; j++ {
			if Prioritas[i][j].nip == nipbayi {
				namaortu = Prioritas[i][j].nama
			}
		} 
	}

	for i:=0; i < 12; i++ {
		for j:=0; j < 4; j++ {
			if Eksekutif[i][j].nip == nipbayi {
				namaortu = Eksekutif[i][j].nama
			}
		} 
	}

	for i:=0; i < 14; i++ {
		for j:=0; j < 4; j++ {
			if Ekonomi[i][j].nip == nipbayi {
				namaortu = Ekonomi[i][j].nama
			}
		} 
	}
	
	if namaortu != "" {
		fmt.Println("Nama Orang Tua :", namaortu)
	} else {
		fmt.Println("Tidak Ditemukan")
	}

	fmt.Println()
	fmt.Print("PRESS ANY KEY AND ENTER TO CONTINUE...")
	fmt.Scan(&enter) // wait for Enter Key
	clear()
}

func cariseatortu(nip string) string{
	var baris,kolom int

	for i:=0; i < 9; i++ {
		for j:=0; j < 4; j++ {
			if Prioritas[i][j].nip == nip {
				baris = i
				kolom = j
			}
		} 
	}

	for i:=0; i < 12; i++ {
		for j:=0; j < 4; j++ {
			if Eksekutif[i][j].nip == nip {
				baris = i
				kolom = j
			}
		} 
	}

	for i:=0; i < 14; i++ {
		for j:=0; j < 4; j++ {
			if Ekonomi[i][j].nip == nip {
				baris = i
				kolom = j
			}
		} 
	}
	
	return kursi(baris,kolom)
}

func tampilbayi(){
	var kelas,isi int
	var enter string

	fmt.Println("\n               LIST BAYI  \n========================================")
	fmt.Println("1. Prioritas | 2. Eksekutif | 3. Ekonomi ")
	fmt.Println()
	fmt.Print("Pilih Gerbong: ")
	fmt.Scan(&kelas)
	for kelas < 1 || kelas > 3 {
		fmt.Print("Pilih Gerbong ( 1 / 2 / 3 ): ")
		fmt.Scan(&kelas)
	}

	fmt.Println()
	if kelas==1 {
		for d:=0; d < 150; d++ {
			if Bayi[d].nip != "" {
				nipbayi := Bayi[d].nip
				namabayi := Bayi[d].nama
				for i:=0; i < 9; i++ {
					for j:=0; j < 4; j++ {
						if Prioritas[i][j].nip == nipbayi {
							isi++
							fmt.Print("NAMA : ",namabayi,"  ||  KURSI : ",kursi(i,j))
						}
					} 
				}
				fmt.Println()
			}
		}
	} else if kelas==2 {
		for d:=0; d < 150; d++ {
			if Bayi[d].nip != "" {
				nipbayi := Bayi[d].nip
				namabayi := Bayi[d].nama
				for i:=0; i < 12; i++ {
					for j:=0; j < 4; j++ {
						if Eksekutif[i][j].nip == nipbayi {
							isi++
							fmt.Print("NAMA : ",namabayi,"  ||  KURSI : ",kursi(i,j))
						}
					} 
				}
				fmt.Println()
			}
		}
	} else {
		for d:=0; d < 150; d++ {
			if Bayi[d].nip != "" {
				nipbayi := Bayi[d].nip
				namabayi := Bayi[d].nama
				for i:=0; i < 14; i++ {
					for j:=0; j < 4; j++ {
						if Ekonomi[i][j].nip == nipbayi {
							isi++
							fmt.Print("NAMA : ",namabayi,"  ||  KURSI : ",kursi(i,j))
						}
					} 
				}
				fmt.Println()
			}
		}
	}
	
	if isi == 0 {
		fmt.Println("DATA BAYI KOSONG")
		
	}
	
	fmt.Println()
	fmt.Print("PRESS ANY KEY AND ENTER TO CONTINUE...")
	fmt.Scan(&enter) // wait for Enter Key'
	clear()
	
}

func hitungkursi() {
	var kelas,jumlah int
	var enter string

	fmt.Println("\n          HITUNG KURSI KOSONG \n========================================")
	fmt.Println("1. Prioritas | 2. Eksekutif | 3. Ekonomi ")
	fmt.Println()
	fmt.Print("Pilih Gerbong: ")
	fmt.Scan(&kelas)
	for kelas < 1 || kelas > 3 {
		fmt.Print("Pilih Gerbong ( 1 / 2 / 3 ): ")
		fmt.Scan(&kelas)
	}

	if kelas==1 {
		for i:=0; i < 9; i++ {
			for j:=0; j < 4; j++ {
				if Prioritas[i][j].noreservasi == 0 {
					jumlah++
				}
			} 
		}
		fmt.Println("\n           GERBONG\n==============================") 
		fmt.Println("   PRIORITAS  |  ",jumlah," Kursi") 
	} else if kelas==2 {
		for i:=0; i < 12; i++ {
			for j:=0; j < 4; j++ {
				if Eksekutif[i][j].noreservasi == 0 {
					jumlah++
				}
			} 
		}
		fmt.Println("           GERBONG\n==============================") 
		fmt.Println("   EKSEKUTIF  |  ",jumlah," Kursi") 
	} else {
		for i:=0; i < 14; i++ {
			for j:=0; j < 4; j++ {
				if Ekonomi[i][j].noreservasi == 0 {
					jumlah++
				}
			} 
		}
		fmt.Println("          GERBONG\n==============================") 
		fmt.Println("    EKONOMI  |  ",jumlah," Kursi") 
	}

	fmt.Println()
	fmt.Print("PRESS ANY KEY AND ENTER TO CONTINUE...")
	fmt.Scan(&enter) // wait for Enter Key
	clear()
}

func sorting(){
	var datasorting[200] Data
	var enter string

	z:=0
	for i:=0; i < 9; i++ {
		for j:=0; j < 4; j++ {
			if Prioritas[i][j].noreservasi != 0 && datasorting[z].noreservasi == 0{
				Prioritas[i][j].kursi = kursi(i,j) 
				datasorting[z] = Prioritas[i][j]
				z++
			}
		} 
	}

	for k:=0; k < 12; k++ {
		for l:=0; l < 4; l++ {
			if Eksekutif[k][l].noreservasi != 0 && datasorting[z].noreservasi == 0{
				Eksekutif[k][l].kursi = kursi(k,l)
				datasorting[z] = Eksekutif[k][l]
				z++
			}
		} 
	}

	for m:=0; m < 14; m++ {
		for n:=0; n < 4; n++ {
			if Ekonomi[m][n].noreservasi != 0 && datasorting[z].noreservasi == 0{
				Ekonomi[m][n].kursi = kursi(m,n)
				datasorting[z] = Ekonomi[m][n]
				z++
			}
		} 
	}

	if z==0{
		fmt.Println("\nDATA KOSONG")
	}else{
		fmt.Print("\n        SELURUH DATA PENUMPANG    \n======================================\n\n")
	}
	
	for q:= 0; q < len(Bayi) ; q++ {
		idx := q
		for r := q + 1; r < len(Bayi); r++ {
			if Bayi[r].usia < Bayi[r].usia {
				idx = r
			}
		}
		temp := Bayi[q]
		Bayi[q]= Bayi[idx]
		Bayi[idx] = temp
	}

	for i:=0; i < len(Bayi) ; i++ {
		if Bayi[i].noreservasi != 0{
			fmt.Println("NO RESERVASI : ",Bayi[i].noreservasi)
			fmt.Println("KELAS        : ",searchkelas(Bayi[i].noreservasi))
			fmt.Println("NIP          : ",Bayi[i].nip)
			fmt.Println("NAMA         : ",Bayi[i].nama)
			fmt.Println("USIA         : ",Bayi[i].usia)
			fmt.Println("SEAT         : ",cariseatortu(Bayi[i].nip))
			fmt.Println()
		}
	}

	for o:= 0; o < z ; o++ {
		idx := o
		for p := o + 1; p < z; p++ {
			if datasorting[p].usia < datasorting[idx].usia {
				idx = p
			}
		}
		temp := datasorting[o]
		datasorting[o]= datasorting[idx]
		datasorting[idx] = temp
	}

	for i:=0; i < z ; i++ {
		if datasorting[i].noreservasi != 0{
			fmt.Println("NO RESERVASI : ",datasorting[i].noreservasi)
			fmt.Println("KELAS        : ",searchkelas(datasorting[i].noreservasi))
			fmt.Println("NIP          : ",datasorting[i].nip)
			fmt.Println("NAMA         : ",datasorting[i].nama)
			fmt.Println("USIA         : ",datasorting[i].usia)
			fmt.Println("SEAT         : ",datasorting[i].kursi)
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Print("PRESS ANY KEY AND ENTER TO CONTINUE...")
	fmt.Scan(&enter) // wait for Enter Key
	clear()
}

func cetakdatapenumpang(){
	var enter string

	f, err := os.Create("penumpang_prioritas.txt") // create
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	for k:=0; k < 9; k++ {
		for l:=0; l < 4; l++ {
			if Prioritas[k][l].noreservasi != 0 {
				_, err = f.WriteString(fmt.Sprintf("===========================================================\n"))
				_, err = f.WriteString(fmt.Sprintf("                    KA. SRILELAWANGSA                     \n"))
				_, err = f.WriteString(fmt.Sprintf("PRIORITAS                                            #%d\n\n",Prioritas[k][l].noreservasi))
				_, err = f.WriteString(fmt.Sprintf("NIP  : %s\n",Prioritas[k][l].nip))
				_, err = f.WriteString(fmt.Sprintf("NAMA : %s\n",Prioritas[k][l].nama))
				_, err = f.WriteString(fmt.Sprintf("USIA : %d\n",Prioritas[k][l].usia))
				_, err = f.WriteString(fmt.Sprintf("SEAT : %s\n",kursi(k,l)))
				_, err = f.WriteString(fmt.Sprintf("===========================================================\n"))
				if err != nil {
					fmt.Printf("error writing string: %v", err)
				}

			}
		} 
	}

	f, err = os.Create("penumpang_eksekutif.txt") // create
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	for k:=0; k < 12; k++ {
		for l:=0; l < 4; l++ {
			if Eksekutif[k][l].noreservasi != 0 {
				_, err = f.WriteString(fmt.Sprintf("===========================================================\n"))
				_, err = f.WriteString(fmt.Sprintf("                    KA. SRILELAWANGSA                     \n"))
				_, err = f.WriteString(fmt.Sprintf("EKSEKUTIF                                            #%d\n\n",Eksekutif[k][l].noreservasi))
				_, err = f.WriteString(fmt.Sprintf("NIP  : %s\n",Eksekutif[k][l].nip))
				_, err = f.WriteString(fmt.Sprintf("NAMA : %s\n",Eksekutif[k][l].nama))
				_, err = f.WriteString(fmt.Sprintf("USIA : %d\n",Eksekutif[k][l].usia))
				_, err = f.WriteString(fmt.Sprintf("SEAT : %s\n",kursi(k,l)))
				_, err = f.WriteString(fmt.Sprintf("===========================================================\n"))
				if err != nil {
					fmt.Printf("error writing string: %v", err)
				}

			}
		} 
	}

	f, err = os.Create("penumpang_ekonomi.txt") // create
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	for k:=0; k < 14; k++ {
		for l:=0; l < 4; l++ {
			if Ekonomi[k][l].noreservasi != 0 {
				_, err = f.WriteString(fmt.Sprintf("===========================================================\n"))
				_, err = f.WriteString(fmt.Sprintf("                    KA. SRILELAWANGSA                     \n"))
				_, err = f.WriteString(fmt.Sprintf("EKONOMI                                              #%d\n\n",Ekonomi[k][l].noreservasi))
				_, err = f.WriteString(fmt.Sprintf("NIP  : %s\n",Ekonomi[k][l].nip))
				_, err = f.WriteString(fmt.Sprintf("NAMA : %s\n",Ekonomi[k][l].nama))
				_, err = f.WriteString(fmt.Sprintf("USIA : %d\n",Ekonomi[k][l].usia))
				_, err = f.WriteString(fmt.Sprintf("SEAT : %s\n",kursi(k,l)))
				_, err = f.WriteString(fmt.Sprintf("===========================================================\n"))
				if err != nil {
					fmt.Printf("error writing string: %v", err)
				}

			}
		} 
	}
	
	f, err = os.Create("penumpang_bayi.txt") // create
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	for k:=0; k < 150; k++ {
		if Bayi[k].noreservasi != 0 {
			_, err = f.WriteString(fmt.Sprintf("===========================================================\n"))
			_, err = f.WriteString(fmt.Sprintf("                    KA. SRILELAWANGSA                     \n"))
			_, err = f.WriteString(fmt.Sprintf("%s                                            #%d\n\n",searchkelas(Bayi[k].noreservasi),Bayi[k].noreservasi))
			_, err = f.WriteString(fmt.Sprintf("NIP  : %s\n",Bayi[k].nip))
			_, err = f.WriteString(fmt.Sprintf("NAMA : %s\n",Bayi[k].nama))
			_, err = f.WriteString(fmt.Sprintf("USIA : %d\n",Bayi[k].usia))
			_, err = f.WriteString(fmt.Sprintf("SEAT : %s\n",cariseatortu(Bayi[k].nip)))
			_, err = f.WriteString(fmt.Sprintf("===========================================================\n"))
			if err != nil {
				fmt.Printf("error writing string: %v", err)
			}

		}
	}

	fmt.Println()
	fmt.Println("DATA BERHASIL DICETAK")
	fmt.Println()
	fmt.Print("PRESS ANY KEY AND ENTER TO CONTINUE...")
	fmt.Scan(&enter) // wait for Enter Key
	clear()

}

func searchkelas(nores int) string{
	for i:=0; i < 9; i++ {
		for j:=0; j < 4; j++ {
			if Prioritas[i][j].noreservasi == nores{
				return "PRIORITAS"
			}
		} 
	}

	for k:=0; k < 12; k++ {
		for l:=0; l < 4; l++ {
			if Eksekutif[k][l].noreservasi == nores{
				return "EKSEKUTIF"
			}
		} 
	}

	for m:=0; m < 14; m++ {
		for n:=0; n < 4; n++ {
			if Ekonomi[m][n].noreservasi == nores{
				return "EKONOMI"
			}
		} 
	}
	return "UNDEFINED"
}

func clear(){
	a:=exec.Command("cmd","/c","cls")
	a.Stdout=os.Stdout
	a.Run()
}

func okupansi(){
	var okuPri, okuEks, okuEko float64
	var enter string

	okuPri = float64(totalpri)/float64(36)*100
	okuEks = float64(totaleks)/float64(48)*100
	okuEko = float64(totaleko)/float64(56)*100
	fmt.Println("\n     PRESENTASI OKUPANSI GERBONG    \n======================================")
	fmt.Println("PRIORITAS : ", okuPri, "%")
	fmt.Println("EKSEKUTIF : ", okuEks, "%")
	fmt.Println("EKONOMI   : ", okuEko, "%")

	fmt.Println()
	fmt.Print("PRESS ANY KEY AND ENTER TO CONTINUE...")
	fmt.Scan(&enter) // wait for Enter Key
	clear()
}
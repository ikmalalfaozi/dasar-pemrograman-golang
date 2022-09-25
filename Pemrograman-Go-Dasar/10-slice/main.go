package main

import "fmt"

func main() {
	// contoh pembuatan slice
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])

	// Perbedaan array dengan slice saat deklarasi variable
	// var fruitsA = []string{"apple", "grape"} // slice
	// var fruitsB = [2]string{"banana", "melon"} // array
	// var fruitsC = [...]string{"papaya", "grape"} // array

	// Hubungan slice dengan array
	// slice adalah referensi tiap elemen array
	var newFruits = fruits[0:2]

	fmt.Println(newFruits)

	// slice merupakan tipe data reference
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]
	fmt.Println(fruits) // [apple grape banana melon]
	fmt.Println(aFruits) // [apple grape banana]
	fmt.Println(bFruits) // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]
	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"
	fmt.Println(fruits) // [apple pinnaple banana melon]
	fmt.Println(aFruits) // [apple pinnaple banana]
	fmt.Println(bFruits) // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	// fungsi len() untuk menghitung jumlah elemen slice
	fmt.Println(len(fruits))

	// fungsi cap() untuk menghitung lebar atau kapasitas maksimum slice
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	// fungsi append() untuk menambahkan elemen baru
	var fruits3 = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits3, "papaya")

	cFruits[0] = "mango"

	fmt.Println(fruits3)
	fmt.Println(cFruits)

	// Ketika jumlah elemen dan lebar kapasitas adalah sama ( len(fruits) == cap(fruits) ), maka elemen baru hasil append() merupakan referensi baru.
	// Ketika jumlah elemen lebih kecil dibanding kapasitas ( len(fruits) < cap(fruits) ), elemen baru tersebut ditempatkan ke dalam cakupan kapasitas, menjadikan semua elemen slice lain yang referensi-nya sama akan berubah nilainya.
	var fruits2 = []string{"apple", "grape", "banana"}
	var bFruits2 = fruits2[0:2]
	fmt.Println(cap(bFruits2)) // 3
	fmt.Println(len(bFruits2)) // 2
	fmt.Println(fruits2) // ["apple", "grape", "banana"]
	fmt.Println(bFruits2) // ["apple", "grape"]
	var cFruits2 = append(bFruits2, "papaya")
	cFruits2[0] = "mango"
	fmt.Println(fruits2) // ["mango", "grape", "papaya"]
	fmt.Println(bFruits2) // ["mango", "grape"]
	fmt.Println(cFruits2) // ["mango", "grape", "papaya"]

	// Fungsi copy() untuk men-copy elements slice pada src ke dst `copy(dst, src)`
	// Jumlah element yang di-copy dari src adalah sejumlah lebar slice dst (atau len(dst) ).
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)
	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n) // 3

	// Pengaksesan elemen slice dengan 3 indeks
	// 3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya. Cara menggunakannnya yaitu dengan menyisipkan angka kapasitas di belakang, seperti fruits[0:1:1]
	// Angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice yang akan di slicing.
	var fruits4 = []string{"apple", "grape", "banana"}
	var aFruits4 = fruits4[0:2]
	var bFruits4 = fruits4[0:2:2]
	fmt.Println(fruits4) // ["apple", "grape", "banana"]
	fmt.Println(len(fruits4)) // len: 3
	fmt.Println(cap(fruits4)) // cap: 3
	fmt.Println(aFruits4) // ["apple", "grape"]
	fmt.Println(len(aFruits4)) // len: 2
	fmt.Println(cap(aFruits4)) // cap: 3
	fmt.Println(bFruits4) // ["apple", "grape"]
	fmt.Println(len(bFruits4)) // len: 2
	fmt.Println(cap(bFruits4)) // cap: 2
}

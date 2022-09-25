package main

import "fmt"

func main() {
	// Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair

	// Contoh penggunaan map
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	// Inisialisasi nilai map
	// variable bertipe map harus diinisialisasi secara explisit nilai awalnya
	var data map[string]int

	/*
	data["one"] = 1
	*/
	// akan muncul error

	data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	// cara lain inisialisasi map
	// cara horizonatal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertikal
	var chicken2 = map[string]int{
		"januari": 50,
		"februari": 40,
	}

	fmt.Println(chicken1)
	fmt.Println(chicken2)

	// Inisialisasi map tanpa nilai awal
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)

	fmt.Println(chicken3)
	fmt.Println(chicken4)
	fmt.Println(chicken5)

	// Iterasi item map menggunakan for - range
	var chicken6 = map[string]int{
		"januari": 50,
		"februari": 49,
		"maret": 34,
		"april": 67,
	}

	for key, val := range chicken6 {
		fmt.Println(key,"\t:", val)
	}

	// menghapus item map dengan fungsi delete
	var chicken7 = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chicken7))
	fmt.Println(chicken7)
	delete(chicken7, "januari")

	fmt.Println(len(chicken7)) // 1
	fmt.Println(chicken7)

	// Deteksi keberadaan item dengan key tertentu
	// Item map sebenarnya mengembailkan 2 nilai, nilai yang ke-2 bersifat opsional yang menyatakan ada atau tidaknya item yang dicari
	var value, isExist = chicken["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// Kombinasi slice dan map
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	// cara lain
	var chickens2 = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens2 {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	// tiap elemen bisa saja memiliki key yang berbeda-beda
	var data2 = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	fmt.Println(data2)

}

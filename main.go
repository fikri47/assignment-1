package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var argument = os.Args[1]

	fmt.Println("No ID:", argument)

	noID, _ := strconv.Atoi(argument)

	var people = []Person{

		{Nama: "Ilham", Alamat: "Kota Jakarta", Pekerjaan: "Mahasiswa", Alasan: "Hobi Mempelajari bahasa pemrograman"},
		{Nama: "Taufik", Alamat: "Kota Bandung", Pekerjaan: "Mahasisawa", Alasan: "Karena Golang itu mudah"},
		{Nama: "Usep", Alamat: "Kota Banjarmasin", Pekerjaan: "Mahasiswa", Alasan: "Tuntutan Zaman"},
		{Nama: "Sholeh", Alamat: "Kota Depok", Pekerjaan: "Software Engineer", Alasan: "Karena Golang Cepat"},
		{Nama: "Nasrul", Alamat: "Kota Garut", Pekerjaan: "Mahasiswa", Alasan: "Golang itu terstruktur"},
	}

	var result Person = CariData(noID, people)

	if noID > len(people) {
		fmt.Printf("No ID:%d tidak ada, jumlah data hanya ada %d", noID, len(people))
	} else {
		fmt.Println("Nama\t:", result.Nama)
		fmt.Println("Alamat\t:", result.Alamat)
		fmt.Println("Pekerjaan:", result.Pekerjaan)
		fmt.Println("Alasan\t:", result.Alasan)
	}
}

func CariData(noId int, data []Person) Person {
	var result Person

	for index, value := range data {
		if noId == index+1 {
			result = value
		}
	}
	return result
}

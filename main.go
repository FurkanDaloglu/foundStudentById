package main

import (
	"fmt"
	"strconv"
)

type student struct {
	firstName string
	lastName  string
	id        int
}

func studentById(students []student, aranacakId int) *student {

	for _, student := range students {
		if student.id == aranacakId {
			return &student
		}
	}
	return nil //öğrenci bulamazsa

}

func main() {
	students := []student{
		{firstName: "Uğur",
			lastName: "özyılmazel",
			id:       1},
		{firstName: "Furkan",
			lastName: "Daloğlu",
			id:       2},
		{firstName: "Mustafa",
			lastName: "Soyad1",
			id:       3},
		{firstName: "Ahmet",
			lastName: "Soyad2",
			id:       4},
		{firstName: "Mehmet",
			lastName: "Soyad2",
			id:       5},
		{firstName: "Merve",
			lastName: "Soyad3",
			id:       6},
		{firstName: "Melisa",
			lastName: "Soyad4",
			id:       7},
		{firstName: "Betül",
			lastName: "Soyad5",
			id:       8},
	}
	fmt.Println("Öğrenci Bilgileri=")
	for _, student := range students {
		fmt.Printf("Id= %d, Ad: %s, Soyad: %s\n", student.id, student.firstName, student.lastName)
	}

	var arananId int
	for {
		fmt.Printf("Bulmak istediğiniz öğrenci idsini giriniz:")
		var istenenId string
		fmt.Scanln(&istenenId)

		id, err := strconv.Atoi(istenenId)
		if err != nil {
			fmt.Println("Lütfen bir id giriniz.")
			continue
		}

		sonuc := false
		for _, student := range students {
			if student.id == id {
				arananId = id
				sonuc = true
				break
			}
		}
		if sonuc == true {
			break
		} else {
			fmt.Println("Bu idye sahip bir kayıt bulunmamaktadır.")
		}
	}

	/* 	fmt.Print("Bulmak istediğiniz öğrenci idsini giriniz:")
	   	var istenenId string
	   	fmt.Scanln(&istenenId)

	   	aranacakId, err := strconv.Atoi(istenenId)
	   	if err != nil {
	   		fmt.Println("Lütfen geçerli bir id giriniz.")
	   		//os.Exit(1)
	   	} */

	arananOgrenci := studentById(students, arananId)
	if arananOgrenci != nil {
		fmt.Printf("\nAradığınız %d numaralı öğrenci: %s %s\n", arananId, arananOgrenci.firstName, arananOgrenci.lastName)
	} else {
		fmt.Printf("\n%d numaralı öğrenci sistemimizde kayıtlı değil\n", arananId)
	}

}

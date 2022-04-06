package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func status(typeIdentity []string, numberIdentity []string) {
	length := len(typeIdentity)
	no := 1

	if cap(typeIdentity) != 0 {
		fmt.Println("No loker\tTipe Identitas\tNo Identitas")
		for i := 0; i < length; i++ {
			fmt.Printf("%d\t\t%s \t\t%s \n", no, typeIdentity[i], numberIdentity[i])
			no++
		}
	} else {
		fmt.Println("Data loker kosong. Masukkan terlebih dahulu dengan perintah 'init [jumlah loker]'")
	}
}

func insert(size int) ([]string, []string) {
	var typeIdentity = make([]string, size)
	var numberIdentity = make([]string, size)
	no := 1

	for i := 0; i < size; i++ {
		fmt.Printf("input ")
		fmt.Scanf("%s %s", &typeIdentity[i], &numberIdentity[i])
		fmt.Printf("Kartu identitas tersimpan di loker nomor %d \n", no)
		no++
	}
	return typeIdentity, numberIdentity
}

func add(typeIdentity []string, numberIdentity []string, addType string, addNumber string) {
	length := len(typeIdentity)
	no := 1
	res := false
	for i := 0; i < length; i++ {
		if numberIdentity[i] == "" {
			typeIdentity[i] = addType
			numberIdentity[i] = addNumber
			no = no + i
			res = true
			break
		}
	}
	if res {
		fmt.Printf("Kartu identitas tersimpan di loker nomor %d \n", no)
	} else {
		fmt.Printf("Maaf Loker sudah penuh \n")
	}

}

func leave(typeIdentity []string, numberIdentity []string, index int) {
	typeIdentity[index-1] = ""
	numberIdentity[index-1] = ""
	fmt.Printf("Loker nomor %d berhasil dikosongkan\n", index)
}

func find(numberIdentity []string, number string) {
	length := len(numberIdentity)
	no := 1
	res := false
	for i := 0; i < length; i++ {
		if numberIdentity[i] == number {
			res = true
			no = no + i
		}
	}
	if res {
		fmt.Printf("Kartu identitas tersebut berada di loker nomor %d \n", no)
	} else {
		fmt.Printf("Nomor identitas tidak ditemukan \n")
	}
}

func search(typeIdentity []string, numberIdentity []string, typeId string) {
	length := len(typeIdentity)
	var checkIdentity []string
	for i := 0; i < length; i++ {
		if typeIdentity[i] == typeId {
			checkIdentity = append(checkIdentity, numberIdentity[i])
		}
	}
	fmt.Println(strings.Join(checkIdentity, ", "))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var typeIdentity []string
	var numberIdentity []string

	for {
		scanner.Scan()
		str := scanner.Text()

		if strings.Contains(str, "status") {
			status(typeIdentity, numberIdentity)
		} else if strings.Contains(str, "init") {
			cmd := strings.Split(str, " ")
			if cap(cmd) == 1 {
				fmt.Println("Masukkan jumlah loker!")
			} else {
				index := cmd[1]
				size, err := strconv.Atoi(index)
				if err == nil {
					typeIdentity, numberIdentity = insert(size)
				}
			}
		} else if strings.Contains(str, "leave") {
			cmd := strings.Split(str, " ")
			index := cmd[1]
			no, err := strconv.Atoi(index)
			if err == nil {
				leave(typeIdentity, numberIdentity, no)
			}
		} else if strings.Contains(str, "input") {
			cmd := strings.Split(str, " ")
			addType := cmd[1]
			addNumber := cmd[2]
			add(typeIdentity, numberIdentity, addType, addNumber)
		} else if strings.Contains(str, "find") {
			cmd := strings.Split(str, " ")
			number := cmd[1]
			find(numberIdentity, number)
		} else if strings.Contains(str, "search") {
			cmd := strings.Split(str, " ")
			typeId := cmd[1]
			search(typeIdentity, numberIdentity, typeId)
		} else if strings.Contains(str, "exit") {
			os.Exit(1)
		} else {
			fmt.Println("Perintah tidak ditemukan!")
		}
	}
}

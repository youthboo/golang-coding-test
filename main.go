package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func findUnionAndDifference(arr1, arr2 []string) ([]string, []string) {
	set := make(map[string]bool)
	difference := []string{}

	for i := 0; i < len(arr1) || i < len(arr2); i++ {
		if i < len(arr1) {
			set[arr1[i]] = true
		}
		if i < len(arr2) {
			set[arr2[i]] = true
		}
	}

	// หาค่าที่อยู่ใน arr1 แต่ไม่มีใน arr2 และค่าที่อยู่ใน arr2 แต่ไม่มีใน arr1
	for i := 0; i < len(arr1); i++ {
		if !contains(arr2, arr1[i]) {
			difference = append(difference, arr1[i])
		}
	}
	for i := 0; i < len(arr2); i++ {
		if !contains(arr1, arr2[i]) {
			difference = append(difference, arr2[i])
		}
	}

	union := make([]string, 0, len(set))
	for key := range set {
		union = append(union, key)
	}

	sort.Strings(union)
	sort.Strings(difference)

	return union, difference
}

func contains(arr []string, target string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}

// รับค่าจากผู้ใช้
func getInputArray(prompt string) []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.Split(strings.TrimSpace(input), ",")
}

func main() {
	arr1 := getInputArray("กรุณาใส่ค่า array 1 (คั่นด้วย comma เช่น a,b,c): ")
	arr2 := getInputArray("กรุณาใส่ค่า array 2 (คั่นด้วย comma เช่น b,c,d): ")

	union, uniqueDifference := findUnionAndDifference(arr1, arr2)

	fmt.Println("\nผลลัพธ์:")
	fmt.Println("รวมข้อมูลโดยไม่มีข้อมูลซ้ำ:", "["+strings.Join(union, ",")+"]")
	fmt.Println("รวมข้อมูลที่รับเข้าด้วยกันโดยตัดข้อมูลซ้ำ:", "["+strings.Join(uniqueDifference, ",")+"]")
}

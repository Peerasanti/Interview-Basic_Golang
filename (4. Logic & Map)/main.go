package main

import (
	"fmt"
)

/*
4. Logic & Map: หาคู่ตัวเลข (Two Sum)
โจทย์: กำหนดให้มี Slice ของตัวเลข nums := []int{2, 7, 11, 15} และค่าเป้าหมาย target := 9
จงเขียนฟังก์ชันที่รับ nums และ target แล้วคืนค่า index ของตัวเลข 2 ตัวใน slice ที่บวกกันแล้วได้เท่ากับ target พอดี

สมมติว่ามีคำตอบที่ถูกต้องแน่นอนเพียง 1 คู่
เงื่อนไข: ห้ามใช้Loop ซ้อน Loop (Double for-loop) ต้องใช้วิธีที่มีประสิทธิภาพ Time Complexity ดีกว่า O(n^2)
*/

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(TwoSum(nums, target))
}

func TwoSum(nums []int, target int) []int {
	var result []int

	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if _, ok := numMap[complement]; ok {
			result = []int{numMap[complement], i}
			break
		}
		numMap[num] = i
	}

	return result
}

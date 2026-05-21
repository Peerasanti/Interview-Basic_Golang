package main

import (
	"fmt"
	"sync"
)

/*
2. Thread-Safety: Safe Counter
โจทย์: จงสร้าง struct ชื่อ SafeCounter ที่ภายในเก็บค่า count (int) และมี Methods 2 ตัวดังนี้:
Inc(): สำ หรับเพิ่มค่า count ทีละ 1
Value(): สำหรับคืนค่า count ปัจจุบันออกมำ
เงื่อนไข: โค้ดนี้ต้องรองรับกรณีที่มี Goroutines จำนวนมำก(เช่น 1,000 routines) เรียกใช้ Inc() พร้อมกัน
โดยที่ค่าสุดท้ายต้องถูกต้องแม่นยำ ห้ามเกิด Race Condition
*/

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (sc *SafeCounter) Inc() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.count++
}

func (sc *SafeCounter) Value() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

func main() {
	sc := SafeCounter{}
	numRoutines := 1000
	result := IncreaseNumber(&sc, numRoutines)

	fmt.Printf("count = %d\n", result)
}

func IncreaseNumber(sc *SafeCounter, numRoutines int) int {
	var wg sync.WaitGroup
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sc.Inc()
		}()
	}

	wg.Wait()
	return sc.Value()
}

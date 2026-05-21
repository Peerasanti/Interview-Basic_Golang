package main

import (
	"fmt"
	"sync"
	"time"
)

/*
1. Basic Concurrency: Worker Pool
โจทย์: จงเขียนฟังก์ชันชื่อ RunWorkers โดยรับ parameter 2 ตัวคือ numWorkers (จำนวนคนงาน) และ numJobs (จำนวนงานทั้งหมด)
ให้สร้ำง Workers จำนวน numWorkers ตัว ที่ทำงานพร้อมๆ กัน (Concurrent)
แต่ละ Worker จะคอยรับงำนจำก Channel แล้วพิมพ์ข้อควำมว่ำ "Worker [ID] processing job [JobID]"
ให้จำลองกำรทำงานด้วย time.Sleep เล็กน้อย(เช่น 1 วินำที)
Main function ตอ้งรอให้ทุกงำนถูกทำ จนเสร็จสิ้นจริงๆ ก่อนถึงจะจบโปรแกรม
*/

func main() {
	numWorkers := 3
	numJobs := 10
	RunWorkers(numWorkers, numJobs)
}

func RunWorkers(numWorkers int, numJobs int) {
	var wg sync.WaitGroup
	jobs := make(chan int, numJobs)
	for worker := 1; worker <= numWorkers; worker++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker [%d] processing job [%d]\n", workerID, job)
				time.Sleep(time.Second)
			}
		}(worker)
	}

	for job := 1; job <= numJobs; job++ {
		jobs <- job
	}
	close(jobs)

	wg.Wait()
	fmt.Println("All jobs have been processed.")
}

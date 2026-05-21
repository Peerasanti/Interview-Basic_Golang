package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
5. HTTP Handler: สร้าง JSON API ง่ายๆ
โจทย์: จงใช้ package net/http เขียน Web Server ที่รันบน port 8080 และมีendpoint ชื่อ/hello
โดยมีเงื่อนไขดังนี้: ต้องรับ Request Method เป็น POST เท่านั้นรับ Body เป็น JSON: {"name": "Somchai"}
แกะค่า name ออกแล้วตอบกลับ (Response) เป็น JSON: {"message": "Hello Somchai"}
กรณีที่ Method ไม่ใช่ POST หรือส่ง JSON ผิด format ให้return HTTP Error Code ที่เหมาะสมกลับไป
*/

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req HelloRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Bad Request: Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := HelloResponse{Message: fmt.Sprintf("Hello %s", req.Name)}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server is running on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		log.Fatal(err)
	}
}

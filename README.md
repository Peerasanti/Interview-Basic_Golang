# Basic Golang Assessment

Repository นี้รวบรวมคำตอบสำหรับแบบทดสอบ Basic Golang โดยครอบคลุมหัวข้อพื้นฐานสำคัญของภาษา Go เช่น Concurrency, Thread Safety, Interface, Algorithm และ HTTP API

## ความต้องการของระบบ

- Go 1.20 ขึ้นไป (1.26.3)
- Git

## โครงสร้างโปรเจกต์

```text
Interview-Basic_Golang/
├── go.mod
├── README.md
│
├── 1.Basic-Concurrency/
│   └── main.go
│
├── 2.Thread-Safety/
│   ├── main.go
│   └── main_test.go
│
├── 3.Interfaces/
│   ├── main.go
│   └── main_test.go
│
├── 4.Logic-and-Map/
│   ├── main.go
│   └── main_test.go
│
└── 5.HTTP-Handler/
    ├── main.go
    └── main_test.go
```

## รายละเอียดโจทย์

### 1. Basic Concurrency: Worker Pool

สร้าง Worker Pool โดยใช้:

- Goroutines
- Channels
- WaitGroup

คุณสมบัติ:

- รองรับ Worker หลายตัวทำงานพร้อมกัน
- จำลองเวลาการทำงานด้วย `time.Sleep()`
- รอให้ทุกงานทำเสร็จก่อนจบโปรแกรม

---

### 2. Thread Safety: Safe Counter

สร้าง Counter ที่รองรับการทำงานแบบ Concurrent โดยใช้:

- `sync.Mutex`
- Goroutines

คุณสมบัติ:

- เพิ่มค่า Counter ได้อย่างปลอดภัย
- ป้องกัน Race Condition
- รองรับการเรียกพร้อมกันจำนวนมาก

---

### 3. Interfaces: Shape Area Calculator

สร้าง:

- Interface `Shape`
- Struct `Rectangle`
- Struct `Circle`

คุณสมบัติ:

- คำนวณพื้นที่ของรูปทรงแต่ละประเภท
- ใช้งานผ่าน Interface

---

### 4. Logic & Map: Two Sum

แก้ปัญหา Two Sum โดยใช้:

- Hash Map (`map[int]int`)

Complexity:

- Time Complexity: `O(n)`
- Space Complexity: `O(n)`

---

### 5. HTTP Handler: JSON API

สร้าง API โดยใช้:

- `net/http`
- JSON Encoding/Decoding

Endpoint:

```http
POST /hello
```

Request:

```json
{
    "name":"Somchai"
}
```

Response:

```json
{
    "message":"Hello Somchai"
}
```

รองรับกรณี:

- Method ไม่ถูกต้อง → `405 Method Not Allowed`
- JSON ไม่ถูกต้อง/ผิด format → `400 Bad Request`
- JSON ว่าง → `400 Bad Request`
- Name ว่าง → `400 Bad Request`
- Field เกิน → `400 Bad Request`

---

## วิธีรันโปรเจกต์

รันเฉพาะโจทย์ที่ต้องการ:

```bash
go run 5.HTTP-Handler/main.go
```

ตัวอย่างการเรียก API:

```bash
curl -X POST http://localhost:8080/hello \
-H "Content-Type: application/json" \
-d '{"name":"Somchai"}'
```

Response:

```json
{
    "message":"Hello Somchai"
}
```

---

## วิธีทดสอบ

รัน Unit Test ทั้งหมด:

```bash
go test ./... -v
```

ตรวจสอบ Race Condition:

```bash
go test ./... -race -v
```

---
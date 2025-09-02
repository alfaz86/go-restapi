# Go REST API with Gin + GORM + PostgreSQL

Proyek ini adalah contoh REST API sederhana menggunakan **Golang** dengan framework **Gin** sebagai HTTP router dan **GORM** sebagai ORM untuk berinteraksi dengan database PostgreSQL.  
Fitur yang tersedia meliputi **User Management**, **Address**, dan **Image Upload**.

---

## 📌 1. Ringkasan Golang Gin
[Gin](https://github.com/gin-gonic/gin) adalah framework web untuk Golang yang ringan dan cepat, cocok untuk membuat REST API.  
Beberapa kelebihan Gin:
- Router berbasis `httprouter` (sangat cepat).
- Middleware built-in untuk logging, recovery, CORS, dll.
- Mudah digunakan untuk JSON binding & response.

---

## 📌 2. Setup & Instalasi

### Prasyarat
- Go `>= 1.21`
- PostgreSQL `>= 14`
- Git

### Cara Clone & Install
```bash
# Clone repository
git clone https://github.com/alfaz86/go-restapi.git
cd go-restapi

# Install dependency
go mod tidy

# Setting .env file
cp .env.example .env

# Jalankan migration (buat tabel di database)
go run database/migrations/migrate.go

# Jalankan aplikasi
go run main.go
```

## 📌 3. Dokumentasi API

### 🔹 Users

#### Create User
- **Endpoint:** `POST /users`
- **Body (JSON):**
```json
{
  "name": "Alfaz",
  "email": "alfaz@example.com",
  "password": "123456"
}
```
- **Response:**
```json
{
  "ID": 1,
  "name": "Alfaz",
  "email": "alfaz@example.com",
  "password": "123456"
}
```
#### Get All Users
- **Endpoint:** `GET /users`
- **Response:**
```json
[
    {
        "ID": 1,
        "name": "Alfaz",
        "email": "alfaz@example.com"
    },
    {
        "ID": 2,
        "name": "Budi",
        "email": "budi@example.com"
    }
]
```

### 🔹 Addresses

#### Create Address
- **Endpoint:** `POST /addresses`
- **Body (JSON):**
```json
{
    "user_id": 1,
    "street": "Jl. Merdeka",
    "city": "Bandung",
    "country": "Indonesia",
}
```
- **Response:**
```json
{
    "ID": 1,
    "user_id": 1,
    "street": "Jl. Merdeka",
    "city": "Bandung",
    "country": "Indonesia",
}
```

#### Get All Addresses
- **Endpoint:** `GET /addresses`
- **Response:**
```json
[
    {
        "ID": 1,
        "user_id": 1,
        "street": "Jl. Merdeka",
        "city": "Bandung",
        "country": "Indonesia",
    }
]
```

### 🔹 Images

#### Create Image
- **Endpoint:** `POST /images`
- **Form Data:**  
    - `user_id`: integer  
    - `image`: file (jpg/png)
- **Response:**
```json
{
    "ID": 1,
    "user_id": 1,
    "path": "/uploads/your-image.jpg"
}
```

#### Get All Images
- **Endpoint:** `GET /images`
- **Response:**
```json
[
    {
        "ID": 1,
        "user_id": 1,
        "path": "/uploads/your-image.jpg",
        "url": "http://localhost:8080/uploads/your-image.jpg"
    }
]
```
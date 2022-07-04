## Practice API

### Soal

Buatlah sebuah API untuk CRUD Product dengan ketentuan sebagai berikut:

#### Model

##### Category
```go
type Category struct {
    CategoryId string `json:"id" binding:"required"`
    CategoryName string `json:"name" binding:"required"`
}
```
##### Category Data
```json
{
    "message": "OK",
    "data": {
        "id":"C001",
        "name":"Food"
    }, 
    {
        "id":"C002",
        "name":"Drink"
    },
    {
        "id":"C003",
        "name":"Other"
    }
}
```
##### Product
```go
type Product struct {
    ProductId string `json:"id" binding:"required"`
    ProductName string `json:"name" binding:"required"`
    Category Category `json:"category" binding:"required"`
    IsActive bool `json:"is_active"`
}
```
##### Storage 
Penyimpanan sementara hanya di `slice`

##### Method
1. POST Product
2. GET Product (params: id)
3. GET Product By Category (List)
4. PUT Product Category
5. DELETE Product (dengan melakukan update is_active)

#### All Request
```go
Use json
```

#### All Response
```go
{
    "message":"string",
    "data":interface{}
}
```
`Note`: `interface{}` disini dapat menampung data sesuai yang di-request, misalnya product atau category, atau sebagai contoh berikut:
```go
{
    "message":"string",
    "data": {
        "id":"string",
        "name":"string",
        "category":{
            "id":"string",
            "name":"string"
        },
        "is_active": true
    },
    {
        "id":"string",
        "name":"string",
        "category":{
            "id":"string",
            "name":"string"
        },
        "is_active": true
    }
}
```

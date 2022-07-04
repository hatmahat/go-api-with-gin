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
##### Product
```go
type Product struct {
    ProductId string `json:"id" binding:"required"`
    ProductName string `json:"id" binding:"required"`
    Category Category
}
```
##### Storage 
Penyimpanan sementara hanya di `slice`

##### Method
1. POST Product
2. GET Product (params: id)
3. GET Product By Category (List)
4. PUT Category
5. DELETE Product (dengan melakukan update is_active)
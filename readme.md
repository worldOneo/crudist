# CRUDist - Model Driven API Development

Automagicaly create CRUD APIs for your gorm models.

## Example
### Model definition
```go
type BaseModel struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	BaseModel
	Username string `json:"username" gorm:"size:100"`
	Password string `json:"password" gorm:"size:128"`
}
```
### CRUD API
```go
crudist.Handle(c, "user/", &User{})
```

Result routs:
```sh
GET    /user/
GET    /user/:id/
POST   /user/
PATCH  /user/
DELETE /user/
```
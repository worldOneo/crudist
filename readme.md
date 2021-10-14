# CRUDist - Model Driven API Development
[![GoReport](https://goreportcard.com/badge/github.com/worldOneo/crudist)](https://goreportcard.com/report/github.com/worldOneo/crudist)  
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

Result routes:
```sh
GET    /user/
GET    /user/:id/
POST   /user/
DELETE /user/
DELETE /user/:id/
PATCH  /user/
```
# CRUDist - Model Driven API Development
[![GoReport](https://goreportcard.com/badge/github.com/worldOneo/crudist)](https://goreportcard.com/report/github.com/worldOneo/crudist)  
Automagicaly create CRUD APIs for your data models.

## Currently supported
To get support for your favourite Web or Model Framework open an issue.
### Web Frameworks:
  * [x] Fiber
  * [x] Gin

### Model Frameworks:
  * [x] GORM

## Example (gorm)
_(Full working examples in demo folder)_
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
	Password string `json:"-" gorm:"size:128"`
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
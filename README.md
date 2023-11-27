# gin-pagination
gin-pagination

## Install
```
go get github.com/xusenlin/gin-pagination@v0.0.1-alpha
```

## Usage

### init
```go
import (
    ginPagination "github.com/xusenlin/gin-pagination"
)
//...Conn is *gorm.DB
ginPagination.Init(&ginPagination.Config{
    PageSizeDefaultVal: "10",
    DB:                 Conn,
})
```
### pagination
```go
import (
    ginPagination "github.com/xusenlin/gin-pagination"
)
//...Conn is *gorm.DB
func List(c *gin.Context) {
    
    model := new(Repository)
    
    pagination := ginPagination.New[*Repository](model, c)
    
    pagination.Eq("id").Like("name")
    pagination.Eq("id").Like("name")
    pagination.Gt("age").Lt("fff")
    pagination.CB(func(db *gorm.DB) {
        db.Where("name2 IN ?", []string{"jinzhu", "jinzhu"})
    })
    
    err := pagination.Query()
    
    if err != nil {
        tools.SendErrJson(c, err)
        return
    }
	//pagination.List is []Repository
    if len(pagination.List) >0{
        fmt.Println(pagination.List[0].Name)
    }
    
    tools.SendOkJson(c, "", pagination)
}
```

## Pagination struct
```go
type Pagination[T any] struct {
	List      []T   `json:"list"`
	Total     int64 `json:"total"`
	PageNum   int   `json:"pageNum"`
	PageSize  int   `json:"pageSize"`
	TotalPage int   `json:"totalPage"`
	query     *gorm.DB
	ctx       *gin.Context
}
```

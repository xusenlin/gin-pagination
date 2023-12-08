# gin-pagination
Pagination for Gin and Gorm

## Install
```
go get github.com/xusenlin/gin-pagination@v0.0.3-alpha
```

## Usage

### init
```golang
import (
    ginPagination "github.com/xusenlin/gin-pagination"
)
//...Conn is *gorm.DB
ginPagination.Init(&ginPagination.Config{
    PageSizeMaxVal:     100,
    PageSizeDefaultVal: 20,
    DB:                 Conn,
})
```
### pagination
```golang

func Find(c *gin.Context) {
    
    model := new(User)
    pagination := ginPagination.New[*User](model, c)
    
    pagination.Like("name").Eq("id") //...more
    
    err := pagination.Query()
    
    if err != nil {
        tools.SendErrJson(c, err)
        return
    }
    
    for idx := range pagination.List {
        pagination.List[idx].Password = "***"
    }
    
    tools.SendOkJson(c, "search successful", pagination)
}
```

## Pagination struct
```golang
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

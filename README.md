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
import (
    ginPagination "github.com/xusenlin/gin-pagination"
)

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

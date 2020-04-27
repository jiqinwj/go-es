package gg

import (
	"context"
	"go-es/AppInit"
	"reflect"
)
func MapToLogs(rsp *elastic.SearchResult) []*LogModel  {
	ret:=[]*LogModel{}
	var t *LogModel
	for _,item:=range rsp.Each(reflect.TypeOf(t)){
		ret=append(ret,item.(*LogModel))
	}
	return ret
}
type LogService struct {
}
func NewLogService() *LogService  {
	return &LogService{}
}
func(this *LogService) Search() ([]*LogModel,error)  {
	rsp,err:=AppInit.GetEsClient().Search().Index("bookslogs").Do(context.Background())
	return MapToLogs(rsp),err
}
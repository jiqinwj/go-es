package AppInit

import (
	"github.com/olivere/elastic/v7"
)

func GetEsClient() *elastic.Client  {
	client,err:=elastic.NewClient(
		elastic.SetURL("http://192.168.199.219:9200/"),
		elastic.SetSniff(false),
	)

	if err!=nil{
		return nil
	}
	return client

}
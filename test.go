package main

import "go-es/logparser"

func main()  {
    //p:=logparser.NewHttpdParser("./logs/apache_log.txt")
    //list:=p.Parse()
	//
	//client:=AppInit.GetEsClient()
	//bulk:=client.Bulk()
	////fmt.Print(list)
	//for _,m:=range list{
	//
	//	fmt.Print(m)
	//	req:=elastic.NewBulkIndexRequest()
	//	req.Index("bookslogs").Doc(m)//直接插入
	//	bulk.Add(req)
	//}
	//_,err:=bulk.Do(context.Background())
	//if err!=nil {
	//	log.Println(err)
	//}

	p:=logparser.NewHttpdParser();
	p.ParseToEs("./logs/apache_log_2.txt")


}


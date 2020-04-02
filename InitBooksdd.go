package main

import (
	"fmt"
	//"github.com/deckarep/golang-set"
	"github.com/satori/go.uuid"
	"github.com/syyongx/php2go"
	"go-es/AppInit"

	//"go-es/AppInit"
	"reflect"
	"strconv"
	"sync"

	//"go-es/AppInit"
	"hash/crc32"
	"math/rand"
	"time"

	//"log"
)

func main() {

	redis := AppInit.Redis

	wg := sync.WaitGroup{}
	//electiveClasses := mapset.NewSet()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 80000; i++ {

				a := uuid.NewV4().String();
				b:=crc3(a);
				//
				////c,_:=strconv.ParseInt(b, 10, 64);
				//
				//
				////fmt.Println(a);
				//fmt.Println(b);






				//fmt.Println(reflect.TypeOf(c));

				err := redis.SetBit("red", int64(b), 1).Err()
				if err != nil {
					panic(err)
				}
				fmt.Println("插入成");
				fmt.Println(b);

				//electiveClasses.Add(b)
			}
		}()
	}
	wg.Wait()

	// strin:=electiveClasses.String()
	//
	//arr:=strings.Split(strin,",")
	//
	//fmt.Println(len(arr))
	//page:=1
	//pagesize:=500
	//wg:=sync.WaitGroup{}
	//for{
	//	book_list:=Models.BookList{}
	//	db:=AppInit.GetDB().Select("book_id,book_name,book_intr,book_price1,book_price2,book_author,book_press,book_kind " +
	//		",if(book_date='','1970-01-01',ltrim(book_date)) as book_date").
	//		Order("book_id desc").Limit(pagesize).Offset((page-1)*pagesize).Find(&book_list)
	//	if db.Error!=nil || len(book_list)==0{
	//		break
	//	}
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		client:=AppInit.GetEsClient()
	//		bulk:=client.Bulk()
	//		for _,book:=range book_list {
	//			req:=elastic.NewBulkIndexRequest()
	//			req.Index("books").Id(strconv.Itoa(book.BookID)).Doc(book)
	//			bulk.Add(req)
	//		}
	//		rsp,err:=bulk.Do(context.Background())
	//		if err!=nil {
	//			log.Println(err)
	//		}else {
	//			fmt.Println(rsp)
	//		}
	//	}()
	//  page=page+1  //必须有
	//}
	//	wg.Wait()

}

func crc3(l string) uint32 {

	return crc32.ChecksumIEEE([]byte(l));
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func ord1(ls string) string {

	var res string;
	for i := 0; i < len(ls); i++ {

		qu := php2go.Substr(ls, uint(i), 1)

		lei := reflect.TypeOf(qu)

		if (lei.Name() == "string") {
			res += strconv.Itoa(php2go.Ord(qu));
		} else {
			res += qu;
		}

	}
	return res;

}

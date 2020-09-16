package main

import(
	"fmt"
	// "github.com/labstack/echo/v4"
	// "github.com/labstack/echo/middleware"
	"time"
	"os"
	// "local.packages/DB"
	"github.com/gomodule/redigo/redis"
	"bufio"
	)
//
// type Table1 struct {
//   Id       int    `json:id sql:AUTO_INCREMENT`
//   Name     string  `json:name`
//   Item     string `json:item`
// }

func Connection() redis.Conn{
	const Addr = "benkyoukai-redis:6379"

	c, err := redis.Dial("tcp", Addr)
    if err != nil {
        panic(err)
    }
    return c
}

func main() {
	c := Connection()

	defer c.Close()

	userName := os.Args[1] //コマンドラインパラメータを受け取り
	userKey := "online." + userName

	val,err := c.Do("SET",userKey,userName,"NX","EX","120")//NX→keyが存在してないならSET EX→keyをTTLつきで設定
	if val == nil{
		fmt.Println("すでにオンラインです")
		os.Exit(1)
	}

	val,err = c.Do("SADD","users",userName)//SADD→メンバーを追加
	tickerChan := time.NewTicker(time.Second * 60).C//周期的なトリガーが設定

	subChan := make(chan string)

	go func(){
		subConn := Connection()
		if err !=nil{
			fmt.Println(err)
			os.Exit(1)
		}
		defer subConn.Close()


		psc := redis.PubSubConn{Conn:subConn}
		psc.Subscribe("messages")

	// ここもよくわかってないから注意
		for{
			switch v := psc.Receive().(type){
			case redis.Message:
				subChan <- string(v.Data)
			case redis.Subscription:
				break
			case error:
				return
			}
		}
	}()

	sayChan := make(chan string)
	go func(){
		prompt := userName + ">"
		bio := bufio.NewReader(os.Stdin)
		for{
			fmt.Println(prompt)
			line,_,err := bio.ReadLine()
			if err != nil{
				fmt.Println(err)
				sayChan <- "/exit"
				return
			}
			sayChan <- string(line)
		}
	}()

	c.Do("PUBLISH","messages",userName + " has joined")

	chatExit := false

	for !chatExit{
		select{
			case msg := <-subChan:
				fmt.Println(msg)
			case <-tickerChan:
				val,err := c.Do("SET",userKey,userName,"XX","EX","120")
				if err != nil || val == nil{
					fmt.Println("Heartbeat set failed")
					chatExit = true
				}
			case line := <-sayChan:
				if line =="/exit"{
					chatExit = true
				} else if line =="/who"{
					names,_ := redis.String(c.Do("SMEMBERS","users"))
					for _,name := range names{
						fmt.Println(name)
					}
				} else {
						c.Do("PUBLISH","messages",userName +":"+ line)
				}
			}
		}

	c.Do("DEL",userKey)
	c.Do("SREM","users",userName)
	c.Do("PUBLISH","messages",userName + " has left")
}



	// db := DB.GormConnect()
	// db.CreateTable(&Table1{})
	//
	// eventEx := Table1{}
	// CreateUser := Table1{}
	// CreateUser.Name = "ケンちゃん"
	// CreateUser.Item = "カステラ"
	// db.Create(&CreateUser)
	// db.First(&eventEx, "id = ?", 3)
	// fmt.Println(eventEx)

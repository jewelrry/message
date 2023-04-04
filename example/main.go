package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"message"
	"message/domain"
)

func main() {
	dsn := "root:1astWeekend@tcp(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:               dsn, // DSN data source name
		DefaultStringSize: 191, // string 类型字段的默认长度
	}

	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		fmt.Printf("err:%+v \n", err)
	}

	base := domain.MessageBase{
		ReceiverRange: 1,
	}
	req := domain.MessageReq{
		MessageBase: base,
	}
	c, err := message.NewMessageClient(db)
	id, err := c.Create(req)
	fmt.Printf("id:%+v, err:%+v \n", id, err)
}

package message

import (
	"github.com/facebookgo/inject"
	"gorm.io/gorm"
	"message/domain"
	"message/service"
)

type Client struct {
	Db      *gorm.DB
	Service *service.MessageService `inject:""`
}

func NewMessageClient(db *gorm.DB) (client *Client, err error) {
	client = &Client{Db: db}

	graph := inject.Graph{}
	if err := graph.Provide(
		&inject.Object{
			Value: db,
		},
		&inject.Object{
			Value: client,
		},
	); err != nil {
		panic(err)
	}
	if err := graph.Populate(); err != nil {
		panic(err)
	}

	err = client.CheckMessageTable()
	return
}

func (c *Client) Create(req domain.MessageReq) (uint, error) {
	return c.Service.Create(req)
}

func (c *Client) Paginate(req domain.MessageReqPaginate) (ret domain.PageData, err error) {
	return c.Service.Paginate(req)
}

func (c *Client) UnreadCount(req domain.MessageScope) (count int64, err error) {
	return c.Service.UnreadCount(req)
}

func (c *Client) OperateRead(req domain.MessageReadReq) (uint, error) {
	return c.Service.OperateRead(req)
}

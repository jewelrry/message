package message

import (
	"errors"
)

func (c *Client) CheckMessageTable() error {
	if c.Db == nil {
		return errors.New("请传入数据库链接")
	}
	messageSql := `
		create table if not exists biz_message
		(
			id             bigint auto_increment
				primary key,
			source varchar(255)  null comment '消息模块来源',
			content        varchar(500)  null comment '内容',
			receiver_range int default 0 not null comment '接收者范围 1:全部 2:个人 3：某角色 4:某项目',
			receiver_id    int           null comment 'receiver_range=1时为0;receiver_range=2时为用户ID;receiver_range=3时为角色ID;receiver_range=4时为项目ID',
			created_at      datetime      null,
			constraint biz_message_id_uindex
				unique (id),
			INDEX receiver_id_index (receiver_id),
			INDEX receiver_range_index (receiver_range)
		);
`
	readSql := `
		create table if not exists biz_message_read
		(
			id         bigint auto_increment
				primary key,
			message_id int default 0 not null comment '消息ID',
			user_id    int           not null comment '用户ID',
			created_at  datetime      null,
			constraint biz_message_read_id_uindex
				unique (id),
			INDEX message_id_index (message_id),
			INDEX user_id_index (user_id)
		)
			comment '存储用户已读的消息';
`
	if err := c.Db.Exec(messageSql).Error; err != nil {
		return err
	}
	if err := c.Db.Exec(readSql).Error; err != nil {
		return err
	}
	return nil
}

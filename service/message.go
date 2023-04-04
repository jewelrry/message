package service

import (
	"message/domain"
	"message/repo"
)

type MessageService struct {
	MessageRepo     *repo.MessageRepo     `inject:""`
	MessageReadRepo *repo.MessageReadRepo `inject:""`
}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (s *MessageService) Create(req domain.MessageReq) (uint, error) {
	return s.MessageRepo.Create(req)
}

func (s *MessageService) Paginate(req domain.MessageReqPaginate) (ret domain.PageData, err error) {
	ret, err = s.MessageRepo.Paginate(req)

	if err != nil {
		return
	}

	return
}

func (s *MessageService) UnreadCount(scope domain.MessageScope) (count int64, err error) {
	count, err = s.MessageRepo.GetUnreadCount(scope)

	if err != nil {
		return
	}

	return
}

func (s *MessageService) OperateRead(req domain.MessageReadReq) (uint, error) {
	return s.MessageReadRepo.Create(req)
}

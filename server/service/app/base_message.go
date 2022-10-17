package app

import (
	"server/global"
	"server/model/app"
)

type BaseMessageService struct{}

func (*BaseMessageService) CreateBaseMessage(baseMessage app.BaseMessage) (err error) {
	err = global.GVA_DB.Create(baseMessage).Error
	return
}

func (*BaseMessageService) UpdateBaseMessage(baseMessage app.BaseMessage) (err error) {
	var base app.BaseMessage
	err = global.GVA_DB.Where("id = ?", baseMessage.ID).First(&base).Error
	if err != nil {
		return
	}

	err = global.GVA_DB.Save(baseMessage).Error
	return
}

func (*BaseMessageService) FindBaseMessage() (app.BaseMessage, error) {
	var base app.BaseMessage
	err := global.GVA_DB.Order("id desc").First(&base).Error
	return base, err
}

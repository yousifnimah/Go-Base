package seeds

import (
	"encoding/json"
	"gateway_api/App/Models"
	"gorm.io/gorm"
)

func InitAutoReply(db *gorm.DB) error {
	var AutoReplyMessage Models.MessageJSONB
	AutoReplyMessageStr, _ := json.Marshal(&AutoReplyMessage)
	AutoReply := Models.GeneralSettings{AutoReplyStatus: false, AutoReplyMessage: string(AutoReplyMessageStr)}
	return db.Create(&AutoReply).Error
}

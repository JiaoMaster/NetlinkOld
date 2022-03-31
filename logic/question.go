package logic

import (
	"NetLinkOld/dao/mysql"
	"NetLinkOld/models"
	"NetLinkOld/pkg/uuid"
	"go.uber.org/zap"
)

func GetQuestionDetail(Qid string) (que *models.Question, err error) {
	//查库
	que, err = mysql.GetQuestionDetail(Qid)
	if err != nil {
		zap.L().Error("GetQuestionDetail(Qid string) err...", zap.Error(err))
		return nil, err
	}
	return que, nil
}

func GetQuestionList(page int, amount int, ch *models.QueCh) (data []*models.QueList, err error) {
	//查库
	data, err = mysql.GetQuestionList(page, amount, ch)
	if err != nil {
		zap.L().Error("mysql.GetQuestionList(page, amount) err ", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func SendQuestion(que *models.Question) (err error) {
	que.ID, err = uuid.Getuuid()
	err = mysql.SendQuestion(que)
	if err != nil {
		zap.L().Error("mysql.SendQuestion(que) err:", zap.Error(err))
		return err
	}
	return nil
}

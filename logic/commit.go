package logic

import (
	"NetLinkOld/dao/mysql"
	"NetLinkOld/models"
)

func SendCommit(com *models.Commit) error {
	err := mysql.SendCommit(com)
	return err
}

func GetCommit(Pid string) ([]*models.Commit, error) {
	com, err := mysql.GetCommit(Pid)
	return com, err
}

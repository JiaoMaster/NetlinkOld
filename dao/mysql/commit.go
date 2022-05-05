package mysql

import (
	"NetLinkOld/models"
)

func SendCommit(com *models.Commit) error {
	sqls := `insert into commit(post_id, content, username) values (?,?,?)`
	_, err := db.Exec(sqls, com.PostId, com.Content, com.UserName)
	return err
}

func GetCommit(Pid string) (com []*models.Commit, err error) {
	sqls := `select id,username,content,create_time from commit where post_id = ? order by create_time`
	err = db.Select(&com, sqls, Pid)
	return com, err
}

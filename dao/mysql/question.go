package mysql

import (
	"NetLinkOld/models"
)

func GetQuestionDetail(Qid string) (*models.Question, error) {
	//查库
	//sql语句
	sqlStr := "select post_id, title , content, username, community_id, create_time from post where post_id=?"
	que := new(models.Question)
	//查询
	err := db.Get(que, sqlStr, Qid)

	if err != nil {
		return nil, err
	}

	//返回
	return que, nil
}

func GetQuestionList(page int, amount int, ch *models.QueCh) ([]*models.QueList, error) {
	sqlStr := ``
	if ch.Ch == 1 {
		sqlStr = `select
	post_id, title,create_time 
	from post
	where community_id = ? and location like CONCAT('%',?,'%')
    order by create_time desc
	limit ?,?
	`
		var data []*models.QueList

		err := db.Select(&data, sqlStr, ch.Ch, ch.Location, (page-1)*amount, amount)
		if err != nil {
			return nil, err
		}

		return data, nil
	} else {
		sqlStr = `select
	post_id, title,create_time 
	from post
    where community_id = ?
    order by create_time desc
	limit ?,?
	`
		var data []*models.QueList

		err := db.Select(&data, sqlStr, ch.Ch, (page-1)*amount, amount)
		if err != nil {
			return nil, err
		}

		return data, nil

	}

}

func SendQuestion(que *models.Question) error {
	//检查标题
	sqlstr := `insert into post(
	post_id,title,content,username,community_id,location)
	values(?,?,?,?,?,?)
`
	_, err := db.Exec(sqlstr, que.ID, que.Title, que.Content, que.UserName, que.CommunityID, que.Location)
	return err
}

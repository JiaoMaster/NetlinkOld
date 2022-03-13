package mysql

import (
	"NetLinkOld/models"
)

func GetQuestionDetail(Qid string) (*models.Question, error) {
	//查库
	//sql语句
	sqlStr := "select post_id, title , content, author_id, community_id, create_time from post where post_id=?"
	que := new(models.Question)
	//查询
	err := db.Get(que, sqlStr, Qid)

	if err != nil {
		return nil, err
	}

	//返回
	return que, nil
}

func GetQuestionList(page int, amount int) ([]*models.QueList, error) {
	sqlStr := `select
	post_id, title,create_time 
	from post
    order by create_time desc
	limit ?,?
	`
	var data []*models.QueList
	err := db.Select(&data, sqlStr, (page-1)*amount, amount)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SendQuestion(que *models.Question) error {
	//检查标题
	sqlstr := `insert into post(
	post_id,title,content,author_id,community_id)
	values(?,?,?,?,?)
`
	_, err := db.Exec(sqlstr, que.ID, que.Title, que.Content, que.AuthorID, que.CommunityID)
	return err
}

package mysql

import "NetLinkOld/models"

func GetQuestionDetail(Qid string) (*models.Question, error) {
	//查库
	//sql语句
	sqlStr := "select post_id, title , content, author_id, community_id, create_time from post where id=?"
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
    ORDER BY id
	limit ?,?
	`
	var data []*models.QueList
	err := db.Select(&data, sqlStr, (page-1)*amount, amount)
	if err != nil {
		return nil, err
	}
	return data, nil
}

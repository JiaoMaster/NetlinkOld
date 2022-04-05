package server

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/NetLinkOld/grpc-demo/proto"
	"time"

	"github.com/NetLinkOld/grpc-demo/bapi"
)

type QueListServer struct {
}

func NewQueListServer() *QueListServer {
	return &QueListServer{}
}

func (q *QueListServer) GetQueList(c context.Context, r *pb.GetQueListRequest) (*pb.GetQueListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8080")
	body, err := api.GetQueList(c, r.GetPage(), r.GetAmount(), r.GetCh(), r.GetLocation())
	if err != nil {
		return nil, err
	}
	type QueList struct {
		ID         int64     `json:"id,string" db:"post_id"` // 帖子id
		Title      string    `json:"title" db:"title"`
		CreateTime time.Time `json:"create_time" db:"create_time"`
	}
	type JBody struct {
		Msg  string     `json:"msg"`
		Data *[]QueList `json:"question_list"`
	}
	queList := pb.GetQueListReply{}
	err = json.Unmarshal(body, &queList)
	fmt.Println(queList.QuestionList)
	fmt.Println(&queList)
	if err != nil {
		return nil, err
	}
	return &queList, nil
}

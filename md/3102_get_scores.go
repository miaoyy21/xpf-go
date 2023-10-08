package md

import (
	"errors"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"psw/ds"
	"psw/pb"
)

type GetScores struct {
}

func (m GetScores) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.GetScoresResp{
		Rt:     SuccessCode.Int32(),
		Scores: make([]*pb.Score, 0),
		Score:  proto.Int64(0),
	}

	limit := pbMsg.(*pb.GetScores).GetLimit()
	offset := pbMsg.(*pb.GetScores).GetOffset()

	// Check Score
	if user.Score < 0 {
		respPbMsg.Rt = ScoreNotEnough.Int32()
		return respPbMsg, errors.New("user's Score is NOT Enough")
	}

	// Get Operates
	scores, err := ds.FindScoresByUserId(tx, user.ID, limit, offset)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Protobuf Operate
	for _, score := range scores {
		pbScore := &pb.Score{
			Action: proto.Int32(int32(score.Action)),
			Cost:   proto.Int64(score.Cost),
			Score:  proto.Int64(score.Score),
			At:     proto.Int32(int32(score.CreatedAt.Unix())),
		}

		respPbMsg.Scores = append(respPbMsg.Scores, pbScore)
	}

	respPbMsg.Score = proto.Int64(user.Score)
	return respPbMsg, nil
}

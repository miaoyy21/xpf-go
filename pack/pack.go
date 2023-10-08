package pack

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io"
	"math/big"
	"net/http"
	"psw/ds"
	"psw/pb"
	"runtime/debug"
	"strconv"
	"time"
)

const contentType = "application/x-protobuf"

const (
	msgNoPackSize  int = 2
	bodyPackSize       = 4
	userIdPackSize     = 8
)

type respRtMsg interface {
	GetRt() int32
}

// NewPlayloadMsg before Protobuf Message Unmarhal Failure
func NewPlayloadMsg(err error) []byte {
	if err == nil {
		logrus.Warnf("arguments is nil .")
		return nil
	}

	pbMsg := &pb.Playload{
		Msg: proto.String(fmt.Sprintf("%s %s", time.Now().Local().String(), err.Error())),
	}

	msg, err := proto.Marshal(pbMsg)
	if err != nil {
		logrus.Warnf("proto.Marshal() %s", err.Error())
		msg = bytes.NewBufferString(time.Now().GoString()).Bytes()
	}

	return msg
}

func fno() {
	var fn []byte

	xbs := bytes.Split(debug.Stack(), []byte{'\n'})
	for _, bs := range xbs {
		if len(bs) > 0 && bs[0] == '\t' {
			if bytes.Contains(bs, []byte("psw")) {
				fns := bytes.Split(bytes.TrimSpace(bs), []byte{' '})
				if len(fns) > 0 {
					fn = fns[0]
				}
			}
		}
	}

	logrus.Debug(string(fn))
}

func Pack(db *gorm.DB, ) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fno()
				logrus.Errorf("Pack() PANIC with RECOVER %s", err)
			}
		}()

		var user *ds.User

		// 设置IP
		ip := c.ClientIP()
		agent := c.Request.UserAgent()

		// 缓冲区
		buf := new(bytes.Buffer)
		_, err := io.Copy(buf, c.Request.Body)
		if err != nil {
			logrus.Warnf("io.Copy() %s", err.Error())
			c.Data(http.StatusOK, contentType, NewPlayloadMsg(err))
			return
		}

		// 获取消息编号（2字节）
		msgNo := pb.MsgNo(binary.BigEndian.Uint16(buf.Next(msgNoPackSize)))
		dealMsg, ok := dealMsgs[msgNo]
		if !ok {
			logrus.Errorf("NOT Found Message Deal with MsgNo [%d]", msgNo)
			c.Data(http.StatusOK, contentType, NewPlayloadMsg(fmt.Errorf("NOT Found Message Deal with MsgNo [%d]", msgNo)))
			return
		}

		// 消息体（4字节）
		bodySize := int(binary.BigEndian.Uint32(buf.Next(bodyPackSize)))
		msg := buf.Next(bodySize)

		// 用户ID 8字节
		if msgNo != pb.MsgNo_Msg9000 {
			userId := int64(binary.BigEndian.Uint64(buf.Next(userIdPackSize)))

			// 获取用户
			user, err = ds.FindUserById(db, userId)
			if err != nil {
				logrus.Warnf("ds.FindUserById(%d) %s", userId, err.Error())
				c.Data(http.StatusOK, contentType, NewPlayloadMsg(err))
				return
			}

			// 获取用户密钥
			key, err := ds.FindUserKeyById(db, userId)
			if err != nil {
				logrus.Warnf("ds.FindUserKeyById(%d) %s", userId, err.Error())
				c.Data(http.StatusOK, contentType, NewPlayloadMsg(err))
				return
			}

			// MD5
			m5Size, _ := buf.ReadByte()
			m5 := buf.Next(int(m5Size))
			if !bytes.EqualFold(m5, key.MD5) {
				logrus.Warnf("User's Key MD5 is NOT Equal %d != %d", m5, key.MD5)
				c.Data(http.StatusOK, contentType, NewPlayloadMsg(err))
				return
			}

			// 是否使用数字签名
			sig, _ := buf.ReadByte()
			if sig == 1 {
				// 使用ECDSA对消息进行认证
				x, _ := new(big.Int).SetString(key.ECX, 62)
				y, _ := new(big.Int).SetString(key.ECY, 62)
				pub := &ecdsa.PublicKey{
					Curve: elliptic.P256(),
					X:     x,
					Y:     y,
				}

				// ECDSA签名 R
				rSize, _ := buf.ReadByte()
				bsR := buf.Next(int(rSize))
				r, _ := new(big.Int).SetString(string(bsR), 10)

				// ECDSA签名 S
				sSize, _ := buf.ReadByte()
				bsS := buf.Next(int(sSize))
				s, _ := new(big.Int).SetString(string(bsS), 10)

				h := sha256.New()
				h.Write(msg)
				res := h.Sum(nil)

				// 认证是否成功
				if ok := ecdsa.Verify(pub, res, r, s); !ok {
					logrus.Warn("ec.Verify() failure")
					c.Data(http.StatusOK, contentType, NewPlayloadMsg(errors.New("ec.Verify() failure")))
					return
				}
			}

			// 跟登录时间不同时，赠送每日积分
			if !ds.IsSameDay(user.LoginAt, time.Now()) {
				score := user.Score + ds.LoginScore
				err := ds.CreateScore(db, userId, ds.LoginScoreAction, ds.LoginScore, score)
				if err != nil {
					logrus.Warnf("ds.CreateScore(%d) %s", userId, err.Error())
					c.Data(http.StatusOK, contentType, NewPlayloadMsg(err))
					return
				}

				// 保存用户积分
				user.Score = score
				user.LoginAt = time.Now()
				if err := db.Save(user).Error; err != nil {
					logrus.Warnf("ds.SaveUser(%d) %s", userId, err.Error())
					c.Data(http.StatusOK, contentType, NewPlayloadMsg(err))
					return
				}
			}
		}

		// 平台信息（1字节）
		bsOsSize, _ := buf.ReadByte()
		osSize := int(bsOsSize)
		os := string(buf.Next(osSize))

		// 平台信息（2字节）
		deviceSize := int(binary.BigEndian.Uint16(buf.Next(2)))
		device := string(buf.Next(deviceSize))

		// 根据消息号，获取对应的消息对象
		reqPbMsgOrg, ok := reqMsgs[msgNo]
		if !ok {
			logrus.Warnf("invalid MsgNo %d", msgNo)
			c.Data(http.StatusOK, contentType, NewPlayloadMsg(fmt.Errorf("invalid MsgNo %d", msgNo)))
			return
		}

		// 消息拷贝
		reqPbMsg := reqPbMsgOrg.ProtoReflect().New().Interface()

		// 消息体解析
		if err := proto.Unmarshal(msg, reqPbMsg); err != nil {
			logrus.Warnf("proto.Unmarshal() %s", err.Error())
			c.Data(http.StatusOK, contentType, NewPlayloadMsg(err))
			return
		}

		c.Request.UserAgent()

		if msgNo == pb.MsgNo_Msg2102 {
			logrus.Debugf("Request %d Size is %d", msgNo, len(msg))
		} else {
			logrus.Debugf("Request %d :: %s", msgNo, reqPbMsg)
		}

		// 事务
		var ope string
		tx := db.Begin()

		// 消息处理
		respPbMsg, err := dealMsg.DealWith(tx, user, reqPbMsg)
		if err != nil {
			tx.Rollback()

			ope = err.Error()
			logrus.Errorf("proto.Marshal() %s", err.Error())
		} else {
			tx.Commit()
		}

		// Implement method of [ GetRt() int32 ]
		rtPbMsg, ok := respPbMsg.(respRtMsg)
		if ok {
			var userId int64
			if msgNo == pb.MsgNo_Msg9000 {
				userId = respPbMsg.(*pb.RegisterResp).GetUserId()
			} else {
				userId = user.ID
			}

			err := ds.CreateOperate(db, userId, msgNo, []string{}, rtPbMsg.GetRt(), ope, ip, os, device, agent)
			if err != nil {
				logrus.Warnf("ds.CreateOperate() %s", err.Error())
			}
		} else {
			// Every Protobuf Request Message Should have field of 'rt'
			logrus.Errorf("%s NOT implement method of [ GetRt() int32 ] ", respPbMsg)
			return
		}

		// 消息返回
		respMsg, err := proto.Marshal(respPbMsg)
		if err != nil {
			logrus.Warnf("proto.Marshal() %s", err.Error())
			c.Data(http.StatusOK, contentType, NewPlayloadMsg(err))
			return
		}

		if msgNo == pb.MsgNo_Msg9002 {
			logrus.Debugf("Response %d Size is %d", msgNo, len(respMsg))
		} else {
			logrus.Debugf("Response %d Size is %d :: %s", msgNo, len(respMsg), respPbMsg)
		}

		c.Writer.Header().Set("content-length", strconv.Itoa(len(respMsg)))
		c.Data(http.StatusOK, contentType, respMsg)
	}
}

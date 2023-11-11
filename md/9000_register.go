package md

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io"
	"xpf/cache"
	"xpf/ds"
	"xpf/pb"
)

type Register struct {
}

func (m Register) DealWith(tx *gorm.DB, _ *ds.User, _ proto.Message) (proto.Message, error) {
	respPbMsg := &pb.RegisterResp{
		Rt: SuccessCode.Int32(),
	}

	// RSA 非对称加密算法
	rsaPriKey, rsaPubKey, err := genRSAKey()
	if err != nil {
		respPbMsg.Rt = InternalErrorCode.Int32()
		return respPbMsg, err
	}

	// AES 高级加密标准
	aesKey, aesIV, err := genAESKey()
	if err != nil {
		respPbMsg.Rt = InternalErrorCode.Int32()
		return respPbMsg, err
	}

	// ECDSA 椭圆曲线数字签名算法
	ecD, ecX, ecY, err := genECKey()
	if err != nil {
		respPbMsg.Rt = InternalErrorCode.Int32()
		return respPbMsg, err
	}

	// 盐
	salt, err := genSaltKey()
	if err != nil {
		respPbMsg.Rt = InternalErrorCode.Int32()
		return respPbMsg, err
	}

	// 生成用户ID
	userId := ds.GenerateId()

	// MD5
	d5, m5 := make([]byte, 0, len(rsaPriKey)+len(aesKey)+len(aesIV)+len(ecD)+len(salt)), make([]byte, 0, 16)
	d5 = append(d5, rsaPriKey...)
	d5 = append(d5, aesKey...)
	d5 = append(d5, aesIV...)
	d5 = append(d5, ecD...)
	d5 = append(d5, salt...)

	sm5 := md5.Sum(d5)
	m5 = append(m5, sm5[:]...)

	// 创建应用分类
	seqs := make([]int64, 0, len(cache.Cache.Categories))
	for _, category := range cache.Cache.Categories {
		cs, err := ds.CreateApplicationCategory(tx, userId, category.ProtoId, category.Name)
		if err != nil {
			respPbMsg.Rt = SqlCreateFailureCode.Int32()
			return respPbMsg, err
		}

		seqs = append(seqs, cs.ID)
	}

	// 创建用户
	user, err := ds.CreateUser(tx, userId, ds.RegisterScore, seqs)
	if err != nil {
		respPbMsg.Rt = SqlCreateFailureCode.Int32()
		return respPbMsg, err
	}

	// 创建用户密钥
	key, err := ds.CreateUserKey(tx, userId, rsaPriKey, rsaPubKey, aesKey, aesIV, ecD, ecX, ecY, salt, m5)
	if err != nil {
		respPbMsg.Rt = SqlCreateFailureCode.Int32()
		return respPbMsg, err
	}

	// 创建存储用户文件的目录
	if err := TargetDirectoryFile.createDirectory(userId); err != nil {
		respPbMsg.Rt = WriteFileErrorCode.Int32()
		return respPbMsg, err
	}

	// 创建存储用户文件缩略图的目录
	if err := TargetDirectoryThumbnail.createDirectory(userId); err != nil {
		respPbMsg.Rt = WriteFileErrorCode.Int32()
		return respPbMsg, err
	}

	// 积分记录
	err = ds.CreateScore(tx, userId, ds.RegisterScoreAction, ds.RegisterScore, ds.RegisterScore)
	if err != nil {
		respPbMsg.Rt = SqlCreateFailureCode.Int32()
		return respPbMsg, err
	}

	// 返回用户ID
	respPbMsg.UserId = proto.Int64(userId)

	// 返回Keys
	respPbMsg.Bytes = key.Bytes()

	// 用户数据
	userData, err := NewUserData(user).Message(tx)
	if err != nil {
		respPbMsg.Rt = InternalErrorCode.Int32()
		return respPbMsg, err
	}
	respPbMsg.UserData = userData

	return respPbMsg, nil
}

// 生成RSA的私钥和公钥
func genRSAKey() ([]byte, []byte, error) {
	// 1024位 私钥
	pri, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}

	// 公钥
	pub := &pri.PublicKey

	// 公钥和私钥的编码
	priKey := x509.MarshalPKCS1PrivateKey(pri)
	pubKey := x509.MarshalPKCS1PublicKey(pub)
	logrus.Debugf("Register RSA Private Key Size is %d", len(priKey))
	logrus.Debugf("Register RSA Public Key Size is %d", len(pubKey))

	return priKey, pubKey, nil
}

// 生成AES的密钥和向量
func genAESKey() ([]byte, []byte, error) {

	// Key
	key := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, nil, err
	}

	// IV
	iv := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}

	return key, iv, nil
}

// 生成ECDSA的私钥和公钥
func genECKey() ([]byte, string, string, error) {
	pri, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, "", "", err
	}

	d := pri.D.String()
	x := pri.X.Text(62)
	y := pri.Y.Text(62)
	return []byte(d), x, y, nil
}

// 生成盐
func genSaltKey() ([]byte, error) {

	// Key
	salt := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}

	return salt, nil
}

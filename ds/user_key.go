package ds

import (
	"bytes"
	"encoding/binary"
	"github.com/jinzhu/gorm"
)

type UserKey struct {
	gorm.Model

	UserId int64 `gorm:"not null;type:bigint"`

	RSAPrivateKey []byte `gorm:"not null;size:2048"` // RSA 私钥
	RSAPublicKey  []byte `gorm:"not null;size:2048"` // RSA 公钥

	AESKey []byte `gorm:"not null;size:16"` // AES Key
	AESIV  []byte `gorm:"not null;size:16"` // AES IV

	ECD []byte `gorm:"not null;size:128"` // ECDSA 私钥 D
	ECX string `gorm:"not null;size:128"` // ECDSA 公钥 X
	ECY string `gorm:"not null;size:128"` // ECDSA 公钥 Y

	Salt []byte `gorm:"not null;size:32"` // 盐

	MD5 []byte `gorm:"not null;size:16"` // MD5(RSAPrivateKey + AESKey + AESIV + ECD + Salt)
}

func (m *UserKey) AfterCreateTable(db *gorm.DB) error {
	tx := db.Model(m).AddIndex("unique_idx_user_key_on_user_id", "user_id")
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (m *UserKey) Bytes() []byte {
	buf := new(bytes.Buffer)

	// 用户ID
	b0 := make([]byte, 8)
	binary.LittleEndian.PutUint64(b0, uint64(m.UserId))
	buf.Write(b0)

	// RSA 私钥
	b1 := make([]byte, 2)
	binary.BigEndian.PutUint16(b1, uint16(len(m.RSAPrivateKey)))
	buf.Write(b1)
	buf.Write(m.RSAPrivateKey)

	// AES Key
	buf.WriteByte(uint8(len(m.AESKey)))
	buf.Write(m.AESKey)

	// AES IV
	buf.WriteByte(uint8(len(m.AESIV)))
	buf.Write(m.AESIV)

	// ECDSA 私钥D
	buf.WriteByte(uint8(len(m.ECD)))
	buf.Write(m.ECD)

	// 盐
	buf.WriteByte(uint8(len(m.Salt)))
	buf.Write(m.Salt)

	// MD5
	buf.WriteByte(uint8(len(m.MD5)))
	buf.Write(m.MD5)

	return buf.Bytes()
}

func CreateUserKey(tx *gorm.DB, userId int64, rsaPriKey, rsaPubKey, aesKey, aesIV, ecD []byte, ecX string, ecY string, salt []byte, m5 []byte) (*UserKey, error) {
	key := &UserKey{
		UserId: userId,

		RSAPrivateKey: rsaPriKey,
		RSAPublicKey:  rsaPubKey,

		AESKey: aesKey,
		AESIV:  aesIV,

		ECD: ecD,
		ECX: ecX,
		ECY: ecY,

		Salt: salt,

		MD5: m5,
	}

	// Create
	db := tx.Create(key)
	if err := db.Error; err != nil {
		return nil, err
	}

	return key, nil
}

func FindUserKeyById(tx *gorm.DB, userId int64) (*UserKey, error) {
	var key UserKey

	db := tx.Where(&UserKey{UserId: userId}).First(&key)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &key, nil
}

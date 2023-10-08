package ds

import (
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"strings"
	"sync/atomic"
	"time"
)

var maxId int64 = 1 << 20

func GenerateUUID(id int64) string {
	// 16 + 1 + 32
	buf := make([]byte, 49)

	// ID
	oid := make([]byte, 8)
	binary.BigEndian.PutUint64(oid[:], uint64(id))
	hex.Encode(buf[:16], oid)

	// -
	buf[16] = '-'

	// UUID
	uid := make([]byte, 16)
	xuid := uuid.NewV4()
	uid = xuid.Bytes()
	hex.Encode(buf[17:], uid)

	return string(buf)
}

func GenerateId() int64 {
	return atomic.AddInt64(&maxId, 1)
}

var Models = []interface{}{
	&User{},
	&UserKey{},
	&ApplicationAccount{},
	&ApplicationAccountField{},
	&ApplicationCategory{},
	&File{},
	&Operate{},
	&Score{},
	&Task{},
	&Purchase{},
}

func Init(db *gorm.DB) error {
	for _, m := range Models {
		if !db.HasTable(m) {
			if err := db.CreateTable(m).Error; err != nil {
				return err
			}

			// After Create Table
			if fn, ok := m.(afterCreateTable); ok {
				if err := fn.AfterCreateTable(db); err != nil {
					return err
				}
			}
		}

		// MaxID
		if _, ok := m.(maxID); ok {
			var id sql.NullInt64

			err := db.Model(m).Unscoped().Select("MAX(id)").Row().Scan(&id)
			if err != nil {
				return err
			}

			if id.Int64 > maxId {
				maxId = id.Int64
			}
		}
	}
	logrus.Debugf("Get MAX ID is %d", maxId)

	// 初始化积分消费任务
	_, err := FindTaskByType(db, CostTask)
	if gorm.IsRecordNotFoundError(err) {
		if _, err := CreateTask(db, CostTask); err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

type maxID interface {
	MaxID()
}

type afterCreateTable interface {
	AfterCreateTable(db *gorm.DB) error
}

type BaseModel struct {
	ID int64 `gorm:"primary_key;type:bigint"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (b *BaseModel) MaxID() {}

func IsSameDay(d1 time.Time, d2 time.Time) bool {
	return strings.EqualFold(d1.Format("20060102"), d2.Format("20060102"))
}

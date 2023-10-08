package ds

import (
	"github.com/jinzhu/gorm"
)

type Purchase struct {
	BaseModel

	UserId int64 `gorm:"not null;type:bigint"`

	Source     string `gorm:"not null;size:256"`
	PurchaseID string `gorm:"not null;size:256"`
	ProductID  string `gorm:"not null;size:256"`

	Receipt         string `gorm:"not null;type:TEXT"`
	TransactionDate string `gorm:"not null;size:256"`

	IsValid bool `gorm:"not null"`
}

func (m *Purchase) AfterCreateTable(db *gorm.DB) error {
	tx := db.Model(m).AddIndex("idx_purchase_on_source_purchase_id_product_id", "source", "purchase_id", "product_id")
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func CreatePurchase(tx *gorm.DB, userId int64, source string, purchaseID string, productID string, receipt string, transactionDate string) error {
	p := &Purchase{
		UserId: userId,
		Source: source,

		PurchaseID: purchaseID,
		ProductID:  productID,

		Receipt:         receipt,
		TransactionDate: transactionDate,
	}
	p.ID = GenerateId()

	// Create
	db := tx.Create(p)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}

func FindPurchaseByPurchaseIDProductID(tx *gorm.DB, source string, purchaseID string, productID string, ) (*Purchase, error) {
	var p Purchase

	db := tx.Where(&Purchase{Source: source, PurchaseID: purchaseID, ProductID: productID}).First(&p)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &p, nil
}

package md

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strings"
	"xpf/cache"
	"xpf/ds"
	"xpf/pb"
)

type InAppPurchase struct {
}

func (m InAppPurchase) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	var err error
	respPbMsg := &pb.InAppPurchaseResp{
		Rt:      SuccessCode.Int32(),
		IsValid: proto.Bool(false),
		Score:   proto.Int64(0),
	}

	purchaseID := pbMsg.(*pb.InAppPurchase).GetPurchaseID()
	productID := pbMsg.(*pb.InAppPurchase).GetProductID()
	receipt := pbMsg.(*pb.InAppPurchase).GetReceipt()
	source := pbMsg.(*pb.InAppPurchase).GetSource()
	transactionDate := pbMsg.(*pb.InAppPurchase).GetTransactionDate()

	logrus.Debugf("Purchase ID is %s", purchaseID)
	logrus.Debugf("Product ID is %s", productID)
	logrus.Debugf("Receipt's Size is %d", len(receipt))
	logrus.Debugf("Source is %s", source)
	logrus.Debugf("Transaction Date is %s", transactionDate)

	// Exists
	purchase, err := ds.FindPurchaseByPurchaseIDProductID(tx, source, purchaseID, productID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	} else if purchase != nil {
		respPbMsg.Score = proto.Int64(user.Score)
		return respPbMsg, nil
	}

	// Verify Purchase
	if strings.EqualFold(source, "app_store") {
		// Receipt URL
		url := "https://buy.itunes.apple.com/verifyReceipt"
		if cache.Cache.Config.Mode == "debug" {
			url = "https://sandbox.itunes.apple.com/verifyReceipt"
		}

		buf := &bytes.Buffer{}
		verifyReceipt := &AppStoreVerifyReceipt{
			ReceiptData:            receipt,
			Password:               cache.Cache.Config.AppStoreVerifyReceiptPassword,
			ExcludeOldTransactions: false,
		}

		// Send Receipt Request
		if err := json.NewEncoder(buf).Encode(verifyReceipt); err != nil {
			respPbMsg.Rt = MarshalJsonFailureCode.Int32()
			return respPbMsg, err
		}

		// Receipt Response
		resp, err := http.Post(url, "application/json", buf)
		if err != nil {
			respPbMsg.Rt = PurchaseReceiptFailureCode.Int32()
			return respPbMsg, err
		}

		logrus.Debugf("Receipt Response :: %#v\n", resp.Body)

		var verifyReceiptResp AppStoreVerifyReceiptResp
		if err := json.NewDecoder(resp.Body).Decode(&verifyReceiptResp); err != nil {
			respPbMsg.Rt = UnmarshalJsonFailureCode.Int32()
			return respPbMsg, err
		}

		logrus.Debugf("AppStoreVerifyReceiptResp :: %#v\n", verifyReceiptResp)
		if verifyReceiptResp.Status != 0 {
			respPbMsg.Rt = PurchaseReceiptFailureCode.Int32()
			return respPbMsg, fmt.Errorf("response status is %d", verifyReceiptResp.Status)
		}
	} else {
		respPbMsg.Rt = ArgumentsErrorCode.Int32()
		return respPbMsg, fmt.Errorf("missing Dealing with Source %q", source)
	}

	// In APP Purchase Recrods
	err = ds.CreatePurchase(tx, user.ID, source, purchaseID, productID, receipt, transactionDate)
	if err != nil {
		respPbMsg.Rt = SqlCreateFailureCode.Int32()
		return respPbMsg, err
	}

	// Get Product Score
	cost, ok := cache.Cache.MapProducts[productID]
	if !ok {
		respPbMsg.Rt = ArgumentsErrorCode.Int32()
		return respPbMsg, fmt.Errorf("missing Product ID %s", productID)
	}

	// User's Score Cost Record
	score := user.Score + cost
	err = ds.CreateScore(tx, user.ID, ds.PurchaseScoreAction, cost, score)
	if err != nil {
		respPbMsg.Rt = SqlCreateFailureCode.Int32()
		return respPbMsg, err
	}

	// Save User's Score
	user.Score = score
	if err := tx.Save(user).Error; err != nil {
		respPbMsg.Rt = SqlSaveFailureCode.Int32()
		return respPbMsg, err
	}

	respPbMsg.IsValid = proto.Bool(true)
	respPbMsg.Score = proto.Int64(user.Score)
	return respPbMsg, nil
}

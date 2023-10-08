package md

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"math"
	"psw/cache"
	"psw/ds"
	"psw/pb"
)

type UserData struct {
	user *ds.User
}

func NewUserData(user *ds.User) *UserData {
	return &UserData{
		user: user,
	}
}

func (m *UserData) Message(tx *gorm.DB) (*pb.UserData, error) {
	pbMsg := &pb.UserData{
		Categories: make([]*pb.ApplicationCategory, 0),
		Snapshots:  make([]*pb.AccountSnapshot, 0),
		Files:      make([]*pb.File, 0),
		Score:      proto.Int64(0),
		Products:   make([]string, 0),
	}

	// Gesture
	pbMsg.Version = proto.String(cache.Cache.Config.Version)
	pbMsg.AssetsVersion = proto.String(cache.Cache.Config.AssetsVersion)
	pbMsg.IsGesture = proto.Bool(m.user.IsGesture)
	pbMsg.Gesture = m.user.Gesture
	pbMsg.ValidityGesture = proto.Int32(m.user.ValidityGesture)

	pbMsg.MaxFileSizeM = proto.Int32(int32(m.user.MaxFileSizeM))
	pbMsg.ImageCompressSizeM = proto.Int32(int32(m.user.ImageCompressSizeM))
	pbMsg.VideoCompressSizeM = proto.Int32(int32(m.user.VideoCompressSizeM))
	pbMsg.CostAccountScore = proto.Int32(int32(math.Abs(float64(ds.CostAccountScore))))
	pbMsg.CostFile10MScore = proto.Int32(int32(math.Abs(float64(ds.CostFile10MScore))))
	pbMsg.AllowCompression = proto.Bool(m.user.AllowCompression)
	pbMsg.ImageCompressPercentage = proto.Int32(int32(m.user.ImageCompressPercentage))
	pbMsg.ImageCompressQuality = proto.Int32(int32(m.user.ImageCompressQuality))
	pbMsg.VideoCompressQuality = m.user.VideoCompressQuality.Enum()

	// Categories
	categories, err := m.getCategories(tx)
	if err != nil {
		return pbMsg, err
	}
	pbMsg.Categories = categories

	// Snapshots
	snapshots, err := m.getSnapshots(tx)
	if err != nil {
		return pbMsg, err
	}
	pbMsg.Snapshots = snapshots

	// Files
	files, err := m.getFiles(tx)
	if err != nil {
		return pbMsg, err
	}
	pbMsg.Files = files

	// Score
	pbMsg.Score = proto.Int64(m.user.Score)

	// MapProducts
	pbMsg.Products = cache.Cache.Products

	// QRCode
	pbMsg.QrCode = proto.String(cache.Cache.Config.QrCode)

	return pbMsg, nil
}

func (m *UserData) getCategories(tx *gorm.DB) ([]*pb.ApplicationCategory, error) {
	pbCategories := make([]*pb.ApplicationCategory, 0)

	// Get Categories Seq
	seqs, err := m.user.GetApplicationCategoriesSeq()
	if err != nil {
		return nil, err
	}

	// Find Categories by User ID
	categories, err := ds.FindApplicationCategoriesByUserId(tx, m.user.ID)
	if err != nil {
		return nil, err
	}

	// Convert to Map Categories
	mapCategories := make(map[int64]*ds.ApplicationCategory)
	for _, category := range categories {
		mapCategories[category.ID] = category
	}

	// Protobuf Categories by Seq
	for _, seq := range seqs {
		category, ok := mapCategories[seq]
		if !ok {
			continue
		}

		name, err := category.GetName()
		if err != nil {
			return nil, err
		}

		pbCategory := &pb.ApplicationCategory{
			Id: proto.Int64(category.ID),

			ProtoId: proto.Int32(category.ProtoId),
			Name:    name,
		}

		pbCategories = append(pbCategories, pbCategory)
	}

	return pbCategories, nil
}

func (m *UserData) getSnapshots(tx *gorm.DB) ([]*pb.AccountSnapshot, error) {
	pbSnapshots := make([]*pb.AccountSnapshot, 0)

	snapshots, err := ds.FindApplicationAccountsByUserId(tx, m.user.ID)
	if err != nil {
		return nil, err
	}

	for _, snapshot := range snapshots {
		if len(snapshot.AccountId) == 0 {
			continue
		}

		appName, err := snapshot.GetAppName()
		if err != nil {
			return nil, err
		}

		pbSnapshot := &pb.AccountSnapshot{
			Id:         proto.Int64(snapshot.ID),
			CategoryId: proto.Int64(snapshot.CategoryId),
			AppId:      proto.String(snapshot.AppId),
			AppName:    appName,
			AccountId:  proto.String(snapshot.AccountId),
			CreateAt:   proto.Int32(int32(snapshot.CreatedAt.Unix())),
			Size:       proto.Int32(snapshot.Size),
		}

		pbSnapshots = append(pbSnapshots, pbSnapshot)
	}

	return pbSnapshots, nil
}

func (m *UserData) getFiles(tx *gorm.DB) ([]*pb.File, error) {
	pbFiles := make([]*pb.File, 0)

	files, err := ds.FindFilesByUserId(tx, m.user.ID)
	if err != nil {
		return nil, err
	}

	for _, fs := range files {
		if fs.Type != pb.FileType_FileCabinet {
			continue
		}

		pbFile := &pb.File{
			Id:           proto.Int64(fs.ID),
			Name:         proto.String(fs.Name),
			Mime:         proto.String(fs.Mime),
			Ext:          proto.String(fs.Ext),
			HasThumbnail: proto.Bool(fs.HasThumbnail),
			Size:         proto.Int32(fs.Size),
			IsShared:     proto.Bool(fs.IsShared),
			UploadAt:     proto.Int32(int32(fs.CreatedAt.Unix())),
		}

		pbFiles = append(pbFiles, pbFile)
	}

	return pbFiles, nil
}

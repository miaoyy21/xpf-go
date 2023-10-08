package pack

import (
	"google.golang.org/protobuf/proto"
	"psw/md"
	"psw/pb"
)

// Request Message
var reqMsgs = map[pb.MsgNo]proto.Message{
	pb.MsgNo_Msg1201: &pb.CreateApplicationCategory{},
	pb.MsgNo_Msg1202: &pb.DeleteApplicationCategory{},
	pb.MsgNo_Msg1203: &pb.SaveApplicationCategory{},
	pb.MsgNo_Msg1204: &pb.SaveApplicationCategoriesSeq{},
	pb.MsgNo_Msg1301: &pb.GetApplicationAccount{},
	pb.MsgNo_Msg1302: &pb.CreateApplicationAccount{},
	pb.MsgNo_Msg1303: &pb.DeleteApplicationAccount{},
	pb.MsgNo_Msg1304: &pb.SaveApplicationAccount{},
	pb.MsgNo_Msg2101: &pb.ShareFile{},
	pb.MsgNo_Msg2102: &pb.UploadFile{},
	pb.MsgNo_Msg2103: &pb.RenameFile{},
	pb.MsgNo_Msg2104: &pb.DeleteFile{},
	pb.MsgNo_Msg3101: &pb.SaveGesture{},
	pb.MsgNo_Msg3102: &pb.GetScores{},
	pb.MsgNo_Msg3103: &pb.GetOperates{},
	pb.MsgNo_Msg3201: &pb.GetTrash{},
	pb.MsgNo_Msg3202: &pb.RestoreTrash{},

	pb.MsgNo_Msg9000: &pb.Register{},
	pb.MsgNo_Msg9001: &pb.GetUserData{},
	pb.MsgNo_Msg9002: &pb.LoadAssets{},
	pb.MsgNo_Msg9009: &pb.InAppPurchase{},
}

// Deal Message
var dealMsgs = map[pb.MsgNo]md.Message{
	pb.MsgNo_Msg1201: md.CreateApplicationCategory{},
	pb.MsgNo_Msg1202: md.DeleteApplicationCategory{},
	pb.MsgNo_Msg1203: md.SaveApplicationCategory{},
	pb.MsgNo_Msg1204: md.SaveApplicationCategoriesSeq{},
	pb.MsgNo_Msg1301: md.GetApplicationAccount{},
	pb.MsgNo_Msg1302: md.CreateApplicationAccount{},
	pb.MsgNo_Msg1303: md.DeleteApplicationAccount{},
	pb.MsgNo_Msg1304: md.SaveApplicationAccount{},
	pb.MsgNo_Msg2101: md.ShareFile{},
	pb.MsgNo_Msg2102: md.UploadFile{},
	pb.MsgNo_Msg2103: md.RenameFile{},
	pb.MsgNo_Msg2104: md.DeleteFile{},
	pb.MsgNo_Msg3101: md.SaveGesture{},
	pb.MsgNo_Msg3102: md.GetScores{},
	pb.MsgNo_Msg3103: md.GetOperates{},
	pb.MsgNo_Msg3201: md.GetTrash{},
	pb.MsgNo_Msg3202: md.RestoreTrash{},

	pb.MsgNo_Msg9000: md.Register{},
	pb.MsgNo_Msg9001: md.GetUserData{},
	pb.MsgNo_Msg9002: md.LoadAssets{},
	pb.MsgNo_Msg9009: md.InAppPurchase{},
}

package protocol

import (
	"errors"
	"time"

	pb "MiniIMServe/protocol/protobuf"

	"google.golang.org/protobuf/proto"
)

var (
	ErrLoginFailed       = errors.New("login verification failed") // 登陆失败
	ErrMessageTypeFailed = errors.New("messages of unknown type")  // 消息类型错误
)

// Decode 解码
func Decode(data []byte) (*pb.PackData, error) {
	// 创建一个空的 PackData 对象
	packData := &pb.PackData{}

	// 使用 proto.Unmarshal 解析数据流
	err := proto.Unmarshal(data, packData)
	if err != nil {
		return nil, err
	}

	return packData, nil
}

// Encode 编码
func Encode(packData *pb.PackData) ([]byte, error) {
	return proto.Marshal(packData)
}

// VerifyLogin 验证登录信息
func VerifyLogin(data []byte) (*pb.LoginPack, error) {

	packData := &pb.PackData{}
	// 使用 proto.Unmarshal 解析数据流
	err := proto.Unmarshal(data, packData)
	if err != nil {
		return nil, err
	}

	// 解码失败
	if err != nil {
		return nil, err
	}

	// 不是登录消息
	if packData.Type != pb.PackType_LOGIN {
		return nil, ErrLoginFailed
	}

	loginInfo := &pb.LoginPack{}
	// 使用 proto.Unmarshal 解析数据流
	err = proto.Unmarshal(packData.Payload, loginInfo)
	if err != nil {
		return nil, err
	}

	if loginInfo.Password == "" {
		return nil, ErrLoginFailed
	}

	return loginInfo, nil
}

// Response 数据响应
func Response(id string, data *pb.ResponsePack) ([]byte, error) {
	payload, err := proto.Marshal(data)
	if err != nil {
		return nil, err
	}

	packData := &pb.PackData{
		Id:        id,
		Type:      pb.PackType_RESPONSE,
		Timestamp: uint64(time.Now().UnixNano()),
		Payload:   payload,
	}

	return proto.Marshal(packData)
}

package service

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// start_ai_generated

type SendMessageReq struct {
	ToUserName   string `xml:"ToUserName" json:"ToUserName"`
	FromUserName string `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime" json:"CreateTime"`
	MsgType      string `xml:"MsgType" json:"MsgType"`
	Content      string `xml:"Content" json:"Content"`
	MsgId        int64  `xml:"MsgId" json:"MsgId"`
	MsgDataId    string `xml:"MsgDataId" json:"MsgDataId"`
	Idx          string `xml:"Idx" json:"Idx"`
}
type ReplyMsgRes struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

// end_ai_generated

// SendMsgHandler 消息发送管理
func SendMsgHandler(w http.ResponseWriter, r *http.Request) {
	// 接收请求
	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("读取请求体失败", err)
	}
	defer r.Body.Close()
	fmt.Println("读取请求体成功", string(body))

	// 尝试解析为JSON
	var msgReq SendMessageReq
	errJSON := json.Unmarshal(body, &msgReq)
	if errJSON != nil {
		fmt.Println("解析为JSON失败:", errJSON)
		return
	}
	fmt.Println("解析为JSON成功:", msgReq)

	// 这里测试 回复消息
	replyMsg := ReplyMsgRes{
		ToUserName:   msgReq.FromUserName,
		FromUserName: msgReq.ToUserName,
		CreateTime:   msgReq.CreateTime,
		MsgType:      "text",
		Content:      "你好，我是机器人:" + msgReq.Content,
	}

	msg, err := xml.Marshal(&replyMsg)
	if err != nil {
		fmt.Println("xml.Marshal失败:", err)
		return
	}

	w.Header().Set("content-type", "application/xml")
	w.Write(msg)

	// res := &JsonResult{}
	// res.Code = 0
	// res.Data = "成功接收到消息" + string(msgReq.Content)
	// msg, err := json.Marshal(res)
	// if err != nil {
	// 	fmt.Fprint(w, "内部错误")
	// 	return
	// }
	// w.Header().Set("content-type", "application/json")
	// w.Write(msg)
}

package service

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type SendMessageReq struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        int64    `xml:"MsgId"`
	MsgDataId    string   `xml:"MsgDataId"`
	Idx          string   `xml:"Idx"`
}

// SendMsgHandler 消息发送管理
func SendMsgHandler(w http.ResponseWriter, r *http.Request) {
	// 读取请求体
	body, err := io.ReadAll(r.Body)

	// 解析XML
	var sendMessageReq SendMessageReq
	err = xml.Unmarshal(body, &sendMessageReq)
	if err != nil {
	}

	sendMessageReqStr, err := json.Marshal(sendMessageReq)

	res := &JsonResult{}
	res.Code = 0
	res.Data = "成功接收到消息" + string(sendMessageReqStr)
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	fmt.Println("fmt接收到消息:", sendMessageReqStr)
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

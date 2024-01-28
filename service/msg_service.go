package service

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type SendMessageReq struct {
	XMLName xml.Name `xml:"xml"`

	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        int64  `xml:"MsgId"`
	MsgDataId    string `xml:"MsgDataId"`
	Idx          string `xml:"Idx"`
}

// SendMsgHandler 消息发送管理
func SendMsgHandler(w http.ResponseWriter, r *http.Request) {
	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("读取请求体失败", err)
	}
	defer r.Body.Close()
	fmt.Println("读取请求体成功", string(body))

	// 输出请求体
	fmt.Println("请求体内容:", string(body))

	// 尝试解析为XML
	var sendMessageReq SendMessageReq
	errXML := xml.Unmarshal(body, &sendMessageReq)
	if errXML == nil {
		fmt.Println("解析为XML成功:", sendMessageReq)
		return
	}
	fmt.Println("解析为XML失败:", errXML)

	// 尝试解析为JSON
	var sendMessageReqJSON map[string]interface{}
	errJSON := json.Unmarshal(body, &sendMessageReqJSON)
	if errJSON == nil {
		fmt.Println("解析为JSON成功:", sendMessageReqJSON)
		return
	}
	fmt.Println("解析为JSON失败:", errJSON)

	// 如果都失败了，返回错误
	fmt.Println("未知的请求格式")

	// // 解析XML
	// var sendMessageReq SendMessageReq
	// err = xml.Unmarshal(body, &sendMessageReq)
	// if err != nil {
	// 	fmt.Println("结构体:解析XML失败", err)
	// }

	// var sendMessageReqV2 map[string]interface{}
	// err = xml.Unmarshal(body, &sendMessageReqV2)
	// if err != nil {
	// 	fmt.Println("map:解析XML失败", err)
	// }

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

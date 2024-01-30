package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

var API_KEY = ""
var SECRET_KEY = ""

type ChatMsgReq struct {
	Messages []Msg `json:"messages"`
}

type Msg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	UserID  string `json:"user_id"`
}

type SendMsgReq struct {
	UserID string `json:"user_id"`
	Msg    string `json:"msg"`
}

type SendMsgResp struct {
	Msg string `json:"msg"`
}

type ChatMsgResp struct {
	ID               string `json:"id"`
	Object           string `json:"object"`
	Created          int64  `json:"created"`
	Result           string `json:"result"`
	IsTruncated      bool   `json:"is_truncated"`
	NeedClearHistory bool   `json:"need_clear_history"`
	Usage            struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func InitWenXin() {
	API_KEY = os.Getenv("API_KEY")
	SECRET_KEY = os.Getenv("SECRET_KEY")
}

func postToWenXin(chatReq ChatMsgReq) (chatMsgResp ChatMsgResp, err error) {
	url := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant?access_token=" + GetAccessToken()

	client := resty.New()
	_, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&chatReq).
		SetResult(&chatMsgResp).
		Post(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	return chatMsgResp, nil
}

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", API_KEY, SECRET_KEY)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}
func createChatMsgReq(wechatMsgReq WeChatSendMsgReq) ChatMsgReq {
	var chatMsgReq ChatMsgReq
	chatMsgReq.Messages = make([]Msg, 0)
	chatMsgReq.Messages = append(chatMsgReq.Messages, Msg{
		Role:    "user",
		Content: wechatMsgReq.Content,
		UserID:  wechatMsgReq.FromUserName,
	})
	return chatMsgReq
}

func createChatMsgReqWithHistory(wechatMsgReq WeChatSendMsgReq, chatMsgReq ChatMsgReq) ChatMsgReq {
	chatMsgReq.Messages = append(chatMsgReq.Messages, Msg{
		Role:    "user",
		Content: wechatMsgReq.Content,
		UserID:  wechatMsgReq.FromUserName,
	})
	return chatMsgReq
}

// test

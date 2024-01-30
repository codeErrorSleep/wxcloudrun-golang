package service

import (
	"encoding/json"

	"github.com/coocood/freecache"
)

var cache *freecache.Cache

var localCacheKey = "local_cache_key:"

func InitLocalCache() {
	cacheSize := 100 * 1024 * 1024
	cache = freecache.NewCache(cacheSize)
}

func getLocalCacheKey(userID string) string {
	return localCacheKey + userID
}

func getMsgHistory(userID string) (ChatMsgReq, error) {
	var chatMsgReq ChatMsgReq
	key := getLocalCacheKey(userID)
	value, err := cache.Get([]byte(key))
	if err != nil && err != freecache.ErrNotFound {
		return chatMsgReq, err
	}
	if value == nil {
		return chatMsgReq, nil
	}

	// 将value解析成ChatMsgReq
	err = json.Unmarshal(value, &chatMsgReq)
	if err != nil {
		return chatMsgReq, err
	}

	return chatMsgReq, nil
}

func setMsgHistory(userID string, chatMsgReq ChatMsgReq) error {
	// 将ChatMsgReq转换为JSON字符串
	value, err := json.Marshal(chatMsgReq)
	if err != nil {
		return err
	}

	// 将JSON字符串存入缓存
	err = cache.Set([]byte(getLocalCacheKey(userID)), value, 3600)
	if err != nil {
		return err
	}

	return nil
}

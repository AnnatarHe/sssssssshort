package src

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func cacheKeyGen(ip string) string {
	return fmt.Sprintf("rate:%s:count", ip)
}

func IPFilter(request *http.Request) (bool, error) {
	ip := request.Header.Get("X-Real-IP")

	if ip == "" {
		ip = strings.Split(request.RemoteAddr, ":")[0]
	}

	redisKey := cacheKeyGen(ip)
	count, err := Cache.Get(redisKey).Result()

	if err != nil {
		// 说明可能是第一次登陆，加入 key
		if err := Cache.Set(redisKey, 0, time.Hour*24).Err(); err != nil {
			log.Println(err)
			return false, err
		}
		count = "0"
	}

	countInt, err := strconv.Atoi(count)
	if err != nil {
		log.Println(err)
		return false, err
	}

	// 每个ip每天限制 3000 个请求
	if countInt > 3000 {
		err := errors.New("api request out of limit")
		return false, err
	}

	go Cache.Incr(redisKey)
	return true, nil

}

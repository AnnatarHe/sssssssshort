package src

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

type ipTimes struct {
	times     uint16
	createdAt time.Time
}

type ipTimesMap map[string]*ipTimes

var ipMap ipTimesMap

func init() {
	ipMap = make(ipTimesMap)
}

func IPFilter(request *http.Request) (bool, error) {
	ip := request.Header.Get("X-Real-IP")

	if ip == "" {
		ip = strings.Split(request.RemoteAddr, ":")[0]
	}

	val, ok := ipMap[ip]

	if !ok {
		// 说明可能是第一次登陆，加入 key
		val = &ipTimes{
			times:     0,
			createdAt: time.Now(),
		}

		ipMap[ip] = val
	}

	if val.createdAt.Add(time.Minute * 60 * 24).Before(time.Now()) {
		delete(ipMap, ip)
		return true, nil
	}

	// 每个ip每天限制 3000 个请求
	if val.times > 3000 {
		err := errors.New("api request out of limit")
		return false, err
	}

	ipMap[ip] = &ipTimes{
		times:     val.times + 1,
		createdAt: val.createdAt,
	}

	return true, nil
}

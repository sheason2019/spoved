package proxy_service

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var debounceHostMap = map[string]*DebounceHostInfo{}

// 获取反向代理Host
func GetHost(ctx *gin.Context, username, projectName string) (string, error) {
	// map 中的 key
	k := fmt.Sprintf("%s/%s", username, projectName)
	// 首先尝试从debounceHostMap获取Host地址
	info, exist := debounceHostMap[k]
	var err error

	// 若info不存在于缓存，则从数据库中检索info信息
	if !exist {
		info, err = findProxyHostInfo(ctx, username, projectName)
		if err != nil {
			info = &DebounceHostInfo{Error: err}
		}

		saveDebounceHostInfo(ctx, k, info)
	}

	if info.Error != nil {
		return "", fmt.Errorf("error: DebounceHostInfoHasError: %w", info.Error)
	}

	for _, match := range info.Miniflow {
		if matchHeader(ctx, &match) {
			return match.HostPath, nil
		}
	}

	return info.Online.HostPath, nil
}

// 将 Host info 的过期时间设置为15秒
const timeoutDuration = time.Second * 15

// 将Info写入缓存
func saveDebounceHostInfo(ctx context.Context, key string, info *DebounceHostInfo) {
	debounceHostMap[key] = info
	timer := time.NewTimer(timeoutDuration)
	go func() {
		<-timer.C
		delete(debounceHostMap, key)
	}()
}

func matchHeader(ctx *gin.Context, match *HostMatch) bool {
	for header, value := range match.HeaderMatch {
		if ctx.GetHeader(header) != value {
			return false
		}
	}

	return true
}

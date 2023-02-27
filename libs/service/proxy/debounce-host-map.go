package proxy_service

import (
	"context"
	"fmt"
	"time"
)

type DebounceHostInfo struct {
	HostPath string
	Error    error
}

var debounceHostMap = map[string]*DebounceHostInfo{}

// 获取反向代理Host
func GetHost(ctx context.Context, username, projectName string) (string, error) {
	// map 中的 key
	k := fmt.Sprintf("%s/%s", username, projectName)
	// 首先尝试从debounceHostMap获取Host地址
	info, exist := debounceHostMap[k]

	// 若info不存在于缓存，则从数据库中检索info信息
	if !exist {
		info = newDebounceHostInfo(findProxyHost(ctx, username, projectName))
		saveDebounceHostInfo(ctx, k, info)
	}

	if info.Error != nil {
		return "", fmt.Errorf("error: DebounceHostInfoHasError: %w", info.Error)
	}

	return info.HostPath, nil
}

func newDebounceHostInfo(hostPath string, err error) *DebounceHostInfo {
	return &DebounceHostInfo{
		HostPath: hostPath,
		Error:    err,
	}
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

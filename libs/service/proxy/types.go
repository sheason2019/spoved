package proxy_service

type HostMatch struct {
	HostPath    string
	HeaderMatch map[string]string
}

type DebounceHostInfo struct {
	Online   HostMatch
	Miniflow []HostMatch
	Error    error
}

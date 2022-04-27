package got

import "sync"

var BroadcastWhiteList = NewWhiteList()

func init() {
	BroadcastWhiteList.AddSign("0x42a23390")
	BroadcastWhiteList.AddSign("0x85052446")
	BroadcastWhiteList.AddSign("0x0ebdf911")
	BroadcastWhiteList.AddSign("0xb007edfd")
	BroadcastWhiteList.AddSign("0xace6975b")
	BroadcastWhiteList.AddSign("0xcb7cbdaa")

	BroadcastWhiteList.AddAddress("0x10ed43c718714eb63d5aa57b78b54704e256024e")
}

type WhiteList struct {
	sync.RWMutex
	signMap    map[string]struct{}
	addressMap map[string]struct{}
}

func NewWhiteList() *WhiteList {
	wl := new(WhiteList)
	wl.signMap = make(map[string]struct{})
	wl.addressMap = make(map[string]struct{})

	return wl
}

func (wl *WhiteList) IncludeSign(sign string) bool {
	wl.RLock()
	defer wl.RUnlock()

	_, ok := wl.signMap[sign]

	return ok
}

func (wl *WhiteList) IncludeRouter(routerAddr string) bool {
	wl.RLock()
	defer wl.RUnlock()

	_, ok := wl.addressMap[routerAddr]

	return ok
}

func (wl *WhiteList) AddSign(sign string) {
	wl.Lock()
	defer wl.Unlock()

	wl.signMap[sign] = struct{}{}
}

func (wl *WhiteList) AddAddress(addr string) {
	wl.Lock()
	defer wl.Unlock()

	wl.addressMap[addr] = struct{}{}
}

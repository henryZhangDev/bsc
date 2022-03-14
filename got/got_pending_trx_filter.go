package got

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"strings"
	"sync"
)

var PendingTxFilter = NewPendingTxFilter()

func init() {
	//add address
	PendingTxFilter.AddTo(strings.ToLower("0x10ed43c718714eb63d5aa57b78b54704e256024e"))

	//add sign
	PendingTxFilter.AddSign("0x3a1f02ba")
	PendingTxFilter.AddSign("0x00001678")
	PendingTxFilter.AddSign("0x00007046")
	PendingTxFilter.AddSign("0x0000dcca")
	PendingTxFilter.AddSign("0xd85e2d20")
	PendingTxFilter.AddSign("0xd92efcc6")
	PendingTxFilter.AddSign("0xe72f160e")
	PendingTxFilter.AddSign("0x620ccfd8")
	PendingTxFilter.AddSign("0x24fbd99a")
	PendingTxFilter.AddSign("0xa0345e93")
	PendingTxFilter.AddSign("0x0000cbb1")
	PendingTxFilter.AddSign("0x0000cbb2")
	PendingTxFilter.AddSign("0x0000cbb3")
	PendingTxFilter.AddSign("0x0000cbb4")
	PendingTxFilter.AddSign("0x62fb1fe6")
	PendingTxFilter.AddSign("0x5f084bb8")
	PendingTxFilter.AddSign("0xf2f7ed83")
	PendingTxFilter.AddSign("0xc5c5c535")
	PendingTxFilter.AddSign("0xebcb1875")
	PendingTxFilter.AddSign("0x1b8b4940")
	PendingTxFilter.AddSign("0x3455b26d")
	PendingTxFilter.AddSign("0x69990e28")
	PendingTxFilter.AddSign("0xaa5e7193")
	PendingTxFilter.AddSign("0xb4865c7e")
	PendingTxFilter.AddSign("0x4da3dd62")
	PendingTxFilter.AddSign("0xeba7903c")
	PendingTxFilter.AddSign("0xbb001eab")
	PendingTxFilter.AddSign("0xef49fef6")
	PendingTxFilter.AddSign("0x4713348b")
	PendingTxFilter.AddSign("0x3e44b7ac")
	PendingTxFilter.AddSign("0x9d625677")
	PendingTxFilter.AddSign("0x1000b904")
	PendingTxFilter.AddSign("0x313ad0d6")
	PendingTxFilter.AddSign("0x960cfe41")
	PendingTxFilter.AddSign("0x3e999aea")
	PendingTxFilter.AddSign("0x06001f0a")
	PendingTxFilter.AddSign("0x90819984")
	PendingTxFilter.AddSign("0xeccacb2b")
	PendingTxFilter.AddSign("0x6f72cbaf")
	PendingTxFilter.AddSign("0x266a1809")
	PendingTxFilter.AddSign("0x000000fe")
	PendingTxFilter.AddSign("0x0000fb82")
	PendingTxFilter.AddSign("0x0712f399")
	PendingTxFilter.AddSign("0x860ca632")
	PendingTxFilter.AddSign("0xc1a96832")
	PendingTxFilter.AddSign("0x6db7b060")
	PendingTxFilter.AddSign("0x9aad4418")
	PendingTxFilter.AddSign("0x00000650")
	PendingTxFilter.AddSign("0xc5efc1c6")
	PendingTxFilter.AddSign("0xe74be8b1")
	PendingTxFilter.AddSign("0xd46fb3cb")
	PendingTxFilter.AddSign("0x7e7e669c")
	PendingTxFilter.AddSign("0xa57ed4fe")
	PendingTxFilter.AddSign("0xa707ad3c")
	PendingTxFilter.AddSign("0x0c961190")
	PendingTxFilter.AddSign("0xc84240ed")
	PendingTxFilter.AddSign("0xb8025adc")
	PendingTxFilter.AddSign("0x84aa0d65")
	PendingTxFilter.AddSign("0xdc42f971")
	PendingTxFilter.AddSign("0x6cc18f47")
	PendingTxFilter.AddSign("0x13727134")
	PendingTxFilter.AddSign("0x76cc3406")
	PendingTxFilter.AddSign("0x9d66dbe1")
	PendingTxFilter.AddSign("0xa63d195e")
	PendingTxFilter.AddSign("0x827a5bfc")
	PendingTxFilter.AddSign("0x25781668")
	PendingTxFilter.AddSign("0xae3f2e74")
	PendingTxFilter.AddSign("0x00000011")
	PendingTxFilter.AddSign("0x0bb9fdd7")
	PendingTxFilter.AddSign("0x7a908eae")
	PendingTxFilter.AddSign("0x2b5acbab")
	PendingTxFilter.AddSign("0xe4a6facd")
	PendingTxFilter.AddSign("0xe0668d4c")
	PendingTxFilter.AddSign("0xbe7b94a6")
	PendingTxFilter.AddSign("0x2871e121")
	PendingTxFilter.AddSign("0xd64f650d")
	PendingTxFilter.AddSign("0x35fcc952")
	PendingTxFilter.AddSign("0xd7e402ea")
}

type TxFilter struct {
	enable     bool
	signLocker sync.RWMutex
	toLocker   sync.RWMutex

	SignMap map[string]struct{}
	ToMap   map[string]struct{}
}

func NewPendingTxFilter() *TxFilter {
	filter := new(TxFilter)
	filter.SignMap = make(map[string]struct{})
	filter.ToMap = make(map[string]struct{})

	return filter
}

func (filter *TxFilter) NeedSendToSubscription(pendingTx *types.Transaction) bool {
	if pendingTx.To() == nil || len(pendingTx.Data()) < 4 {
		return false
	}

	input := pendingTx.Data()
	sign := hexutil.Encode(input[:4])
	to := strings.ToLower(pendingTx.To().String())

	return filter.IncludeTo(to) || filter.IncludeSign(sign)
}

func (filter *TxFilter) IncludeSign(sign string) bool {
	filter.signLocker.RLock()
	defer filter.signLocker.RUnlock()

	_, ok := filter.SignMap[sign]

	return ok
}

func (filter *TxFilter) AddSign(sign string) {
	filter.signLocker.Lock()
	defer filter.signLocker.Unlock()

	filter.SignMap[sign] = struct{}{}
}

func (filter *TxFilter) ListSign() ([]byte, error) {
	filter.signLocker.RLock()
	defer filter.signLocker.RUnlock()

	return json.Marshal(filter.SignMap)
}

func (filter *TxFilter) IncludeTo(to string) bool {
	filter.toLocker.RLock()
	defer filter.toLocker.RUnlock()

	_, ok := filter.ToMap[to]

	return ok
}

func (filter *TxFilter) IsEnable() bool {
	return filter.enable
}

func (filter *TxFilter) SetEnable() {
	filter.enable = true
}

func (filter *TxFilter) SetDisable() {
	filter.enable = false
}

func (filter *TxFilter) AddTo(to string) {
	filter.toLocker.Lock()
	defer filter.toLocker.Unlock()

	filter.ToMap[to] = struct{}{}
}

func (filter *TxFilter) ListTo() ([]byte, error) {
	filter.toLocker.RLock()
	defer filter.toLocker.RUnlock()

	return json.Marshal(filter.ToMap)
}

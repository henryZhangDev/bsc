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
	//add sign

	//add address
}

type TxFilter struct {
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

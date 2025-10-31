package environment

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NonceManager struct {
	mu     sync.Mutex
	next   uint64
	addr   common.Address
	client *ethclient.Client
	log    *slog.Logger

	lastSync  time.Time
	syncEvery time.Duration
}

func NewNonceManager(client *ethclient.Client, addr common.Address, log *slog.Logger) (*NonceManager, error) {
	n, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return nil, err
	}

	return &NonceManager{
		next:      n,
		addr:      addr,
		client:    client,
		log:       log,
		syncEvery: 5 * time.Second,
	}, nil
}

func (m *NonceManager) NextBlock(ctx context.Context, k uint64) (start int64, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if time.Since(m.lastSync) > m.syncEvery {
		n, err := m.client.PendingNonceAt(ctx, m.addr)
		if err != nil {
			return 0, err
		}

		if n > m.next {
			m.next = n
		}
		m.lastSync = time.Now()
	}

	start = int64(m.next)
	m.next += k
	m.log.Debug("allocated nonces", "start", start, "count", k)
	return start, nil
}

func (m *NonceManager) ForceSync(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	n, err := m.client.PendingNonceAt(ctx, m.addr)
	if err != nil {
		return err
	}
	if n > m.next {
		m.next = n
	}
	m.lastSync = time.Now()
	m.log.Debug("forced nonce sync", "next", m.next)
	return nil
}

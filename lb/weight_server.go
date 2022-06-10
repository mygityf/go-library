package lb

import (
	"fmt"
	"strings"
	"sync"
)

// BackendNode backend server
type BackendNode struct {
	// weight
	Weight int
	// currentWeight, to set to Weight at first time
	currentWeight int
	// node name
	Name string
}

// LoadBalanceServer
type LoadBalanceServer interface {
	// AddBackendServer add
	Add(backendServer *BackendNode)
	// GetBackendServer get
	Get() *BackendNode
}

// weightServerRoundRobin weight robin struct
type weightServerRoundRobin struct {
	// sum of effective weight
	effectiveWeight int
	// backend server list
	backendServerList []*BackendNode
	// mutex
	mutex sync.Mutex
}

// NewWeightServerRoundRobin create a robin load balance.
func NewWeightServerRoundRobin() *weightServerRoundRobin {
	return &weightServerRoundRobin{
		effectiveWeight: 0,
	}
}

// Add add backend with weight and name
func (r *weightServerRoundRobin) Add(backendServer *BackendNode) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.effectiveWeight += backendServer.Weight
	r.backendServerList = append(r.backendServerList, backendServer)
}

// Get get backend by weight
func (r *weightServerRoundRobin) Get() *BackendNode {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	var expectBackendServer *BackendNode
	for _, backendServer := range r.backendServerList {
		// add initial weight to all node
		backendServer.currentWeight += backendServer.Weight
		if expectBackendServer == nil {
			expectBackendServer = backendServer
		}
		if backendServer.currentWeight > expectBackendServer.currentWeight {
			expectBackendServer = backendServer
		}
	}
	// r.Visit()
	// decrease weight of expect node
	expectBackendServer.currentWeight -= r.effectiveWeight
	return expectBackendServer
}

// Visit 打印后端服务的当前权重变化
func (r *weightServerRoundRobin) Visit() {
	var serverListForLog []string
	for _, backendServer := range r.backendServerList {
		serverListForLog = append(serverListForLog,
			fmt.Sprintf("%v", backendServer.currentWeight))
	}
	fmt.Printf("(%v)\n", strings.Join(serverListForLog, ", "))
}

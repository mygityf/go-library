package lb

import (
	"fmt"
	"testing"
)

func TestNewWeightServerRoundRobin(t *testing.T) {
	weightServerRoundRobin := NewWeightServerRoundRobin()
	weightServerRoundRobin.Add(&BackendNode{
		Name:   "ServerA",
		Weight: 5,
	})
	weightServerRoundRobin.Add(&BackendNode{
		Name:   "ServerB",
		Weight: 3,
	})
	weightServerRoundRobin.Add(&BackendNode{
		Name:   "ServerC",
		Weight: 1,
	})

	expectServerNameList := []string{
		"ServerA", "ServerB", "ServerA", "ServerC", "ServerA", "ServerB", "ServerA", "ServerB", "ServerA",
		//"ServerA", "ServerB", "ServerA", "ServerC", "ServerA", "ServerB", "ServerA", "ServerB", "ServerA",
	}
	fmt.Printf("(A, B, C)\n")
	for ii, expectServerName := range expectServerNameList {
		weightServerRoundRobin.Visit()
		backendServer := weightServerRoundRobin.Get()
		if backendServer.Name != expectServerName {
			t.Errorf("%v.%v.expect:%v, actual:%v", t.Name(), ii, expectServerName, backendServer.Name)
			return
		}
	}
}

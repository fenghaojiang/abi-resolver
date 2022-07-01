package resolve

import "testing"

func TestOnResolve(t *testing.T) {
	t.Run("resolve from etherscan", func(t *testing.T) {
		resolver := NewResolver()
		content := resolver.FetchABIFromEtherscan("0x8d12a197cb00d4747a1fe03395095ce2a5cc6819")
		resolver.SerializeABI(content, "0x8d12a197cb00d4747a1fe03395095ce2a5cc6819")
	})
}

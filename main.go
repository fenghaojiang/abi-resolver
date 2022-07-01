package main

import (
	"github.com/fenghaojiang/abi-resolver/resolve"
)

func main() {
	abis := []string{"./abis/seaport.abi", "./abis/uniswap_v3_factory.abi"}
	resolver := resolve.NewResolver().WithAbis(abis)
	resolver.Resolve("")
}

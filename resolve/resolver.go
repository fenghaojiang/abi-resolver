package resolve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/fenghaojiang/abi-resolver/config"
)

type Resolver struct {
	client *http.Client
	abis   []string
}

func NewResolver() *Resolver {
	return &Resolver{
		client: &http.Client{},
	}
}

func (r *Resolver) WithAbis(abis []string) *Resolver {
	r.abis = append(r.abis, abis...)
	return r
}

func (r *Resolver) RemoveAbis(abis []string) *Resolver {
	for _, abi := range abis {
		for i, a := range r.abis {
			if a == abi {
				r.abis = append(r.abis[:i], r.abis[i+1:]...)
			}
		}
	}
	return r
}

func (r *Resolver) Resolve(contractAddress string) {
	if len(r.abis) == 0 {
	}
}

func (r *Resolver) ResolveByABI(abi string) *Resolver {
	return r
}

func (r *Resolver) FetchABIFromEtherscan(contractAddress string) string {
	request := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "https",
			Host:   "api.etherscan.io",
			Path:   "/api",
			RawQuery: url.Values{
				"module":  {"contract"},
				"action":  {"getabi"},
				"address": {contractAddress},
			}.Encode(),
		},
	}
	resp, err := r.client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	abiContent := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&abiContent); err != nil {
		fmt.Println(err.Error())
		return ""
	}
	fmt.Println(abiContent)
	return abiContent["result"].(string)
}

func (r *Resolver) SerializeABI(abiContent string, contractAddress string) {
	f, err := os.OpenFile(config.DefaultConfig().OuputDir+contractAddress+".abi", os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("failed to open file, err:", err.Error())
		return
	}
	defer f.Close()
	_, err = f.WriteString(abiContent)
	if err != nil {
		fmt.Println("write content failed, err:", err.Error())
	}
}

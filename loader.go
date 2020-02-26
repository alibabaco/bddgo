package bddgo

import (
	"fmt"
	"net/http"
	"plugin"
)

func LoadPackageBinary(packageName string, functionName string) (
	handler http.Handler,
	err error,
) {

	p, err := plugin.Open(fmt.Sprintf("%s.so", packageName))
	if err != nil {
		return
	}

	sym, err := p.Lookup(functionName)
	if err != nil {
		return
	}

	handler = sym.(func() http.Handler)()
	return
}

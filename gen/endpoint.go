package gen

import (
	"github.com/worldOneo/crudist/core"
	"github.com/worldOneo/crudist/gen/backends"
)

// GenerateEndpoints generates crud endpoints
func GenerateEndpoints(framework string, endpoints []core.CrudEndpoint) error {
	for _, endpoint := range endpoints {
		err := backends.Generate(framework, endpoint)
		if err != nil {
			return err
		}
	}
	return nil
}

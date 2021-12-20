package gen

import "github.com/worldOneo/crudist/core"

type errorCollector struct {
	err error
}

func (e *errorCollector) run(a func() error) {
	if e.err != nil {
		return
	}
	e.err = a()
}

// Generate generates every resource from the config
func Generate(cmd core.Meta, conf core.CrudConfig) error {
	coll := errorCollector{}
	coll.run(func() error { return GenerateErrors(conf.Errors) })
	coll.run(func() error { return GenerateEndpoints(conf.Framework, conf.Endpoints) })
	return coll.err
}

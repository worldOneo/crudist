package backends

import (
	"bytes"
	// Importing templates
	_ "embed"
	"html/template"

	"github.com/worldOneo/crudist/core"
	"github.com/worldOneo/crudist/errors"
)

var registered map[string]*template.Template = make(map[string]*template.Template)

// Generate generates an endpoint for the given framework
func Generate(framework string, endpoint core.CrudEndpoint) error {
	backend, ok := registered[framework]
	if !ok {
		return errors.NewFrameworkMissingError(framework)
	}
	res := bytes.Buffer{}
	err := backend.Execute(&res, endpoint)
	if err != nil {
		return errors.NewTemplateError(err)
	}
	return core.WriteToPackageSuffixed(endpoint.Package, "_" + endpoint.Name, res.String())
}

//go:embed fiber.gotemplate
var fiber string

//go:embed gin.gotemplate
var gin string

//go:embed gochi.gotemplate
var gochi string

func makeTemplate(n, inp string) *template.Template {
	return template.Must(template.New(n).Parse(inp))
}

func init() {
	registered["fiber"] = makeTemplate("fiber", fiber)
	registered["gin"] = makeTemplate("gin", gin)
	registered["gochi"] = makeTemplate("gochi", gochi)
}

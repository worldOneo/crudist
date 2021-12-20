package gen

import (
	"bytes"
	// Include templates
	_ "embed"
	"html/template"

	"github.com/worldOneo/crudist/core"
	"github.com/worldOneo/crudist/errors"
)

//go:embed errors.gotemplate
var errorTemplateRaw string
var errorTemplate = template.Must(template.New("error").Parse(errorTemplateRaw))

// GenerateErrors an errors file from the given config
// If config is nil this function is no-op
func GenerateErrors(errorConfig *core.Errors) error {
	if (errorConfig == nil) {
		return nil
	}
	res := bytes.Buffer{}
	err := errorTemplate.Execute(&res, errorConfig)
	if err != nil {
		return errors.NewTemplateError(err)
	}
	return core.WriteToPackage(errorConfig.Package, res.String())
}

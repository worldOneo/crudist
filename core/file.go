package core

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/worldOneo/crudist/errors"
)

// WriteToFile writes content to a file or overwrites if it exists
func WriteToFile(file, content string) error {
	err := os.MkdirAll(filepath.Dir(file), 0o666)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, []byte(content), 0o644)
	if err != nil {
		return errors.NewFileWriteError(file, err)
	}
	return nil
}

// WriteToPackage writes content to pkg and
// a file with the name of the last package
// e.g: html/template writes to html/template/template.crudist.go
func WriteToPackage(pkg, content string) error {
	return WriteToPackageSuffixed(pkg, "", content)
}

// WriteToPackageSuffixed writes content to pkg and
// a file with the name of the last package + suffix
// e.g: "html/template", "_myfile" writes to html/template/template_myfile.gen.go
func WriteToPackageSuffixed(pkg, suffix, content string) error {
	last := LastPackagePart(pkg)
	path := filepath.Join(pkg, last + suffix + ".crudist.go")
	return WriteToFile(path, content)
}

// LastPackagePart returns the last part of a package
// e.g: html/template -> template
func LastPackagePart(pkg string) string {
	parts := strings.Split(pkg, "/")
	return parts[len(parts)-1]
}

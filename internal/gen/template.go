// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"fmt"
	"github.com/domgolonka/ent2proto/internal/gen/internal"
	"text/template"
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o=internal/bindata.go -pkg=internal -modtime=1 ./template/...

type (
	// TypeTemplate specifies a template that is executed with
	// each Type object of the graph.
	TypeTemplate struct {
		Name   string             // template name.
		Format func(*Type) string // file name format.
	}
	// GraphTemplate specifies a template that is executed with
	// the Graph object.
	GraphTemplate struct {
		Name   string            // template name. Skip   func(*Graph) bool // skip condition.
		Format string            // file name format.
	}
)

var (
	// Templates holds the template information for a file that the graph is generating.
	Templates = []TypeTemplate{
		{
			Name:   "protobuf",
			//Format: pkgf("%s_create.go"),
		},
		//{
		//	Name:   "mapper",
		//	//Format: pkgf("%s_update.go"),
		//},
		//{
		//	Name:   "delete",
		//	Format: pkgf("%s_delete.go"),
		//},
		//{
		//	Name:   "query",
		//	Format: pkgf("%s_query.go"),
		//},
		//{
		//	Name:   "model",
		//	Format: pkgf("%s.go"),
		//},
		//{
		//	Name:   "where",
		//	Format: pkgf("%s/where.go"),
		//},
		//{
		//	Name: "meta",
		//	Format: func(t *Type) string {
		//		return fmt.Sprintf("%s/%s.go", t.Package(), t.Package())
		//	},
		//},
	}
	// GraphTemplates holds the templates applied on the graph.
	MapperTemplate = []GraphTemplate{
		{
			Name:   "mapper",
			//Format: "ent.go",
		},
		//{
		//	Name:   "client",
		//	Format: "client.go",
		//},
		//{
		//	Name:   "context",
		//	Format: "context.go",
		//},
		//{
		//	Name:   "tx",
		//	Format: "tx.go",
		//},
		//{
		//	Name:   "config",
		//	Format: "config.go",
		//},
		//{
		//	Name:   "mutation",
		//	Format: "mutation.go",
		//},
		//{
		//	Name:   "predicate",
		//	Format: "predicate/predicate.go",
		//},
		//{
		//	Name:   "hook",
		//	Format: "hook/hook.go",
		//},
		//{
		//	Name:   "privacy",
		//	Format: "privacy/privacy.go",
		//},
		//{
		//	Name:   "runtime/ent",
		//	Format: "runtime.go",
		//},
		//{
		//	Name:   "enttest",
		//	Format: "enttest/enttest.go",
		//},
		//{
		//	Name:   "runtime/pkg",
		//	Format: "runtime/runtime.go",
		//},
	}
	// templates holds the Go templates for the code generation.
	// the init function below initializes the templates and its
	// funcs to avoid initialization loop.
	templates = template.New("templates")
	// importPkg are the import packages used for code generation.
	importPkg = make(map[string]string)
)

func init() {
	templates.Funcs(Funcs)
	for _, asset := range internal.AssetNames() {
		templates = template.Must(templates.Parse(string(internal.MustAsset(asset))))
	}

	//b := bytes.NewBuffer([]byte("package main\n"))
	//check(templates.ExecuteTemplate(b, "import", Type{Config: &Config{}}), "load imports")
	//f, err := parser.ParseFile(token.NewFileSet(), "", b, parser.ImportsOnly)
	//check(err, "parse imports")
	//for _, spec := range f.Imports {
	//	path, err := strconv.Unquote(spec.Path.Value)
	//	check(err, "unquote import path")
	//	importPkg[filepath.Base(path)] = path
	//}

}

func pkgf(s string) func(t *Type) string {
	return func(t *Type) string { return fmt.Sprintf(s, t.Package()) }
}
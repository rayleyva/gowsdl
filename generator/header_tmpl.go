package generator

var headerTmpl = `
package {{.}}
// Generated by https://github.com/c4milo/gowsdl
// Do not modify
// Copyright (c) 2014, Cloudescape. All rights reserved.
import (
	"time"
	gowsdl "github.com/c4milo/gowsdl/generator"
	{{/*range .Imports*/}}
		{{/*.*/}}
	{{/*end*/}}
)
`

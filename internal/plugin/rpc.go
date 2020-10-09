/*
@Author: nullzz
@Date: 2020/9/30 12:06 下午
@Version: 1.0
*/
package plugin

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func init() {
	generator.RegisterPlugin(new(RpcPlugin))
}

type RpcPlugin struct {
	*generator.Generator
}

func (RpcPlugin) Name() string {
	return "nrpc"
}

func (p *RpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *RpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

func (p *RpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) <= 0 {
		return
	}
	p.genImportCode(file)
}

const (
	reflectPath    = "reflect"
	contextPkgPath = "context"
	stringsPkgPath = "strings"
	//messagePkgPath = "bitbucket.org/funplus/sandwich/message"
)

func (p *RpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P("// TODO: import code")
	p.P(`import "net"`)
	p.P(`import "net/rpc"`)
	//p.P(`import "protoc-gen/internal/plugin"`)

}

/*
@Author: nullzz
@Date: 2020/9/30 4:29 下午
@Version: 1.0
*/
package plugin

import (
	"bytes"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"log"
	"text/template"
)

func (p *RpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	p.P("// TODO: service code, Name = " + svc.GetName())
	spec := p.buildServiceSpec(svc)
	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService))
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}

	p.P(buf.String())
}

type ServiceSpec struct {
	ServiceName string
	Methods     []*ServiceMethodSpec
}

type ServiceMethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
}

func (p *RpcPlugin) buildServiceSpec(svc *descriptor.ServiceDescriptorProto) *ServiceSpec {
	spec := &ServiceSpec{
		ServiceName: generator.CamelCase(svc.GetName()),
	}
	for _, m := range svc.Method {
		spec.Methods = append(spec.Methods, &ServiceMethodSpec{
			MethodName:     generator.CamelCase(m.GetName()),
			InputTypeName:  p.TypeName(p.ObjectNamed(m.GetInputType())),
			OutputTypeName: p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
	}
	return spec
}

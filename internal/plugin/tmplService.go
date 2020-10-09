/*
@Author: nullzz
@Date: 2020/9/30 4:41 下午
@Version: 1.0
*/
package plugin

var tmplService = `
{{$root := .}}

type {{.ServiceName}}Interface interface {
    {{- range $_, $m := .Methods}}
    {{$m.MethodName}}(*{{$m.InputTypeName}}, *{{$m.OutputTypeName}}) error
    {{- end}}
}

func Register{{.ServiceName}}(
    srv *rpc.Server, x {{.ServiceName}}Interface,
) error {
    if err := srv.RegisterName("{{.ServiceName}}", x); err != nil {
        return err
    }
    return nil
}

type {{.ServiceName}}Client struct {
    *rpc.Client
}

var _ {{.ServiceName}}Interface = (*{{.ServiceName}}Client)(nil)

func Dial{{.ServiceName}}(network, address string) (
    *{{.ServiceName}}Client, error,
) {
    c, err := rpc.Dial(network, address)
    if err != nil {
        return nil, err
    }
    return &{{.ServiceName}}Client{Client: c}, nil
}

{{range $_, $m := .Methods}}
func (p *{{$root.ServiceName}}Client) {{$m.MethodName}}(
    in *{{$m.InputTypeName}}, out *{{$m.OutputTypeName}},
) error {
    return p.Client.Call("{{$root.ServiceName}}.{{$m.MethodName}}", in, out)
}
{{end}}


`

//type HelloServiceInterface interface {
//	Hello(in *string, out *string) error
//}
//
//func RegisterHelloService(srv *rpc.Server, x HelloService) error {
//	if err := srv.RegisterName("HelloService", x); err != nil {
//		return err
//	}
//	return nil
//}
//
//type HelloServiceClient struct {
//	*rpc.Client
//}
//
//var _ HelloServiceInterface = (*HelloServiceClient)(nil)
//
//func DialHelloService(network, address string) (*HelloServiceClient, error) {
//	c, err := rpc.Dial(network, address)
//	if err != nil {
//		return nil, err
//	}
//	return &HelloServiceClient{Client: c}, nil
//}
//
//func (p *HelloServiceClient) Hello(in *string, out *string) error {
//	return p.Client.Call("HelloService.Hello", in, out)
//}

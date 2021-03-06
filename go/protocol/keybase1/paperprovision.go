// Auto-generated by avdl-compiler v1.3.21 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/paperprovision.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type PaperProvisionArg struct {
	SessionID  int    `codec:"sessionID" json:"sessionID"`
	Username   string `codec:"username" json:"username"`
	DeviceName string `codec:"deviceName" json:"deviceName"`
	PaperKey   string `codec:"paperKey" json:"paperKey"`
}

type PaperprovisionInterface interface {
	// Performs paper provision.
	// If the current device isn't provisioned, this function will
	// provision it.
	PaperProvision(context.Context, PaperProvisionArg) error
}

func PaperprovisionProtocol(i PaperprovisionInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.paperprovision",
		Methods: map[string]rpc.ServeHandlerDescription{
			"paperProvision": {
				MakeArg: func() interface{} {
					ret := make([]PaperProvisionArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PaperProvisionArg)
					if !ok {
						err = rpc.NewTypeError((*[]PaperProvisionArg)(nil), args)
						return
					}
					err = i.PaperProvision(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type PaperprovisionClient struct {
	Cli rpc.GenericClient
}

// Performs paper provision.
// If the current device isn't provisioned, this function will
// provision it.
func (c PaperprovisionClient) PaperProvision(ctx context.Context, __arg PaperProvisionArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.paperprovision.paperProvision", []interface{}{__arg}, nil)
	return
}

package endpoint

import (
	"context"
	"go-microservice/client"
	"go-microservice/model"

	"github.com/go-kit/kit/endpoint"
)

func MakeVMPostEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		vm := request.(model.VirtualMachine)
		re_vm := client.CreateItem(vm)
		return re_vm, nil
	}
}

func MakeDiskEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		disk := request.(model.Disk)
		re_disk := client.CreateItem(disk)
		return re_disk, nil
	}
}

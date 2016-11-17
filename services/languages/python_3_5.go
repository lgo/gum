package datastores

import (
	"github.com/xLegoz/gum/registry"
)

func apply(registry.ServiceOptions) {

}

func init() {
	registry.RegisterService(Service{
		apply: apply,
	})
}

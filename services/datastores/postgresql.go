package datastores

import (
	"github.com/xLegoz/gum/registry"
)

func GetService() {
	registry.Service{apply}
}

func apply(registry.ServiceOptions) {

}

package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"fmt"

	"github.com/rezafirhan-kargo/kargo-trucks/graph/model"
)

type Resolver struct {
	Trucks    []*model.Truck
	Shipments []*model.Shipment
}

// in graph/resolver.go

func (r *Resolver) Init() {
	for i := 0; i < 20; i++ {
		truck := &model.Truck{
			ID:      fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
			PlateNo: fmt.Sprintf("B %d CD", len(r.Trucks)+1),
		}
		r.Trucks = append(r.Trucks, truck)
	}
}

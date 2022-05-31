package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rezafirhan-kargo/kargo-trucks/graph/generated"
	"github.com/rezafirhan-kargo/kargo-trucks/graph/model"
)

func (r *mutationResolver) SaveTruck(ctx context.Context, id *string, plateNo string) (*model.Truck, error) {
	// panic(fmt.Errorf("not implemented"))
	truck := &model.Truck{
		ID:      fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
		PlateNo: plateNo,
	}
	r.Trucks = append(r.Trucks, truck)
	return truck, nil
}

func (r *mutationResolver) SaveShipment(ctx context.Context, id *string, name string, origin string, destination string, deliveryDate string, truckID *string) (*model.Shipment, error) {
	// panic(fmt.Errorf("not implemented"))
	shipment := &model.Shipment{
		ID:           fmt.Sprintf("SHIPMENT-%d", len(r.Shipments)+1),
		Name:         name,
		Origin:       origin,
		Destination:  destination,
		DeliveryDate: deliveryDate,
		Truck:        &model.Truck{ID: *truckID},
	}
	r.Shipments = append(r.Shipments, shipment)
	return shipment, nil
}

func (r *mutationResolver) DeleteTruck(ctx context.Context, id string) (bool, error) {
	// panic(fmt.Errorf("not implemented"))
	for i, t := range r.Trucks {
		if t.ID == id {
			r.Trucks = append(r.Trucks[:i], r.Trucks[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (r *mutationResolver) DeleteShipment(ctx context.Context, id string) (bool, error) {
	// panic(fmt.Errorf("not implemented"))
	for i, s := range r.Shipments {
		if s.ID == id {
			r.Shipments = append(r.Shipments[:i], r.Shipments[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (r *queryResolver) PaginatedTruck(ctx context.Context, page int, first int) ([]*model.Truck, error) {
	// panic(fmt.Errorf("not implemented"))
	return r.Trucks, nil
}

func (r *queryResolver) PaginatedShipment(ctx context.Context, page int, first int) ([]*model.Shipment, error) {
	// panic(fmt.Errorf("not implemented"))
	return r.Shipments, nil
}

func (r *queryResolver) ExportTruck(ctx context.Context, page int, first int) ([]*model.Truck, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

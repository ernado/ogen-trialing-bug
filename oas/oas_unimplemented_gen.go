// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateCluster implements createCluster operation.
//
// POST /clusters/
func (UnimplementedHandler) CreateCluster(ctx context.Context, req *Cluster) (r *Cluster, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteCluster implements deleteCluster operation.
//
// DELETE /clusters/{cluster_id}
func (UnimplementedHandler) DeleteCluster(ctx context.Context, params DeleteClusterParams) error {
	return ht.ErrNotImplemented
}

// GetCluster implements getCluster operation.
//
// GET /clusters/{cluster_id}
func (UnimplementedHandler) GetCluster(ctx context.Context, params GetClusterParams) (r *Cluster, _ error) {
	return r, ht.ErrNotImplemented
}

// ListClusters implements listClusters operation.
//
// GET /clusters/
func (UnimplementedHandler) ListClusters(ctx context.Context) (r []Cluster, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateCluster implements updateCluster operation.
//
// PUT /clusters/{cluster_id}
func (UnimplementedHandler) UpdateCluster(ctx context.Context, req *Cluster, params UpdateClusterParams) error {
	return ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}

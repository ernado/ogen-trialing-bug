// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteClusterParams is parameters of deleteCluster operation.
type DeleteClusterParams struct {
	// Cluster ID.
	ClusterID uuid.UUID
}

func unpackDeleteClusterParams(packed middleware.Parameters) (params DeleteClusterParams) {
	{
		key := middleware.ParameterKey{
			Name: "cluster_id",
			In:   "path",
		}
		params.ClusterID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeDeleteClusterParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteClusterParams, _ error) {
	// Decode path: cluster_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "cluster_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ClusterID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "cluster_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetClusterParams is parameters of getCluster operation.
type GetClusterParams struct {
	// Cluster ID.
	ClusterID uuid.UUID
}

func unpackGetClusterParams(packed middleware.Parameters) (params GetClusterParams) {
	{
		key := middleware.ParameterKey{
			Name: "cluster_id",
			In:   "path",
		}
		params.ClusterID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeGetClusterParams(args [1]string, argsEscaped bool, r *http.Request) (params GetClusterParams, _ error) {
	// Decode path: cluster_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "cluster_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ClusterID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "cluster_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateClusterParams is parameters of updateCluster operation.
type UpdateClusterParams struct {
	// Cluster ID.
	ClusterID uuid.UUID
}

func unpackUpdateClusterParams(packed middleware.Parameters) (params UpdateClusterParams) {
	{
		key := middleware.ParameterKey{
			Name: "cluster_id",
			In:   "path",
		}
		params.ClusterID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeUpdateClusterParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateClusterParams, _ error) {
	// Decode path: cluster_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "cluster_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ClusterID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "cluster_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

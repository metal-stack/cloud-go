// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1beta1Kubernetes v1beta1 kubernetes
// swagger:model v1beta1.Kubernetes
type V1beta1Kubernetes struct {

	// allow privileged containers
	AllowPrivilegedContainers bool `json:"allowPrivilegedContainers,omitempty"`

	// cluster autoscaler
	ClusterAutoscaler *V1beta1ClusterAutoscaler `json:"clusterAutoscaler,omitempty"`

	// kube API server
	KubeAPIServer *V1beta1KubeAPIServerConfig `json:"kubeAPIServer,omitempty"`

	// kube controller manager
	KubeControllerManager *V1beta1KubeControllerManagerConfig `json:"kubeControllerManager,omitempty"`

	// kube proxy
	KubeProxy *V1beta1KubeProxyConfig `json:"kubeProxy,omitempty"`

	// kube scheduler
	KubeScheduler *V1beta1KubeSchedulerConfig `json:"kubeScheduler,omitempty"`

	// kubelet
	Kubelet *V1beta1KubeletConfig `json:"kubelet,omitempty"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this v1beta1 kubernetes
func (m *V1beta1Kubernetes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterAutoscaler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubeAPIServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubeControllerManager(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubeProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubeScheduler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubelet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Kubernetes) validateClusterAutoscaler(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterAutoscaler) { // not required
		return nil
	}

	if m.ClusterAutoscaler != nil {
		if err := m.ClusterAutoscaler.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterAutoscaler")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Kubernetes) validateKubeAPIServer(formats strfmt.Registry) error {

	if swag.IsZero(m.KubeAPIServer) { // not required
		return nil
	}

	if m.KubeAPIServer != nil {
		if err := m.KubeAPIServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeAPIServer")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Kubernetes) validateKubeControllerManager(formats strfmt.Registry) error {

	if swag.IsZero(m.KubeControllerManager) { // not required
		return nil
	}

	if m.KubeControllerManager != nil {
		if err := m.KubeControllerManager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeControllerManager")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Kubernetes) validateKubeProxy(formats strfmt.Registry) error {

	if swag.IsZero(m.KubeProxy) { // not required
		return nil
	}

	if m.KubeProxy != nil {
		if err := m.KubeProxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeProxy")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Kubernetes) validateKubeScheduler(formats strfmt.Registry) error {

	if swag.IsZero(m.KubeScheduler) { // not required
		return nil
	}

	if m.KubeScheduler != nil {
		if err := m.KubeScheduler.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeScheduler")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Kubernetes) validateKubelet(formats strfmt.Registry) error {

	if swag.IsZero(m.Kubelet) { // not required
		return nil
	}

	if m.Kubelet != nil {
		if err := m.Kubelet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubelet")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Kubernetes) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1Kubernetes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1Kubernetes) UnmarshalBinary(b []byte) error {
	var res V1beta1Kubernetes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
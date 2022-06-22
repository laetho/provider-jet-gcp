/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this CAPool
func (mg *CAPool) GetTerraformResourceType() string {
	return "google_privateca_ca_pool"
}

// GetConnectionDetailsMapping for this CAPool
func (tr *CAPool) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CAPool
func (tr *CAPool) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CAPool
func (tr *CAPool) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CAPool
func (tr *CAPool) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CAPool
func (tr *CAPool) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CAPool
func (tr *CAPool) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CAPool using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CAPool) LateInitialize(attrs []byte) (bool, error) {
	params := &CAPoolParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CAPool) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CAPoolIAMBinding
func (mg *CAPoolIAMBinding) GetTerraformResourceType() string {
	return "google_privateca_ca_pool_iam_binding"
}

// GetConnectionDetailsMapping for this CAPoolIAMBinding
func (tr *CAPoolIAMBinding) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CAPoolIAMBinding
func (tr *CAPoolIAMBinding) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CAPoolIAMBinding
func (tr *CAPoolIAMBinding) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CAPoolIAMBinding
func (tr *CAPoolIAMBinding) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CAPoolIAMBinding
func (tr *CAPoolIAMBinding) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CAPoolIAMBinding
func (tr *CAPoolIAMBinding) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CAPoolIAMBinding using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CAPoolIAMBinding) LateInitialize(attrs []byte) (bool, error) {
	params := &CAPoolIAMBindingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CAPoolIAMBinding) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CAPoolIAMMember
func (mg *CAPoolIAMMember) GetTerraformResourceType() string {
	return "google_privateca_ca_pool_iam_member"
}

// GetConnectionDetailsMapping for this CAPoolIAMMember
func (tr *CAPoolIAMMember) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CAPoolIAMMember
func (tr *CAPoolIAMMember) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CAPoolIAMMember
func (tr *CAPoolIAMMember) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CAPoolIAMMember
func (tr *CAPoolIAMMember) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CAPoolIAMMember
func (tr *CAPoolIAMMember) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CAPoolIAMMember
func (tr *CAPoolIAMMember) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CAPoolIAMMember using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CAPoolIAMMember) LateInitialize(attrs []byte) (bool, error) {
	params := &CAPoolIAMMemberParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CAPoolIAMMember) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CAPoolIAMPolicy
func (mg *CAPoolIAMPolicy) GetTerraformResourceType() string {
	return "google_privateca_ca_pool_iam_policy"
}

// GetConnectionDetailsMapping for this CAPoolIAMPolicy
func (tr *CAPoolIAMPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CAPoolIAMPolicy
func (tr *CAPoolIAMPolicy) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CAPoolIAMPolicy
func (tr *CAPoolIAMPolicy) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CAPoolIAMPolicy
func (tr *CAPoolIAMPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CAPoolIAMPolicy
func (tr *CAPoolIAMPolicy) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CAPoolIAMPolicy
func (tr *CAPoolIAMPolicy) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CAPoolIAMPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CAPoolIAMPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &CAPoolIAMPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CAPoolIAMPolicy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Certificate
func (mg *Certificate) GetTerraformResourceType() string {
	return "google_privateca_certificate"
}

// GetConnectionDetailsMapping for this Certificate
func (tr *Certificate) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Certificate
func (tr *Certificate) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Certificate
func (tr *Certificate) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Certificate
func (tr *Certificate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Certificate
func (tr *Certificate) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Certificate
func (tr *Certificate) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Certificate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Certificate) LateInitialize(attrs []byte) (bool, error) {
	params := &CertificateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Certificate) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CertificateAuthority
func (mg *CertificateAuthority) GetTerraformResourceType() string {
	return "google_privateca_certificate_authority"
}

// GetConnectionDetailsMapping for this CertificateAuthority
func (tr *CertificateAuthority) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CertificateAuthority
func (tr *CertificateAuthority) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CertificateAuthority
func (tr *CertificateAuthority) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CertificateAuthority
func (tr *CertificateAuthority) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CertificateAuthority
func (tr *CertificateAuthority) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CertificateAuthority
func (tr *CertificateAuthority) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CertificateAuthority using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CertificateAuthority) LateInitialize(attrs []byte) (bool, error) {
	params := &CertificateAuthorityParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CertificateAuthority) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CertificateTemplate
func (mg *CertificateTemplate) GetTerraformResourceType() string {
	return "google_privateca_certificate_template"
}

// GetConnectionDetailsMapping for this CertificateTemplate
func (tr *CertificateTemplate) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CertificateTemplate
func (tr *CertificateTemplate) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CertificateTemplate
func (tr *CertificateTemplate) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CertificateTemplate
func (tr *CertificateTemplate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CertificateTemplate
func (tr *CertificateTemplate) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CertificateTemplate
func (tr *CertificateTemplate) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CertificateTemplate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CertificateTemplate) LateInitialize(attrs []byte) (bool, error) {
	params := &CertificateTemplateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CertificateTemplate) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CertificateTemplateIAMBinding
func (mg *CertificateTemplateIAMBinding) GetTerraformResourceType() string {
	return "google_privateca_certificate_template_iam_binding"
}

// GetConnectionDetailsMapping for this CertificateTemplateIAMBinding
func (tr *CertificateTemplateIAMBinding) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CertificateTemplateIAMBinding
func (tr *CertificateTemplateIAMBinding) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CertificateTemplateIAMBinding
func (tr *CertificateTemplateIAMBinding) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CertificateTemplateIAMBinding
func (tr *CertificateTemplateIAMBinding) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CertificateTemplateIAMBinding
func (tr *CertificateTemplateIAMBinding) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CertificateTemplateIAMBinding
func (tr *CertificateTemplateIAMBinding) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CertificateTemplateIAMBinding using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CertificateTemplateIAMBinding) LateInitialize(attrs []byte) (bool, error) {
	params := &CertificateTemplateIAMBindingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CertificateTemplateIAMBinding) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CertificateTemplateIAMMember
func (mg *CertificateTemplateIAMMember) GetTerraformResourceType() string {
	return "google_privateca_certificate_template_iam_member"
}

// GetConnectionDetailsMapping for this CertificateTemplateIAMMember
func (tr *CertificateTemplateIAMMember) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CertificateTemplateIAMMember
func (tr *CertificateTemplateIAMMember) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CertificateTemplateIAMMember
func (tr *CertificateTemplateIAMMember) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CertificateTemplateIAMMember
func (tr *CertificateTemplateIAMMember) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CertificateTemplateIAMMember
func (tr *CertificateTemplateIAMMember) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CertificateTemplateIAMMember
func (tr *CertificateTemplateIAMMember) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CertificateTemplateIAMMember using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CertificateTemplateIAMMember) LateInitialize(attrs []byte) (bool, error) {
	params := &CertificateTemplateIAMMemberParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CertificateTemplateIAMMember) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CertificateTemplateIAMPolicy
func (mg *CertificateTemplateIAMPolicy) GetTerraformResourceType() string {
	return "google_privateca_certificate_template_iam_policy"
}

// GetConnectionDetailsMapping for this CertificateTemplateIAMPolicy
func (tr *CertificateTemplateIAMPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CertificateTemplateIAMPolicy
func (tr *CertificateTemplateIAMPolicy) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CertificateTemplateIAMPolicy
func (tr *CertificateTemplateIAMPolicy) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CertificateTemplateIAMPolicy
func (tr *CertificateTemplateIAMPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CertificateTemplateIAMPolicy
func (tr *CertificateTemplateIAMPolicy) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CertificateTemplateIAMPolicy
func (tr *CertificateTemplateIAMPolicy) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CertificateTemplateIAMPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CertificateTemplateIAMPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &CertificateTemplateIAMPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CertificateTemplateIAMPolicy) GetTerraformSchemaVersion() int {
	return 0
}

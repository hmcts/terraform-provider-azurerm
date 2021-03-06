package templatespecs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2019-06-01-preview/templatespecs"

// BasicArtifact represents a Template Spec artifact.
type BasicArtifact interface {
	AsTemplateArtifact() (*TemplateArtifact, bool)
	AsArtifact() (*Artifact, bool)
}

// Artifact represents a Template Spec artifact.
type Artifact struct {
	// Path - A filesystem safe relative path of the artifact.
	Path *string `json:"path,omitempty"`
	// Kind - Possible values include: 'KindTemplateSpecArtifact', 'KindTemplate'
	Kind Kind `json:"kind,omitempty"`
}

func unmarshalBasicArtifact(body []byte) (BasicArtifact, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["kind"] {
	case string(KindTemplate):
		var ta TemplateArtifact
		err := json.Unmarshal(body, &ta)
		return ta, err
	default:
		var a Artifact
		err := json.Unmarshal(body, &a)
		return a, err
	}
}
func unmarshalBasicArtifactArray(body []byte) ([]BasicArtifact, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	aArray := make([]BasicArtifact, len(rawMessages))

	for index, rawMessage := range rawMessages {
		a, err := unmarshalBasicArtifact(*rawMessage)
		if err != nil {
			return nil, err
		}
		aArray[index] = a
	}
	return aArray, nil
}

// MarshalJSON is the custom marshaler for Artifact.
func (a Artifact) MarshalJSON() ([]byte, error) {
	a.Kind = KindTemplateSpecArtifact
	objectMap := make(map[string]interface{})
	if a.Path != nil {
		objectMap["path"] = a.Path
	}
	if a.Kind != "" {
		objectMap["kind"] = a.Kind
	}
	return json.Marshal(objectMap)
}

// AsTemplateArtifact is the BasicArtifact implementation for Artifact.
func (a Artifact) AsTemplateArtifact() (*TemplateArtifact, bool) {
	return nil, false
}

// AsArtifact is the BasicArtifact implementation for Artifact.
func (a Artifact) AsArtifact() (*Artifact, bool) {
	return &a, true
}

// AsBasicArtifact is the BasicArtifact implementation for Artifact.
func (a Artifact) AsBasicArtifact() (BasicArtifact, bool) {
	return &a, true
}

// AzureResourceBase common properties for all Azure resources.
type AzureResourceBase struct {
	// ID - READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty"`
	// SystemData - READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty"`
}

// Error template Specs error response.
type Error struct {
	Error *ErrorResponse `json:"error,omitempty"`
}

// ErrorAdditionalInfo the resource management error additional info.
type ErrorAdditionalInfo struct {
	// Type - READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty"`
	// Info - READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty"`
}

// ErrorResponse common error response for all Azure Resource Manager APIs to return error details for
// failed operations. (This also follows the OData error response format.)
type ErrorResponse struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The error target.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; The error details.
	Details *[]ErrorResponse `json:"details,omitempty"`
	// AdditionalInfo - READ-ONLY; The error additional info.
	AdditionalInfo *[]ErrorAdditionalInfo `json:"additionalInfo,omitempty"`
}

// ListResult list of Template Specs.
type ListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of Template Specs.
	Value *[]TemplateSpec `json:"value,omitempty"`
	// NextLink - READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for ListResult.
func (lr ListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if lr.Value != nil {
		objectMap["value"] = lr.Value
	}
	return json.Marshal(objectMap)
}

// ListResultIterator provides access to a complete listing of TemplateSpec values.
type ListResultIterator struct {
	i    int
	page ListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListResultIterator) Response() ListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListResultIterator) Value() TemplateSpec {
	if !iter.page.NotDone() {
		return TemplateSpec{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ListResultIterator type.
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return ListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (lr ListResult) IsEmpty() bool {
	return lr.Value == nil || len(*lr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (lr ListResult) hasNextLink() bool {
	return lr.NextLink != nil && len(*lr.NextLink) != 0
}

// listResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lr ListResult) listResultPreparer(ctx context.Context) (*http.Request, error) {
	if !lr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lr.NextLink)))
}

// ListResultPage contains a page of TemplateSpec values.
type ListResultPage struct {
	fn func(context.Context, ListResult) (ListResult, error)
	lr ListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.lr)
		if err != nil {
			return err
		}
		page.lr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListResultPage) NotDone() bool {
	return !page.lr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListResultPage) Response() ListResult {
	return page.lr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListResultPage) Values() []TemplateSpec {
	if page.lr.IsEmpty() {
		return nil
	}
	return *page.lr.Value
}

// Creates a new instance of the ListResultPage type.
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return ListResultPage{
		fn: getNextPage,
		lr: cur,
	}
}

// Properties template Spec properties.
type Properties struct {
	// Description - Template Spec description.
	Description *string `json:"description,omitempty"`
	// DisplayName - Template Spec display name.
	DisplayName *string `json:"displayName,omitempty"`
	// Versions - READ-ONLY; High-level information about the versions within this Template Spec. The keys are the version names. Only populated if the $expand query parameter is set to 'versions'.
	Versions map[string]*VersionInfo `json:"versions"`
}

// MarshalJSON is the custom marshaler for Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.Description != nil {
		objectMap["description"] = p.Description
	}
	if p.DisplayName != nil {
		objectMap["displayName"] = p.DisplayName
	}
	return json.Marshal(objectMap)
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The type of identity that last modified the resource.
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// TemplateArtifact represents a Template Spec artifact containing an embedded Azure Resource Manager
// template.
type TemplateArtifact struct {
	// Template - The Azure Resource Manager template.
	Template interface{} `json:"template,omitempty"`
	// Path - A filesystem safe relative path of the artifact.
	Path *string `json:"path,omitempty"`
	// Kind - Possible values include: 'KindTemplateSpecArtifact', 'KindTemplate'
	Kind Kind `json:"kind,omitempty"`
}

// MarshalJSON is the custom marshaler for TemplateArtifact.
func (ta TemplateArtifact) MarshalJSON() ([]byte, error) {
	ta.Kind = KindTemplate
	objectMap := make(map[string]interface{})
	if ta.Template != nil {
		objectMap["template"] = ta.Template
	}
	if ta.Path != nil {
		objectMap["path"] = ta.Path
	}
	if ta.Kind != "" {
		objectMap["kind"] = ta.Kind
	}
	return json.Marshal(objectMap)
}

// AsTemplateArtifact is the BasicArtifact implementation for TemplateArtifact.
func (ta TemplateArtifact) AsTemplateArtifact() (*TemplateArtifact, bool) {
	return &ta, true
}

// AsArtifact is the BasicArtifact implementation for TemplateArtifact.
func (ta TemplateArtifact) AsArtifact() (*Artifact, bool) {
	return nil, false
}

// AsBasicArtifact is the BasicArtifact implementation for TemplateArtifact.
func (ta TemplateArtifact) AsBasicArtifact() (BasicArtifact, bool) {
	return &ta, true
}

// TemplateSpec template Spec object.
type TemplateSpec struct {
	autorest.Response `json:"-"`
	// Location - The location of the Template Spec. It cannot be changed after Template Spec creation. It must be one of the supported Azure locations.
	Location *string `json:"location,omitempty"`
	// Properties - Template Spec properties.
	*Properties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// ID - READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty"`
	// SystemData - READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty"`
}

// MarshalJSON is the custom marshaler for TemplateSpec.
func (ts TemplateSpec) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ts.Location != nil {
		objectMap["location"] = ts.Location
	}
	if ts.Properties != nil {
		objectMap["properties"] = ts.Properties
	}
	if ts.Tags != nil {
		objectMap["tags"] = ts.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for TemplateSpec struct.
func (ts *TemplateSpec) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				ts.Location = &location
			}
		case "properties":
			if v != nil {
				var properties Properties
				err = json.Unmarshal(*v, &properties)
				if err != nil {
					return err
				}
				ts.Properties = &properties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ts.Tags = tags
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ts.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ts.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ts.Type = &typeVar
			}
		case "systemData":
			if v != nil {
				var systemData SystemData
				err = json.Unmarshal(*v, &systemData)
				if err != nil {
					return err
				}
				ts.SystemData = &systemData
			}
		}
	}

	return nil
}

// UpdateModel template Spec properties to be updated (only tags are currently supported).
type UpdateModel struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// ID - READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty"`
	// SystemData - READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty"`
}

// MarshalJSON is the custom marshaler for UpdateModel.
func (um UpdateModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if um.Tags != nil {
		objectMap["tags"] = um.Tags
	}
	return json.Marshal(objectMap)
}

// VersionInfo high-level information about a Template Spec version.
type VersionInfo struct {
	// Description - READ-ONLY; Template Spec version description.
	Description *string `json:"description,omitempty"`
	// TimeCreated - READ-ONLY; The timestamp of when the version was created.
	TimeCreated *date.Time `json:"timeCreated,omitempty"`
	// TimeModified - READ-ONLY; The timestamp of when the version was last modified.
	TimeModified *date.Time `json:"timeModified,omitempty"`
}

// VersionProperties template Spec Version properties.
type VersionProperties struct {
	// Artifacts - An array of Template Spec artifacts.
	Artifacts *[]BasicArtifact `json:"artifacts,omitempty"`
	// Description - Template Spec version description.
	Description *string `json:"description,omitempty"`
	// Template - The Azure Resource Manager template content.
	Template interface{} `json:"template,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for VersionProperties struct.
func (vp *VersionProperties) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "artifacts":
			if v != nil {
				artifacts, err := unmarshalBasicArtifactArray(*v)
				if err != nil {
					return err
				}
				vp.Artifacts = &artifacts
			}
		case "description":
			if v != nil {
				var description string
				err = json.Unmarshal(*v, &description)
				if err != nil {
					return err
				}
				vp.Description = &description
			}
		case "template":
			if v != nil {
				var templateVar interface{}
				err = json.Unmarshal(*v, &templateVar)
				if err != nil {
					return err
				}
				vp.Template = templateVar
			}
		}
	}

	return nil
}

// VersionsListResult list of Template Specs versions
type VersionsListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of Template Spec versions.
	Value *[]VersionTemplatespecs `json:"value,omitempty"`
	// NextLink - READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for VersionsListResult.
func (vlr VersionsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if vlr.Value != nil {
		objectMap["value"] = vlr.Value
	}
	return json.Marshal(objectMap)
}

// VersionsListResultIterator provides access to a complete listing of VersionTemplatespecs values.
type VersionsListResultIterator struct {
	i    int
	page VersionsListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *VersionsListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VersionsListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *VersionsListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter VersionsListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter VersionsListResultIterator) Response() VersionsListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter VersionsListResultIterator) Value() VersionTemplatespecs {
	if !iter.page.NotDone() {
		return VersionTemplatespecs{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the VersionsListResultIterator type.
func NewVersionsListResultIterator(page VersionsListResultPage) VersionsListResultIterator {
	return VersionsListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (vlr VersionsListResult) IsEmpty() bool {
	return vlr.Value == nil || len(*vlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (vlr VersionsListResult) hasNextLink() bool {
	return vlr.NextLink != nil && len(*vlr.NextLink) != 0
}

// versionsListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (vlr VersionsListResult) versionsListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !vlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(vlr.NextLink)))
}

// VersionsListResultPage contains a page of VersionTemplatespecs values.
type VersionsListResultPage struct {
	fn  func(context.Context, VersionsListResult) (VersionsListResult, error)
	vlr VersionsListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *VersionsListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VersionsListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.vlr)
		if err != nil {
			return err
		}
		page.vlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *VersionsListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page VersionsListResultPage) NotDone() bool {
	return !page.vlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page VersionsListResultPage) Response() VersionsListResult {
	return page.vlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page VersionsListResultPage) Values() []VersionTemplatespecs {
	if page.vlr.IsEmpty() {
		return nil
	}
	return *page.vlr.Value
}

// Creates a new instance of the VersionsListResultPage type.
func NewVersionsListResultPage(cur VersionsListResult, getNextPage func(context.Context, VersionsListResult) (VersionsListResult, error)) VersionsListResultPage {
	return VersionsListResultPage{
		fn:  getNextPage,
		vlr: cur,
	}
}

// VersionTemplatespecs template Spec Version object.
type VersionTemplatespecs struct {
	autorest.Response `json:"-"`
	// Location - The location of the Template Spec Version. It must match the location of the parent Template Spec.
	Location *string `json:"location,omitempty"`
	// VersionProperties - Template Spec Version properties.
	*VersionProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// ID - READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty"`
	// SystemData - READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty"`
}

// MarshalJSON is the custom marshaler for VersionTemplatespecs.
func (vt VersionTemplatespecs) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if vt.Location != nil {
		objectMap["location"] = vt.Location
	}
	if vt.VersionProperties != nil {
		objectMap["properties"] = vt.VersionProperties
	}
	if vt.Tags != nil {
		objectMap["tags"] = vt.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for VersionTemplatespecs struct.
func (vt *VersionTemplatespecs) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				vt.Location = &location
			}
		case "properties":
			if v != nil {
				var versionProperties VersionProperties
				err = json.Unmarshal(*v, &versionProperties)
				if err != nil {
					return err
				}
				vt.VersionProperties = &versionProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				vt.Tags = tags
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				vt.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				vt.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				vt.Type = &typeVar
			}
		case "systemData":
			if v != nil {
				var systemData SystemData
				err = json.Unmarshal(*v, &systemData)
				if err != nil {
					return err
				}
				vt.SystemData = &systemData
			}
		}
	}

	return nil
}

// VersionUpdateModel template Spec Version properties to be updated (only tags are currently supported).
type VersionUpdateModel struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// ID - READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty"`
	// SystemData - READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty"`
}

// MarshalJSON is the custom marshaler for VersionUpdateModel.
func (vum VersionUpdateModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if vum.Tags != nil {
		objectMap["tags"] = vum.Tags
	}
	return json.Marshal(objectMap)
}

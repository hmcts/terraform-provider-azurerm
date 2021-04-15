package web

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CertificateRegistrationProviderClient is the webSite Management Client
type CertificateRegistrationProviderClient struct {
	BaseClient
}

// NewCertificateRegistrationProviderClient creates an instance of the CertificateRegistrationProviderClient client.
func NewCertificateRegistrationProviderClient(subscriptionID string) CertificateRegistrationProviderClient {
	return NewCertificateRegistrationProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCertificateRegistrationProviderClientWithBaseURI creates an instance of the CertificateRegistrationProviderClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewCertificateRegistrationProviderClientWithBaseURI(baseURI string, subscriptionID string) CertificateRegistrationProviderClient {
	return CertificateRegistrationProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListOperations description for Implements Csm operations Api to exposes the list of available Csm Apis under the
// resource provider
func (client CertificateRegistrationProviderClient) ListOperations(ctx context.Context) (result CsmOperationCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateRegistrationProviderClient.ListOperations")
		defer func() {
			sc := -1
			if result.coc.Response.Response != nil {
				sc = result.coc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listOperationsNextResults
	req, err := client.ListOperationsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificateRegistrationProviderClient", "ListOperations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.coc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificateRegistrationProviderClient", "ListOperations", resp, "Failure sending request")
		return
	}

	result.coc, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificateRegistrationProviderClient", "ListOperations", resp, "Failure responding to request")
		return
	}
	if result.coc.hasNextLink() && result.coc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListOperationsPreparer prepares the ListOperations request.
func (client CertificateRegistrationProviderClient) ListOperationsPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.CertificateRegistration/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOperationsSender sends the ListOperations request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateRegistrationProviderClient) ListOperationsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListOperationsResponder handles the response to the ListOperations request. The method always
// closes the http.Response Body.
func (client CertificateRegistrationProviderClient) ListOperationsResponder(resp *http.Response) (result CsmOperationCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listOperationsNextResults retrieves the next set of results, if any.
func (client CertificateRegistrationProviderClient) listOperationsNextResults(ctx context.Context, lastResults CsmOperationCollection) (result CsmOperationCollection, err error) {
	req, err := lastResults.csmOperationCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.CertificateRegistrationProviderClient", "listOperationsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.CertificateRegistrationProviderClient", "listOperationsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificateRegistrationProviderClient", "listOperationsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListOperationsComplete enumerates all values, automatically crossing page boundaries as required.
func (client CertificateRegistrationProviderClient) ListOperationsComplete(ctx context.Context) (result CsmOperationCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateRegistrationProviderClient.ListOperations")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListOperations(ctx)
	return
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsLogRequestBuilder struct{ BaseRequestBuilder }

// Log action undocumented
func (b *WorkbookFunctionsRequestBuilder) Log(reqObj *WorkbookFunctionsLogRequestParameter) *WorkbookFunctionsLogRequestBuilder {
	bb := &WorkbookFunctionsLogRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/log"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsLogRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsLogRequestBuilder) Request() *WorkbookFunctionsLogRequest {
	return &WorkbookFunctionsLogRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsLogRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsLogNorm_DistRequestBuilder struct{ BaseRequestBuilder }

// LogNorm_Dist action undocumented
func (b *WorkbookFunctionsRequestBuilder) LogNorm_Dist(reqObj *WorkbookFunctionsLogNorm_DistRequestParameter) *WorkbookFunctionsLogNorm_DistRequestBuilder {
	bb := &WorkbookFunctionsLogNorm_DistRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/logNorm_Dist"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsLogNorm_DistRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsLogNorm_DistRequestBuilder) Request() *WorkbookFunctionsLogNorm_DistRequest {
	return &WorkbookFunctionsLogNorm_DistRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsLogNorm_DistRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsLogNorm_InvRequestBuilder struct{ BaseRequestBuilder }

// LogNorm_Inv action undocumented
func (b *WorkbookFunctionsRequestBuilder) LogNorm_Inv(reqObj *WorkbookFunctionsLogNorm_InvRequestParameter) *WorkbookFunctionsLogNorm_InvRequestBuilder {
	bb := &WorkbookFunctionsLogNorm_InvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/logNorm_Inv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsLogNorm_InvRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsLogNorm_InvRequestBuilder) Request() *WorkbookFunctionsLogNorm_InvRequest {
	return &WorkbookFunctionsLogNorm_InvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsLogNorm_InvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
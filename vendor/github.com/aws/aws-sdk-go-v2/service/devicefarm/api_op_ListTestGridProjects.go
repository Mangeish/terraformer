// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTestGridProjectsInput struct {
	_ struct{} `type:"structure"`

	// Return no more than this number of results.
	MaxResult *int64 `locationName:"maxResult" min:"1" type:"integer"`

	// From a response, used to continue a paginated listing.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListTestGridProjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTestGridProjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTestGridProjectsInput"}
	if s.MaxResult != nil && *s.MaxResult < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResult", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTestGridProjectsOutput struct {
	_ struct{} `type:"structure"`

	// Used for pagination. Pass into ListTestGridProjects to get more results in
	// a paginated request.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// The list of TestGridProjects, based on a ListTestGridProjectsRequest.
	TestGridProjects []TestGridProject `locationName:"testGridProjects" type:"list"`
}

// String returns the string representation
func (s ListTestGridProjectsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTestGridProjects = "ListTestGridProjects"

// ListTestGridProjectsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets a list of all Selenium testing projects in your account.
//
//    // Example sending a request using ListTestGridProjectsRequest.
//    req := client.ListTestGridProjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListTestGridProjects
func (c *Client) ListTestGridProjectsRequest(input *ListTestGridProjectsInput) ListTestGridProjectsRequest {
	op := &aws.Operation{
		Name:       opListTestGridProjects,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResult",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTestGridProjectsInput{}
	}

	req := c.newRequest(op, input, &ListTestGridProjectsOutput{})
	return ListTestGridProjectsRequest{Request: req, Input: input, Copy: c.ListTestGridProjectsRequest}
}

// ListTestGridProjectsRequest is the request type for the
// ListTestGridProjects API operation.
type ListTestGridProjectsRequest struct {
	*aws.Request
	Input *ListTestGridProjectsInput
	Copy  func(*ListTestGridProjectsInput) ListTestGridProjectsRequest
}

// Send marshals and sends the ListTestGridProjects API request.
func (r ListTestGridProjectsRequest) Send(ctx context.Context) (*ListTestGridProjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTestGridProjectsResponse{
		ListTestGridProjectsOutput: r.Request.Data.(*ListTestGridProjectsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTestGridProjectsRequestPaginator returns a paginator for ListTestGridProjects.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTestGridProjectsRequest(input)
//   p := devicefarm.NewListTestGridProjectsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTestGridProjectsPaginator(req ListTestGridProjectsRequest) ListTestGridProjectsPaginator {
	return ListTestGridProjectsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTestGridProjectsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListTestGridProjectsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTestGridProjectsPaginator struct {
	aws.Pager
}

func (p *ListTestGridProjectsPaginator) CurrentPage() *ListTestGridProjectsOutput {
	return p.Pager.CurrentPage().(*ListTestGridProjectsOutput)
}

// ListTestGridProjectsResponse is the response type for the
// ListTestGridProjects API operation.
type ListTestGridProjectsResponse struct {
	*ListTestGridProjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTestGridProjects request.
func (r *ListTestGridProjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
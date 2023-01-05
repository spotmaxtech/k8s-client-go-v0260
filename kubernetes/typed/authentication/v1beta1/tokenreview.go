/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1beta1 "github.com/spotmaxtech/k8s-api-v0260/authentication/v1beta1"
	v1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
	scheme "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes/scheme"
	rest "github.com/spotmaxtech/k8s-client-go-v0260/rest"
)

// TokenReviewsGetter has a method to return a TokenReviewInterface.
// A group's client should implement this interface.
type TokenReviewsGetter interface {
	TokenReviews() TokenReviewInterface
}

// TokenReviewInterface has methods to work with TokenReview resources.
type TokenReviewInterface interface {
	Create(ctx context.Context, tokenReview *v1beta1.TokenReview, opts v1.CreateOptions) (*v1beta1.TokenReview, error)
	TokenReviewExpansion
}

// tokenReviews implements TokenReviewInterface
type tokenReviews struct {
	client rest.Interface
}

// newTokenReviews returns a TokenReviews
func newTokenReviews(c *AuthenticationV1beta1Client) *tokenReviews {
	return &tokenReviews{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a tokenReview and creates it.  Returns the server's representation of the tokenReview, and an error, if there is any.
func (c *tokenReviews) Create(ctx context.Context, tokenReview *v1beta1.TokenReview, opts v1.CreateOptions) (result *v1beta1.TokenReview, err error) {
	result = &v1beta1.TokenReview{}
	err = c.client.Post().
		Resource("tokenreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tokenReview).
		Do(ctx).
		Into(result)
	return
}

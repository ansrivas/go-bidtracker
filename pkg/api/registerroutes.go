//
// Copyright (c) 2020 Ankur Srivastava
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package api

import (
	"net/url"
	"path"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// RegisterRoutesOption will be a set of routes registry options
type RegisterRoutesOption struct {
	setup func(ro *routesOptions)
}

type routesOptions struct {
	apiVersion  string
	proxyPrefix string
}

func prepareRoutes(baseURL, suffix string) string {
	return path.Join(baseURL, suffix)
}

// RegisterWithAPIVersion returns a RegisterRoutesOption that configures prefix of the API
// This must look like "/api/v1"
func RegisterWithAPIVersion(apiVersion string) RegisterRoutesOption {
	return RegisterRoutesOption{func(ro *routesOptions) {
		ro.apiVersion = apiVersion
	}}
}

// RegisterWithAPIProxyPrefix returns a RegisterRoutesOption that configures proxy-prefix for the API
// This must look like "/production/prefix/stuff"
func RegisterWithAPIProxyPrefix(proxyPrefix string) RegisterRoutesOption {
	return RegisterRoutesOption{func(ro *routesOptions) {
		ro.proxyPrefix = proxyPrefix
	}}
}

// RegisterRoutes registers all the routes available in this application
func RegisterRoutes(api *API, options ...RegisterRoutesOption) error {
	ro := &routesOptions{}
	for _, option := range options {
		option.setup(ro)
	}

	var err error

	var apiVersion *url.URL
	if ro.apiVersion != "" {
		apiVersion, err = url.Parse(ro.apiVersion)
		if err != nil {
			return errors.WithMessage(err, "Failed to register API version")
		}
	}

	var proxyPrefix *url.URL
	if ro.proxyPrefix != "" {
		proxyPrefix, err = url.Parse(ro.proxyPrefix)
		if err != nil {
			return errors.WithMessage(err, "Failed to register proxy URL")
		}
	}

	var finalURL string
	if proxyPrefix != nil {
		finalURL = path.Join(proxyPrefix.String(), apiVersion.String())
	} else {
		finalURL = apiVersion.String()
	}

	log.Info().Msgf("Now registering %s", finalURL)

	api.server.Post(prepareRoutes(finalURL, URLBidItem), api.PostHandlerBidNew)
	api.server.Get(prepareRoutes(finalURL, URLBidGetAll), api.GetHandlerBids)
	api.server.Get(prepareRoutes(finalURL, URLBidGetWinning), api.GetHandlerCurrentWinningBid)
	api.server.Get(prepareRoutes(finalURL, URLUserGetAllBids), api.GetHandlerUserBidGetAll)

	return nil
}

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
	"github.com/ansrivas/bid-tracker/pkg/bidtracker"
	"github.com/gofiber/fiber/v2"
)

// API is the base struct for this implementation
type API struct {
	itemsBid *bidtracker.BidManagement
	server   *fiber.App
}

// NewAPI returns the pointer to a new api instance
// Users can override the settings after invoking the API
func NewAPI() *API {
	return &API{}
}

// NewAPIWithSettings returns the pointer to a new api instance
func NewAPIWithSettings(itemsBid *bidtracker.BidManagement, app *fiber.App) *API {
	return &API{
		itemsBid: itemsBid,
		server:   app,
	}
}

// FiberApp returns the underlying fiber app instance
func (api *API) FiberApp() *fiber.App {
	return api.server
}

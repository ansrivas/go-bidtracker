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
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// GetHandlerUserBidGetAll godoc
// @Summary Get all the bids of a user
// @Description Get all the bids of a user by its uuid
// @Tags User
// @Accept  json
// @Produce  json
// @Param useruuid path string true "useruuid"
// @Success 200 {object} ResponseGetBids
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /users/{useruuid}/bids [get]
// GetHandlerUserBidGetAll handles GET request to get all the bids of a user
func (api *API) GetHandlerUserBidGetAll(c *fiber.Ctx) {

	var useruuid uuid.UUID
	var err error

	if useruuid, err = uuid.FromString(c.Params("useruuid")); err != nil {
		msg := errors.WithMessage(err, "useruuid can not be parsed successfully").Error()
		SendJSON(c, fiber.StatusBadRequest, msg, EmptyResponse)
		return
	}

	bids, err := api.itemsBid.GetBidsByUser(useruuid)
	if err != nil {
		msg := errors.WithMessage(err, "Failed to fetch the bids for given useruuid").Error()
		SendJSON(c, fiber.StatusInternalServerError, msg, EmptyResponse)
		return
	}
	SendJSON(c, fiber.StatusOK, "Success", bids)
	return

}

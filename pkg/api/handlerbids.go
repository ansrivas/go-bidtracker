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
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// PostHandlerBidNew godoc
// @Summary Post a new bid
// @Description get string by ID
// @Tags Bids
// @Accept  json
// @Produce  json
// @Param itemuuid path string true "itemuuid"
// @Param  Bid body bidtracker.Bid true  "Bid"
// @Success 200 {object} ResponseBid
// @Failure 400 {object} Response
// @Failure 422 {object} Response
// @Router /bids [post]
// PostHandlerBidNew handles all the POST requests regarding creation of new bids
func (api *API) PostHandlerBidNew(c *fiber.Ctx) {

	// var itemuuid uuid.UUID
	// var err error

	// if itemuuid, err = uuid.FromString(c.Params("itemuuid")); err != nil {
	// 	msg := errors.WithMessage(err, "itemuuid can not be parsed successfully").Error()
	// 	SendJSON(c, fiber.StatusBadRequest, msg, EmptyResponse)
	// 	return
	// }
	// Associate the current user bid with the itemuuid
	// userBid.ItemUUID = itemuuid

	userBid := new(bidtracker.Bid)
	if err := c.BodyParser(userBid); err != nil {
		msg := errors.WithMessage(err, "json body can not be parsed successfully").Error()
		SendJSON(c, fiber.StatusBadRequest, msg, EmptyResponse)
		return
	}

	if err := api.itemsBid.InsertBid(userBid); err != nil {
		msg := errors.WithMessage(err, "Failed to insert the bid").Error()
		SendJSON(c, fiber.StatusUnprocessableEntity, msg, EmptyResponse)
		return
	}

	SendJSON(c, fiber.StatusOK, "Updated the bid", userBid)
	return
}

// GetHandlerBids godoc
// @Summary Get all current bids on an item
// @Description get string by ID
// @Tags Bids
// @Accept  json
// @Produce  json
// @Param itemuuid path string true "itemuuid"
// @Success 200 {object} ResponseGetBids
// @Failure 400 {object} Response
// @Failure 422 {object} Response
// @Router /bids/{itemuuid} [get]
// GetHandlerBids handles all the GET requests regarding creation of new bids
func (api *API) GetHandlerBids(c *fiber.Ctx) {

	var itemuuid uuid.UUID
	var err error

	if itemuuid, err = uuid.FromString(c.Params("itemuuid")); err != nil {
		msg := errors.WithMessage(err, "itemuuid can not be parsed successfully").Error()
		SendJSON(c, fiber.StatusBadRequest, msg, EmptyResponse)
		return
	}

	bids, err := api.itemsBid.GetBids(itemuuid)
	if err != nil {
		msg := errors.WithMessage(err, "Failed to fetch the list of bids").Error()
		SendJSON(c, fiber.StatusUnprocessableEntity, msg, EmptyResponse)
		return
	}
	SendJSON(c, fiber.StatusOK, "Success", bids)
	return

}

// GetHandlerCurrentWinningBid godoc
// @Summary Get currently winning bids
// @Description get string by ID
// @Tags Bids
// @Accept  json
// @Produce  json
// @Param itemuuid path string true "itemuuid"
// @Success 200 {object} ResponseBid
// @Failure 400 {object} Response
// @Failure 422 {object} Response
// @Router /bids/{itemuuid}/winning [get]
// GetHandlerCurrentWinningBid handles all the GET requests to get currently winning bids
func (api *API) GetHandlerCurrentWinningBid(c *fiber.Ctx) {

	var itemuuid uuid.UUID
	var err error

	if itemuuid, err = uuid.FromString(c.Params("itemuuid")); err != nil {
		msg := errors.WithMessage(err, "itemuuid can not be parsed successfully").Error()
		SendJSON(c, fiber.StatusBadRequest, msg, EmptyResponse)
		return
	}

	bid, err := api.itemsBid.CurrentWinningBid(itemuuid)
	if err != nil {
		msg := errors.WithMessage(err, "Failed to fetch the current winning bid").Error()
		SendJSON(c, fiber.StatusUnprocessableEntity, msg, EmptyResponse)
		return
	}
	SendJSON(c, fiber.StatusOK, "Success", bid)
	return

}

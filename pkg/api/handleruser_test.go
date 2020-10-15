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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ansrivas/bid-tracker/pkg/bidtracker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

type responseAllUserBids struct {
	Status  int
	Message string
	Data    []bidtracker.Bid
}

func TestGetHandlerUserBidGetAll(t *testing.T) {
	assert := assert.New(t)

	itemUUID := uuid.Must(uuid.FromString("b2f9ee6d-79fe-4b14-9c19-35a69a89219a"))
	biddableItems := []uuid.UUID{itemUUID}
	api := NewAPI()
	api.itemsBid = bidtracker.NewBidManagement(biddableItems...)
	api.server = fiber.New()

	api.server.Post(URLBidItem, api.PostHandlerBidNew)
	api.server.Get(URLUserGetAllBids, api.GetHandlerUserBidGetAll)

	tsBidding, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	jsonData := fmt.Sprintf(`{"useruuid":"ae8f7716-867b-4479-b455-c5769e7475ba","itemuuid":"b2f9ee6d-79fe-4b14-9c19-35a69a89219a","timestamp":%d, "amount":%f}`, tsBidding.Unix(), 30.0)
	req1 := httptest.NewRequest("POST", "/bids", bytes.NewBuffer([]byte(jsonData)))
	req1.Header.Add("Content-Type", "application/json")

	tsBidding2, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	jsonData2 := fmt.Sprintf(`{"useruuid":"f475091b-a8f1-4679-83bd-483b616e5260","itemuuid":"b2f9ee6d-79fe-4b14-9c19-35a69a89219a", "timestamp":%d, "amount":%f}`, tsBidding2.Unix(), 31.0)
	req2 := httptest.NewRequest("POST", "/bids", bytes.NewBuffer([]byte(jsonData2)))
	req2.Header.Add("Content-Type", "application/json")

	tsBidding3, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	jsonData3 := fmt.Sprintf(`{"useruuid":"f475091b-a8f1-4679-83bd-483b616e5260", "itemuuid":"b2f9ee6d-79fe-4b14-9c19-35a69a89219a","timestamp":%d, "amount":%f}`, tsBidding3.Unix(), 32.0)
	req3 := httptest.NewRequest("POST", "/bids", bytes.NewBuffer([]byte(jsonData3)))
	req3.Header.Add("Content-Type", "application/json")

	// WHEN
	api.server.Test(req1)
	api.server.Test(req2)
	api.server.Test(req3)

	reqGetAll := httptest.NewRequest("GET", "/users/f475091b-a8f1-4679-83bd-483b616e5260/bids", nil)
	resp, _ := api.server.Test(reqGetAll)

	// THEN
	response := new(responseAllUserBids)
	err := json.NewDecoder(resp.Body).Decode(response)

	// body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.Fail("Failed to read the response from server")
	}

	// want := "f475091b-a8f1-4679-83bd-483b616e5260"
	want := &responseAllUserBids{
		Status:  200,
		Message: "Success",
		Data: []bidtracker.Bid{
			{
				ItemUUID:  itemUUID,
				UserUUID:  uuid.Must(uuid.FromString("f475091b-a8f1-4679-83bd-483b616e5260")),
				Amount:    31.0,
				Timestamp: 1351807721,
			},
			{
				ItemUUID:  itemUUID,
				UserUUID:  uuid.Must(uuid.FromString("f475091b-a8f1-4679-83bd-483b616e5260")),
				Amount:    32.0,
				Timestamp: 1351807721,
			},
		},
	}

	// Do something with results:
	if resp.StatusCode == 200 {
		got := response
		assert.Equal(got, want, "Failed to fetch all the bids for the given user")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		assert.Fail(fmt.Sprintf("Failed response from the server %d. %s", resp.StatusCode, string(body)))
	}
}

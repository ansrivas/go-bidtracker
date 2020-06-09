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
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ansrivas/bid-tracker/pkg/bidtracker"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPostHandlerBidNew(t *testing.T) {
	assert := assert.New(t)

	biddableItems := []uuid.UUID{
		uuid.Must(uuid.FromString("b2f9ee6d-79fe-4b14-9c19-35a69a89219a")),
	}
	api := NewAPI()
	api.itemsBid = bidtracker.NewBidManagement(biddableItems...)
	api.server = fiber.New()

	api.server.Post(URLBidItem, api.PostHandlerBidNew)

	tsBidding, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	jsonData := fmt.Sprintf(`{"useruuid":"ae8f7716-867b-4479-b455-c5769e7475ba", "itemuuid":"b2f9ee6d-79fe-4b14-9c19-35a69a89219a", "timestamp":%d, "amount":%f}`, tsBidding.Unix(), 30.0)
	req := httptest.NewRequest("POST", "/bids", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Add("Content-Type", "application/json")

	// http.Response
	resp, _ := api.server.Test(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.Fail("Failed to read the response from server")
	}

	want := `{"Status":200,"Message":"Updated the bid","Data":{"itemuuid":"b2f9ee6d-79fe-4b14-9c19-35a69a89219a","useruuid":"ae8f7716-867b-4479-b455-c5769e7475ba","timestamp":1351807721,"amount":30}}`
	// Do something with results:
	if resp.StatusCode == 200 {
		got := string(body)
		assert.Equal(want, got, "hello equals hello")
	} else {
		assert.Fail(fmt.Sprintf("Failed response from the server %d. %s", resp.StatusCode, string(body)))
	}
}

func TestPostHandlerBidNewNonProcessable(t *testing.T) {
	assert := assert.New(t)

	biddableItems := []uuid.UUID{
		uuid.Must(uuid.FromString("b2f9ee6d-79fe-4b14-9c19-35a69a89219a")),
	}
	api := NewAPI()
	api.itemsBid = bidtracker.NewBidManagement(biddableItems...)
	api.server = fiber.New()

	api.server.Post(URLBidItem, api.PostHandlerBidNew)

	tsBidding, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	wrongUUID := "wrong-uuid"
	jsonData := fmt.Sprintf(`{"useruuid":"ae8f7716-867b-4479-b455-c5769e7475ba", "itemuuid":"%s", "timestamp":%d, "amount":%f}`, wrongUUID, tsBidding.Unix(), 30.0)
	req := httptest.NewRequest("POST", "/bids", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Add("Content-Type", "application/json")

	// http.Response
	resp, _ := api.server.Test(req)
	assert.Equal(400, resp.StatusCode)
}

func TestGetHandlerBids(t *testing.T) {
	assert := assert.New(t)

	biddableItems := []uuid.UUID{
		uuid.Must(uuid.FromString("b2f9ee6d-79fe-4b14-9c19-35a69a89219a")),
	}
	api := NewAPI()
	api.itemsBid = bidtracker.NewBidManagement(biddableItems...)
	api.server = fiber.New()

	api.server.Post(URLBidItem, api.PostHandlerBidNew)
	api.server.Get(URLBidGetAll, api.GetHandlerBids)

	tsBidding, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	jsonData := fmt.Sprintf(`{"useruuid":"ae8f7716-867b-4479-b455-c5769e7475ba", "itemuuid":"b2f9ee6d-79fe-4b14-9c19-35a69a89219a", "timestamp":%d, "amount":%f}`, tsBidding.Unix(), 30.0)
	req1 := httptest.NewRequest("POST", "/bids", bytes.NewBuffer([]byte(jsonData)))
	req1.Header.Add("Content-Type", "application/json")

	// WHEN
	api.server.Test(req1)
	reqGetAll := httptest.NewRequest("GET", "/bids/b2f9ee6d-79fe-4b14-9c19-35a69a89219a", nil)
	resp, _ := api.server.Test(reqGetAll)

	// THEN
	// response := new(Response)
	// err := json.NewDecoder(resp.Body).Decode(response)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.Fail("Failed to read the response from server")
	}

	want :=
		"{\"Status\":200,\"Message\":\"Success\",\"Data\":[{\"itemuuid\":\"b2f9ee6d-79fe-4b14-9c19-35a69a89219a\",\"useruuid\":\"ae8f7716-867b-4479-b455-c5769e7475ba\",\"timestamp\":1351807721,\"amount\":30}]}"

	// Do something with results:
	if resp.StatusCode == 200 {
		got := string(body)
		assert.Equal(want, got, "hello equals hello")
	} else {
		assert.Fail(fmt.Sprintf("Failed response from the server %d. %s", resp.StatusCode, string(body)))
	}
}

func TestGetHandlerBidsUnprocessable(t *testing.T) {
	assert := assert.New(t)

	biddableItems := []uuid.UUID{
		uuid.Must(uuid.FromString("b2f9ee6d-79fe-4b14-9c19-35a69a89219a")),
	}
	api := NewAPI()
	api.itemsBid = bidtracker.NewBidManagement(biddableItems...)
	api.server = fiber.New()

	api.server.Post(URLBidItem, api.PostHandlerBidNew)
	api.server.Get(URLBidGetAll, api.GetHandlerBids)

	// WHEN
	req1 := httptest.NewRequest("GET", "/bids/bad-uuid", nil)
	resp, _ := api.server.Test(req1)
	assert.Equal(fiber.StatusBadRequest, resp.StatusCode)

	nonExistentUUID := "cef31b6b-cdeb-4035-8d42-a4f33b2d02fe"
	reqGetAll := httptest.NewRequest("GET", fmt.Sprintf("/bids/%s", nonExistentUUID), nil)
	resp, _ = api.server.Test(reqGetAll)

	// THEN
	assert.Equal(fiber.StatusUnprocessableEntity, resp.StatusCode)

}

func TestGetHandlerCurrentWinningBid(t *testing.T) {
	assert := assert.New(t)

	biddableItems := []uuid.UUID{
		uuid.Must(uuid.FromString("b2f9ee6d-79fe-4b14-9c19-35a69a89219a")),
	}
	api := NewAPI()
	api.itemsBid = bidtracker.NewBidManagement(biddableItems...)
	api.server = fiber.New()

	api.server.Post(URLBidItem, api.PostHandlerBidNew)
	api.server.Get(URLBidGetWinning, api.GetHandlerCurrentWinningBid)

	tsBidding, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	jsonData := fmt.Sprintf(`{"useruuid":"ae8f7716-867b-4479-b455-c5769e7475ba", "itemuuid":"b2f9ee6d-79fe-4b14-9c19-35a69a89219a", "timestamp":%d, "amount":%f}`, tsBidding.Unix(), 30.0)
	req1 := httptest.NewRequest("POST", "/bids", bytes.NewBuffer([]byte(jsonData)))
	req1.Header.Add("Content-Type", "application/json")

	tsBidding2, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	jsonData2 := fmt.Sprintf(`{"useruuid":"f475091b-a8f1-4679-83bd-483b616e5260","itemuuid":"b2f9ee6d-79fe-4b14-9c19-35a69a89219a", "timestamp":%d, "amount":%f}`, tsBidding2.Unix(), 31.0)
	req2 := httptest.NewRequest("POST", "/bids", bytes.NewBuffer([]byte(jsonData2)))
	req2.Header.Add("Content-Type", "application/json")

	// WHEN
	api.server.Test(req1)
	api.server.Test(req2)
	reqGetAll := httptest.NewRequest("GET", "/bids/b2f9ee6d-79fe-4b14-9c19-35a69a89219a/winning", nil)
	resp, _ := api.server.Test(reqGetAll)

	// THEN
	// response := new(Response)
	// err := json.NewDecoder(resp.Body).Decode(response)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.Fail("Failed to read the response from server")
	}

	want := "f475091b-a8f1-4679-83bd-483b616e5260"

	// Do something with results:
	if resp.StatusCode == 200 {
		got := string(body)
		assert.Contains(got, want, "Failed to find the winning userid")
	} else {
		assert.Fail(fmt.Sprintf("Failed response from the server %d. %s", resp.StatusCode, string(body)))
	}
}

func TestGetHandlerCurrentWinningBidUnprocessable(t *testing.T) {
	assert := assert.New(t)

	biddableItems := []uuid.UUID{
		uuid.Must(uuid.FromString("b2f9ee6d-79fe-4b14-9c19-35a69a89219a")),
	}
	api := NewAPI()
	api.itemsBid = bidtracker.NewBidManagement(biddableItems...)
	api.server = fiber.New()

	api.server.Post(URLBidItem, api.PostHandlerBidNew)
	api.server.Get(URLBidGetWinning, api.GetHandlerCurrentWinningBid)

	// WHEN
	req1 := httptest.NewRequest("GET", "/bids/bad-uuid/winning", nil)
	resp, _ := api.server.Test(req1)
	assert.Equal(fiber.StatusBadRequest, resp.StatusCode)

	reqGetAll := httptest.NewRequest("GET", "/bids/b2f9ee6d-79fe-4b14-9c19-35a69a89219a/winning", nil)
	resp, _ = api.server.Test(reqGetAll)

	// THEN
	want := fiber.StatusUnprocessableEntity
	assert.Equal(want, resp.StatusCode)
}

//
// Copyright (c) 2019 Ankur Srivastava
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
package bidtracker

import (
	"fmt"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert := assert.New(t)

	itemUUID1 := uuid.Must(uuid.FromString("6aa04324-8aea-4a42-a948-e1da58c86148"))
	itemUUID2 := uuid.Must(uuid.FromString("ae8f7716-867b-4479-b455-c5769e7475ba"))

	allowedItems := []uuid.UUID{itemUUID1, itemUUID2}

	userUUID := uuid.Must(uuid.FromString("8f2f2a79-9091-44fb-9fe3-3eb5f0d76746"))
	bid1 := Bid{
		ItemUUID:  itemUUID1,
		UserUUID:  userUUID,
		Timestamp: time.Now().UTC().Unix(),
		Amount:    30.0,
	}

	bid2 := Bid{
		ItemUUID:  itemUUID2,
		UserUUID:  userUUID,
		Timestamp: time.Now().UTC().Unix(),
		Amount:    30.0,
	}
	items := NewBidManagement(allowedItems...)
	assert.Equal(2, len(items.itemsMap))

	err := items.InsertBid(&bid1)
	assert.Nil(err, fmt.Sprintf("Failed to insert new bid"))
	assert.Equal(1, len(items.itemsMap[itemUUID1].Bids))

	err = items.InsertBid(&bid2)
	assert.Nil(err, fmt.Sprintf("Failed to insert new bid"))
	assert.Equal(1, len(items.itemsMap[itemUUID2].Bids))

	err = items.InsertBid(&bid2)
	assert.Nil(err, fmt.Sprintf("Failed to insert new bid"))
	assert.Equal(2, len(items.itemsMap[itemUUID2].Bids))
}

func TestGetBids(t *testing.T) {
	assert := assert.New(t)

	itemUUID1 := uuid.Must(uuid.FromString("6aa04324-8aea-4a42-a948-e1da58c86148"))
	itemUUID2 := uuid.Must(uuid.FromString("ae8f7716-867b-4479-b455-c5769e7475ba"))

	allowedItems := []uuid.UUID{itemUUID1, itemUUID2}

	userUUID := uuid.Must(uuid.FromString("8f2f2a79-9091-44fb-9fe3-3eb5f0d76746"))
	bid1 := Bid{
		ItemUUID:  itemUUID1,
		UserUUID:  userUUID,
		Timestamp: time.Now().UTC().Unix(),
		Amount:    30.0,
	}

	bid2 := Bid{
		ItemUUID:  itemUUID2,
		UserUUID:  userUUID,
		Timestamp: time.Now().UTC().Unix(),
		Amount:    30.0,
	}

	items := NewBidManagement(allowedItems...)
	assert.Equal(2, len(items.itemsMap))

	err := items.InsertBid(&bid1)
	assert.Nil(err, fmt.Sprintf("Failed to insert new bid"))
	assert.Equal(1, len(items.itemsMap[itemUUID1].Bids))

	bids, err := items.GetBids(itemUUID1)
	assert.Nil(err, fmt.Sprintf("Failed to fetch all the bids for itemUUID1"))
	assert.Equal(1, len(bids))

	err = items.InsertBid(&bid2)
	assert.Nil(err, fmt.Sprintf("Failed to insert new bid"))
	assert.Equal(1, len(items.itemsMap[itemUUID2].Bids))
	bids, err = items.GetBids(itemUUID2)
	assert.Nil(err, fmt.Sprintf("Failed to fetch all the bids for itemUUID2"))
	assert.Equal(1, len(bids))

	err = items.InsertBid(&bid2)
	assert.Nil(err, fmt.Sprintf("Failed to insert new bid"))
	bids, err = items.GetBids(itemUUID2)
	assert.Nil(err, fmt.Sprintf("Failed to fetch all the bids for itemUUID2"))
	assert.Equal(2, len(bids))

}

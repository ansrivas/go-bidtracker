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

// Package bidtracker ...
package bidtracker

import (
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
)

// ItemBidState represents the current state of an item
type ItemBidState struct {
	ItemID             uuid.UUID
	Bids               []Bid
	currentWinndingBid *Bid
}

// UserBids represents the state of bids for a user
type UserBids struct {
	Bids []Bid
}

// BidManagement is a wrapper to store current items bid statte
type BidManagement struct {
	sync.Mutex
	itemsMap   map[uuid.UUID]ItemBidState
	userBidMap map[uuid.UUID]UserBids
}

// NewBidManagement creates a new instance of BidManagement struct
func NewBidManagement(allowedItemUUIDs ...uuid.UUID) *BidManagement {
	itemsMap := make(map[uuid.UUID]ItemBidState, len(allowedItemUUIDs))
	useBidMap := make(map[uuid.UUID]UserBids)
	for i := 0; i < len(allowedItemUUIDs); i++ {
		itemID := allowedItemUUIDs[i]
		itemsMap[itemID] = ItemBidState{
			ItemID: itemID,
			Bids:   []Bid{},
		}
	}
	return &BidManagement{
		itemsMap:   itemsMap,
		userBidMap: useBidMap,
	}
}

// CurrentWinningBid will return the current winning bid for the given itemuuid
func (ibm *BidManagement) CurrentWinningBid(itemuuid uuid.UUID) (*Bid, error) {
	ibm.Lock()
	defer ibm.Unlock()

	itemMetaInfo, ok := ibm.itemsMap[itemuuid]
	if !ok {
		return nil, fmt.Errorf("Requested item is not available for bidding. %s", itemuuid)
	}

	if itemMetaInfo.currentWinndingBid != nil {
		return itemMetaInfo.currentWinndingBid, nil
	}

	return nil, fmt.Errorf("No currentbid found for requested uuid %s", itemuuid)
}

// InsertBid a new bid for the provided item.
func (ibm *BidManagement) InsertBid(bid *Bid) error {
	ibm.Lock()
	defer ibm.Unlock()

	itemMetaInfo, ok := ibm.itemsMap[bid.ItemUUID]
	if !ok {
		return fmt.Errorf("Requested item is not available for bidding. %s", bid.ItemUUID)
	}

	// Update the current winning bid
	if itemMetaInfo.currentWinndingBid == nil {
		itemMetaInfo.currentWinndingBid = bid
	} else {
		if itemMetaInfo.currentWinndingBid.Amount < bid.Amount {
			itemMetaInfo.currentWinndingBid = bid
		}
	}

	// Update the user-section
	userBidInfo, ok := ibm.userBidMap[bid.UserUUID]
	if !ok {
		// Insert the value first time if its not found
		ibm.userBidMap[bid.UserUUID] = UserBids{
			Bids: []Bid{
				*bid,
			},
		}
	} else {
		userBidInfo.Bids = append(userBidInfo.Bids, *bid)
		ibm.userBidMap[bid.UserUUID] = userBidInfo

	}

	itemMetaInfo.Bids = append(itemMetaInfo.Bids, *bid)
	ibm.itemsMap[bid.ItemUUID] = itemMetaInfo
	return nil
}

// GetBids get bids for a given item
func (ibm *BidManagement) GetBids(itemuuid uuid.UUID) ([]Bid, error) {
	ibm.Lock()
	defer ibm.Unlock()

	itemMetaInfo, ok := ibm.itemsMap[itemuuid]
	if !ok {
		return nil, fmt.Errorf("Requested item is not available for bidding. %s", itemuuid)
	}
	return itemMetaInfo.Bids, nil
}

// GetBidsByUser fetches all the bids for a given useruuid
func (ibm *BidManagement) GetBidsByUser(useruuid uuid.UUID) ([]Bid, error) {
	ibm.Lock()
	defer ibm.Unlock()

	userBidInfo, ok := ibm.userBidMap[useruuid]
	if !ok {
		return nil, fmt.Errorf("No user found with uuid%v", useruuid)
	}

	return userBidInfo.Bids, nil
}

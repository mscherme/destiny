package bungie

import "errors"

type equipItem struct {
	ItemID         string `json:"itemId"`
	CharacterID    string `json:"characterId"`
	MembershipType int64  `json:"membershipType"`
}

func (b *API) EquipItem(c *Character, i *Item) error {
	return b.post("EquipItem/", &equipItem{
		i.ItemID,
		c.CharacterBase.CharacterID,
		c.CharacterBase.MembershipType})
}

type transferItem struct {
	ItemReferenceHash int64  `json:"itemReferenceHash"`
	StackSize         int64  `json:"stackSize"`
	TransferToVault   bool   `json:"transferToVault"`
	ItemID            string `json:"itemId"`
	CharacterID       string `json:"characterId"`
	MembershipType    int64  `json:"membershipType"`
}

func (b *API) TransferItem(c *Character, i *Item, toVault bool) error {
	if i.Quantity > 1 {
		return errors.New("Cannot transfer stackable item")
	}
	return b.post("TransferItem/", &transferItem{
		ItemReferenceHash: i.ItemHash,
		ItemID:            i.ItemID,
		CharacterID:       c.CharacterBase.CharacterID,
		MembershipType:    c.CharacterBase.MembershipType,
		TransferToVault:   toVault})
}

func (b *API) TransferStack(c *Character, i *Item, toVault bool, stackSize int64) error {
	if i.ItemID != "0" {
		return errors.New("Item is not stackable")
	}
	if i.Quantity < stackSize {
		return errors.New("Not enough of item")
	}
	return b.post("TransferItem/", &transferItem{
		ItemReferenceHash: i.ItemHash,
		ItemID:            i.ItemID,
		CharacterID:       c.CharacterBase.CharacterID,
		MembershipType:    c.CharacterBase.MembershipType,
		TransferToVault:   toVault,
		StackSize:         stackSize})
}

func (b *API) TransferStackFromItemHash(c *Character, itemHash int64, toVault bool, stackSize int64) error {
	return b.post("TransferItem/", &transferItem{
		ItemReferenceHash: itemHash,
		ItemID:            "0",
		CharacterID:       c.CharacterBase.CharacterID,
		MembershipType:    c.CharacterBase.MembershipType,
		TransferToVault:   toVault,
		StackSize:         stackSize})
}

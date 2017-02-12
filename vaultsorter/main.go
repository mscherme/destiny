package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"destiny/bungie"
)

var (
	listItems      = flag.Bool("list_items_only", true, "set to false to actually sort")
	vaultSection   = flag.String("vault_section", "", "Armor or Weapon")
	engramsToEnd   = flag.Bool("engrams_to_end", false, "")
	cookieLocation = flag.String("cookie_path", "/home/mscherme/bungie/cookie", "")
	xcsrfLocation  = flag.String("xcsrf_path", "/home/mscherme/bungie/xcsrf", "")
	gamertag       = flag.String("gamertag", "mscherme", "")

	b *bungie.API
)

var bucketTypeHashToSection = map[int64]string{
	2465295065: "Weapon",
	953998645:  "Weapon",
	1498876634: "Weapon",

	14239492:   "Armor",
	1585787867: "Armor",
	20886954:   "Armor",
	3448274439: "Armor",
	3551918588: "Armor",
}

func readCookie() {
	data, err := ioutil.ReadFile(*cookieLocation)
	if err != nil {
		log.Fatal(err)
	}
	b.SetCookie(string(data))

	data, err = ioutil.ReadFile(*xcsrfLocation)
	if err != nil {
		log.Fatal(err)
	}
	b.SetXCSRF(string(data))
}

func listAllItemsInInventory(inv *bungie.Inventory, section string) {
	for _, i := range inv.Items {
		if i.ItemID == "0" {
			continue
		}
		info := lookup(i.ItemHash)
		if *listItems && section == bucketTypeHashToSection[info.BucketTypeHash] {
			fmt.Printf("%d, // %s\n", i.ItemHash, info.ItemName)
		}
	}
}

func listAllItems(account *bungie.Account) {
	for _, section := range []string{"Armor", "Weapon"} {
		fmt.Println(section)
		listAllItemsInInventory(account.Inventory, section)
		for _, c := range account.Characters {
			inv, err := b.CharacterInventory(c)
			if err != nil {
				log.Fatal(err)
			}
			listAllItemsInInventory(inv, section)
		}
		fmt.Println()
	}
}

func main() {
	flag.Parse()

	var err error
	b, err = bungie.New()
	if err != nil {
		log.Fatal(err)
	}
	readCookie()
	account, err := b.LookupAccount(bungie.XBOX, *gamertag)
	if err != nil {
		log.Fatal(err)
	}

	if *listItems {
		listAllItems(account)
		return
	}

	storage := account.Characters[0]

	inv := account.Inventory
	var itemMap = map[int64][]*bungie.Item{}
	for _, i := range inv.Items {
		itemMap[i.ItemHash] = append(itemMap[i.ItemHash], i)
		if i.ItemID == "0" {
			continue
		}
		info := lookup(i.ItemHash)
		section := bucketTypeHashToSection[info.BucketTypeHash]
		if *listItems && section == *vaultSection {
			fmt.Printf("%d, // %s\n", i.ItemHash, info.ItemName)
		}
	}

	if _, ok := order[*vaultSection]; ok {
		for _, hash := range order[*vaultSection] {
			// Ignore error if the item doesn't exist.
			moveAllToEnd(storage, itemMap[hash])
		}
		for _, i := range inv.Items {
			info := lookup(i.ItemHash)
			if *engramsToEnd && strings.Contains(info.ItemName, "Engram") {
				continue
			}
			if bucketTypeHashToSection[info.BucketTypeHash] == *vaultSection &&
				location(*vaultSection, i.ItemHash) == -1 {
				err = moveToEnd(storage, i)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	if *engramsToEnd {
		for _, i := range inv.Items {
			info := lookup(i.ItemHash)
			if strings.Contains(info.ItemName, "Engram") {
				err := moveToEnd(storage, i)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
func moveAllToEnd(c *bungie.Character, items []*bungie.Item) error {
	for _, i := range items {
		err := moveToEnd(c, i)
		if err != nil {
			return err
		}
	}
	return nil
}

func moveToEnd(c *bungie.Character, i *bungie.Item) error {
	err := b.TransferItem(c, i, false)
	if err != nil {
		return err
	}
	return b.TransferItem(c, i, true)
}

func location(t string, hash int64) int {
	for i, x := range order[t] {
		if x == hash {
			return i
		}
	}
	return -1
}

func lookup(itemHash int64) *bungie.InventoryItem {
	i, err := b.ManifestInventoryItem(itemHash)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

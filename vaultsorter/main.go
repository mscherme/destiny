package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"destiny/bungie"
	"destiny/gear"
)

var (
	listItemsOnly  = flag.Bool("list_items_only", true, "set to false to actually sort")
	vaultSection   = flag.String("vault", "Weapon", "Armor, Weapon or Both")
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

	storage := account.Characters[0]

	inv := account.Inventory.Items
	sort.Sort(sortableItemList(inv))
	for _, i := range inv {
		info := lookup(i.ItemHash)
		section := bucketTypeHashToSection[info.BucketTypeHash]
		if section == "" {
			continue
		}

		if section == *vaultSection || *vaultSection == "Both" {
			fmt.Print(info.ItemName)
			if s := gear.SetForItem(info); s != 0 {
				fmt.Printf(" (%s)", s)
			}
			fmt.Println()

			if !*listItemsOnly {
				err = moveToEnd(storage, i)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func moveToEnd(c *bungie.Character, i *bungie.Item) error {
	err := b.TransferItem(c, i, false)
	if err != nil {
		return err
	}
	return b.TransferItem(c, i, true)
}

func lookup(i int64) *bungie.InventoryItem {
	info, err := b.ManifestInventoryItem(i)
	if err != nil {
		log.Fatal(err)
	}
	return info
}

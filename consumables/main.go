package main

// This is a tool that organizes consumables.
// It makes sure that each character has:
// - 50 of each synth type,
// - 200 of each resource type
// - 50 motes of light
// - 50 SIVA Key Fragments

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"destiny/bungie"
)

var (
	cookie, xcsrf  string
	cookieLocation = flag.String("cookie_path", "/home/mscherme/bungie/cookie", "")
	xcsrfLocation  = flag.String("xcsrf_path", "/home/mscherme/bungie/xcsrf", "")
	gamertag       = flag.String("gamertag", "mscherme", "")
)

func readCookie() {
	data, err := ioutil.ReadFile(*cookieLocation)
	if err != nil {
		log.Fatal(err)
	}
	cookie = string(data)

	data, err = ioutil.ReadFile(*xcsrfLocation)
	if err != nil {
		log.Fatal(err)
	}
	xcsrf = string(data)
}

func main() {
	flag.Parse()
	readCookie()

	var err error
	b, err := bungie.New()
	if err != nil {
		log.Fatal(err)
	}
	b.SetCookie(cookie)
	b.SetXCSRF(xcsrf)
	account, err := b.LookupAccount(bungie.XBOX, *gamertag)
	if err != nil {
		log.Fatal(err)
	}
	balanceConsumables(b, account)
}

var targetSizes = map[int64]int64{
	1898539128: 200, // Weapon Parts
	3164836592: 200, // Wormspore
	2254123540: 200, // Spirit Bloom
	1542293174: 200, // Armor Materials
	3242866270: 200, // Relic Iron
	2882093969: 200, // Spinmetal
	1797491610: 200, // Helium Filaments
	937555249:  50,  // Mote of Light
	2206724918: 50,  // SIVA Key Fragments
	211861343:  50,  // Heavy Ammo Synthesis
	928169143:  50,  // Special Ammo Synthesis
	2180254632: 50,  // Ammo Synthesis
	1500229041: 5,   // Crucible Reputation Booster
	2220921114: 5,   // Vanguard Reputation Booster
	1603376703: 5,   // House of Judgment Reputation Booster
	1826822442: 4,   // Iron Lords' Legacy
}

var ignore = map[int64]bool{
	110672197:  true, // Silver Dust
	1389966135: true, // Perfected SIVA Offering
	2575095886: true, // Splicer Intel Relay
	2575095887: true, // Splicer Intel Relay
	2634463554: true, // Perfected Ornament
	27411484:   true, // Dead Orbit
	2954371221: true, // New Monarchy
	3026483582: true, // Skyburners Command Beacon
	3345355735: true, // Days of Iron Ornament
	342707700:  true, // Stolen Rune
	342707701:  true, // Reciprocal Rune
	342707703:  true, // Antiquated Rune
	3558442274: true, // Paper Fortune
	3705287264: true, // Yellow Chroma
	3705287265: true, // White Chroma
	3705287266: true, // Red Chroma
	3705287267: true, // Blue Chroma
	3771657596: true, // Perfected SIVA Offering
	4244618453: true, // Splicer Key
	614056762:  true, // Skeleton Key
	75513258:   true, // Perfected SIVA Offering
	894761024:  true, // Green Chroma
	894761026:  true, // Orange Chroma
	894761027:  true, // Magenta Chroma
	969832704:  true, // Future War Cult
}

func name(b *bungie.API, itemHash int64) string {
	i, err := b.ManifestInventoryItem(itemHash)
	if err != nil {
		log.Fatal(err)
	}
	return i.ItemName
}

func abs(x int64) int64 {
	if x < 0 {
		x = -x
	}
	return x
}

func balanceConsumables(b *bungie.API, account *bungie.Account) {
	for _, c := range account.Characters {
		inventory, err := b.CharacterInventory(c)
		if err != nil {
			log.Fatal(err)
		}
		targets := map[int64]int64{}
		for k, v := range targetSizes {
			targets[k] = v
		}
		sizes := map[int64]int64{}
		for _, i := range inventory.Items {
			if ignore[i.ItemHash] || i.ItemID != "0" {
				continue
			}
			targets[i.ItemHash] = targetSizes[i.ItemHash]
			sizes[i.ItemHash] += i.Quantity

			if _, ok := targetSizes[i.ItemHash]; !ok {
				fmt.Printf("%d: 0, // %s\n", i.ItemHash, name(b, i.ItemHash))
			}
		}
		for hash, size := range targets {
			current := sizes[hash]
			if delta := abs(current - size); delta > 0 {
				err = b.TransferStackFromItemHash(c, hash, current > size, delta)
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
}

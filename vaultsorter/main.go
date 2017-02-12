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
	vaultSection   = flag.String("vault_section", "Armor", "Armor or Weapon")
	engramsToEnd   = flag.Bool("engrams_to_end", false, "")
	cookie, xcsrf   string
	cookieLocation = flag.String("cookie_path", "/home/mscherme/bungie/cookie", "")
	xcsrfLocation   = flag.String("xcsrf_path", "/home/mscherme/bungie/xcsrf", "")
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
	b, err = bungie.New()
	if err != nil {
		log.Fatal(err)
	}
	b.SetCookie(cookie)
	b.SetXCSRF(xcsrf)
	account, err := b.LookupAccount(bungie.XBOX, *gamertag)
	if err != nil {
		log.Fatal(err)
	}

	storage := account.Characters[0]

	inv := account.Inventory
	var itemMap = map[int64][]*bungie.Item{}
	for _, i := range inv.Items {
		itemMap[i.ItemHash] = append(itemMap[i.ItemHash], i)
		if i.ItemID != "0" {
			info := lookup(i.ItemHash)
			section := bucketTypeHashToSection[info.BucketTypeHash]
			if *listItems && section == *vaultSection {
				fmt.Printf("%d, // %s\n", i.ItemHash, info.ItemName)
			}
		}
	}
	if !*listItems {
		for _, hash := range order[*vaultSection] {
			err = moveAllToEnd(storage, itemMap[hash])
			if err != nil {
				log.Fatal(err)
			}
		}
		for _, i := range inv.Items {
			info := lookup(i.ItemHash)
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
		err = moveEngramsToEnd(inv, storage)
		if err != nil {
			log.Fatal(err)
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

var order = map[string][]int64{
	"Weapon": {
		3074713346, // Atheon's Epilogue
		3892023023, // Fatebringer
		1603229152, // Praedyth's Timepiece
		2149012811, // Vision of Confluence
		1267053937, // Praetorian Foil
		892741686,  // Found Verdict
		3695068318, // Praedyth's Revenge
		152628833,  // Corrective Measure
		3807770941, // Hezen Vengeance

		4144666151, // Abyss Defiant
		437329200,  // Word of Crota
		868574327,  // Oversoul Edict
		4252504452, // Fang of Ir Yût
		1267147308, // Light of the Abyss
		3615265777, // Swordbreaker
		560601823,  // Black Hammer
		788203480,  // Song of Ir Yût
		2361858758, // Hunger of Crota

		2235712584, // Wolfslayer's Claw
		270944924,  // Six Dreg Pride II
		1548878643, // Aegis of the Kell II
		3093674676, // Wolves' Leash II
		2763938995, // Servant of Aksor
		391452304,  // Wolfborne Oath
		190731588,  // Shadow of Veils
		1409833631, // Chain of Orbiks-Fel
		3772051159, // The Last Rebellion

		346443849,  // Vex Mythoclast
		2809229973, // Necrochasm
		346443850,  // Pocket Infinity
		261727865,  // The Stranger's Rifle
		2609120348, // Murmur
		713097484,  // Vestian Dynasty
		1011266581, // Three Little Words
		3384077431, // LDR 5001
		2853794413, // Jolder's Hammer
	},
	"Armor": {
		1096028869, // Prime Zealot Helm
		1835128980, // Prime Zealot Gloves
		3833808556, // Prime Zealot Cuirass
		1698410142, // Prime Zealot Greaves
		2237496545, // Shattered Vault Cloak
		2147998057, // Battlecage of Kabr
		3851493600, // Kabr's Brazen Grips
		3367833896, // Kabr's Wrath
		2504856474, // Kabr's Forceful Greaves
		774963973,  // Light of the Great Prism
		1883484055, // Gloves of the Hezen Lords
		3267664569, // Tread of the Hezen Lords
		991704636,  // Fragment of the Prime

		1311326450, // Unyielding Casque
		1736102875, // Dogged Gage
		1261228341, // Relentless Harness
		186143053,  // Tireless Striders
		4253790216, // Shroud of Flies
		1898281764, // Willbreaker's Watch
		1462595581, // Willbreaker's Fists
		2450884227, // Willbreaker's Resolve
		3786747679, // Willbreaker's Greaves
		1349707258, // Mark of the Pit
		2477121987, // Deathsinger's Gaze
		3148626578, // Deathsinger's Grip
		3009953622, // Deathsinger's Mantle
		3549968172, // Deathsinger's Herald
		2339580799, // Bone Circlet

		2303881503, // Kellhunter's Sight
		2662103142, // Kellhunter's Rally
		3661659402, // Kellhunter's Blood
		2221235552, // Kellhunter's Path
		3509787,    // Kellhunter's Hood
		2884887211, // Kellslayer's Helm
		1165567642, // Kellslayer's Grips
		4007053294, // Kellslayer's Cuirass
		2887970516, // Kellslayer's Greaves
		454534839,  // Mark of a Kellslayer
		3620910776, // Kellbreaker's Mind
		3660153609, // Kellbreaker's Gloves
		4013007887, // Kellbreaker's Cloak
		3910621915, // Kellbreaker's Path
		1006775966, // "Dread's Enmity"
	},
}

func lookup(itemHash int64) *bungie.InventoryItem {
	i, err := b.ManifestInventoryItem(itemHash)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func moveEngramsToEnd(inv *bungie.Inventory, c *bungie.Character) error {
	for _, i := range inv.Items {
		info := lookup(i.ItemHash)
		if strings.Contains(info.ItemName, "Engram") {
			err := moveToEnd(c, i)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

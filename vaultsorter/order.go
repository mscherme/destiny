package main

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
		2486746566, // Façade of the Hezen Lords
		1883484055, // Gloves of the Hezen Lords
		4079606241, // Cuirass of the Hezen Lords
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
/*
621603243, // GENESIS CHAIN~ (Wrath of the Machine)
2154053164, // FEVER AND REMEDY~ (Wrath of the Machine)
3632330099, // STEEL MEDULLA~ (Wrath of the Machine)
2542033072, // CHAOS DOGMA~ (Wrath of the Machine)
3598793896, // ETHER NOVA~ (Wrath of the Machine)
2125403517, // QUANTIPLASM~ (Wrath of the Machine)
3569444312, // ZEAL VECTOR~ (Wrath of the Machine)
2001493563, // EX MACHINA~ (Wrath of the Machine)
242628276, // IF MATERIA~ (Wrath of the Machine)
1784034858, // SOUND AND FURY~ (Wrath of the Machine)

1047302226, // Spliced Nanomania Cloak (Wrath of the Machine)
1299613399, // Spliced Red Miasma Hood (Wrath of the Machine)
155972845, // Spliced Red Miasma Bond (Wrath of the Machine)
1680970216, // Spliced Red Miasma Boots (Wrath of the Machine)
1823306243, // Spliced Nanomania Vest (Wrath of the Machine)
2011530291, // Spliced Cosmoclast Greaves (Wrath of the Machine)
2503707046, // Spliced Red Miasma Gloves (Wrath of the Machine)
2679967696, // Spliced Red Miasma Robes (Wrath of the Machine)
2903199471, // Spliced Nanomania Grasps (Wrath of the Machine)
307172369, // Spliced Cosmoclast Plate (Wrath of the Machine)
3138343136, // Spliced Cosmoclast Helm (Wrath of the Machine)
3529294924, // Spliced Cosmoclast Mark (Wrath of the Machine)
3976262649, // Spliced Cosmoclast Gauntlets (Wrath of the Machine)
661529958, // Spliced Nanomania Mask (Wrath of the Machine)
691310313, // Spliced Nanomania Boots (Wrath of the Machine)

3471865173, // Harrowed Darkhollow Mask (King's Fall)
2302693612, // Harrowed Darkhollow Grasps (King's Fall)
3907799186, // Harrowed Darkhollow  Chiton (King's Fall)
2549035182, // Harrowed Darkhollow Treads (King's Fall)
2242715339, // Cloak of Seven Bones (King's Fall)
1846107924, // Harrowed Mouth of Ur (King's Fall)
521951205, // Harrowed Grasp of Eir (King's Fall)
372855005, // Harrowed Chasm of Yuul (King's Fall)
2028036495, // Harrowed Path of Xol (King's Fall)
1658688592, // "Worm Gods' Boon" (King's Fall)
1245063911, // Harrowed War Numen's Crown (King's Fall)
217447094, // Harrowed War Numen's Fist (King's Fall)
3176903680, // Harrowed War Numen's Chest (King's Fall)
1601524312, // Harrowed War Numen's Boots (King's Fall)
130578781, // Mark of the Old Challenge (King's Fall)
1457207756, // Harrowed Anguish of Drystan (King's Fall)
962497239, // Harrowed Zaouli's Bane (King's Fall)
2918358303, // Harrowed Doom of Chelchis (King's Fall)
3042333087, // Harrowed Midha's Reckoning (King's Fall)
2536361592, // Harrowed Smite of Merain (King's Fall)
1551744703, // Harrowed Qullim's Terminus (King's Fall)
1397524041, // Harrowed Elulim's Frenzy (King's Fall)
3919765140, // Harrowed Defiance of Yasmin (King's Fall)
2201079122, // Harrowed Silence of A'arn (King's Fall)
*/

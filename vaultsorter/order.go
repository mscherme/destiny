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

		//346443849,  // Vex Mythoclast
		//2809229973, // Necrochasm
		//346443850,  // Pocket Infinity
		//261727865,  // The Stranger's Rifle
		//2609120348, // Murmur
		//713097484,  // Vestian Dynasty
		//1011266581, // Three Little Words
		//3384077431, // LDR 5001
		//2853794413, // Jolder's Hammer

		1457207756, // Harrowed Anguish of Drystan (King's Fall)
		962497239,  // Harrowed Zaouli's Bane (King's Fall)
		2536361592, // Harrowed Smite of Merain (King's Fall)
		2918358303, // Harrowed Doom of Chelchis (King's Fall)
		3042333087, // Harrowed Midha's Reckoning (King's Fall)
		2201079122, // Harrowed Silence of A'arn (King's Fall)
		3919765140, // Harrowed Defiance of Yasmin (King's Fall)
		1551744703, // Harrowed Qullim's Terminus (King's Fall)
		1397524041, // Harrowed Elulim's Frenzy (King's Fall)

		621603243,  // GENESIS CHAIN~ (Wrath of the Machine)
		2154053164, // FEVER AND REMEDY~ (Wrath of the Machine)
		3632330099, // STEEL MEDULLA~ (Wrath of the Machine)
		2542033072, // CHAOS DOGMA~ (Wrath of the Machine)
		3598793896, // ETHER NOVA~ (Wrath of the Machine)
		2125403517, // QUANTIPLASM~ (Wrath of the Machine)
		3569444312, // ZEAL VECTOR~ (Wrath of the Machine)
		2001493563, // EX MACHINA~ (Wrath of the Machine)
		242628276,  // IF MATERIA~ (Wrath of the Machine)
		1784034858, // SOUND AND FURY~ (Wrath of the Machine)
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
		454534839,  // Mark of a Kellslayer
		3620910776, // Kellbreaker's Mind
		3660153609, // Kellbreaker's Gloves
		4013007887, // Kellbreaker's Cloak
		3910621915, // Kellbreaker's Path
		1006775966, // "Dread's Enmity"

		3471865173, // Harrowed Darkhollow Mask (King's Fall)
		2302693612, // Harrowed Darkhollow Grasps (King's Fall)
		3907799186, // Harrowed Darkhollow Chiton (King's Fall)
		2549035182, // Harrowed Darkhollow Treads (King's Fall)
		2242715339, // Cloak of Seven Bones (King's Fall)
		1846107924, // Harrowed Mouth of Ur (King's Fall)
		521951205,  // Harrowed Grasp of Eir (King's Fall)
		372855005,  // Harrowed Chasm of Yuul (King's Fall)
		2028036495, // Harrowed Path of Xol (King's Fall)
		1658688592, // "Worm Gods' Boon" (King's Fall)
		1245063911, // Harrowed War Numen's Crown (King's Fall)
		217447094,  // Harrowed War Numen's Fist (King's Fall)
		3176903680, // Harrowed War Numen's Chest (King's Fall)
		1601524312, // Harrowed War Numen's Boots (King's Fall)
		130578781,  // Mark of the Old Challenge (King's Fall)

		661529958,  // Spliced Nanomania Mask (Wrath of the Machine)
		2903199471, // Spliced Nanomania Grasps (Wrath of the Machine)
		1823306243, // Spliced Nanomania Vest (Wrath of the Machine)
		691310313,  // Spliced Nanomania Boots (Wrath of the Machine)
		1047302226, // Spliced Nanomania Cloak (Wrath of the Machine)
		1299613399, // Spliced Red Miasma Hood (Wrath of the Machine)
		2503707046, // Spliced Red Miasma Gloves (Wrath of the Machine)
		2679967696, // Spliced Red Miasma Robes (Wrath of the Machine)
		1680970216, // Spliced Red Miasma Boots (Wrath of the Machine)
		155972845,  // Spliced Red Miasma Bond (Wrath of the Machine)
		3138343136, // Spliced Cosmoclast Helm (Wrath of the Machine)
		3976262649, // Spliced Cosmoclast Gauntlets (Wrath of the Machine)
		307172369,  // Spliced Cosmoclast Plate (Wrath of the Machine)
		2011530291, // Spliced Cosmoclast Greaves (Wrath of the Machine)
		3529294924, // Spliced Cosmoclast Mark (Wrath of the Machine)
	},
}

var otherItems = []int64{
	//Armor
	773948710,  // Snow Angel Vest
	2364921276, // Overtaker Helmet
	3909166968, // Speeder Gloves
	1519376145, // Apotheosis Veil
	2188828888, // Venom of Ikaheka
	1458254034, // Khepri's Sting
	16664395,   // Desolate Grips
	2360689228, // "Emerald Light"
	1054763958, // Skyburners Annex
	825392783,  // Wolfswood Mark
	1520434778, // ATS/8 ARACHNID
	3335833364, // Nightstalker's Cloak
	1184900752, // Speeder Suit
	941890987,  // Eternal Warrior
	498794675,  // Snow Angel Horns
	1944434610, // Speeder Boots
	3841951818, // Wolfswood Bond
	3907475891, // Iron Companion Helm
	2620256215, // Arc Flayer Mantle
	975923469,  // Bond of the Desolate
	2405148796, // Mark of the Nexus Undone
	2055339186, // Mantle of Gheleon
	1866413378, // Iron Companion Gauntlets
	13565936,   // Spektar's Mark
	13565938,   // Mark of Desolation
	2300914892, // Nightstalker Cloak
	1722532281, // Wolfswood Cloak
	3335833366, // Gunslinger's Cloak
	3335833367, // Bladedancer's Cloak
	3847828218, // Desolate Cover
	591060260,  // Empyrean Bellicose
	941890990,  // Helm of Inmost Light
	2778128366, // Astrocyte Verse
	501878012,  // Snow Angel Boots
	2747259661, // Iron Companion Mask
	1520434779, // Mask of the Third Man
	4223010274, // Circuit Defender
	1520434777, // Knucklehead Radar
	1520434781, // Celestial Nighthawk
	1054763959, // Graviton Forfeit
	1520434776, // Achlyophage Symbiote
	1458254032, // Young Ahamkara's Spine
	2241750289, // Lion's Vigil Grips
	2217280775, // Sealed Ahamkara Grasps
	1458254033, // Don't Touch Me
	2072689472, // Circuit Gauntlets
	2217280774, // Shinobu's Vow
	3399822332, // Iron Companion Sleeves
	2882684153, // Lucky Raspberry
	2555395778, // Circuit Chestplate
	3324900039, // Lion's Vigil Vest
	554244708,  // Iron Companion Vest
	105485105,  // ATS/8 Tarantella
	2882684152, // Crest of Alpha Lupi
	3809851981, // Desolate Vest
	1706217754, // Circuit Striders
	2402750214, // Iron Companion Boots
	1775312682, // Radiant Dance Machines
	308606233,  // Desolate Legs
	1775312683, // Bones of Eao
	1648442595, // Lion's Vigil Strides
	1394543945, // Fr0st-EE5
	2563486541, // Racer's Scarf
	1516397455, // Cloak of Snows
	2620256214, // Cloak of Taniks
	2300914895, // Cloak of the Rising
	2300914893, // Cloak of Oblivion
	2300914894, // Chaos Cloak
	1186397660, // Cloak of Desolation
	1519376146, // Light Beyond Nemesis
	870077908,  // Iron Companion Hood
	1519376148, // The Ram
	2778128367, // THE STAG
	1519376147, // Obsidian Mind
	3689370929, // Desolate Veil
	1519376144, // Skull of Dire Ahamkara
	1275480035, // Nothing Manacles
	1062853750, // The Impossible Machines
	1275480032, // Claws of Ahamkara
	3995873645, // Iron Companion Gloves
	1062853751, // Ophidian Aspect
	1648308416, // Desolate Gloves
	1275480033, // Sunbreakers
	233474724,  // Desolate Robes
	3574778423, // Heart of the Praxic Fire
	3574778422, // Voidfang Vestments
	3574778420, // Purifier Robes
	1212068371, // Iron Companion Vestments
	3574778421, // Starfire Protocol
	2898542650, // Alchemist's Raiment
	4242215215, // Iron Companion Legs
	3692454270, // Desolate Boots
	2275132880, // Transversive Steps
	2194008687, // Stormcaller's Bond
	2122538505, // "Circle of War"
	2194008684, // Sunsinger's Bond
	2122538507, // Stormcaller Bond
	2194008685, // Voidwalker's Bond
	2122538506, // "Light Beyond"
	2122538504, // "The Age to Come"
	591060261,  // The Taikonaut
	941890989,  // An Insurmountable Skullfort
	941890991,  // Helm of Saint-14
	2928398154, // Spektar Helmet
	941890988,  // The Glasshouse
	2928398152, // Desolate Helm
	1176452006, // Scarab's Vigil Helm
	155374077,  // Immolation Fists
	2975536019, // Spektar Gauntlets
	3055446327, // Ruin Wings
	3055446326, // No Backup Plans
	155374076,  // Thagomizers
	3055446324, // ACD/0 Feedback Fence
	2975536017, // Desolate Gauntlets
	2661471738, // Crest of Alpha Lupi
	2804968637, // Spektar Plate
	3921595523, // Twilight Garrison
	2661471739, // The Armamentarium
	82944038,   // Iron Companion Plate
	2804968639, // Desolate Plate
	1627572929, // Scarab's Vigil Plate
	2479526175, // Dunemarchers
	2710490085, // Spektar Greaves
	1998946585, // Scarab's Vigil Greaves
	4267828624, // Mk. 44 Stand Asides
	4267828625, // Peregrine Greaves
	3910559228, // Iron Companion Greaves
	2710490087, // Desolate Greaves
	2820418552, // Mark of the Circle
	2820418554, // Mark of the Sunforged
	687119878,  // Sunbreaker's Mark
	2820418553, // Mark of the Executor
	2820418555, // Mark of Oblivion
	687119876,  // Striker's Mark
	687119877,  // Defender's Mark

	//Weapon
	2968802931, // Stellar Vestige
	3452625744, // Haakon's Hatchet
	1623420384, // The Young Wolf's Howl
	3497087277, // Colovance's Duty
	1994742696, // Stolen Will
	4097026463, // No Time to Explain
	2748609458, // Fabian Strategy
	2008951974, // Immobius
	3227022823, // Hereafter
	1287343925, // Party Crasher +1
	2413549607, // The Lingering Song
	3688594189, // Touch of Malice
	3835813881, // No Land Beyond
	3677466371, // The Hothead
	2055601061, // SUROS Regime
	2297964387, // Parthian Shot
	2447423792, // Hawkmoon
	2055601060, // Hard Light
	3768542598, // The Laughing Heart
	4100639363, // Void Edge
	2897238917, // Finnala's Peril
	99462853,   // Universal Remote
	3688594188, // Boolean Gemini
	3675783241, // The Chaperone
	3835813880, // Patience and Time
	1742712674, // Stillpiercer
	1050258874, // The Unbent Tree
	987423912,  // The First Curse
	2443083323, // Ashraven's Flight
	100397241,  // Hopscotch Pilgrim
	4068577415, // Grasp of Malok
	1272989272, // Deidris's Retort
	1982014077, // Dreg's Promise
	2205574383, // The Binding Blaze
	958238921,  // The Branded Lord
	1264422556, // Weyloran's March
	1026578963, // The Distant Star
	1758882169, // Khvostov 7G-0X
	1173766590, // Doctrine of Passing
	1346849289, // MIDA Multi-Tool
	2999797736, // The Clever Dragon
	1132793057, // Eyasluna
	3742521821, // Outbreak Prime
	255654879,  // Zhalo Supercell
	552354419,  // Ace of Spades
	1475134443, // Hung Jury SR4
	3622856851, // Matador 64
	99462852,   // Invective
	1703777169, // 1000-Yard Stare
	3012398148, // Telesto
	1200540135, // Y-09 Longbow Synthesis
	4242230174, // Ice Breaker
	3564229425, // Gjallarhorn
	4100639364, // Dark-Drinker
	3012398149, // Sleeper Simulant
	2808364178, // Truth
	3536592559, // Tormod's Bellows
	3052681344, // Imago Loop
	803312564,  // Tlaloc
	2055601062, // Monte Carlo
	1177550374, // Bad Juju
	3904536202, // Nirwen's Mercy
	2447423793, // The Last Word
	3366907656, // Vision Stone
	4257542172, // Susanoo
	99462854,   // The 4th Horseman
	3938709034, // Trespasser
	1689897198, // The Proud Spire
	99462855,   // Lord of Wolves
	3227022822, // Black Spindle
	4100639365, // Raze-Lighter
	1170904292, // Iron Gjallarhorn
	2121113047, // Bretomart's Stand
	4100639361, // Arc Edge
	2931351952, // Dreadfang
	2604291456, // Nova Mortis
	57660786,   // Super Good Advice
	3172428255, // The Palindrome
	2790109143, // SUROS PDX-45
	3490124917, // Burning Eye
	3904467563, // Thorn
	1177550375, // Red Death
	3491886959, // SUROS DIS-43
	1419914235, // Saladin's Vigil
	3078564839, // Plan C
	1200540134, // LDR 5001
	3078564838, // Queenbreakers' Bow
	486279087,  // Conspiracy Theory-D
	3068424913, // Ironwreath-D
	2878293129, // The Silvered Dread
	57660787,   // Thunderlord
	3851373522, // Nemesis Star
	4100639362, // Bolt-Caster
	2604291457, // Abbadon
	2808364179, // Dragon's Breath
	330048677,  // The Titanium Orchid
	4100639360, // Sol Edge
}

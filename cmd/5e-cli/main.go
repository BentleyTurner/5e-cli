package main

import (
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
)

var ROLL_RANGE_CEILINGS = map[int]func() error{
	1:  lowgold,
	2:  magicItem,
	3:  mediumgold,
	4:  mediumgold,
	5:  amulet,
	6:  ring,
	7:  shrine,
	8:  relic,
	9:  func() error { log.Println("Dream Mirror"); return nil },
	100:  func() error { log.Println("Glyph"); return nil },
} 	

var COMMAND_MAP = map[string]func() error{
	"exit":          func() error { os.Exit(0); return nil },
	"q":             func() error { os.Exit(0); return nil },
	"colour":        colour,
	"wep":           weaponEnchant,
	"arm":           armourEnchant,
	"glyph":         glyph,
	"relic":         upgradeRelic,
	"skill":         skill,
	"dmg type":      dmgType,
	"creature type": creatureType,
	"ability":       ability,
	"loot":          loot,
	"harvest":       harvest,
	"condi":         condition,
	"mutate":        mutation,
	"encounter":     randomEncounter,
	"insight":       insight,
	"dmg polarity":  dmgPolarity,
	"party member":  partyMember,
	"npc":           npc,
	"xiloan":        xiloan,
	"positive":      positiveReward,
	"weapon class":  weaponClass,
	"phys type":     physType,
	"non-phys type": nonPhysType,
	"challenge":     challenge,
	"class":         class,
	"tarot":         tarot,
	"gem":           gem,
	"tag":           tag,
	"craft":         craft,
	"target craft":  targetCraft,
	"dmg upgrade":   dmgUpgrade,
	"activity":      activity,
	"amulet":        amulet,
	"relic new":     relic,
	"chaos":         chaos,
	"wondrous":      wondrous,
	"ring":          ring,
	"combat":        combat,
	"travel":        travel,
	"ja":            journeyActivity,
	"feat":          feat,
	"simple wep":    simpleWeapon,
	"martial wep":   martialWeapon,
	"posi enc":      posiEnc,
	"language":      language,
}

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	rand.Seed(time.Now().UnixNano())

	for {
		log.Println()
		baseP := promptui.Prompt{
			Label:    "Command",
			Validate: validateBase,
		}
		input, err := baseP.Run()
		if err != nil {
			log.Fatal(err)
			return
		}

		inputInt, err := strconv.Atoi(input)
		if err == nil {
			ceilings := make([]int, len(ROLL_RANGE_CEILINGS))
			i := 0
			for k := range ROLL_RANGE_CEILINGS {
				ceilings[i] = k
				i++
			}
			sort.Ints(ceilings)
			for _, c := range ceilings {
				if inputInt <= c {
					err = ROLL_RANGE_CEILINGS[c]()
					break
				}
			}
		} else {
			err = COMMAND_MAP[input]()
		}
		if err != nil {
			log.Printf("Error occurred during running of %s", input)
			log.Fatal(err)
		}
		log.Println()
	}
}

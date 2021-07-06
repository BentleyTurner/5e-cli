package main

import (
	"log"
	"math/rand"
	"regexp"
	"strings"
)

var SUBSTITUTION_MAP = map[string][]string{
	"$dmgType":         DAMAGE_TYPES,
	"$abilityScore":    ABILITY_SCORES,
	"$hitForm":         HIT_FORMS,
	"$physType":        PHYS_TYPES,
	"$willAbility":     WILL_ABILITIES,
	"$weaponHands":     WEAPON_HANDS,
	"$weaponClass":     WEAPON_CLASSES,
	"$skill":           SKILLS,
	"$nonPhysType":     NON_PHYS_TYPES,
	"$dmgPolarity":     DAMAGE_POLARITIES,
	"$partyMember":     PARTY_MEMBERS,
	"$lightType":       LIGHT_TYPES,
	"$condition":       CONDITIONS,
	"$aoeShape":        AOE_SHAPES,
	"$school":          SCHOOLS,
	"$creatureType":    CREATURE_TYPES,
	"$healthStatus":    HEALTH_STATUSES,
	"$weaponTrait":     WEAPON_TRAITS,
	"$sizeDifference":  SIZE_DIFFERENCES,
	"$armourWeight":    ARMOUR_WEIGHTS,
	"$enemyArmourForm": ENEMY_ARMOUR_FORMS,
	"$rangeType":       RANGE_TYPES,
}

func processMod(modString string) string {
	log.Printf("[DEBUG] modString: %s", modString)
	matcher := regexp.MustCompile(`\$\w+`)
	subs := matcher.FindAllString(modString, -1)
	ret := modString
	for _, s := range subs {
		sub := SUBSTITUTION_MAP[s][rand.Intn(len(SUBSTITUTION_MAP[s]))]
		ret = strings.Replace(ret, s, sub, 1)
	}
	return ret
}
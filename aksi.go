package main

import (
	"fmt"
)

const NMAX = 100

type character struct {
	name         string
	health       int
	attacks      [NMAX]int
	attackCount  int
	totalDamage  int
}

func main() {
	var link, ganon character
	link.name = "Link"
	ganon.name = "Ganon"

	fmt.Scan(&link.health, &ganon.health)

	for link.health > 0 && ganon.health > 0 {
		var name, action string
		var damage int

		_, err := fmt.Scan(&name, &action)
		if err != nil {
			break
		}

		var attacker, defender *character
		if name == "Link" {
			attacker = &link
			defender = &ganon
		} else {
			attacker = &ganon
			defender = &link
		}

		if action == "ATTACK" {
			fmt.Scan(&damage)
			defender.health -= damage
			attacker.attacks[attacker.attackCount] = damage
			attacker.attackCount++
			attacker.totalDamage += damage
		} else if action == "DEFENSE" {
			if defender.attackCount > 0 {
				lastDamage := defender.attacks[defender.attackCount-1]
				attacker.health += lastDamage
			}
		} else if action == "PARRY" {
			if defender.attackCount > 0 {
				lastDamage := defender.attacks[defender.attackCount-1]
				defender.health -= lastDamage
				attacker.totalDamage += lastDamage
			}
		}

		if link.health < 0 {
			link.health = 0
		}
		if ganon.health < 0 {
			ganon.health = 0
		}
	}

	if ganon.health <= 0 {
		fmt.Println("Link menang! putri Zelda berhasil diselamatkan.")
	} else {
		fmt.Println("Ganon menang! putri Zelda dalam bahaya.")
	}

	fmt.Printf("Link: health %d, attacks %d, damage dealt %d\n", link.health, link.attackCount, link.totalDamage)
	fmt.Printf("Ganon: health %d, attacks %d, damage dealt %d\n", ganon.health, ganon.attackCount, ganon.totalDamage)
}
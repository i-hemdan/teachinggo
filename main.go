package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

type Person struct {
	name   string
	health int
	attack int
}

type Item struct {
	Name   string `json:"name"`
	Damage int    `json:"damage"`
}

type Player struct {
	Person
	equipment *Item
}

type Goblin struct {
	Person
	callout string
	drop    *Item
}

func (p *Player) init() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("what is your name")
	scanner.Scan()
	p.name = scanner.Text()

	p.health = 20
	p.attack = 1
}
func (p *Player) Attack(goblin *Goblin) {
	tempAttack := p.attack
	if p.equipment != nil {
		tempAttack += p.equipment.Damage
	}
	goblin.health -= tempAttack
	fmt.Printf(
		"%v attacks %v for %v damage, %v has %v health left\n",
		p.name,
		goblin.name,
		tempAttack,
		goblin.name,
		goblin.health,
	)
}

func (g *Goblin) Attack(person *Player) {
	person.health -= g.attack
	fmt.Printf(
		"%v attacks %v for %v damage, %v has %v health left\n",
		g.name,
		person.name,
		g.attack,
		person.name,
		person.health,
	)
}

func makeItems() []Item {
	var items []Item
	jsonItems, err := ioutil.ReadFile("resources\\items.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonItems, &items)
	if err != nil {
		log.Fatal(err)
	}
	return items
}

func main() {

	items := makeItems()

	commons := items[:2]
	fmt.Println(commons)

	uncommons := items[2:4]
	fmt.Println(uncommons)

	rares := items[4:]
	fmt.Println(rares)

	var player Player
	player.init()
	player.equipment = &items[0]

	goblin := Goblin{
		Person:  Person{name: "billy", health: 5, attack: 2},
		callout: "fuck you m8",
		drop:    &rares[rand.Intn(len(rares))],
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Oh no a goblin named %v is attacking, he says \"%v\"\n", goblin.name, goblin.callout)
	for {
		if player.health > 0 {
			fmt.Println("what do you do?")
			scanner.Scan()
			input := scanner.Text()
			if input == "attack" {
				fmt.Println("you attack the goblin")
				player.Attack(&goblin)
			} else {
				fmt.Println("you cant do that")
			}
			if goblin.health > 0 {
				goblin.Attack(&player)
			} else {
				fmt.Printf("holy shit you killed %v! but at least you got a %v", goblin.name, goblin.drop.Name)
				player.equipment = goblin.drop
				break
			}
		} else {
			fmt.Println("you fuckin died!")
			break
		}
	}
}

package main

import (
	"errors"
)

type Item struct {
	Name       string
	ItemForUse int
	Contains   []int
}

func FindItemByName(itemName string) (error, int, *Item) {
	for index, itm := range Items {
		if itm.Name == itemName {
			return nil, index, itm
		}
	}
	return errors.New("Item not found"), -1, nil
}

func OpenItem(p *Actor, itemName string) {
	loc := LocationMap[p.CurrentLocation]
	for _, itm := range loc.Items {
		if Items[itm].Name == itemName {
			if Items[itm].ItemForUse != 0 && PlayerHasItem(p, Items[itm].ItemForUse) {
				loc.Items = append(loc.Items, Items[itm].Contains...)
				Items[itm].Contains = *new([]int)
			} else {
				Output("red", "Could not open the "+itemName)
				return
			}
		} else {
			Output("red", "Could not open the "+itemName)
		}
	}
}

func PlayerHasItem(pla *Actor, itm int) bool {
	for _, pitm := range pla.Items {
		if pitm == itm {
			return true
		}
	}
	return false
}

func (it *Item) RemoveItemFromRoom(loc *Location) {
	for index, itm := range loc.Items {
		if Items[itm].Name == it.Name {
			loc.Items = append(loc.Items[:index], loc.Items[index+1:]...)
		}
	}
}

//To be refactored on a location struct
func (it *Item) ItemInRoom(loc *Location) bool {
	for _, itm := range loc.Items {
		if Items[itm].Name == it.Name {
			return true
		}
	}
	return false
}

//To be refactored on a character struct
func (it *Item) ItemOnPlayer(pla *Actor) bool {
	for _, itm := range pla.Items {
		if Items[itm].Name == it.Name {
			return true
		}
	}
	return false
}

//To be refactored on a location struct
func describeItems(player Actor) {
	l := LocationMap[player.CurrentLocation]

	Output("You see:")
	for _, itm := range l.Items {
		Outputf("\t%s\n", Items[itm].Name)
	}
}

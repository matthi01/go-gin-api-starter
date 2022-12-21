package db

import "github.com/matthi01/go-gin-api-starter/util"

// TODO: hook this up to a db rather than just storing this in memory
type Item struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type List []Item

func New() *List {
	return &List{
		Item{Id: 12321, Name: "test", Description: "test"},
	}
}

func (l *List) GetAll() []Item {
	return *l
}

func (l *List) Get(id int) (Item, bool) {
	for _, item := range *l {
		if item.Id == id {
			return item, true
		}
	}
	return Item{}, false
}

func (l *List) Add(name string, description string) (Item, bool) {
	newItem := Item{
		Id:          util.GenerateId(),
		Name:        name,
		Description: description,
	}
	*l = append(*l, newItem)
	return newItem, true
}

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
	return &List{}
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

func (l *List) Update(id int, name string, description string) (Item, bool) {
	itemIndex := l.findIndex(id)
	if itemIndex < 0 {
		return Item{}, false
	}
	(*l)[itemIndex].Name = name
	(*l)[itemIndex].Description = description
	return (*l)[itemIndex], true
}

func (l *List) Delete(id int) (Item, bool) {
	itemIndex := l.findIndex(id)
	if itemIndex < 0 {
		return Item{}, false
	}
	removedItem := (*l)[itemIndex]
	*l = append((*l)[:itemIndex], (*l)[itemIndex+1:]...)
	return removedItem, true
}

func (l *List) findIndex(id int) int {
	for index, item := range *l {
		if item.Id == id {
			return index
		}
	}
	return -1
}

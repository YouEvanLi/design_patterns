package main

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

type IShoe struct {
}

type IShirt struct {
}

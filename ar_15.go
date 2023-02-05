package mypackage

import "fmt"

type Ar struct {
	Magazine_capacity int
	Material          string
	Scope             bool
	Scope_zoom        int
}

func NewAr() *Ar {
	return &Ar{Magazine_capacity: 30, Material: "titan", Scope: false, Scope_zoom: 1}
}
func AddMagazineCapacity(ar *Ar, newAmount int) {
	ar.Magazine_capacity = newAmount
}
func AddScope(ar *Ar, zoom int) {
	if !ar.Scope {
		ar.Scope = true
		ar.Scope_zoom = zoom
		return
	}
	fmt.Println("Error: You already have scope!")
}
func ChangeScope(ar *Ar, zoom int) {
	if ar.Scope {
		if zoom < 2 {
			fmt.Println("Scope removed!")
			ar.Scope = false
			ar.Scope_zoom = 1
			return
		}
		ar.Scope_zoom = zoom
		return
	}
	fmt.Println("Error: You dont have scope!")
}

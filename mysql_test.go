package main

import "testing"

var users = []User{
	{Name: "tom"},
	{Name: "bob"},
}

func Setup() {
	DeleteUserAll()
}

func Test_Create(t *testing.T) {
	Setup()
	err := users[0].Create()
	if err != nil {
		t.Error("Cannot Create error")
	}
	if users[0].Id == 0 {
		t.Error("Failed Create id error")
	}
}

func Test_Delete(t *testing.T) {
	Setup()
	err := users[0].Create()
	if err != nil {
		t.Errorf("Cannot Create error")
	}
	err = Delete(1)
	if err != nil {
		t.Errorf("Failed Delete id Error")
	}
}

func Test_Users(t *testing.T) {
	Setup()
	err := users[0].Create()
	if err != nil {
		t.Errorf("Cannot Create error")
	}
	err = users[1].Create()
	u, err := Users()
	if err != nil {

	}
	if len(u) != 2 {
		t.Errorf("Failed All users get")
	}
}

func Test_Update(t *testing.T) {
	Setup()
	err := users[0].Create()
	if err != nil {
		t.Errorf("Cannot Create error")
	}
	users[0].Name = "fish"
	err = users[0].Update()
	if err != nil {
		t.Errorf("Update Error")
	}
	u := User{Id: 1}
	err = u.UpdateCheck()
	if err != nil {
		t.Errorf("Failed UpdateCheck error")
	}
	if users[0].Name != u.Name {
		t.Errorf("Name Update check Error")
	}
}

package userdb

import (
	"testing"
)

func TestLogin(t *testing.T) {
	LoadUsers()
	if ValidateUser("vkroll2", []byte("mypassword")) {
		t.Logf("ok for valid data %s", "vkroll2")

	} else {
		t.Fatalf("not ok for valid data %s", "vkroll2")

	}
	if !ValidateUser("foo", []byte("bar")) {
		t.Logf("no login for bad data %s", "foo")
	} else {
		t.Fatalf("not ok for valid data %s", "foo")

	}

}

func TestUserExists(t *testing.T) {
	LoadUsers()
	if UserExists("foo") {
		t.Fatal("foo does not exist")
	}
	if UserExists("vkroll2") {
		t.Log("vkroll2 exists")
	} else {
		t.Fatal("vkroll2 should exist but doesn't")
	}
}

func TestCreateUser(t *testing.T) {
	LoadUsers()
	ok, _ := CreateUser("vkroll2", []byte("foo"))
	if !ok {
		t.Log("vkroll2 exists will not created")
	} else {
		t.Fatal("vkroll2 would be recreated even though it already exists")
	}
	ok, err := CreateUser("foo", []byte("fii"))

	if ok {
		t.Log("user correctly created")
	} else {
		t.Fatalf("User not created %s ", err.Error())
	}
}

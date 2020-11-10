package config

import "testing"

func TestUser(t *testing.T) {
	config := Database{}
	if config.User() != "root" {
		t.Fatal()
	}
}

func TestPassword(t *testing.T) {
	config := Database{}
	if config.Password() != "test" {
		t.Fatal()
	}
}

func TestHost(t *testing.T) {
	config := Database{}
	if config.Host() != "localhost" {
		t.Fatal()
	}
}

func TestPort(t *testing.T) {
	config := Database{}
	if config.Port() != "3306" {
		t.Fatal()
	}
}

func TestName(t *testing.T) {
	config := Database{}
	if config.Name() != "dropit" {
		t.Fatal()
	}
}

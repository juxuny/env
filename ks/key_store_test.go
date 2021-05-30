package ks

import (
	"testing"
)

func TestInitKeyName(t *testing.T) {
	var k struct {
		DbHost string
		DbPwd  string
		DbUser string
		DbPort string
	}
	InitKeyName(&k, true)
	if k.DbHost != "DB_HOST" ||
		k.DbPort != "DB_PORT" ||
		k.DbUser != "DB_USER" ||
		k.DbPwd != "DB_PWD" {
		t.Fatal("wrong data:", k)
	}
}

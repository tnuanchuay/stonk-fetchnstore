package mysql

import (
	"net/url"
	"testing"
)

func TestNewDb(t *testing.T) {
	u := url.URL{
		User: url.UserPassword("root", "password"),
		Host: "localhost",
	}

	db, err := New(u)
	if err != nil {
		t.Error(err)
	}

	err = db.Ping()
	if err != nil {
		t.Error(err)
	}
}

package structTag

import "testing"

type Person struct {
	Address *Address `json:"address" bson:"address"`
	Age     int
	Name    string `json:"name"`
}

type Address struct {
	Nation string `json:"nation" bson:"nation"`
	City   string `json:"city" bson:"city"`
	Street string `json:"street" bson:"-"`
}

func TestTag(t *testing.T) {
	cfg := Person{}
	s := New(cfg)
	t.Log("tag:", s.Get("bson", "Address", "Nation"))
	t.Log("tag:", s.Get("bson", "Address", "City"))
	t.Log("tag:", s.Get("bson", "Address", "Street"))
	t.Log("tag:", s.Get("bson", "Age"))
	t.Log("tag:", s.Get("bson", "Name"))
}

func TestTagPoint(t *testing.T) {
	cfg := Person{}
	s := New(cfg)
	t.Log("tag:", s.GetPoint("bson", "Address.Nation"))
	t.Log("tag:", s.GetPoint("bson", "Address.City"))
	t.Log("tag:", s.GetPoint("bson", "Address.Street"))
	t.Log("tag:", s.GetPoint("bson", "Age"))
	t.Log("tag:", s.GetPoint("bson", "Name"))
}

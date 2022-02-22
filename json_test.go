package doener

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {
	d := &Doener{}
	b, _ := json.Marshal(d)

	if string(b) != "{\"Filling\":\"\",\"Sauces\":null,\"Toppings\":null}" {
		t.Fatalf("Empty Doener is %s", string(b))
	}
}

func TestUnmarshal(t *testing.T) {
	fullDoener := WithAll()

	b, _ := json.Marshal(fullDoener)

	var nD Doener
	json.Unmarshal(b, &nD)

	if fullDoener.String() != nD.String() {
		t.Fatal("Unmarshalled Doener does not match")
	}
}

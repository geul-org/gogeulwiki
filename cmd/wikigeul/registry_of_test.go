//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestRegistryOf — ENUM names keyed by lower-case tag to their code; empty, a 2-WORD or unprefixed code and a tag twice (case ignored) fail
package main

import "testing"

func TestRegistryOf(t *testing.T) {
	got, err := registryOf(map[string]string{"0x0784": "en", "0x4130": "pt-BR"})
	if err != nil || len(got) != 2 || got["pt-br"] != 0x4130 || got["en"] != 0x0784 {
		t.Fatalf("%v %v", got, err)
	}
	for _, bad := range []map[string]string{
		nil,
		{"0x00010002": "en"},
		{"0784": "en"},
		{"0x0784": "en", "0x0785": "EN"},
	} {
		if _, err := registryOf(bad); err == nil {
			t.Errorf("%v: want error", bad)
		}
	}
}

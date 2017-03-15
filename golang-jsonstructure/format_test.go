package jsonstructure

import "testing"

func TestDateTimeApply(t *testing.T) {
	err := dateTimeApply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

func TestEmailApply(t *testing.T) {
	err := emailApply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

func TestHostnameApply(t *testing.T) {
	err := hostnameApply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

func TestIPV4Apply(t *testing.T) {
	err := ipv4Apply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

func TestIPV6Apply(t *testing.T) {
	err := ipv6Apply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

func TestURIApply(t *testing.T) {
	err := uriApply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

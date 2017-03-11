package jsonstructure

import "testing"

func TestAcceptString(t *testing.T) {
	if !acceptString("string") {
		t.Error("acceptString failure")
	}
	if acceptString("strings") {
		t.Error("acceptString failure")
	}
}

func TestDateTimeApply(t *testing.T) {
	err := dateTimeApply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = dateTimeApply("2000-01-01T00:00:00+00:00", "string")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = dateTimeApply("2000-01-01", "string")
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
}

func TestEmailApply(t *testing.T) {
	err := emailApply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = emailApply("user@example.com", "string")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = emailApply("user@example@com", "string")
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
}

func TestHostnameApply(t *testing.T) {
	err := hostnameApply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = hostnameApply("server.example.com", "string")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = hostnameApply("", "string")
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	name := "a"
	for i := 0; i < 7; i++ {
		name = name + name
	}
	err = hostnameApply(name, "string")
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	err = hostnameApply("!", "string")
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	name = "a"
	for i := 0; i < 10; i++ {
		name = name + name
	}
	err = hostnameApply(name, "string")
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
}

func TestIPV4Apply(t *testing.T) {
	err := ipv4Apply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = ipv4Apply("1.2.3.4", "string")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = ipv4Apply("!", "string")
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	err = ipv4Apply("2001:0db8:85a3:0000:0000:8a2e:0370:7334", "string")
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
}

func TestIPV6Apply(t *testing.T) {
	err := ipv6Apply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = ipv6Apply("2001:0db8:85a3:0000:0000:8a2e:0370:7334", "string")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = ipv6Apply("!", "string")
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
}

func TestURIApply(t *testing.T) {
	err := uriApply(1, "integer")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = uriApply("http://www.example.com", "string")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = ipv6Apply("!", "string")
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
}

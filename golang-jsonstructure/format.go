package jsonstructure

import (
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type Format struct {
	Accept func(typ string) bool
	Apply  func(val interface{}, typ string) error
}

var hostnameIllegalChars = regexp.MustCompile(`[^0-9a-zA-Z.\-]`)

var defaultFormats = map[string]*Format{
	"date-time": &Format{acceptString, dateTimeApply},
	"email":     &Format{acceptString, emailApply},
	"hostname":  &Format{acceptString, hostnameApply},
	"ipv4":      &Format{acceptString, ipv4Apply},
	"ipv6":      &Format{acceptString, ipv6Apply},
	"uri":       &Format{acceptString, uriApply},
}

func acceptString(typ string) bool {
	return typ == "string"
}

func dateTimeApply(val interface{}, _ string) error {
	str, ok := val.(string)
	if !ok {
		return nil
	}
	_, err := time.Parse(time.RFC3339, str)
	return err
}

func emailApply(val interface{}, _ string) error {
	str, ok := val.(string)
	if !ok {
		return nil
	}
	_, err := mail.ParseAddress(str)
	return err
}

func hostnameApply(val interface{}, _ string) error {
	str, ok := val.(string)
	if !ok {
		return nil
	}
	if len(str) > 253 {
		return errors.New("hostname is greater than 253 characters")
	}
	labels := strings.Split(str, ".")
	for i, label := range labels {
		if len(label) < 1 {
			return fmt.Errorf("label %d of hostname is 0 characters", i+1)
		}
		if len(label) > 63 {
			return fmt.Errorf("label %d of hostname > 63 characters", i+1)
		}
	}
	char := hostnameIllegalChars.FindString(str)
	if char != "" {
		return fmt.Errorf(`hostname has illegal character "%s"`, char)
	}
	return nil
}

func ipv4Apply(val interface{}, _ string) error {
	str, ok := val.(string)
	if !ok {
		return nil
	}
	ip := net.ParseIP(str)
	if ip == nil {
		return errors.New("invalid IPv4 address")
	}
	if ip.To4() == nil {
		return errors.New("invalid IPv4 address")
	}
	return nil
}

func ipv6Apply(val interface{}, _ string) error {
	str, ok := val.(string)
	if !ok {
		return nil
	}
	ip := net.ParseIP(str)
	if ip == nil {
		return errors.New("invalid IPv6 address")
	}
	if ip.To4() != nil {
		return errors.New("invalid IPv4 address")
	}
	if ip.To16() == nil {
		return errors.New("invalid IPv6 address")
	}
	return nil
}

func uriApply(val interface{}, _ string) error {
	str, ok := val.(string)
	if !ok {
		return nil
	}
	_, err := url.Parse(str)
	return err
}

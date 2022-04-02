package conv

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

const (
	BackgroundOptionName = "BACKGROUND"
	ForegroundOptionName = "FOREGROUND"
	DefaultAttributeName = "TEXT"
)

var escapedNameExpr = regexp.MustCompile("[^a-zA-Z0-9]+")

type Scheme struct {
	XMLName    xml.Name `xml:"scheme"`
	Name       string   `xml:"name,attr"`
	Colors     []Option `xml:"colors>option"`
	Attributes []Option `xml:"attributes>option"`
}

func (s Scheme) Attr(name string) (Option, error) {

	for _, o := range s.Attributes {
		if o.Name == name {
			return o, nil
		}
	}

	return Option{}, fmt.Errorf("attribute %q was not found", name)
}

func (s Scheme) Color(name string) (string, error) {

	for _, o := range s.Colors {
		if o.Name == name {
			return fmt.Sprintf("#%s", o.Value), nil
		}
	}

	return "", fmt.Errorf("no color named %q found", name)
}

func (s Scheme) BG() (string, error) {
	o, err := s.Attr(DefaultAttributeName)
	if err != nil {
		return "", err
	}

	return o.InnerValue.BG()
}

func (s Scheme) FG() (string, error) {
	o, err := s.Attr(DefaultAttributeName)
	if err != nil {
		return "", err
	}

	return o.InnerValue.FG()
}

func (s Scheme) EscapedName() string {
	str := escapedNameExpr.ReplaceAllString(s.Name, "")
	return strings.ToLower(str)
}

type Option struct {
	Name       string `xml:"name,attr"`
	Value      string `xml:"value,attr"`
	InnerValue Value  `xml:"value"`
}

type Value struct {
	Options []Option `xml:"option"`
}

func (v Value) ValueOf(key string) (string, error) {
	for _, o := range v.Options {
		if o.Name == key {
			return o.Value, nil
		}
	}

	return "", fmt.Errorf("no option named %q found", key)
}

func (v Value) BG() (string, error) {

	bg, err := v.ValueOf(BackgroundOptionName)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("#%s", bg), nil
}

func (v Value) FG() (string, error) {

	fg, err := v.ValueOf(ForegroundOptionName)
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("#%s", fg), nil
}

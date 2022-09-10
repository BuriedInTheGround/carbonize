// Package carbonize implements utilities to deal with [Carbon] configurations
// and ease interactions with its website.
//
// [Carbon]: https://carbon.now.sh
package carbonize

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
)

// A Config is a configuration for Carbon.
type Config struct {
	BackgroundColor      string     `json:"backgroundColor" url:"bg"`
	Theme                string     `json:"theme" url:"t"`
	WindowTheme          string     `json:"windowTheme" url:"wt"`
	Language             string     `json:"language" url:"l"`
	Width                int        `json:"width" url:"width"`
	DropShadow           bool       `json:"dropShadow" url:"ds"`
	DropShadowOffsetY    string     `json:"dropShadowOffsetY" url:"dsyoff"`
	DropShadowBlurRadius string     `json:"dropShadowBlurRadius" url:"dsblur"`
	WindowControls       bool       `json:"windowControls" url:"wc"`
	WidthAdjustment      bool       `json:"widthAdjustment" url:"wa"`
	PaddingVertical      string     `json:"paddingVertical" url:"pv"`
	PaddingHorizontal    string     `json:"paddingHorizontal" url:"ph"`
	LineNumbers          bool       `json:"lineNumbers" url:"ln"`
	FirstLineNumber      int        `json:"firstLineNumber" url:"fl"`
	FontFamily           string     `json:"fontFamily" url:"fm"`
	FontSize             string     `json:"fontSize" url:"fs"`
	LineHeight           string     `json:"lineHeight" url:"lh"`
	SquaredImage         bool       `json:"squaredImage" url:"si"`
	ExportSize           string     `json:"exportSize" url:"es"`
	Watermark            bool       `json:"watermark" url:"wm"`
	Highlights           Highlights `json:"highlights" url:"highlights"`
}

// A Highlights is encapsulated inside a Config. It contains information about
// the syntax highlighting that Carbon should use for each different token.
type Highlights struct {
	Background string `json:"background" url:"background"`
	Text       string `json:"text" url:"text"`
	Attribute  string `json:"attribute" url:"attribute"`
	Keyword    string `json:"keyword" url:"keyword"`
	Variable   string `json:"variable" url:"variable"`
	Definition string `json:"definition" url:"definition"`
	Property   string `json:"property" url:"property"`
	String     string `json:"string" url:"string"`
	Meta       string `json:"meta" url:"meta"`
	Comment    string `json:"comment" url:"comment"`
	Number     string `json:"number" url:"number"`
	Operator   string `json:"operator" url:"operator"`
}

// QueryString builds a URL-encoded query string corresponding to the Config
// held by c.
func (c *Config) QueryString() (string, error) {
	query := url.Values{}

	fields := reflect.ValueOf(c).Elem()
	for i := 0; i < fields.NumField(); i++ {
		urlTag := fields.Type().Field(i).Tag.Get("url")

		if fields.Field(i).Type() == reflect.TypeOf(Highlights{}) {
			highlights, err := json.Marshal(fields.Field(i).Addr().Elem().Interface().(Highlights))
			if err != nil {
				return "", fmt.Errorf("cannot marshal highlights: %v", err)
			}
			query.Add(urlTag, string(highlights))
			continue
		}

		query.Add(urlTag, fmt.Sprintf("%v", fields.Field(i).Addr().Elem()))
	}

	return query.Encode(), nil
}

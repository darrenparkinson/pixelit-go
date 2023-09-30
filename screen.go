package pixelit

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Screen struct {
	SleepMode       *bool            `json:"sleepMode,omitempty"`
	Brightness      *bool            `json:"brightness,omitempty"`
	Text            *Text            `json:"text,omitempty"`
	SwitchAnimation *SwitchAnimation `json:"switchAnimation,omitempty"`
	Clock           *Clock           `json:"clock,omitempty"`
	Bitmap          *Bitmap          `json:"bitmap,omitempty"`
	Bitmaps         []Bitmap         `json:"bitmaps,omitempty"` // When displaying multiple bitmaps, animated bitmaps, scrolling or text are not supported!
	BitmapAnimation *BitmapAnimation `json:"bitmapAnimation,omitempty"`
}

type Text struct {
	TextString      *string   `json:"textString,omitempty"`      // Required. Displayed Text.
	BigFont         *bool     `json:"bigFont,omitempty"`         // Required. Big Font.
	CenterText      *bool     `json:"centerText,omitempty"`      // Required.
	ScrollText      *string   `json:"scrollText,omitempty"`      // One of true, false or auto.
	ScrollTextDelay *int      `json:"scrollTextDelay,omitempty"` // Required if ScrollText is set. 1-9999.
	Position        *Position `json:"position,omitempty"`        // Required.
	Color           *Color    `json:"color,omitempty"`           // Required unless hex color specified.
	HexColor        *string   `json:"hexColor,omitempty"`        // Alternative to Color.
}

type SwitchAnimation struct {
	Aktiv     *bool   `json:"aktiv,omitempty"`
	Animation *string `json:"animation,omitempty"`
	Data      *[]int  `json:"data,omitempty"`
	Width     *int    `json:"width,omitempty"`
}

type Clock struct {
	Show         *bool  `json:"show,omitempty"`         // Required
	SwitchAktiv  *bool  `json:"switchAktiv,omitempty"`  // Switch clock / date
	WithSeconds  *bool  `json:"withSeconds,omitempty"`  // Show seconds
	SwitchSec    *int   `json:"switchSec,omitempty"`    // Switch clock / date in seconds. Required when SwitchAktiv is true
	DrawWeekDays *bool  `json:"drawWeekDays,omitempty"` // Draw weekday blocks at the bottom
	Color        *Color `json:"color,omitempty"`        // Color of clock
	HexColor     string `json:"hexColor,omitempty"`     // Alternative to Color, e.g. #FFFFFF
}

type Bitmap struct {
	Data     []int     `json:"data,omitempty"`     // Required
	Position *Position `json:"position,omitempty"` // Required
	Size     *Size     `json:"size,omitempty"`     // Required
}

type BitmapAnimation struct {
	Data           [][]int `json:"data"`           //Required. Only 8x8 BMPs are supported here!
	AnimationDelay int     `json:"animationDelay"` // Required. Higher is a slower animation
	Rubberbanding  bool    `json:"rubberbanding"`  // Should the animation run back and forth
	LimitLoops     int     `json:"limitLoops"`     // If the repetition is to be limited
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Color struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (c *Client) SendScreen(screen *Screen) error {
	url := fmt.Sprintf("%s/screen", c.BaseURL)
	payload, err := json.Marshal(screen)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		return err
	}
	var r interface{}
	err = c.makeRequest(context.Background(), req, r)
	if err != nil {
		return err
	}
	return nil
}

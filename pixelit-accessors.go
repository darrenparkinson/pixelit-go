// Copyright 2021 The pixelit-go AUTHORS. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
// Code generated by gen-accessors; DO NOT EDIT.
package pixelit

// GetPosition returns the Position field.
func (b *Bitmap) GetPosition() *Position {
	if b == nil {
		return nil
	}
	return b.Position
}

// GetSize returns the Size field.
func (b *Bitmap) GetSize() *Size {
	if b == nil {
		return nil
	}
	return b.Size
}

// GetColor returns the Color field.
func (c *Clock) GetColor() *Color {
	if c == nil {
		return nil
	}
	return c.Color
}

// GetDrawWeekDays returns the DrawWeekDays field if it's non-nil, zero value otherwise.
func (c *Clock) GetDrawWeekDays() bool {
	if c == nil || c.DrawWeekDays == nil {
		return false
	}
	return *c.DrawWeekDays
}

// GetShow returns the Show field if it's non-nil, zero value otherwise.
func (c *Clock) GetShow() bool {
	if c == nil || c.Show == nil {
		return false
	}
	return *c.Show
}

// GetSwitchAktiv returns the SwitchAktiv field if it's non-nil, zero value otherwise.
func (c *Clock) GetSwitchAktiv() bool {
	if c == nil || c.SwitchAktiv == nil {
		return false
	}
	return *c.SwitchAktiv
}

// GetSwitchSec returns the SwitchSec field if it's non-nil, zero value otherwise.
func (c *Clock) GetSwitchSec() int {
	if c == nil || c.SwitchSec == nil {
		return 0
	}
	return *c.SwitchSec
}

// GetWithSeconds returns the WithSeconds field if it's non-nil, zero value otherwise.
func (c *Clock) GetWithSeconds() bool {
	if c == nil || c.WithSeconds == nil {
		return false
	}
	return *c.WithSeconds
}

// GetBitmap returns the Bitmap field.
func (s *Screen) GetBitmap() *Bitmap {
	if s == nil {
		return nil
	}
	return s.Bitmap
}

// GetBitmapAnimation returns the BitmapAnimation field.
func (s *Screen) GetBitmapAnimation() *BitmapAnimation {
	if s == nil {
		return nil
	}
	return s.BitmapAnimation
}

// GetBrightness returns the Brightness field if it's non-nil, zero value otherwise.
func (s *Screen) GetBrightness() bool {
	if s == nil || s.Brightness == nil {
		return false
	}
	return *s.Brightness
}

// GetClock returns the Clock field.
func (s *Screen) GetClock() *Clock {
	if s == nil {
		return nil
	}
	return s.Clock
}

// GetSleepMode returns the SleepMode field if it's non-nil, zero value otherwise.
func (s *Screen) GetSleepMode() bool {
	if s == nil || s.SleepMode == nil {
		return false
	}
	return *s.SleepMode
}

// GetSwitchAnimation returns the SwitchAnimation field.
func (s *Screen) GetSwitchAnimation() *SwitchAnimation {
	if s == nil {
		return nil
	}
	return s.SwitchAnimation
}

// GetText returns the Text field.
func (s *Screen) GetText() *Text {
	if s == nil {
		return nil
	}
	return s.Text
}

// GetAktiv returns the Aktiv field if it's non-nil, zero value otherwise.
func (s *SwitchAnimation) GetAktiv() bool {
	if s == nil || s.Aktiv == nil {
		return false
	}
	return *s.Aktiv
}

// GetAnimation returns the Animation field if it's non-nil, zero value otherwise.
func (s *SwitchAnimation) GetAnimation() string {
	if s == nil || s.Animation == nil {
		return ""
	}
	return *s.Animation
}

// GetData returns the Data field if it's non-nil, zero value otherwise.
func (s *SwitchAnimation) GetData() []int {
	if s == nil || s.Data == nil {
		return nil
	}
	return *s.Data
}

// GetWidth returns the Width field if it's non-nil, zero value otherwise.
func (s *SwitchAnimation) GetWidth() int {
	if s == nil || s.Width == nil {
		return 0
	}
	return *s.Width
}

// GetBigFont returns the BigFont field if it's non-nil, zero value otherwise.
func (t *Text) GetBigFont() bool {
	if t == nil || t.BigFont == nil {
		return false
	}
	return *t.BigFont
}

// GetCenterText returns the CenterText field if it's non-nil, zero value otherwise.
func (t *Text) GetCenterText() bool {
	if t == nil || t.CenterText == nil {
		return false
	}
	return *t.CenterText
}

// GetColor returns the Color field.
func (t *Text) GetColor() *Color {
	if t == nil {
		return nil
	}
	return t.Color
}

// GetHexColor returns the HexColor field if it's non-nil, zero value otherwise.
func (t *Text) GetHexColor() string {
	if t == nil || t.HexColor == nil {
		return ""
	}
	return *t.HexColor
}

// GetPosition returns the Position field.
func (t *Text) GetPosition() *Position {
	if t == nil {
		return nil
	}
	return t.Position
}

// GetScrollText returns the ScrollText field if it's non-nil, zero value otherwise.
func (t *Text) GetScrollText() string {
	if t == nil || t.ScrollText == nil {
		return ""
	}
	return *t.ScrollText
}

// GetScrollTextDelay returns the ScrollTextDelay field if it's non-nil, zero value otherwise.
func (t *Text) GetScrollTextDelay() int {
	if t == nil || t.ScrollTextDelay == nil {
		return 0
	}
	return *t.ScrollTextDelay
}

// GetTextString returns the TextString field if it's non-nil, zero value otherwise.
func (t *Text) GetTextString() string {
	if t == nil || t.TextString == nil {
		return ""
	}
	return *t.TextString
}

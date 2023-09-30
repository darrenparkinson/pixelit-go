// Copyright 2021 The pixelit-go AUTHORS. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
// Code generated by gen-accessors; DO NOT EDIT.
package pixelit

import "testing"

func TestBitmap_GetPosition(tt *testing.T) {
	b := &Bitmap{}
	b.GetPosition()
	b = nil
	b.GetPosition()
}

func TestBitmap_GetSize(tt *testing.T) {
	b := &Bitmap{}
	b.GetSize()
	b = nil
	b.GetSize()
}

func TestClock_GetColor(tt *testing.T) {
	c := &Clock{}
	c.GetColor()
	c = nil
	c.GetColor()
}

func TestClock_GetDrawWeekDays(tt *testing.T) {
	var zeroValue bool
	c := &Clock{DrawWeekDays: &zeroValue}
	c.GetDrawWeekDays()
	c = &Clock{}
	c.GetDrawWeekDays()
	c = nil
	c.GetDrawWeekDays()
}

func TestClock_GetShow(tt *testing.T) {
	var zeroValue bool
	c := &Clock{Show: &zeroValue}
	c.GetShow()
	c = &Clock{}
	c.GetShow()
	c = nil
	c.GetShow()
}

func TestClock_GetSwitchAktiv(tt *testing.T) {
	var zeroValue bool
	c := &Clock{SwitchAktiv: &zeroValue}
	c.GetSwitchAktiv()
	c = &Clock{}
	c.GetSwitchAktiv()
	c = nil
	c.GetSwitchAktiv()
}

func TestClock_GetSwitchSec(tt *testing.T) {
	var zeroValue int
	c := &Clock{SwitchSec: &zeroValue}
	c.GetSwitchSec()
	c = &Clock{}
	c.GetSwitchSec()
	c = nil
	c.GetSwitchSec()
}

func TestClock_GetWithSeconds(tt *testing.T) {
	var zeroValue bool
	c := &Clock{WithSeconds: &zeroValue}
	c.GetWithSeconds()
	c = &Clock{}
	c.GetWithSeconds()
	c = nil
	c.GetWithSeconds()
}

func TestScreen_GetBitmap(tt *testing.T) {
	s := &Screen{}
	s.GetBitmap()
	s = nil
	s.GetBitmap()
}

func TestScreen_GetBitmapAnimation(tt *testing.T) {
	s := &Screen{}
	s.GetBitmapAnimation()
	s = nil
	s.GetBitmapAnimation()
}

func TestScreen_GetBrightness(tt *testing.T) {
	var zeroValue bool
	s := &Screen{Brightness: &zeroValue}
	s.GetBrightness()
	s = &Screen{}
	s.GetBrightness()
	s = nil
	s.GetBrightness()
}

func TestScreen_GetClock(tt *testing.T) {
	s := &Screen{}
	s.GetClock()
	s = nil
	s.GetClock()
}

func TestScreen_GetSleepMode(tt *testing.T) {
	var zeroValue bool
	s := &Screen{SleepMode: &zeroValue}
	s.GetSleepMode()
	s = &Screen{}
	s.GetSleepMode()
	s = nil
	s.GetSleepMode()
}

func TestScreen_GetSwitchAnimation(tt *testing.T) {
	s := &Screen{}
	s.GetSwitchAnimation()
	s = nil
	s.GetSwitchAnimation()
}

func TestScreen_GetText(tt *testing.T) {
	s := &Screen{}
	s.GetText()
	s = nil
	s.GetText()
}

func TestSwitchAnimation_GetAktiv(tt *testing.T) {
	var zeroValue bool
	s := &SwitchAnimation{Aktiv: &zeroValue}
	s.GetAktiv()
	s = &SwitchAnimation{}
	s.GetAktiv()
	s = nil
	s.GetAktiv()
}

func TestSwitchAnimation_GetAnimation(tt *testing.T) {
	var zeroValue string
	s := &SwitchAnimation{Animation: &zeroValue}
	s.GetAnimation()
	s = &SwitchAnimation{}
	s.GetAnimation()
	s = nil
	s.GetAnimation()
}

func TestSwitchAnimation_GetData(tt *testing.T) {
	var zeroValue []int
	s := &SwitchAnimation{Data: &zeroValue}
	s.GetData()
	s = &SwitchAnimation{}
	s.GetData()
	s = nil
	s.GetData()
}

func TestSwitchAnimation_GetWidth(tt *testing.T) {
	var zeroValue int
	s := &SwitchAnimation{Width: &zeroValue}
	s.GetWidth()
	s = &SwitchAnimation{}
	s.GetWidth()
	s = nil
	s.GetWidth()
}

func TestText_GetBigFont(tt *testing.T) {
	var zeroValue bool
	t := &Text{BigFont: &zeroValue}
	t.GetBigFont()
	t = &Text{}
	t.GetBigFont()
	t = nil
	t.GetBigFont()
}

func TestText_GetCenterText(tt *testing.T) {
	var zeroValue bool
	t := &Text{CenterText: &zeroValue}
	t.GetCenterText()
	t = &Text{}
	t.GetCenterText()
	t = nil
	t.GetCenterText()
}

func TestText_GetColor(tt *testing.T) {
	t := &Text{}
	t.GetColor()
	t = nil
	t.GetColor()
}

func TestText_GetHexColor(tt *testing.T) {
	var zeroValue string
	t := &Text{HexColor: &zeroValue}
	t.GetHexColor()
	t = &Text{}
	t.GetHexColor()
	t = nil
	t.GetHexColor()
}

func TestText_GetPosition(tt *testing.T) {
	t := &Text{}
	t.GetPosition()
	t = nil
	t.GetPosition()
}

func TestText_GetScrollText(tt *testing.T) {
	var zeroValue string
	t := &Text{ScrollText: &zeroValue}
	t.GetScrollText()
	t = &Text{}
	t.GetScrollText()
	t = nil
	t.GetScrollText()
}

func TestText_GetScrollTextDelay(tt *testing.T) {
	var zeroValue int
	t := &Text{ScrollTextDelay: &zeroValue}
	t.GetScrollTextDelay()
	t = &Text{}
	t.GetScrollTextDelay()
	t = nil
	t.GetScrollTextDelay()
}

func TestText_GetTextString(tt *testing.T) {
	var zeroValue string
	t := &Text{TextString: &zeroValue}
	t.GetTextString()
	t = &Text{}
	t.GetTextString()
	t = nil
	t.GetTextString()
}
package pixelit

func (c *Client) SendText(text string) error {
	t := Text{
		TextString:      String(text),
		BigFont:         Bool(false),
		CenterText:      Bool(false),
		ScrollText:      String("auto"),
		ScrollTextDelay: Int(120),
		Position: &Position{
			X: 8,
			Y: 1,
		},
		Color: &Color{
			R: 255,
			G: 255,
			B: 255,
		},
	}
	s := Screen{
		Text: &t,
	}
	return c.SendScreen(&s)

}

func (c *Client) SendClock() error {
	s := Screen{
		SwitchAnimation: &SwitchAnimation{
			Aktiv:     Bool(true),
			Animation: String("coloredBarWipe"),
		},
		Clock: &Clock{
			Show:         Bool(true),
			SwitchAktiv:  Bool(true),
			WithSeconds:  Bool(false),
			SwitchSec:    Int(7),
			DrawWeekDays: Bool(true),
			Color: &Color{
				R: 255,
				G: 255,
				B: 255,
			},
		},
	}
	return c.SendScreen(&s)
}

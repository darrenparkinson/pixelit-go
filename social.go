package pixelit

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func (c *Client) SendInstagramFollowers(username string) error {
	followers, err := c.instagramFollowers(username)
	if err != nil {
		return err
	}
	s := Screen{
		SwitchAnimation: &SwitchAnimation{
			Aktiv:     Bool(true),
			Animation: String("coloredBarWipe"),
		},
		Bitmap: &Bitmap{
			Data: []int{0, 14783, 25023, 37375, 37375, 49598, 49598, 0, 35294, 35294, 65535, 65535, 65535, 65535, 63999, 63928, 53791, 65535, 63928, 63928, 63928, 65535, 65535, 63928, 61983, 65535, 63928, 65535, 65535, 63928, 65535, 63928, 59885, 65535, 59885, 65535, 65535, 63928, 65535, 63928, 64518, 65535, 64518, 64518, 64518, 59885, 65535, 63928, 64901, 65260, 65535, 65535, 65535, 65535, 59885, 63928, 0, 65461, 65461, 65166, 64967, 62502, 64166, 0},
			Position: &Position{
				X: 0,
				Y: 0,
			},
			Size: &Size{
				Width:  8,
				Height: 8,
			},
		},
		Text: &Text{
			TextString:      String(followers),
			BigFont:         Bool(false),
			CenterText:      Bool(true),
			ScrollText:      String("auto"),
			ScrollTextDelay: Int(120),
			Position: &Position{
				X: 0,
				Y: 0,
			},
			Color: &Color{
				R: 255,
				G: 255,
				B: 255,
			},
		},
	}
	c.SendScreen(&s)
	return nil
}

// func (c *Client) SendTwitterFollowers(username string) error {
// 	followers, err := c.twitterFollowers(username)
// 	if err != nil {
// 		return err
// 	}
// 	s := Screen{
// 		SwitchAnimation: &SwitchAnimation{
// 			Aktiv:     Bool(true),
// 			Animation: String("coloredBarWipe"),
// 		},
// 		Bitmap: &Bitmap{
// 			Data: []int{1119, 1119, 1119, 1119, 1119, 1119, 1119, 1119, 1119, 65535, 65535, 1119, 1119, 1119, 1119, 1119, 1119, 65535, 65535, 65535, 65535, 65535, 65535, 1119, 1119, 65535, 65535, 65535, 65535, 65535, 65535, 1119, 1119, 65535, 65535, 1119, 1119, 1119, 1119, 1119, 1119, 65535, 65535, 65535, 65535, 65535, 65535, 1119, 1119, 1119, 65535, 65535, 65535, 65535, 65535, 1119, 1119, 1119, 1119, 1119, 1119, 1119, 1119, 1119},
// 			Position: &Position{
// 				X: 0,
// 				Y: 0,
// 			},
// 			Size: &Size{
// 				Width:  8,
// 				Height: 8,
// 			},
// 		},
// 		Text: &Text{
// 			TextString:      String(followers),
// 			BigFont:         Bool(false),
// 			CenterText:      Bool(true),
// 			ScrollText:      String("auto"),
// 			ScrollTextDelay: Int(120),
// 			Position: &Position{
// 				X: 0,
// 				Y: 0,
// 			},
// 			Color: &Color{
// 				R: 255,
// 				G: 255,
// 				B: 255,
// 			},
// 		},
// 	}
// 	c.SendScreen(&s)
// 	return nil
// }

func (c *Client) SendTiktokFollowers(username string) error {
	followers, err := c.tiktokFollowers(username)
	if err != nil {
		return err
	}
	s := Screen{
		SwitchAnimation: &SwitchAnimation{
			Aktiv:     Bool(true),
			Animation: String("coloredBarWipe"),
		},
		Bitmap: &Bitmap{
			Data: []int{0, 0, 0, 0, 2047, 63488, 0, 0, 0, 0, 0, 0, 2047, 65535, 63488, 0, 0, 0, 0, 0, 2047, 65535, 65535, 63488, 0, 2047, 2047, 0, 2047, 63488, 0, 0, 2047, 65535, 65535, 0, 2047, 63488, 0, 0, 2047, 65535, 0, 0, 2047, 63488, 0, 0, 2047, 65535, 65535, 65535, 65535, 63488, 0, 0, 0, 2047, 65535, 65535, 63488, 0, 0, 0},
			Position: &Position{
				X: 0,
				Y: 0,
			},
			Size: &Size{
				Width:  8,
				Height: 8,
			},
		},
		Text: &Text{
			TextString:      String(followers),
			BigFont:         Bool(false),
			CenterText:      Bool(true),
			ScrollText:      String("auto"),
			ScrollTextDelay: Int(120),
			Position: &Position{
				X: 0,
				Y: 0,
			},
			Color: &Color{
				R: 255,
				G: 255,
				B: 255,
			},
		},
	}
	c.SendScreen(&s)
	return nil
}

func (c *Client) instagramFollowers(username string) (string, error) {
	if username == "" {
		return "", errors.New("no username provided")
	}
	url := fmt.Sprintf("https://instagram.com/%s", username)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "5.0 (iPhone; CPU iPhone OS 14_8 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Mobile/15E148 Safari/604.1")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, _ := io.ReadAll(resp.Body)
	data := getMetaContent(string(bytes))
	var followers string
	var found bool
	var parts []string
	for _, content := range data {
		parts = strings.SplitN(content, ", ", 3)
		if len(parts) == 3 && strings.HasSuffix(parts[0], "Followers") && strings.HasSuffix(parts[1], "Following") && strings.Contains(parts[2], "Posts") {
			found = true
			followers = strings.TrimSuffix(parts[0], " Followers")
			break
		}
	}
	if found {
		return followers, nil
	}
	return "", fmt.Errorf("no followers found for specified username: %s", username)
}

// func (c *Client) twitterFollowers(username string) (string, error) {
// 	url := fmt.Sprintf("https://www.twitter.com/%s", username)
// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Set("User-Agent", "5.0 (iPhone; CPU iPhone OS 14_8 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Mobile/15E148 Safari/604.1")
// 	resp, err := c.HTTPClient.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// 	bytes, _ := io.ReadAll(resp.Body)
// 	fmt.Println(string(bytes))
// 	data := getATags(string(bytes))
// 	if len(data) < 1 {
// 		return "", errors.New("no follower count found")
// 	}
// 	return data[0], nil
// }

func (c *Client) tiktokFollowers(username string) (string, error) {
	// strong with title "Followers" and data-e2e of "followers-count"
	// note this works for now due to the user agent. We may have more to do to identify captcha responses
	// example here: https://github.com/davidteather/TikTok-Api/blob/master/TikTokApi/tiktok.py
	url := fmt.Sprintf("https://www.tiktok.com/@%s", username)
	req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Set("User-Agent", "5.0 (iPhone; CPU iPhone OS 14_8 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Mobile/15E148 Safari/604.1")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, _ := io.ReadAll(resp.Body)
	data := getStrongTags(string(bytes))
	if len(data) < 1 {
		return "", errors.New("no follower count found")
	}
	return data[0], nil
}

func getStrongTags(text string) (data []string) {
	tkn := html.NewTokenizer(strings.NewReader(text))
	var vals []string
	var isStrong bool
	var isFollowers bool
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return vals // reached the end of the data
		case tt == html.StartTagToken:
			t := tkn.Token()
			isStrong = t.Data == "strong"
			if isStrong {
				for _, a := range t.Attr {
					if a.Key == "title" && a.Val == "Followers" {
						isFollowers = true
					}
				}
			}
		case tt == html.TextToken:
			t := tkn.Token()
			if isFollowers {
				vals = append(vals, t.Data)
			}
			isFollowers = false
			isStrong = false
		}
	}
}
func getATags(text string) (data []string) {
	tkn := html.NewTokenizer(strings.NewReader(text))
	var vals []string
	var isA bool
	// var isFollowers bool
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return vals // reached the end of the data
		case tt == html.StartTagToken:
			t := tkn.Token()
			isA = t.Data == "a"
			if isA {
				fmt.Println("beginning of a")
			}
		case tt == html.EndTagToken:
			t := tkn.Token()
			isA = t.Data == "a"
			if isA {
				fmt.Println("end of a")
				isA = false
			}
			// case tt == html.TextToken:
			// 	t := tkn.Token()
			// 	if isFollowers {
			// 		vals = append(vals, t.Data)
			// 	}
			// 	isFollowers = false
			// 	isStrong = false
		}
	}
}

func getMetaContent(text string) (data []string) {
	tkn := html.NewTokenizer(strings.NewReader(text))
	var vals []string
	var isMeta bool
	for {
		tt := tkn.Next()
		switch {
		case tt == html.SelfClosingTagToken:
			t := tkn.Token()
			isMeta = t.Data == "meta"
			if isMeta {
				for _, a := range t.Attr {
					if a.Key == "content" {
						vals = append(vals, a.Val)
					}
				}
			}
			isMeta = false
		case tt == html.ErrorToken:
			return vals // reached the end of the data
		}
	}
}

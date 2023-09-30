package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/darrenparkinson/pixelit-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var bitmapCmd = &cobra.Command{
	Use:   "bitmap",
	Short: "Send a bitmap to PixelIt by ID or as data",
	Run:   sendBitmap,
	PreRun: func(cmd *cobra.Command, args []string) {
		// https://github.com/spf13/viper/issues/233
		viper.BindPFlag("id", cmd.PersistentFlags().Lookup("id"))
		viper.BindPFlag("data", cmd.PersistentFlags().Lookup("data"))
		viper.BindPFlag("text", cmd.PersistentFlags().Lookup("text"))
		viper.BindPFlag("width", cmd.PersistentFlags().Lookup("width"))
		viper.BindPFlag("height", cmd.PersistentFlags().Lookup("height"))
	},
}

func init() {
	rootCmd.AddCommand(bitmapCmd)

	bitmapCmd.PersistentFlags().Int("id", 0, "pixel gallery id - either this or data must be provided")
	bitmapCmd.PersistentFlags().IntSlice("data", nil, "pixel data - comma separated ints")
	bitmapCmd.PersistentFlags().String("text", "", "additional text to display (optional)")
	bitmapCmd.PersistentFlags().Int("width", 0, "width of bitmap (calculated if not provided)")
	bitmapCmd.PersistentFlags().Int("height", 0, "height of bitmap (calculated if not provided)")
}

func sendBitmap(cmd *cobra.Command, args []string) {

	id := viper.GetInt("id")
	data := viper.GetIntSlice("data")
	text := viper.GetString("text")

	if id == 0 && len(data) == 0 {
		log.Fatal("data or id required")
	}

	height, width := viper.GetInt("height"), viper.GetInt("width")

	b := &pixelit.Bitmap{
		Position: &pixelit.Position{
			X: 0,
			Y: 0,
		},
	}
	ab := &pixelit.BitmapAnimation{
		AnimationDelay: 150,
		Rubberbanding:  false,
		LimitLoops:     0,
	}

	s := pixelit.Screen{}

	// data will be priority over id to avoid unnecessary external api calls
	if len(data) > 0 {
		s.Bitmap = b
		s.Bitmap.Data = data
		if height == 0 || width == 0 {
			height = 8
			width = len(data) / 8
		}
		s.Bitmap.Size = &pixelit.Size{
			Width:  width,
			Height: height,
		}
	} else if id != 0 {
		bmp, err := getBMPByID(id)
		if err != nil {
			log.Fatal(err)
		}
		height = bmp.SizeY
		width = bmp.SizeX
		if bmp.Animated {
			s.BitmapAnimation = ab
			prepData := fmt.Sprintf("[%s]", bmp.Data) // need to add extra square brackets for unmarshal
			var animatedData [][]int
			if err := json.Unmarshal([]byte(prepData), &animatedData); err != nil {
				log.Fatal("error decoding data", err)
			}
			s.BitmapAnimation.Data = animatedData
		}
		if !bmp.Animated {
			s.Bitmap = b
			if err := json.Unmarshal([]byte(bmp.Data), &data); err != nil {
				log.Fatal("error decoding data", err)
			}
			s.Bitmap.Size = &pixelit.Size{
				Width:  width,
				Height: height,
			}
			s.Bitmap.Data = data
		}
	}

	if text != "" {
		t := &pixelit.Text{
			TextString:      pixelit.String(text),
			BigFont:         pixelit.Bool(false),
			CenterText:      pixelit.Bool(false),
			ScrollText:      pixelit.String("auto"),
			ScrollTextDelay: pixelit.Int(120),
			Position: &pixelit.Position{
				X: width,
				Y: 1,
			},
			Color: &pixelit.Color{
				R: 255,
				G: 255,
				B: 255,
			},
		}
		s.Text = t
	}
	pxc.SendScreen(&s)
}

type bitmapResponse struct {
	ID       int       `json:"id"`
	DateTime time.Time `json:"dateTime"`
	Name     string    `json:"name"`
	Data     string    `json:"rgB565Array"`
	Animated bool      `json:"animated"`
	SizeX    int       `json:"sizeX"`
	SizeY    int       `json:"sizeY"`
	UserID   int       `json:"userID"`
	Username string    `json:"username"`
	HitCount int       `json:"hitCount"`
}

func getBMPByID(id int) (bitmapResponse, error) {
	var bmpresp bitmapResponse
	u := fmt.Sprintf("https://pixelit.bastelbunker.de/API/GetBMPByID/%d", id)
	resp, err := http.Get(u)
	if err != nil {
		return bmpresp, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&bmpresp)
	if err != nil {
		return bmpresp, err
	}
	return bmpresp, nil
}

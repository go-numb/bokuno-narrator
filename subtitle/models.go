package subtitle

import (
	"fmt"
	"strings"
	"time"

	"github.com/wargarblgarbl/libgosubs/ass"
)

type Page struct {
	PageID       string  `csv:"page_id"`
	Title        string  `csv:"title"`
	Href         string  `csv:"href"`
	Date         string  `csv:"date"`
	Category     string  `csv:"category"`
	CategoryLink string  `csv:"category_link"`
	Force        float64 `csv:"force"`
	Posts        []Post  `csv:"-"`
}

func New(title string) *Page {
	return &Page{
		Title: title,
		Posts: []Post{},
	}
}

type Post struct {
	PageID string `csv:"page_id"`
	// .uid
	UID int `csv:"uid"`
	// .name
	Name string `csv:"name"`
	// .date
	Date string `csv:"date"`
	// .message > .escaped
	Text string `csv:"text"`

	// 字幕付与用（第一発言が0
	Start int64 `csv:"start"`
	End   int64 `csv:"end"`
}

// ToAss from pages struct
// isDialogue is setting the color of the matching hand
func (p *Page) ToAss(isDialogue bool, filename string) error {
	var (
		as     = new(ass.Ass)
		events = make([]ass.Event, 0)
	)

	as.ScriptInfo.Body.Title = p.Title // "Auto stl"
	as.ScriptInfo.Body.ScriptType = "v4.00+"
	as.ScriptInfo.Body.PlayResX = 384
	as.ScriptInfo.Body.PlayResY = 288
	as.ScriptInfo.Body.SBaShadow = "yes"
	as.Styles.Body = p.SetStyle()

	for i, post := range p.Posts {
		style := "Blue"
		if isDialogue { // 掛け合いにする
			if i%2 == 0 {
				style = "Yellow"
			}
		}

		events = append(events, ass.Event{
			Format: "Dialogue",
			Style:  style,
			Start:  _unixMilliToAssTimeFormat(post.Start),
			End:    _unixMilliToAssTimeFormat(post.End),
			Text:   post.Text,
		})
	}

	as.Events.Body = events

	ass.Setheaders(as)

	if err := ass.WriteAss(as, fmt.Sprintf("%s.ass", filename)); err != nil {
		return err
	}

	return nil
}

func _unixMilliToAssTimeFormat(t int64) string {
	d := time.Duration(t) * time.Millisecond
	temp := d.String()
	if !strings.Contains(temp, "m") {
		temp = fmt.Sprintf("0m%s", temp)
	}
	if !strings.Contains(temp, "h") {
		temp = fmt.Sprintf("0h%s", temp)
	}
	if !strings.Contains(temp, ".") {
		temp = fmt.Sprintf("%s.00", temp)
	}
	temp = strings.Replace(temp, "h", ":", 1)
	temp = strings.Replace(temp, "m", ":", 1)
	temp = strings.Replace(temp, "s", "", 1)
	return temp
}

func (p *Page) SetStyle() []ass.Style {
	return []ass.Style{
		{
			// Style: Default,Arial,16,&Hffffff,&Hffffff,&H0,&H0,0,0,0,0,100,100,0,0,1,1,0,2,10,10,10,0
			Format:          "Style",
			Name:            "Default",
			Fontname:        "Arial",
			Fontsize:        18,
			PrimaryColour:   "&Hffffff",
			SecondaryColour: "&Hffffff",
			OutlineColour:   "&H0",
			Backcolour:      "&H0",
			ScaleX:          100,
			ScaleY:          100,
			BorderStyle:     1,
			Outline:         1,
			Shadow:          0,
			Alignment:       2,
			MarginL:         10,
			MarginR:         10,
			MarginV:         10,
			Encoding:        0,
		},
		{
			Format:          "Style",
			Name:            "Black",
			Fontname:        "Arial",
			Fontsize:        18,
			PrimaryColour:   "&H0",
			SecondaryColour: "&H0",
			OutlineColour:   "&Hffffff",
			Backcolour:      "&Hffffff",
			ScaleX:          100,
			ScaleY:          100,
			BorderStyle:     1,
			Outline:         1,
			Shadow:          0,
			Alignment:       2,
			MarginL:         10,
			MarginR:         10,
			MarginV:         10,
			Encoding:        0,
		},
		{
			Format:          "Style",
			Name:            "Blue",
			Fontname:        "Arial",
			Fontsize:        18,
			PrimaryColour:   "&Hf4a903&",
			SecondaryColour: "&Hf4a903&",
			OutlineColour:   "&Hfef5e1&",
			Backcolour:      "&Hfef5e1&",
			ScaleX:          100,
			ScaleY:          100,
			BorderStyle:     1,
			Outline:         1,
			Shadow:          0,
			Alignment:       2,
			MarginL:         10,
			MarginR:         10,
			MarginV:         10,
			Encoding:        0,
		},
		{
			Format:          "Style",
			Name:            "Cyan",
			Fontname:        "Arial",
			Fontsize:        18,
			PrimaryColour:   "&Hd4bc00&",
			SecondaryColour: "&Hd4bc00&",
			OutlineColour:   "&Hfaf7e0&",
			Backcolour:      "&Hfaf7e0&",
			ScaleX:          100,
			ScaleY:          100,
			BorderStyle:     1,
			Outline:         1,
			Shadow:          0,
			Alignment:       2,
			MarginL:         10,
			MarginR:         10,
			MarginV:         10,
			Encoding:        0,
		},
		{
			Format:          "Style",
			Name:            "Green",
			Fontname:        "Arial",
			Fontsize:        18,
			PrimaryColour:   "&H4ac38b&",
			SecondaryColour: "&H4ac38b&",
			OutlineColour:   "&He9f8f1&",
			Backcolour:      "&He9f8f1&",
			ScaleX:          100,
			ScaleY:          100,
			BorderStyle:     1,
			Outline:         1,
			Shadow:          0,
			Alignment:       2,
			MarginL:         10,
			MarginR:         10,
			MarginV:         10,
			Encoding:        0,
		},
		{
			Format:          "Style",
			Name:            "Yellow",
			Fontname:        "Arial",
			Fontsize:        18,
			PrimaryColour:   "&H25a8f9",
			SecondaryColour: "&H25a8f9",
			OutlineColour:   "&He7fdff",
			Backcolour:      "&He7fdff",
			ScaleX:          100,
			ScaleY:          100,
			BorderStyle:     1,
			Outline:         1,
			Shadow:          0,
			Alignment:       2,
			MarginL:         10,
			MarginR:         10,
			MarginV:         10,
			Encoding:        0,
		},
		{
			Format:          "Style",
			Name:            "Red",
			Fontname:        "Arial",
			Fontsize:        18,
			PrimaryColour:   "&H3539e5",
			SecondaryColour: "&H3539e5",
			OutlineColour:   "&Heeebff",
			Backcolour:      "&Heeebff",
			ScaleX:          100,
			ScaleY:          100,
			BorderStyle:     1,
			Outline:         1,
			Shadow:          0,
			Alignment:       2,
			MarginL:         10,
			MarginR:         10,
			MarginV:         10,
			Encoding:        0,
		},
	}
}

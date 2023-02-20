# BOKUNO-NARRATOR ぼくのナレーターさん

_data/narrae.txtを行ごとに読み上げ、字幕ファイルを作ります。  
バックグラウンドでOBSなどで録画したファイルが有れば下記コマンドで字幕ファイルと統合できます。
```sh
$ ffmpeg -i input.mp4 -vf "ass=字幕.ass" -c:v h264_nvenc output.mp4
```


# Usage
```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"bokuno-narrator/subtitle"

	"github.com/go-numb/go-bouyomichan"
)

const (
	// 改行してセリフを書き込む
	narratefile = "./_data/narrate.txt"
)

var (
	// 指定のワードが含まれたコメント（投稿）を除く
	EXCLUDE = []string{}
	// 置き換える
	REPLACEWORDS = []string{"死", "殺"}

    // 掛け合いにする
    isDialogue = true
)

func init() {
	f, _ := os.Open("./_data/excludes.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		EXCLUDE = append(EXCLUDE, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(EXCLUDE)
}

func main() {
	bouyomi := bouyomichan.New("localhost:50001")
	bouyomi.Voice = bouyomichan.VoiceNeutral
	bouyomi.Tone = 110

	words := Read(narratefile)

	content := subtitle.New("market maker bot")
	content.PageID = "mm"

	var (
		start = time.Now()
		in    int64
	)
	for i := 0; i < len(words); i++ {
		if isDialogue {
            if i%2 == 0 {
                bouyomi.Voice = bouyomichan.VoiceDefault
                bouyomi.Tone = 105
            } else {
                bouyomi.Voice = bouyomichan.VoiceNeutral
                bouyomi.Tone = 110
            }
        }

		in = time.Now().UnixMilli() - start.UnixMilli()
		bouyomi.Speaking(words[i])
		for {
			if !bouyomi.IsNowPlayng() {
				break
			}
		}
		time.Sleep(200 * time.Millisecond)
		log.Println("[INFO] end one sentence")

		content.Posts = append(content.Posts, subtitle.Post{
			PageID: content.PageID,
			UID:    i,
			Text:   words[i],
			Start:  in,
			End:    time.Now().UnixMilli() - start.UnixMilli(),
		})
	}

	content.ToAss(false, filepath.Clean(narratefile))
}

func Read(filename string) (words []string) {
	f, _ := os.Open(filepath.Clean(narratefile))
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		w := scanner.Text()
		for j := 0; j < len(EXCLUDE); j++ {
			if strings.Contains(w, EXCLUDE[j]) {
				continue
			}
		}
		for j := 0; j < len(REPLACEWORDS); j++ {
			w = strings.ReplaceAll(w, REPLACEWORDS[j], "◯")
		}
		words = append(words, w)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

```
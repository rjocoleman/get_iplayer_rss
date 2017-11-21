// Copyright © 2016 Robert Coleman <github@robert.net.nz>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gogap/types"
	"github.com/goware/urlx"
	"github.com/kennygrant/sanitize"
	"github.com/rjocoleman/get_iplayer_rss/utils"
	"github.com/spf13/cobra"
)

var iplayerDir string
var outputPath string
var url string

var showName string
var showWeb string
var showChannel string

// DownloadHistory get_iplayer
type DownloadHistory struct {
	PID        string
	Name       string
	Episode    string
	Type       string
	TimeAdded  int64
	Mode       string
	Filename   string
	Versions   string
	Duration   int
	Desc       string
	Channel    string
	Categories string
	Thumbnail  string
	Guidance   string
	Web        string
	EpisodeNum int
	SeriesNum  int
}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate an RSS feed",
	Long:  `Generate an RSS feed from a get_iplayer download_history file`,
	Run: func(cmd *cobra.Command, args []string) {
		downloadHistoryFile, err := os.Open(path.Join(iplayerDir, "download_history"))

		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		defer downloadHistoryFile.Close()

		reader := csv.NewReader(downloadHistoryFile)
		reader.Comma = '|'
		reader.LazyQuotes = true
		reader.FieldsPerRecord = -1

		downloadHistory, err := reader.ReadAll()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		var episode DownloadHistory
		shows := make(map[string][]DownloadHistory)

		// pid name episode type timeadded mode filename versions duration desc channel categories thumbnail guidance web episodenum seriesnum
		for _, entry := range downloadHistory {
			episode.PID = entry[0]
			episode.Name = entry[1]
			episode.Episode = entry[2]
			episode.Type = entry[3]
			episode.TimeAdded, _ = strconv.ParseInt(entry[4], 10, 64)
			episode.Mode = entry[5]
			episode.Filename = entry[6]
			episode.Versions = entry[7]
			episode.Duration, _ = strconv.Atoi(entry[8])
			episode.Desc = entry[9]
			episode.Channel = entry[10]
			episode.Categories = entry[11]
			episode.Thumbnail = entry[12]
			episode.Guidance = entry[13]
			episode.Web = entry[14]
			episode.EpisodeNum, _ = strconv.Atoi(entry[15])
			episode.SeriesNum, _ = strconv.Atoi(entry[16])

			sanatizedShowName := sanitize.Path(episode.Name)
			shows[sanatizedShowName] = append(shows[sanatizedShowName], episode)
		}

		fmt.Println("Parsed Shows:", len(shows))

		for _, show := range shows {
			items := utils.PodcastItems{}

			fmt.Println("Show:", showName)
			fmt.Println("Parsed Episdoes:", len(show))

			for _, episode := range show {
				var item utils.PodcastItem

				showName = episode.Name
				showWeb = episode.Web
				showChannel = episode.Channel

				item.Title = episode.Episode
				item.ITunesAuthor = episode.Name
				item.ITunesSummary = episode.Desc
				item.ITunesImage.Href = episode.Thumbnail
				item.PubDate = types.DateTime(time.Unix(episode.TimeAdded, 0))

				_, filename := filepath.Split(episode.Filename)
				webUrl, _ := urlx.Parse(url)
				webUrl.Path = path.Join(webUrl.Path, filename)

				item.Enclosure.URL = webUrl.String()
				item.Enclosure.Type = "audio/mp4"
				item.Enclosure.Length = episode.Duration

				items = append(items, item)
			}

			rss := utils.NewPodcastRSS()
			rss.Channel = utils.PodcastChannel{
				Title:        showName,
				Link:         showWeb,
				ITunesAuthor: showChannel,
				Items:        items,
			}

			var file bytes.Buffer
			file.WriteString(xml.Header)
			enc := xml.NewEncoder(&file)
			enc.Indent("", "    ")
			if err := enc.Encode(rss); err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}

			sanatizedShowName := sanitize.Path(showName)
			outputFilepath := path.Join(outputPath, sanatizedShowName) + ".rss"
			fmt.Println("Writing:", outputFilepath)
			err := ioutil.WriteFile(outputFilepath, file.Bytes(), 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(&iplayerDir, "directory", "d", "/etc/get_iplayer", "Path to get_iplayer directory")
	genCmd.Flags().StringVarP(&outputPath, "output-path", "o", "", "RSS file output path e.g. /var/www")
	genCmd.Flags().StringVarP(&url, "url", "u", "", "URL to webroot e.g. https://example.com/path/to/dir")
}

package utils

import (
	"encoding/xml"

	"github.com/gogap/types"
)

// PodcastRSS ...
type PodcastRSS struct {
	XMLName     xml.Name `xml:"rss"`
	XmlnsItunes string   `xml:"xmlns:itunes,attr,omitempty"`
	Version     string   `xml:"version,attr,omitempty"`
	Channel     PodcastChannel
}

// PodcastChannel ...
type PodcastChannel struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title,omitempty"`
	Link           string   `xml:"link,omitempty"`
	Language       string   `xml:"language,omitempty"`
	Copyright      string   `xml:"copyright,omitempty"`
	ITunesSubtitle string   `xml:"itunes:subtitle,omitempty"`
	ITunesAuthor   string   `xml:"itunes:author,omitempty"`
	ITunesSummary  string   `xml:"itunes:summary,omitempty"`
	Description    string   `xml:"description,omitempty"`
	ITunesOwner    struct {
		ITunesName  string `xml:"itunes:name,omitempty"`
		ITunesEmail string `xml:"itunes:email,omitempty"`
	} `xml:"itunes:owner,omitempty"`
	ITunesImage struct {
		Href string `xml:"href,attr,omitempty"`
	} `xml:"itunes:image,omitempty"`
	ITunesCategory struct {
		Text string `xml:"text,attr,omitempty"`
	} `xml:"itunes:category,omitempty"`
	Items PodcastItems
}

// PodcastItem ...
type PodcastItem struct {
	XMLName        xml.Name `xml:"item"`
	Title          string   `xml:"title,omitempty"`
	ITunesAuthor   string   `xml:"itunes:author,omitempty"`
	ITunesSubtitle string   `xml:"itunes:subtitle,omitempty"`
	ITunesSummary  string   `xml:"itunes:summary,omitempty"`
	ITunesImage    struct {
		Href string `xml:"href,attr,omitempty"`
	} `xml:"itunes:image,omitempty"`
	Enclosure struct {
		URL    string `xml:"url,attr,omitempty"`
		Length int    `xml:"length,attr,omitempty"`
		Type   string `xml:"type,attr,omitempty"`
	} `xml:"enclosure,omitempty"`
	GUID           string         `xml:"guid,omitempty"`
	PubDate        types.DateTime `xml:"pubDate,omitempty"`
	ITunesDuration string         `xml:"itunes:duration,omitempty"`
}

// PodcastItems ...
type PodcastItems []PodcastItem

// NewPodcastRSS ...
func NewPodcastRSS() *PodcastRSS {
	return &PodcastRSS{
		XmlnsItunes: "http://www.itunes.com/dtds/podcast-1.0.dtd",
		Version:     "2.0",
	}
}

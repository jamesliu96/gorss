/*
 * gorss.go
 * Copyright (C) 2015  James Liu
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package gorss

import (
    "bytes"
    "encoding/xml"
)

type Channel struct {
    Title          string   `xml:"title"`
    Link           string   `xml:"link"`
    Description    string   `xml:"description"`

    Language       string   `xml:"language,omitempty"`
    Copyright      string   `xml:"copyright,omitempty"`
    ManagingEditor string   `xml:"managingEditor,omitempty"`
    WebMaster      string   `xml:"webMaster,omitempty"`
    PubDate        string   `xml:"pubDate,omitempty"`
    LastBuildDate  string   `xml:"lastBuildDate,omitempty"`
    Category       string   `xml:"category,omitempty"`
    Generator      string   `xml:"generator,omitempty"`
    Docs           string   `xml:"docs,omitempty"`
    TTL            string   `xml:"ttl,omitempty"`
    SkipHours      string   `xml:"skiphours,omitempty"`
    SkipDays       string   `xml:"skipdays,omitempty"`

    Items          []*Item

    XMLName        xml.Name `xml:"channel"`
}

type Item struct {
    Link        string   `xml:"link"`
    Description string   `xml:"description"`

    Title       string   `xml:"title,omitempty"`
    PubDate     string   `xml:"pubDate,omitempty"`
    Author      string   `xml:"author,omitempty"`
    Guid        string   `xml:"guid,omitempty"`
    Comments    string   `xml:"comments,omitempty"`

    XMLName     xml.Name `xml:"item"`
}

const Generator string = "GoRSS - https://github.com/jamesliu96/gorss"

func (c *Channel) AddItem(i *Item) {
    c.Items = append(c.Items, i)
}

func (c *Channel) Publish() []byte {
    c.Generator = Generator
    out, err := xml.MarshalIndent(c, "  ", "    ")
    if err != nil {
        panic(err)
    }
    return bytes.Join([][]byte{
        []byte(`<?xml version="1.0"?>`),
        []byte(`<rss version="2.0">`),
        out,
        []byte(`</rss>`),
    }, []byte("\n"))
}

func (c *Channel) PublishCompressed() []byte {
    c.Generator = Generator
    out, err := xml.Marshal(c)
    if err != nil {
        panic(err)
    }
    return bytes.Join([][]byte{
        []byte(`<?xml version="1.0"?>`),
        []byte(`<rss version="2.0">`),
        out,
        []byte(`</rss>`),
    }, []byte(""))
}

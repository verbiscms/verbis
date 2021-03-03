// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wordpress

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"time"
)

type WpXML struct {
	Channel        Channel `xml:"channel"`
	CreatorCounts  map[string]int
	CreatorToIndex map[string]int
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Authors []Author `xml:"author"`
	Items   []Item   `xml:"item"`
}

// Author is the WordPress XML author object.
type Author struct {
	AuthorID          int        `xml:"author_id"`
	AuthorLogin       string     `xml:"author_login"`
	AuthorEmail       string     `xml:"author_email"`
	AuthorDisplayName string     `xml:"author_display_name"`
	AuthorFirstName   string     `xml:"author_first_name"`
	AuthorLastName    string     `xml:"author_last_name"`
	AuthorArticles    []ItemThin `xml:"-"`
}

// Item is a WordPress XML item which can be a post, page or other object.
type Item struct {
	ID           int        `xml:"post_id"`
	Title        string     `xml:"title"`
	Creator      string     `xml:"creator"`
	Encoded      []string   `xml:"encoded"`
	IsSticky     int        `xml:"is_sticky"`
	Link         string     `xml:"link"`
	PubDate      string     `xml:"pubDate"`
	Description  string     `xml:"description"`
	PostDate     string     `xml:"post_date"`
	PostDateGmt  string     `xml:"post_date_gmt"`
	PostName     string     `xml:"post_name"`
	PostType     string     `xml:"post_type"`
	Status       string     `xml:"status"`
	Categories   []Category `xml:"category"`
	Comments     []Comment  `xml:"comment"`
	Meta         []Meta     `xml:"postmeta"`
	Content      string
	PostDatetime time.Time
	PubDatetime  time.Time
}

// ItemThin is a WordPress XML item that is used as additional
// metadata in the Author object.
type ItemThin struct {
	Title string
	Index int
}

type Category struct {
	Domain      string `xml:"domain,attr"`
	DisplayName string `xml:",chardata"`
	URLSlug     string `xml:"nicename,attr"`
}

type Comment struct {
	ID          int    `xml:"comment_id"`
	Parent      int    `xml:"comment_parent"`
	Author      string `xml:"comment_author"`
	AuthorEmail string `xml:"comment_author_email"`
	AuthorURL   string `xml:"comment_author_url"`
	DateGmt     string `xml:"comment_date_gmt"`
	Content     string `xml:"comment_content"`
	IndentLevel int    `xml:"-"`
}

type Meta struct {
	Text      string `xml:",chardata"`
	MetaKey   string `xml:"meta_key"`
	MetaValue string `xml:"meta_value"`
}

func NewWordpressXML() WpXML {
	return WpXML{
		CreatorCounts:  map[string]int{},
		CreatorToIndex: map[string]int{}}
}

// ReadXml reads a WordPress XML file from the provided path.
func (wpxml *WpXML) ReadFile(filepath string) error {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(bytes, &wpxml)
	if err != nil {
		return err
	}
	err = wpxml.inflate()
	if err != nil {
		return err
	}
	return nil
}

func (wpxml *WpXML) inflate() error {
	creatorMap := map[string]int{}
	for i, item := range wpxml.Channel.Items {
		if len(item.Creator) > 0 {
			_, ok := creatorMap[item.Creator]
			if ok {
				creatorMap[item.Creator]++
			} else {
				creatorMap[item.Creator] = 1
			}
		}

		item = wpxml.inflateItem(item)
		wpxml.Channel.Items[i] = item
	}
	wpxml.CreatorCounts = creatorMap
	err := wpxml.inflateAuthors()
	if err != nil {
		return err
	}
	return nil
}

func (wpxml *WpXML) inflateItem(item Item) Item {
	if len(item.Encoded) > 0 && len(item.Encoded[0]) > 0 {
		item.Content = item.Encoded[0]
		item.Encoded[0] = ""
	}
	if len(item.PostDate) > 0 {
		dt, err := time.Parse("Mon Jan 02 2006 15:04:05 GMT-0700", item.PostDate)
		if err == nil {
			item.PostDatetime = dt
		}
	}
	if len(item.PubDate) > 0 {
		dt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err == nil {
			item.PubDatetime = dt
		}
	}
	return item
}

func (wpxml *WpXML) inflateAuthors() error {
	a2i := wpxml.AuthorsToIndex()
	for i, item := range wpxml.Channel.Items {
		if len(item.Creator) > 0 {
			authorLogin := item.Creator
			if _, ok := a2i[authorLogin]; ok {
				authorIndex := a2i[authorLogin]
				itemThin := ItemThin{Title: item.Title, Index: i}
				if wpxml.Channel.Authors[authorIndex].AuthorArticles == nil {
					wpxml.Channel.Authors[authorIndex].AuthorArticles = []ItemThin{}
				}
				wpxml.Channel.Authors[authorIndex].AuthorArticles = append(wpxml.Channel.Authors[authorIndex].AuthorArticles, itemThin)
			}
		}
	}
	wpxml.CreatorToIndex = a2i
	return nil
}

func (wpxml *WpXML) AuthorsToIndex() map[string]int {
	a2i := map[string]int{}
	for i, author := range wpxml.Channel.Authors {
		a2i[author.AuthorLogin] = i
	}
	return a2i
}

// AuthorForLogin returns the Author object for a given AuthorLogin
// or username.
func (wpxml *WpXML) AuthorForLogin(authorLogin string) (Author, error) {
	a2i := wpxml.CreatorToIndex
	if index, ok := a2i[authorLogin]; ok {
		author := wpxml.Channel.Authors[index]
		return author, nil
	}
	return Author{}, errors.New("author Not Found")
}

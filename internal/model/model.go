package model

import "time"

type Item struct {
	Title      string
	Categories []string
	Link       string
	Date       time.Time
	Summary    string
	SourceName string
}

type Source struct {
	ID        int64
	Name      string
	FeedUrl   string
	CreatedAt time.Time
}

type Article struct {
	Id          int64
	SourceId    int64
	Title       string
	Link        string
	Summary     string
	PublishedAt time.Time
	CreatedAt   time.Time
	PostedAt    time.Time
}

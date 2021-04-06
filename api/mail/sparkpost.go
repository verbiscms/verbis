// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mail

import (
	sp "github.com/SparkPost/gosparkpost"
)

type sparkPost struct {
	cfg    Config
	client sp.Client
}

const (
	SparkAPIVersion = 1
)

// newSparkPost
//
//
func newSparkPost(cfg Config) (*sparkPost, error) {
	config := &sp.Config{
		BaseUrl:    cfg.URL,
		ApiKey:     cfg.ApiKey,
		ApiVersion: SparkAPIVersion,
	}

	var client sp.Client
	err := client.Init(config)
	if err != nil {
		return nil, err
	}

	return &sparkPost{
		cfg:    cfg,
		client: client,
	}, nil
}

// Send
//
//
func (s *sparkPost) Send(t Transmission) (Response, string, error) {
	content := sp.Content{
		HTML:    t.HTML,
		From:    s.cfg.FromAddress,
		Subject: t.Subject,
	}

	if len(t.Attachments) != 0 {
		var att []sp.Attachment
		for _, v := range t.Attachments {
			att = append(att, sp.Attachment{
				MIMEType: v.MIMEType,
				Filename: v.Filename,
				B64Data:  *v.B64Data,
			})
		}
		content.Attachments = att
	}

	tx := &sp.Transmission{
		Recipients: t.To,
		Content:    content,
	}

	id, response, err := s.client.Send(tx)
	if err != nil {
		return Response{}, "", err
	}

	return Response{
		StatusCode: response.HTTP.StatusCode,
		Body:       string(response.Body),
		Headers:    response.HTTP.Header,
		Id:         id,
	}, id, nil
}

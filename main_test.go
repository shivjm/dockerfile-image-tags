package main

import (
	"testing"

	"github.com/asottile/dockerfile"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	expected := []Image{
		{Name: "golang", Tag: "1.17.0-alpine"},
		{Name: "common", Tag: " * "},
		{Name: "common", Tag: " * "},
		{Name: "common", Tag: " * "},
		{Name: "viaductoss/ksops", Tag: "v3.0.0"},
		{Name: "quay.io/argoproj/argocd", Tag: "$ARGOCD_VERSION"},
	}

	commands, err := dockerfile.ParseFile("tests/Dockerfile.1")

	if err != nil {
		t.Errorf("Could not open Dockerfile.1: %s", err)
	}

	tags := getImages(commands, " * ")

	assert.Equal(t, expected, tags)
}

func TestQuery(t *testing.T) {
	cases := []struct {
		query string
		match bool
		tag   string
	}{
		{query: "foo", match: false, tag: ""},
		{query: "viaductoss/ksops", match: true, tag: "v3.0.0"},
		{query: "golang", match: true, tag: "1.17.0-alpine"},
		{query: "common", match: true, tag: "?"},
	}

	commands, err := dockerfile.ParseFile("tests/Dockerfile.1")

	if err != nil {
		t.Errorf("Could not open Dockerfile.1: %s", err)
	}

	tags := getImages(commands, "?")

	for _, c := range cases {
		result, err := getSingleTag(tags, c.query)

		if c.match {
			assert.NoError(t, err, "must match %v", c.query)
			assert.Equal(t, result, c.tag)
		} else {
			assert.Error(t, err, "must not match %v", c.query)
			assert.Equal(t, result, "")
		}
	}
}

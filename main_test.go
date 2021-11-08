package main

import (
	"testing"

	"github.com/asottile/dockerfile"
	"github.com/stretchr/testify/assert"

	"github.com/shivjm/dockerfile-image-tags/pkg/images"
)

func TestParsing(t *testing.T) {
	expected := []images.Image{
		{Name: "golang", Tag: "1.17.0-alpine"},
		{Name: "common", Tag: " * "},
		{Name: "common", Tag: "fixme"},
		{Name: "common", Tag: " * "},
		{Name: "viaductoss/ksops", Tag: "v3.0.0"},
		{Name: "quay.io/argoproj/argocd", Tag: "$ARGOCD_VERSION"},
	}

	commands, err := dockerfile.ParseFile("tests/Dockerfile.1")

	if err != nil {
		t.Errorf("Could not open Dockerfile.1: %s", err)
	}

	tags := images.GetImages(commands, " * ")

	assert.Equal(t, expected, tags)
}

func TestQuery(t *testing.T) {
	cases := []struct {
		query      string
		occurrence int
		match      bool
		tag        string
	}{
		{query: "foo", occurrence: 0, match: false, tag: ""},
		{query: "viaductoss/ksops", occurrence: 0, match: true, tag: "v3.0.0"},
		{query: "golang", occurrence: 0, match: true, tag: "1.17.0-alpine"},
		{query: "common", occurrence: 0, match: true, tag: "?"},
		{query: "foo", occurrence: 1, match: false, tag: ""},
		{query: "viaductoss/ksops", occurrence: 1, match: true, tag: "v3.0.0"},
		{query: "golang", occurrence: 1, match: true, tag: "1.17.0-alpine"},
		{query: "common", occurrence: 1, match: true, tag: "?"},
		{query: "viaductoss/ksops", occurrence: 2, match: false, tag: ""},
		{query: "common", occurrence: 2, match: true, tag: "fixme"},
	}

	commands, err := dockerfile.ParseFile("tests/Dockerfile.1")

	if err != nil {
		t.Errorf("Could not open Dockerfile.1: %s", err)
	}

	tags := images.GetImages(commands, "?")

	for _, c := range cases {
		result, err := images.GetSingleTag(tags, c.query, c.occurrence)

		if c.match {
			assert.NoError(t, err, "must match %v", c.query)
			assert.Equal(t, result, c.tag)
		} else {
			assert.Error(t, err, "must not match %v", c.query)
			assert.Equal(t, result, "")
		}
	}
}

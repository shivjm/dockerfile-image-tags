package main

import (
	"testing"

	"github.com/asottile/dockerfile"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	expected := []Image{
		{Name: "golang", Tag: "1.17.0-alpine"},
		{Name: "common", Tag: "?"},
		{Name: "common", Tag: "?"},
		{Name: "common", Tag: "?"},
		{Name: "viaductoss/ksops", Tag: "v3.0.0"},
		{Name: "quay.io/argoproj/argocd", Tag: "$ARGOCD_VERSION"},
	}

	commands, err := dockerfile.ParseFile("tests/Dockerfile.1")

	if err != nil {
		t.Errorf("Could not open Dockerfile.1: %s", err)
	}

	tags := getTags(commands)

	assert.Equal(t, expected, tags)
}

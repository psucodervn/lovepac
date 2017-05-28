package packer_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/RaniSputnik/lovepac/packer"
)

func TestRunOutputsAtlasAndDescriptor(t *testing.T) {
	files := []string{
		"button_active.png",
		"button_hover.png",
		"button.png",
		"character_evil.png",
		"character_hero.png",
	}
	expected := map[string]string{
		"myatlas-1.png": "",
		"myatlas-1.lua": "",
	}

	outputRecorder := packer.NewOutputRecorder()
	params := &packer.Params{
		Name:   "myatlas",
		Format: packer.FormatLove,
		Input:  packer.NewFilenameStream("./fixtures", files...),
		Output: outputRecorder,
		Width:  1024,
		Height: 1024,
	}

	err := packer.Run(context.Background(), params)
	got := outputRecorder.Got()

	if err != nil {
		t.Errorf("Expected run to succeed without error but got '%s'", err)
	}

	for gotFile := range got {
		if _, ok := expected[gotFile]; !ok {
			t.Errorf("Got unexpected file '%s'", gotFile)
		}
	}

	for expect := range expected {
		if _, ok := got[expect]; !ok {
			t.Errorf("Expected file '%s' to be outputted", expect)
		}
	}
}

func TestRunWithoutParamsSpecifiedUsesSensibleDefaults(t *testing.T) {
	files := []string{"button.png"}
	expected := map[string]string{
		fmt.Sprintf("%s-1.png", packer.DefaultAtlasName): "",
		fmt.Sprintf("%s-1.lua", packer.DefaultAtlasName): "",
	}

	outputRecorder := packer.NewOutputRecorder()
	params := &packer.Params{
		Input:  packer.NewFilenameStream("./fixtures", files...),
		Output: outputRecorder,
	}

	err := packer.Run(context.Background(), params)
	got := outputRecorder.Got()

	if err != nil {
		t.Errorf("Expected run to succeed without error but got '%s'", err)
	}

	for gotFile := range got {
		if _, ok := expected[gotFile]; !ok {
			t.Errorf("Got unexpected file '%s'", gotFile)
		}
	}

	for expect := range expected {
		if _, ok := got[expect]; !ok {
			t.Errorf("Expected file '%s' to be outputted", expect)
		}
	}
}
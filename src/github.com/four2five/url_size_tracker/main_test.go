package main

import(
  "testing"
)

const (
	url1 string = "www.one.com" // len 11
	url2 string = "www.onetwo.com" // len 14
	url3 string = "www.onethree.com" // len 16   
	url4 string = "www.onethreefour.com" // len 20
	url5 string = "www.onethreefourfive.com" // len 24
)

func TestThreeAddsWithOneGetMedian(t *testing.T) {
	tracker := NewURLLengthTracker()

	tracker.AddURL(url1)
	tracker.AddURL(url2)
	tracker.AddURL(url3)

	medianLength := tracker.GetMedianURLLength()

	if 14 != medianLength {
		t.Errorf("Expected length: 14, actual %d", medianLength)
	}
}

func TestThreeAddsGetMedianTwoAddsOfLowerLengthsGetMedian(t *testing.T) {
	tracker := NewURLLengthTracker()

	tracker.AddURL(url2)
	tracker.AddURL(url3)
	tracker.AddURL(url4)

	medianLength := tracker.GetMedianURLLength()

	if 16 != medianLength {
		t.Errorf("Expected length: 16, actual %d", medianLength)
	}

	tracker.AddURL(url1)
	tracker.AddURL(url1)

	medianLength = tracker.GetMedianURLLength()

	if 14 != medianLength {
		t.Errorf("Expected length: 14, actual %d", medianLength)
	}
}

func TestThreeAddsGetMedianTwoAddsOfLongerLengthsGetMedian(t *testing.T) {
	tracker := NewURLLengthTracker()

	tracker.AddURL(url2)
	tracker.AddURL(url3)
	tracker.AddURL(url4)

	medianLength := tracker.GetMedianURLLength()

	if 16 != medianLength {
		t.Errorf("Expected length: 16, actual %d", medianLength)
	}

	tracker.AddURL(url5)
	tracker.AddURL(url5)

	medianLength = tracker.GetMedianURLLength()

	if 20 != medianLength {
		t.Errorf("Expected length: 20, actual %d", medianLength)
	}
}

func TestThreeAddsGetMedianTwoAddsOfOffsettingLengthsGetMedian(t *testing.T) {
	tracker := NewURLLengthTracker()

	tracker.AddURL(url2)
	tracker.AddURL(url3)
	tracker.AddURL(url4)

	medianLength := tracker.GetMedianURLLength()

	if 16 != medianLength {
		t.Errorf("Expected length: 16, actual %d", medianLength)
	}

	tracker.AddURL(url1)
	tracker.AddURL(url5)

	medianLength = tracker.GetMedianURLLength()

	if 16 != medianLength {
		t.Errorf("Expected length: 16, actual %d", medianLength)
	}
}

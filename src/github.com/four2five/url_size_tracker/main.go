package main

import(
	"fmt"
)

const (
  MAX_URL_SIZE int = 32
  DEBUG bool = false
)


type URLLengthTracker struct {
  TrackerEmpty bool
  URLLengths []int
  lastMedianIndex int
  offsetWithinIndex int
  totalURLCount int
  recentValuesLessThanLastMedian int
  recentValuesEqualToLastMedian int
  recentValuesGreaterThanLastMedian int
}

func NewURLLengthTracker() *URLLengthTracker {
  tracker := &URLLengthTracker{}
  tracker.URLLengths = make([]int, MAX_URL_SIZE)
  tracker.TrackerEmpty = true

  return tracker
}

func (u *URLLengthTracker) AddURL(url string){
  urlLength := len(url)
  // Update the counter corresponding to the length of this URL
  if urlLength >= MAX_URL_SIZE {
    urlLength = MAX_URL_SIZE
  }

  u.URLLengths[urlLength]++

  if u.TrackerEmpty {
    u.lastMedianIndex = urlLength
    u.TrackerEmpty = false
  } else {
    // Update the appropriate value indicating which side of the current
    // median this URL length fell on
    if urlLength < u.lastMedianIndex {
      u.recentValuesLessThanLastMedian++
    } else if urlLength == u.lastMedianIndex {
      u.recentValuesEqualToLastMedian++
    } else { // implicitly, urlLength > u.lastMedianIndex
      u.recentValuesGreaterThanLastMedian++
    }
  }

}

// GetMedianURLLength updates the median index based
// on the data that has come in since the last call to
// GetMedianURLLength
func (u *URLLengthTracker) GetMedianURLLength() int {
  if DEBUG {
	  fmt.Printf("current counts: %+v\n", u.URLLengths)
  }
  totalChange := u.recentValuesGreaterThanLastMedian - u.recentValuesLessThanLastMedian
	// divide by half and round up to get the net change
	netOffsetChange := totalChange/2
	if totalChange % 2 > 0 {
		netOffsetChange++
	}

  if DEBUG {
	  fmt.Printf("netOffsetChange %d\n", netOffsetChange)
  }

  // TODO: Adjust by the number of recent values that were equal to the index
  /*
  if netOffsetChange > 0 {
    netOffsetChange -= u.u.recentValuesEqualToLastMedian
  } else if netOffsetChange < 0 {
    netOffsetChange += u.u.recentValuesEqualToLastMedian
  }
  */

  // Get to the end of the currend index, if appropriate
  // Now, update the median index accordingly
  if netOffsetChange > 0 {
    // First, move through the remaining values of url length pointed to by the
    // current median index
    if u.URLLengths[u.lastMedianIndex] - (u.offsetWithinIndex+1) < netOffsetChange  {
      netOffsetChange -= (u.URLLengths[u.lastMedianIndex] - (u.offsetWithinIndex+1))
    } else {
      u.offsetWithinIndex += netOffsetChange
      netOffsetChange = 0
    }

    for netOffsetChange > 0 {
      u.lastMedianIndex++
      netOffsetChange -= u.URLLengths[u.lastMedianIndex]
    }

    // If we over shot, move back within the last index URL count
    for netOffsetChange < 0 {
      u.offsetWithinIndex = u.URLLengths[u.lastMedianIndex] - 1
      u.offsetWithinIndex--
      netOffsetChange++
    }
  } else if netOffsetChange < 0 {
    // First, move through the remaining values of url length pointed to by the
    // current median index
    if (0 - u.offsetWithinIndex) > netOffsetChange {
      netOffsetChange += (u.offsetWithinIndex)
    } else {
      u.offsetWithinIndex -= netOffsetChange
      netOffsetChange = 0
    }

    for netOffsetChange < 0 {
      u.lastMedianIndex--
      netOffsetChange += u.URLLengths[u.lastMedianIndex]
    }

    // If we over shot, move back within the last index URL count
    for netOffsetChange > 0 {
      u.offsetWithinIndex = 0
      u.offsetWithinIndex++
      netOffsetChange--
    }
  }

	u.recentValuesGreaterThanLastMedian = 0
	u.recentValuesLessThanLastMedian = 0
	u.recentValuesEqualToLastMedian = 0

  return u.lastMedianIndex
}

func main() {
}

package timespan

//////////////////////
//
// A struct for representing time spans,
// a factory method that does some input validation,
// and a function for determining whether two spans
// overlap.
//
//////////////////////

type Timespan struct {
  StartTime int64
  EndTime int64
}

// NewTimespan enforces the invariant that a Timespan's
// end time cannot be prior to the start time.
func NewTimespan(startTime, endTime int64) *Timespan {
  // Validate the inputs
  if endTime < startTime {
    return nil
  } else {
    return &Timespan{startTime, endTime}
  }
}

// TimeSpansOverlap tests whether the two Timespan parameters
// overlap at all. The invariant that a given Timespan's
// start time is <= its end time is leveraged.
func TimeSpansOverlap(spanA, spanB *Timespan) bool {
  // Test whether spanA ends before spanB begins or if
  // spanA starts after spanB ends. If either is true, the spans do not overlap.
  if spanA.EndTime < spanB.StartTime || spanA.StartTime > spanB.EndTime {
    return false
  } else {
    return true
  }
}



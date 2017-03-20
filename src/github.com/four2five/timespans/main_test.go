package main_test

import(
  "testing"

  "github.com/four2five/timespans/timespan"
)

const(
  ts1 int64 = 1489995000
  ts2 int64 = 1489995010
  ts3 int64 = 1489995020
  ts4 int64 = 1489995030
  ts5 int64 = 1489995040
)

func TestSpanConstructionWithBadInputs(t *testing.T) {
  spanA := timespan.NewTimespan(ts2, ts1)

  if spanA != nil {
    t.Errorf("Able to create a Timespan with end time < start time")
  }
}

func TestSpanADoesNotOverlapAndIsPriorToB(t *testing.T) {
  spanA := timespan.NewTimespan(ts1, ts2)
  spanB := timespan.NewTimespan(ts3, ts4)

  if timespan.TimeSpansOverlap(spanA, spanB) {
    t.Errorf("Non-overlapping spans were claimed to overlap")
  }
}

func TestSpanADoesNotOverlapAndIsAfterB(t *testing.T) {
  spanA := timespan.NewTimespan(ts3, ts4)
  spanB := timespan.NewTimespan(ts1, ts2)

  if timespan.TimeSpansOverlap(spanA, spanB) {
    t.Errorf("Non-overlapping spans were claimed to overlap")
  }
}

func TestSpanAOverlapsBAtStart(t *testing.T) {
  spanA := timespan.NewTimespan(ts1, ts2)
  spanB := timespan.NewTimespan(ts2, ts4)

  if !timespan.TimeSpansOverlap(spanA, spanB) {
    t.Errorf("Wverlapping spans were claimed to not overlap")
  }
}

func TestSpanAIsASubsetOfSpanB(t *testing.T) {
  spanA := timespan.NewTimespan(ts2, ts4)
  spanB := timespan.NewTimespan(ts1, ts4)

  if !timespan.TimeSpansOverlap(spanA, spanB) {
    t.Errorf("Wverlapping spans were claimed to not overlap")
  }
}


func TestSpanAOverlapsBAtEnd(t *testing.T) {
  spanA := timespan.NewTimespan(ts3, ts5)
  spanB := timespan.NewTimespan(ts1, ts4)

  if !timespan.TimeSpansOverlap(spanA, spanB) {
    t.Errorf("Wverlapping spans were claimed to not overlap")
  }
}


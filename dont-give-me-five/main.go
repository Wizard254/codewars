package kata

// Adapted from https://www.codewars.com/kata/5813d19765d81c592200001a/train/go


func DontGiveMeFive(start int, end int) int {
  if start < 0 && end < 0 {
    // Handle both start and end are negatives
    // -4 -3 -2 -1 
    return DontGiveMeFivePositives(-end, -start)
  } else if start < 0 {
    // Handle case when start is negative
    // -3 -2 -1 0 1 2 3
    return DontGiveMeFivePositives(1, -start) + DontGiveMeFivePositives(1, end) + 1
  }
  
  return DontGiveMeFivePositives(start, end)
}

// DontGiveMeFivePositives same as DontGiveMeFive but only handles positive values
func DontGiveMeFivePositives(start, end int) int {
  // 5 15 25 35 45
  // 1  3  5  7  9
  count := (end - start) + 1
  return count - (HowMany5(end) - HowMany5(start))
}

// HowMany5 Given an `end` value, returns the count of numbers with 5 up to and including `end`
func HowMany5(end int) int {
  odd := end / 5
  n := (odd + 1) / 2
  return n
}



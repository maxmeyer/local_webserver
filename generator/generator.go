package generator

import "time"
import "strconv"
import "math/rand"

func RandomInt(min, max int) int {
  rand.Seed(time.Now().Unix())

  return rand.Intn(max - min) + min
}

func RandomStr(min, max int) string {
  return strconv.Itoa(RandomInt(min, max))
}

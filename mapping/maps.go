package mapping

import (
  "fmt"
)

// maps are unordered lists
// they are flexible
// they are references

// building maps
// map[key type]value type
// make(map[key type]value type, size)
// key type must be a comparable type - bools, strings, numerical
// all keys should be unique

func Mapping() {
  myMap := make(map[int]string)
  fmt.Println(myMap)
  myMap[1] = "Kenya"
  fmt.Println(myMap)

  maplips := map[string]int{
    "love": 1,
    "hate": 2,
  }
  fmt.Println(maplips)

  // insert, update and delete items in a map
  // mapname[key] = value
  delete(maplips, "hate")
  fmt.Println(maplips)

// maps are not preferred for concurrency
// specify size of large maps
}

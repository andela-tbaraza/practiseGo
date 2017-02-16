package main

import (
  "fmt"
  "strings"
)

var cars = []map[string]interface{}{
  map[string]interface{}{
  "label": "Autos",
  "subs":[]map[string]interface{}{
    map[string]interface{}{"label": "SUVs", "subs": []map[string]interface{}{}},
    map[string]interface{}{
        "label": "Trucks",
        "subs": []map[string]interface{}{
          map[string]interface{}{"label": "2 Wheel Drive", "subs": []map[string]interface{}{}},
          map[string]interface{}{
            "label": "4 Wheel Drive",
            "subs": []map[string]interface{}{
              map[string]interface{}{"label": "Ford","subs": []map[string]interface{}{}},
              map[string]interface{}{"label": "Chevrolet", "subs": []map[string]interface{}{}},
            },
          },
        },
      },
    map[string]interface{}{"label": "Sedan", "subs": []map[string]interface{}{}},
  },
  },
}

// func deepSearch(labels []string){
//
// }

func search(label string, items []map[string]interface{}) (bool, map[string]interface{}){
  for _, item := range items {
    if item["label"] == label {
      return true, item
    }
   }
   return  false, nil
}
// if given a string of parent/child/.....
// it should return the attributes under the last child

func display(results []map[string]interface{}){
  for _, resItem := range results {
    reslabel := resItem["label"]
    fmt.Println(reslabel)
  }

}

func main () {
  testString := "Autos/Trucks/Wheel Drive"

  searchLabels := strings.Split(testString, "/")
  labelItemsInFocus :=  cars
  currentLevelLabel := "cars"

  for _, label := range searchLabels {
    if found, labelItems := search(label, labelItemsInFocus); found{
      fmt.Printf("Found %s in %s\n", label, currentLevelLabel)
      labelItemsInFocus = labelItems["subs"].([]map[string]interface{})
      currentLevelLabel = labelItems["label"].(string)
    } else {
      /// here label is not in this current level
      fmt.Printf("Could not find %s in %s\n", label, currentLevelLabel)
    }
  }

  fmt.Printf("\n\n\n%s sub items are : \n", currentLevelLabel)
  display(labelItemsInFocus)
  fmt.Println("Currrent level parents: ", testString)
  fmt.Println("Currrent level stuff: ", labelItemsInFocus)

  // for _, car := range cars {
  //   fmt.Println(car)
  // }
}

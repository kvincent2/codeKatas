package main

func FindUniq(arr []float32) float32 {
  var unique float32
    item_count := dup_count(arr)
    for k, v := range item_count {
      if v == 1 {
      unique = k
      }
    }
  return unique
}

 func dup_count(list []float32) map[float32]int {

 	duplicate_frequency := make(map[float32]int)

 	for _, item := range list {
 		// check if the item/element exist in the duplicate_frequency map

 		_, exist := duplicate_frequency[item]

 		if exist {
 			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
 		} else {
 			duplicate_frequency[item] = 1 // else start counting from 1
 		}
 	}
 	return duplicate_frequency
 }
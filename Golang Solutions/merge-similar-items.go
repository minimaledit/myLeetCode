/*
https://leetcode.com/problems/merge-similar-items/description/

You are given two 2D integer arrays, items1 and items2, representing two sets of items. Each array items has the following properties:
- items[i] = [valuei, weighti] where valuei represents the value and weighti represents the weight of the ith item.
- The value of each item in items is unique.

Return a 2D integer array ret where ret[i] = [valuei, weighti], with weighti being the sum of weights of all items with value valuei.

Note: ret should be returned in ascending order by value.
*/
func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    numCounter := make([]int, 1001)
	for i := 0; i < len(items1); i++ {
		numCounter[items1[i][0]] += items1[i][1]
	}
	for i := 0; i < len(items2); i++ {
		numCounter[items2[i][0]] += items2[i][1]
	}
	var mergedArray [][]int
	for i := 1; i <= 1000; i++ {
		if numCounter[i] != 0 {
			mergedArray = append(mergedArray, []int{i, numCounter[i]})
		}
	}
	return mergedArray
}

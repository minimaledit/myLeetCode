/*
https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/description/

You have information about n different recipes. You are given a string array recipes and a 2D string array ingredients.
The ith recipe has the name recipes[i], and you can create it if you have all the needed ingredients from ingredients[i].
A recipe can also be an ingredient for other recipes, i.e., ingredients[i] may contain a string that is in recipes.

You are also given a string array supplies containing all the ingredients that you initially have, and you have an infinite supply of all of them.

Return a list of all the recipes that you can create. You may return the answer in any order.

!!! Note that two recipes may contain each other in their ingredients.

Example:
Input: recipes = ["bread","sandwich"], ingredients = [["yeast","flour"],["bread","meat"]], supplies = ["yeast","flour","meat"]  Output: ["bread","sandwich"]
Explanation:
We can create "bread" since we have the ingredients "yeast" and "flour".
We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
*/

//shitty solution but it works 450ms(dfs or bfs would be more faster):
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	mp := make(map[string]bool)
	for _, supply := range supplies {
		mp[supply] = true
	}

	ans := make([]string, 0)
	for i := 0; i < len(ingredients); i++ {
		canDo := true
		if !mp[recipes[i]] {
			for _, ingredient := range ingredients[i] {
				if !mp[ingredient] {
					canDo = false
				}
			}
			if canDo {
				mp[recipes[i]] = true
				i = -1
			}
		}
	}
	for _, recipe := range recipes {
		if mp[recipe] {
			ans = append(ans, recipe)
		}
	}
	return ans
}

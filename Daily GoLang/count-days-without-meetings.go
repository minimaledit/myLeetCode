/*
https://leetcode.com/problems/count-days-without-meetings/description/

You are given a positive integer days representing the total number of days an employee is available for work (starting from day 1).
You are also given a 2D array meetings of size n where, meetings[i] = [start_i, end_i] represents the starting and ending days of meeting i (inclusive).

Return the count of days when the employee is available for work but no meetings are scheduled.
Note: The meetings may overlap.

Example 1:
Input: days = 10, meetings = [[5,7],[1,3],[9,10]]  Output: 2
Explanation:
There is no meeting scheduled on the 4th and 8th days.

Example 2:
Input: days = 5, meetings = [[2,4],[1,3]]  Output: 1
Explanation:
There is no meeting scheduled on the 5th day.
*/

import (
	"sort"
)

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	mergedMeetings := [][]int{meetings[0]}
	for i := 1; i < len(meetings); i++ {
		if mergedMeetings[len(mergedMeetings)-1][1] < meetings[i][0] {
			mergedMeetings = append(mergedMeetings, meetings[i])
		} else {
			mergedMeetings[len(mergedMeetings)-1][1] = max(mergedMeetings[len(mergedMeetings)-1][1], meetings[i][1])
		}
	}

	for _, meet := range mergedMeetings {
		days -= meet[1] - meet[0] + 1
	}

	return days
}

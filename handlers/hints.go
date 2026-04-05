package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"sde-tracker-backend/database"
	"sde-tracker-backend/models"

	"github.com/gin-gonic/gin"
)

func GetHints(c *gin.Context) {
	problemID := c.Param("id")

	var problem models.Problem
	if result := database.DB.First(&problem, problemID); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
		return
	}

	hints := generateHints(problem)
	c.JSON(http.StatusOK, gin.H{
		"problem_id": problem.ID,
		"title":      problem.Title,
		"hints":      hints,
	})
}

func generateHints(p models.Problem) []map[string]string {
	title := strings.ToLower(p.Title)
	diff := p.Difficulty

	hint1 := "Think about the brute force approach first. What is the simplest way to solve this?"
	hint2 := "Consider the time and space complexity. Can you optimize the brute force?"
	hint3 := "Try to code the solution step by step. Handle edge cases."

	// Pattern-based hints
	if strings.Contains(title, "two sum") || strings.Contains(title, "pair") {
		hint1 = "Think about using a Hash Map to store elements you've seen so far."
		hint2 = "For each element, check if (target - element) exists in the map. This gives O(n) time."
		hint3 = "Handle edge cases: duplicate values, empty array, no valid pair exists."
	} else if strings.Contains(title, "sort") || strings.Contains(title, "sorting") {
		hint1 = "What sorting algorithm fits best here? Consider time complexity requirements."
		hint2 = "For O(n) sorting with limited range, think Counting Sort. For general: Merge Sort or Quick Sort."
		hint3 = "Consider in-place sorting to optimize space. Dutch National Flag for 3-way partition."
	} else if strings.Contains(title, "linked list") || strings.Contains(title, "reverse") && strings.Contains(title, "list") {
		hint1 = "Use pointer manipulation. Think about what happens when you change the next pointer."
		hint2 = "Try using two or three pointers (prev, curr, next). Draw it out on paper."
		hint3 = "Handle edge cases: empty list, single node, cycle detection with Floyd's algorithm."
	} else if strings.Contains(title, "binary search") || strings.Contains(title, "search") && strings.Contains(title, "sorted") {
		hint1 = "The array is sorted — Binary Search can find elements in O(log n)."
		hint2 = "Define your search space carefully. What should low and high represent?"
		hint3 = "Be careful with mid calculation (use low + (high-low)/2 to avoid overflow)."
	} else if strings.Contains(title, "tree") || strings.Contains(title, "bst") || strings.Contains(title, "binary tree") {
		hint1 = "Think about tree traversals: Inorder, Preorder, Postorder. Which one helps here?"
		hint2 = "Consider using recursion. What is the base case? What do you return at each node?"
		hint3 = "For BST problems, use the property: left < root < right."
	} else if strings.Contains(title, "graph") || strings.Contains(title, "island") || strings.Contains(title, "cycle") || strings.Contains(title, "path") {
		hint1 = "Think about BFS or DFS traversal. Which one is more suitable for this problem?"
		hint2 = "Use a visited array/set to avoid revisiting nodes. Think about the graph representation."
		hint3 = "For shortest path: BFS (unweighted) or Dijkstra (weighted). For cycle: track visiting state."
	} else if strings.Contains(title, "dynamic") || strings.Contains(title, "dp") || strings.Contains(title, "subsequence") || strings.Contains(title, "knapsack") {
		hint1 = "Can you identify overlapping subproblems? Think about what state you need to track."
		hint2 = "Define dp[i] or dp[i][j]. What does each cell represent? Write the recurrence relation."
		hint3 = "Start with recursive solution + memoization, then convert to tabulation for optimization."
	} else if strings.Contains(title, "stack") || strings.Contains(title, "parenthes") || strings.Contains(title, "bracket") {
		hint1 = "Stack is useful for matching pairs and tracking recent elements. Think LIFO."
		hint2 = "Push opening brackets, pop on closing. For next greater/smaller, use monotonic stack."
		hint3 = "At the end, check if stack is empty (all matched) or process remaining elements."
	} else if strings.Contains(title, "sliding window") || strings.Contains(title, "substring") || strings.Contains(title, "subarray") {
		hint1 = "Think about using the Sliding Window technique with two pointers."
		hint2 = "Expand the window (right pointer) and shrink it (left pointer) based on conditions."
		hint3 = "Use a HashMap or array to track frequency of elements in the current window."
	} else if strings.Contains(title, "matrix") || strings.Contains(title, "2d") {
		hint1 = "Think about how to traverse the matrix efficiently. Row by row? Column by column?"
		hint2 = "For search in sorted matrix: treat it as a 1D sorted array or use staircase search."
		hint3 = "For matrix rotation/spiral: use layer-by-layer approach with boundary variables."
	} else if strings.Contains(title, "trie") || strings.Contains(title, "prefix") {
		hint1 = "Trie is perfect for prefix-based operations. Each node represents a character."
		hint2 = "Insert by creating nodes for each character. Search by following the path."
		hint3 = "For XOR problems, build a bit-trie. For prefix matching, check the isEnd flag."
	} else if strings.Contains(title, "greedy") || strings.Contains(title, "meeting") || strings.Contains(title, "platform") || strings.Contains(title, "job") {
		hint1 = "Greedy approach: make the locally optimal choice at each step."
		hint2 = "Sort the input based on a key criterion (end time, deadline, profit)."
		hint3 = "Prove your greedy choice is correct. Can you find a counter-example?"
	}

	// Adjust complexity hint based on difficulty
	complexityHint := ""
	switch diff {
	case "Easy":
		complexityHint = fmt.Sprintf("Target: O(n) time, O(1) or O(n) space.")
	case "Medium":
		complexityHint = fmt.Sprintf("Target: O(n log n) or O(n) time. Think about optimizing with the right data structure.")
	case "Hard":
		complexityHint = fmt.Sprintf("Target: Optimal solution. Consider advanced techniques like segment trees, binary indexed trees, or complex DP.")
	}

	return []map[string]string{
		{"level": "Approach", "hint": hint1, "icon": "💡"},
		{"level": "Algorithm", "hint": hint2 + " " + complexityHint, "icon": "⚡"},
		{"level": "Implementation", "hint": hint3, "icon": "🔧"},
	}
}

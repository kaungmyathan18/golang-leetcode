package main

import "fmt"

func main() {
	// nums1 := []int{1, 1, 2}
	// (removeDuplicates(nums1))
	// nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// (removeDuplicates(nums2))
	// fmt.Println(nums2)
	// fmt.Println(problems.TotalMoney(4))
	// fmt.Println(problems.TotalMoney(10))
	// fmt.Println(problems.TotalMoney(20))
	// fmt.Println(problems.LargestOddNumber("52"))
	// fmt.Println(problems.LargestOddNumber("4206"))
	// fmt.Println(problems.LargestOddNumber("35427"))
	// fmt.Println(problems.LargestOddNumber("10133890"))
	// fmt.Println(problems.LargestOddNumber("239537672423884969653287101"))
	// fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	// fmt.Println(findWords([]string{"omk"}))
	// fmt.Println(findWords([]string{"adsdf", "sfd"}))
	// fmt.Println(binaryTreePaths(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	// fmt.Println(binaryTreePaths(&TreeNode{Val: 1, Left: nil, Right: nil}))
	// fmt.Println(InorderTraversal(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	// fmt.Println(InorderTraversal(&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	// fmt.Println(InorderTraversal(&TreeNode{Val: 1, Left: nil, Right: nil}))
	// fmt.Println(isSameTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	// fmt.Println(isSameTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: nil}, &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2}}))
	// fmt.Println(IsSymmetric(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	// fmt.Println(IsSymmetric(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}))
	// fmt.Println(ReverseInteger(123))
	// fmt.Println(ReverseInteger(-123))
	// fmt.Println(ReverseInteger(1534236469))
	// fmt.Println(ReverseInteger(-2147483648))
	// fmt.Println(ReverseInteger(120))
	// fmt.Println(StringToInteger("+1")) // 1
	// fmt.Println(StringToInteger("42"))
	// fmt.Println(StringToInteger("   -042"))
	// fmt.Println(StringToInteger("-42")) // -42
	// fmt.Println(StringToInteger("4193 with words"))
	// fmt.Println(StringToInteger("words and 987"))
	// fmt.Println(StringToInteger("-91283472332")) // -2147483648
	// fmt.Println(IntegerToRoman(3))    // III
	// fmt.Println(IntegerToRoman(58))   // LVIII
	// fmt.Println(IntegerToRoman(1994)) // MCMXCIV
	// fmt.Println(IntegerToRoman(3749)) // MMMDCCXLIX
	// fmt.Println(GenerateParentheses(3))
	// fmt.Println(GenerateParentheses(1))
	// fmt.Println(MaxProfit([]int{7, 1, 5, 3, 6, 4}))              // 5
	// fmt.Println(MaxProfit([]int{7, 6, 4, 3, 1}))                 // 0
	// fmt.Println(MaxProfit([]int{1, 2}))                          // 1
	// fmt.Println(MaxProfit([]int{2, 1, 2, 1, 0, 1, 2}))           // 1
	// fmt.Println(MaxProfit([]int{2, 4, 1}))                       // 2
	// fmt.Println(MaxProfit([]int{3, 2, 6, 5, 0, 3}))              // 4
	// fmt.Println(MaxProfit([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // 9
	// fmt.Println(MaxProfit([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})) // 0
	// fmt.Println(MaxProfit([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // 9
	// fmt.Println(minMirrorPairDistance([]int{12, 21, 45, 33, 54})) // 9
	// fmt.Println(minMirrorPairDistance([]int{10, 11, 12, 13}))     // 1
	// fmt.Println(minMirrorPairDistance([]int{120, 210}))           // 90
	// fmt.Println(minMirrorPairDistance([]int{21, 120})) // 90
	// fmt.Println(solveQueries([]int{1, 3, 1, 4, 1, 3, 2}, []int{0, 3, 5})) // [2, -1, 3]
	// fmt.Println(solveQueries([]int{1, 2, 3, 4}, []int{0, 1, 2, 3}))       // [-1,-1,-1,-1]
	// fmt.Println(findFinalValue([]int{5, 3, 6, 1, 12}, 3))     // 24
	// fmt.Println(findFinalValue([]int{2, 7, 9}, 4))            // 4
	// fmt.Println(findFinalValue([]int{8, 19, 4, 2, 15, 3}, 2)) // 16
	// fmt.Println(mirrorDistance(123)) // 123
	// fmt.Println(mirrorDistance(25))  // 123
	// fmt.Println(mirrorDistance(10))  // 123
	// fmt.Println(mirrorDistance(7))   // 123
	// fmt.Println(getMinDistance([]int{1, 2, 3, 4, 5}, 3, 2)) // 1
	// fmt.Println(getMinDistance([]int{1, 2, 3, 4, 5}, 5, 3)) // 1
	// fmt.Println(addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}})) // 7 -> 0 -> 8
	// fmt.Println(addTwoNumbers(&ListNode{Val: 0}, &ListNode{Val: 0}))
	// fmt.Println(addTwoNumbers(&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}, &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}})) // 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1
	// fmt.Println(ClosestTarget([]string{"apple", "orange", "banana", "pear", "pineapple"}, "pear", 2)) // 2
	// fmt.Println(ClosestTarget([]string{"hello", "i", "am", "leetcode", "hello"}, "hello", 1))         // 2
	// fmt.Println(ClosestTarget([]string{"a", "b", "leetcode"}, "leetcode", 0))                         // 2
	// fmt.Println(ClosestTarget([]string{"i", "eat", "leetcode"}, "ate", 0))                            // 2
	// fmt.Println(MinimumDistanceBetweenThreeEqualElements([]int{1, 2, 1, 1, 3}))
	// fmt.Println(MinimumDistanceBetweenThreeEqualElements([]int{1, 1, 2, 3, 2, 1, 2}))
	// fmt.Println(MinimumDistanceBetweenThreeEqualElements([]int{1}))
	// fmt.Println(MinimumHammingDistance([]int{1, 2, 3, 4}, []int{2, 1, 4, 5}, [][]int{{0, 1}, {2, 3}}))
	// fmt.Println(MinOperations("0100"))
	// fmt.Println(MinOperations("10"))
	// fmt.Println(MinOperations("1111"))
	// fmt.Println(SpecialPositionsInABinaryMatrix([][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}))
	// fmt.Println(SpecialPositionsInABinaryMatrix([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
	// fmt.Println(twoEditWords([]string{"word", "note", "ants", "wood"}, []string{"wood", "joke", "moat"}))
	// fmt.Println(twoEditWords([]string{"yes"}, []string{"not"}))
	// fmt.Println(robotReturnToOrigin("UD"))
	// fmt.Println(robotReturnToOrigin("LL"))
	// fmt.Println(robotReturnToOrigin("RRDD"))
	// fmt.Println(robotReturnToOrigin("LDRRLRUULR"))
	// fmt.Println(canBeEqual("abc", "bca"))
	// fmt.Println(canBeEqual("abcd", "cdab"))
	// fmt.Println(canBeEqual("abcd", "dacb"))
	// fmt.Println(distance([]int{1, 3, 1, 1, 2}))
	// fmt.Println(distance([]int{0, 5, 3}))
	// fmt.Println(findRotation([][]int{{0, 1}, {1, 0}}, [][]int{{1, 0}, {0, 1}}))
	// fmt.Println(findRotation([][]int{{0, 1}, {1, 1}}, [][]int{{1, 0}, {0, 1}}))
	// fmt.Println(findRotation([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}))
	// fmt.Println(reverseSquareSubmatrix([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, 1, 0, 3))
	// fmt.Println(reverseSquareSubmatrix([][]int{{3, 4, 2, 3}, {2, 3, 4, 2}}, 0, 2, 2))
	// fmt.Println(furthestDistanceFromOrigin("L_RL__R"))
	// fmt.Println(furthestDistanceFromOrigin("_R__LL_"))
	// fmt.Println(furthestDistanceFromOrigin("_______"))
	// fmt.Println(constructTransformedArray([]int{3, -2, 1, 1}))
	// fmt.Println(constructTransformedArray([]int{-1, 4, -1}))
	// fmt.Println(isTrionic([]int{1, 3, 5, 4, 2, 6}))
	// fmt.Println(isTrionic([]int{2, 1, 3}))
	// fmt.Println(maxDistance(3, [][]int{{1, 2}, {2, 1}, {3, 3}}, 1))
	// fmt.Println(maxDistance(2, [][]int{{0, 2}, {2, 0}, {2, 2}, {0, 0}}, 4))
	// fmt.Println(minimumPairRemoval([]int{1, 2, 3, 4, 5}))
	// fmt.Println(minimumPairRemoval([]int{5, 2, 3, 1}))
	// fmt.Println(minimumPairRemoval([]int{1, 2, 2}))
	// fmt.Println(minBitwiseArray([]int{2, 3, 5, 7}))
	// fmt.Println(minBitwiseArray([]int{11, 13, 31}))
	// fmt.Println(Constructor([]int{1, 2, 3, 4, 5}))
	// fmt.Println(isPowerOfThree(27))
	// fmt.Println(isPowerOfThree(0))
	// fmt.Println(isPowerOfThree(-1))
	// fmt.Println(isPowerOfThree(9))
	// fmt.Println(countBits(2))
	// fmt.Println(countBits(5))
	// fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	// debug(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	// debug(sortedArrayToBST([]int{1, 3}))
	// fmt.Println(minDepth(sortedArrayToBST([]int{-10, -3, 0, 5, 9})))
	// fmt.Println(minDepth(sortedArrayToBST([]int{1, 3})))
	// fmt.Println(hasCycle(&ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}))
	// fmt.Println(hasCycle(&ListNode{Val: 1, Next: &ListNode{Val: 2}}))
	// fmt.Println(minOperations([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 3))
	// fmt.Println(minOperations([][]int{{2, 4}, {6, 8}}, 2))
	// fmt.Println(minOperations([][]int{{1, 5}, {2, 3}}, 1))
	// fmt.Println(missingNumber([]int{3, 0, 1}))
	// fmt.Println(missingNumber([]int{0, 1}))
	// fmt.Println(addStrings("123", "456"))
	// fmt.Println(addStrings("123", "11"))
	// fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	// fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	// fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	// fmt.Println(islandPerimeter([][]int{{1}}))
	// fmt.Println(islandPerimeter([][]int{{1, 0}}))
	// fmt.Print(maxPathScore([][]int{{0, 3, 1, 0}, {3, 4, 2, 5}, {5, 2, 0, 3}, {4, 3, 0, 2}}, 2))
	//fmt.Println(maxPathScore([][]int{{0, 1}, {2, 0}}, 1))
	// fmt.Println(repeatedSubstringPattern("abab"))
	//fmt.Println(findComplement(5))
	//fmt.Println(maxRotateFunction([]int{4, 3, 2, 6}))
	//fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
	//fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
	//fmt.Println(fib(2))
	//fmt.Println(fib(3))
	//fmt.Println(fib(4))
	// fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	//fmt.Println(rotateDigit(10))
	// fmt.Println("\nExample 1:")
	// box1 := [][]byte{{'#', '.', '#'}}
	// result1 := rotateTheBox(box1)
	// printGrid(result1)

	// fmt.Println("\nExample 2:")
	// box2 := [][]byte{
	// 	{'#', '.', '*', '.'},
	// 	{'#', '#', '*', '.'},
	// }
	// result2 := rotateTheBox(box2)
	// printGrid(result2)

	// fmt.Println("\nExample 3:")
	// box3 := [][]byte{
	// 	{'#', '#', '*', '.', '*', '.'},
	// 	{'#', '#', '#', '*', '.', '.'},
	// 	{'#', '#', '#', '.', '#', '.'},
	// }
	// result3 := rotateTheBox(box3)
	// printGrid(result3)
	fmt.Println(isBalanced(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	fmt.Println(isBalanced(&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	fmt.Println(isBalanced(&TreeNode{Val: 1, Left: nil, Right: nil}))
}

// func printGrid(grid [][]byte) {
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[i]); j++ {
// 			fmt.Printf("%c ", grid[i][j])
// 			// fmt.Print(grid[i][j])
// 		}
// 		fmt.Println()
// 	}
// }

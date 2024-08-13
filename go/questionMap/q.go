package q

import (
	"leetCode/go/solutions/easy"
	"leetCode/go/solutions/hard"
	"leetCode/go/solutions/medium"
)

var (
	Q1    = easy.TwoSum                                            // easy
	Q4    = hard.FindMedianSortedArrays                            //. . . . . . . . . hard
	Q5    = medium.LongestPalindrome                               //. . . . medium
	Q30=hard.FindSubstring										   //. . . . . . . . . hard				Substring with Concatenation of All Words
	Q32   = hard.LongestValidParentheses                           //. . . . . . . . . hard
	Q41   = hard.FirstMissingPositive                              //. . . . . . . . . hard				First Missing Positive
	Q42   = hard.Trap                                              //. . . . . . . . . hard				Trapping Rain Water
	Q48   = []func(matrix [][]int){medium.Rotate1, medium.Rotate2} //. . . . medium
	Q51   = hard.SolveNQueens                                      //. . . . . . . . . hard				N-Queens-I
	Q52   = hard.TotalNQueens                                      //. . . . . . . . . hard				N-Queens-II
	Q75   = medium.SortColors                                      //. . . . medium
	Q214  = hard.ShortestPalindrome                                //. . . . . . . . . hard
	Q217  = easy.ContainsDuplicate                                 // easy
	Q268  = easy.MissingNumber                                     // easy
	Q273  = hard.NumberToWords                                     //. . . . . . . . . hard				Integer to English Words
	Q344  = easy.ReverseString                                     // easy
	Q350  = easy.Intersect                                         // easy
	Q409  = easy.LongestPalindrome                                 // easy
	Q648  = medium.ReplaceWords                                    //. . . . medium
	Q726  = hard.CountOfAtoms                                      //. . . . . . . . . hard				Number of Atoms
	Q840  = medium.NumMagicSquaresInside                           //. . . . medium						Magic Squares In Grid
	Q846  = medium.IsNStraightHand                                 //. . . . medium
	Q885  = medium.SpiralMatrixIII                                 //. . . . medium						Spiral Matrix III
	Q912  = medium.SortArray                                       //. . . . medium						Sort an Array
	Q945  = medium.MinIncrementForUnique                           //. . . . medium
	Q1002 = easy.CommonChars                                       // easy
	Q1105 = medium.MinHeightShelves                                //. . . . medium						Filling Bookcase Shelves
	Q1110 = medium.DelNodes                                        //. . . . medium						Delete Nodes And Return Forest
	Q1190 = medium.ReverseParentheses                              //. . . . medium
	Q1122 = easy.RelativeSortArray                                 // easy
	Q1334 = medium.FindTheCity                                     //. . . . medium						Find the City With the Smallest Number of Neighbors at a Threshold Distance
	Q1380 = easy.LuckyNumbers                                      // easy								Lucky Numbers in a Matrix
	Q1395 = medium.NumTeams                                        //. . . . medium						Count Number of Teams
	Q1460 = easy.CanBeEqual                                        // easy								Make Two Arrays Equal by Reversing Sub-arrays
	Q1508 = medium.RangeSum                                        //. . . . medium						Range Sum of Sorted Subarray Sums
	Q1509 = medium.MinDifference                                   //. . . . medium
	Q1518 = easy.NumWaterBottles                                   // easy
	Q1530 = medium.CountPairs                                      //. . . . medium						Number of Good Leaf Nodes Pairs
	Q1550 = easy.ThreeConsecutiveOdds                              // easy
	Q1598 = easy.MinOperations                                     // easy
	Q1605 = medium.RestoreMatrix                                   //. . . . medium						Find Valid Matrix Given Row and Column Sums
	Q1636 = easy.FrequencySort                                     // easy								Sort Array by Increasing Frequency
	Q1653 = medium.MinimumDeletions                                //. . . . medium						Minimum Deletions to Make String Balanced
	Q1701 = medium.AverageWaitingTime                              //. . . . medium
	Q1717 = medium.MaximumGain                                     //. . . . medium
	Q1823 = medium.FindTheWinner                                   //. . . . medium
	Q2045 = hard.SecondMinimum                                     //. . . . . . . . . hard				Second Minimum Time to Reach Destination
	Q2055 = easy.KthDistinct                                       // easy								Kth Distinct String in an Array
	Q2058 = medium.NodesBetweenCriticalPoints                      //. . . . medium
	Q2096 = medium.GetDirections                                   //. . . . medium						Step-By-Step Directions From a Binary Tree Node to Another
	Q2134 = medium.MinSwapsII                                      //. . . . medium					    Minimum Swaps to Group All 1's Together II
	Q2181 = medium.MergeNodes                                      //. . . . medium
	Q2191 = medium.SortJumbled                                     //. . . . medium						Sort the Jumbled Numbers
	Q2196 = medium.CreateBinaryTree                                //. . . . medium						Create Binary Tree From Descriptions
	Q2418 = easy.SortPeople                                        // easy								Sort the People
	Q2486 = medium.AppendCharacters                                //. . . . medium
	Q2544 = easy.AlternateDigitSum                                 // easy								Alternate Digit Sum
	Q2582 = easy.PassThePillow                                     // easy
	Q2678 = easy.CountSeniors                                      // easy							    Number of Senior Citizens
	Q2976 = medium.MinimumCost                                     //. . . . medium						Minimum Cost to Convert String I
	Q3014 = medium.MinimumPushes                                   //. . . . easy						Minimum Number of Pushes to Type Word I (Same as 3016)
	Q3016 = medium.MinimumPushes                                   //. . . . medium						Minimum Number of Pushes to Type Word II (Same as 3014)
	Q3110 = easy.ScoreOfString                                     // easy
	Q3168 = easy.MinimumChairs                                     // easy
	Q3169 = medium.CountDays                                       //. . . . medium
	Q3170 = medium.ClearStars                                      //. . . . medium
	Q3174 = easy.ClearDigits                                       // easy
	Q3175 = medium.FindWinningPlayer                               //. . . . medium
	Q3178 = easy.NumberOfChild                                     // easy
	Q3179 = medium.ValueAfterKSeconds                              //. . . . medium
	Q3180 = medium.MaxTotalReward                                  //. . . . medium
	Q3184 = easy.CountCompleteDayPairs_I                           //. . . . medium
	Q3185 = medium.CountCompleteDayPairs_II                        //. . . . medium
	Q3190 = easy.MinimumOperations                                 // easy
	Q3191 = medium.MinOperations_I                                 //. . . . medium
	Q3192 = medium.MinOperations_II                                //. . . . medium
	Q3194 = easy.MinimumAverage                                    // easy
	Q3195 = medium.MinimumArea                                     //. . . . medium
	Q3196 = easy.GetSmallestString                                 // easy								Lexicographically Smallest String After a Swap
	Q3217 = medium.ModifiedList                                    //. . . . medium						Delete Nodes From Linked List Present in Array
	Q3232 = easy.CanAliceWin                                       // easy								Find if Digit Game Can Be Won
)

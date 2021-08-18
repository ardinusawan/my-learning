"""
Leetcode: https://leetcode.com/problems/product-of-array-except-self/

https://app.sparkmailapp.com/web-share/43IgCvF1HFIRCUW3NsLwvXVCWupZAxspvwFgGCey

Good morning! Here's your coding interview problem for today.

This problem was asked by Uber.

Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
"""
class Solution:
    def productExceptSelf(self, nums):
        result = []
        for i, n in enumerate(nums):
            if i == 0:
                result.append(n)
            else:
                value = result[i-1]*n
                result.append(value)

        product = 1
        for i in range(len(nums), 0, -1):
            if i == len(nums):
                result[i-1] = result[i-2]
                product = nums[i-1]
            elif i == 1:
                result[0] = product
            else:
                result[i-1] = result[i-2]*product
                product=product*nums[i-1]

        return result

s = Solution()
print(s.productExceptSelf([1,2,3,4]))

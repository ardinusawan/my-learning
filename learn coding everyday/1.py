"""
https://app.sparkmailapp.com/web-share/vw59_6rp5LWOVKJGFJIiLf_QI9QhDb66t8EIfIb5

Good morning! Here's your coding interview problem for today.

This problem was recently asked by Google.

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
"""

def solution(nums, k):
    s = set()
    for n in nums:
        if k-n in s:
            print("true")
        s.add(n)

nums = [10, 15, 3, 7]
k = 17
solution(nums, k)

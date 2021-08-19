"""
https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

For example, given the following Node class

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
The following test should pass:

node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'
"""
class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def serialize(root):
    """Encodes a tree to a single string.

    :type root: TreeNode
    :rtype: str
    """
    result = []
    def dfs(node):
        if not node:
            result.append("N")
            return
        result.append(node.val)
        dfs(node.left)
        dfs(node.right)
    dfs(root)

    return ",".join(result)

i = 0
def deserialize(data):
    """Decodes your encoded data to tree.

    :type data: str
    :rtype: TreeNode
    """
    vals = data.split(",")
    def dfs():
        global i
        if vals[i] == 'N':
            i+=1
            return
        n = Node(vals[i])
        i+=1
        n.left = dfs()
        n.right = dfs()
        return n

    return dfs()

node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'

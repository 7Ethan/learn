
class Node(object):
    def  __init__(self, data, left=None, right=None):
        self.data = data
        self.left = left
        self.right = right

tree = Node(1, Node(3, Node(7, Node(0)), Node(6)), Node(2, Node(5), Node(4)))

def lookup(root):
    row = [root]
    while row:
        print(row)
        row = [kid for item in row for kid in (item.left, item.right) if kid]

def deep(root):
    if not root:
        return
    print(root.data)
    deep(root.left)
    deep(root.right)


if  __name__ == '__main__':
    lookup(tree)
    deep(tree)
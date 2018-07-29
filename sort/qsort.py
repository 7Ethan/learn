def qsort(arr):
    if len(arr) < 1:
        return arr
    mark = arr[0]
    mid = []
    left = []
    right = []

    for n in arr:
        if n == mark:
            mid.append(n)
        if n > mark:
            right.append(n)
        if n < mark:
            left.append(n)
    left = qsort(left)
    right = qsort(right)

    return left + mid + right

arr = [345, 89, -456, 23, 0, 45, 654, 23, 7, 52, 769]
print(qsort(arr))
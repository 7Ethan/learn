import random

def QuickSort(arr,start,end):
    if start < end:
        key = arr[random.randint(0,len(arr)-1)]
        i,j = start,end

        while arr[i] < key:
            i += 1
        while arr[j] > key:
            j -= 1
        if i <= j:
            arr[i],arr[j] = arr[j],arr[i]
            i += 1
            j -= 1
        if start < j:
            QuickSort(arr,start,j)
        if end > i:
            QuickSort(arr,i,end)

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
QuickSort(arr,0,len(arr)-1)
print(arr)
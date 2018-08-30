def shift_down(array,start,end):
    """
    调整成大顶堆，初始堆时，从下往上；交换堆顶与堆尾后，从上往下调整
    :param array: 列表的引用
    :param start: 父结点
    :param end: 结束的下标
    :return : none
    """
    while True:
        left_child = start * 2 + 1
        if left_child > end:
            break

        if left_child + 1 <= end and array[left_child] < array[left_child + 1]:
            left_child += 1
        if array[left_child] > array[start]:
            array[left_child],array[start] = array[start],array[left_child]
            start = left_child # 交换之后以交换子结点为根的堆可能不是大顶堆，需重新调整
        else:    # 若父结点大于左右孩子，则退出循环
            break
        
        # print("->",array)

def heap_sort(array):
    first = len(array) // 2 -1
    for i in range(first,-1,-1):
        print(array[i])
        shift_down(array,i,len(array) - 1)

    print("初始化大顶堆结果:", array)

    for head_end in range(len(array) -1 ,0,-1):
        array[head_end],array[0] = array[0],array[head_end]
        shift_down(array,0,head_end - 1)

if __name__ == "__main__":
    array = [16, 7, 3, 20, 17, 8,34, 678, -57, 768, 0, 67, 895, 24, 789, 234]
    print(array)
    heap_sort(array)
    print(array)
    
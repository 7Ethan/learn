package main

import (
	"fmt"
)

func main() {
	arr := []int{34, 678, -57, 768, 0, 67, 895, 24, 789, 234, 6, 348, 357}

	HeapSort(arr)
	fmt.Println(arr)
}

func HeapSort(arr []int) {
	//1.构建大顶堆
	//默认数组已经是一颗完全二叉树.
	for i := len(arr)/2 - 1; i >= 0; i-- { // arr[len(arr)/2 - 1] 表示第一个非叶子节点
		//从第一个非叶子结点从下至上，从右至左调整结构
		adjustHeap(arr, i, len(arr))
	}

	//2.调整堆结构+交换堆顶元素与末尾元素
	//真正排序的实现在这里,数组从后往前填充.
	/**
	<--------------
	| | | | | | MAX|
	*/
	for j := len(arr) - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0] //将堆顶元素与末尾元素进行交换,这一步将已经确定了的最大(最小)放到了数组的末尾
		adjustHeap(arr, 0, j)           //重新对堆进行调整,n - 1 个元素保持为最大(最小)堆
	}
}

//sink
func adjustHeap(arr []int, i, length int) {
	temp := arr[i]                              //先取出当前元素i
	for k := i*2 + 1; k < length; k = k*2 + 1 { //从i结点的左子结点开始，也就是2i+1处开始
		if k+1 < length && arr[k] < arr[k+1] { //如果左子结点小于右子结点，k指向右子结点
			k++
		}
		if arr[k] > temp { //如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
			arr[i] = arr[k]

			i = k
		} else {
			break
		}
	}
	arr[i] = temp

	/**
	* 每次调整都是从父节点、左孩子节点、右孩子节点三者中选择最大者跟父节点进行交换，
	* 交换之后都可能造成被交换的孩子节点不满足堆的性质，
	* 因此每次交换之后要重新对被交换的孩子节点进行调整。
	*
	 */
}

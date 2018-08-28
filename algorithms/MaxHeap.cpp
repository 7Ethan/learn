#include <iostream>
#include <ctime>
#include <algorithm>
#include <string>
#include <cmath>
#include <cassert>

 using namespace std;

 template <typename item>
 class MaxHeap{
     private:
     item* data;
     int count;
     int capacity;

     void siftUp(int);
     void siftDown(int);

    public:
    MaxHeap(int capacity);
    //使用已经存在的自定义数组构造堆
    MaxHeap(item arr[],int n);

    ~MaxHeap(){
        delete[] data;
        cout << "A heap object has been delete " << endl;
    }

    int size();
    bool isEmpty();
    void insert(item);
    item extractMax();
    item getMax();
 };

template <typename item>
MaxHeap<item>::MaxHeap(int capacity){
    data = new item[capacity + 1];
    this->capacity = capacity;
    count = 0;
}

//队列索引从 1 开始
template <typename item>
MaxHeap<item>::MaxHeap(item arr[],int n){
    data = new item[n];
    this->count = n;
    this->capacity = n + 1;
    for (int i = 0;i < n;i++){
        data[i + 1] = arr[i];
    }

    for(int i = count / 2;i >= 1;i--){
        siftDown(i);
    }
}

template <typename item>
void MaxHeap<item>::siftUp(int k){
    while (k > 1 && data[k/2] < data[k]) {
            swap( data[k/2], data[k] );
            k /= 2;
        }
}

template <typename item>
void MaxHeap<item>::siftDown(int k){
    while (k * 2 < count) {
        int j = 2 * k;
        if (j + 1 <= count && data[j] < data[j + 1]){
            j += 1;
        }
        if (data[k] > data[j]){
            break;
        }
        swap(data[k],data[j]);
        k = j;
    }
}

template <typename item>
int MaxHeap<item>::size(){
    return count;
}

template <typename item>
bool MaxHeap<item>::isEmpty(){
    return count == 0;
}

template <typename item>
void MaxHeap<item>::insert(item i){
    assert(count + 1 <= capacity);
    data[count + 1] = i;
    count++;
    siftUp(count);
}

template <typename item>
item MaxHeap<item>::extractMax(){
    assert(count > 0);
    item ret = data[1];
    swap(data[1],data[count]);
    count--;

    siftDown(1);
    return ret;
}

template <typename item>
item MaxHeap<item>::getMax(){
    assert(count > 0);
    return data[1];
}

int main(int argc, char const *argv[])
{
    /* code */
    MaxHeap<int> *h = new MaxHeap<int>(50);
    int N = 50;
    int M = 100;
    //随机数填充数据.
    for (int i = 0; i < N;i++){
        h->insert(i);
    }

    int *arr = new int[N];
    // 将maxHeap中的数据使用extractMax()取出来
        // 取出来的顺序应该是从大到小排序的
    for (int i = 0;i < N;i++){
        arr[i] = h->extractMax();
        cout << arr[i] << " " ;
    }
    return 0;

    // 确保arr数组是从大到小排序的
    for (int i = 1;i < N;i++){
        assert( arr[i -1] >= arr[i]);
    }
}

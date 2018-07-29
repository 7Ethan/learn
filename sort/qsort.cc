#include <iostream>

using namespace std;

class QSORT
{
    public:
    int partition(int a[],int low,int high);
    void qsort(int arr[],int start,int end);
    void print(int arr[],int length);

};

int QSORT::partition(int arr[],int low,int high)
{
    int pivot = arr[low];

    while( low < high)
    {
        while(low < high && arr[high] > pivot)
        {
            high--;
        }
        arr[low] = arr[high];

        while(low < high && arr[low] < pivot)
        {
            low++;
        }
        arr[high] = arr[low];
    }

    arr[low] = pivot;
    return low;
}

void QSORT::qsort(int arr[],int start,int end)
{
    if( start < end)
    {
        int pivotPos = partition(arr,start,end);

        qsort(arr,start,pivotPos-1);
        qsort(arr,pivotPos+1,end);
    }

}

void QSORT::print(int arr[],int length)
{
    for(int i = 0; i < length;i++)
    {
        cout << arr[i] << endl;
    }
    printf("\n");
}



int main()
{
    int arr[] = {70,75,69,32,88,4,56,-235,6,678,18,16,58,-45};

    QSORT q;
    q.print(arr,14);
    printf("============\n");
    q.qsort(arr,0,14);
    q.print(arr,14);
    return 0;
}
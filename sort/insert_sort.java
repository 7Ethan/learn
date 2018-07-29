import java.util.Arrays;

public class insert_sort{

      /**
     * 交换数组元素
     * @param arr
     * @param a
     * @param b
     */
    public static void swap(int []arr,int a,int b){
        arr[a] = arr[a]+arr[b];
        arr[b] = arr[a]-arr[b];
        arr[a] = arr[a]-arr[b];
    }

    public static void insert_sort(int[] arr){
        for(int i = 1;i < arr.length;i++){
            int j = 0;
            while(j < i && arr[j] > arr[j + 1]){
                swap(arr,j,j + 1);
            }
        }
    }

    public static void main(String args[])
    {
        int array[] = {23,65,8,56,0,67,2,6,-67,34,5,546,34};
        System.out.println(Arrays.toString(array));
        insert_sort(array);
        System.out.println(Arrays.toString(array));
    }
}
import java.util.Random;
import java.util.Arrays;

public class QuickSort {
    public static void main(String[] args) {
        int []arr = {34,678,-57,768,0,67,895,24,789,234,6,348,357};
        qsort(arr,0,arr.length - 1);
        System.out.println(Arrays.toString(arr));
    }

    public static void qsort(int []arr,int start,int end){
        if (start < end){
            int i = start;
            int j = end;
            Random rand = new Random();
            int length = arr.length;
            int key = arr[rand.nextInt(length)];

            while(arr[i] < key){
                i++;
            }

            while(arr[j] > key){
                j--;
            }

            if (i <= j){
                swap(arr,i,j);
                i++;
                j--;
            }

            if (start < j){
                qsort(arr,start,j);
            }
            if (end > i){
                qsort(arr,i,end);
            }
        }
    }

    private static void swap(int []arr,int i,int j){
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }
}
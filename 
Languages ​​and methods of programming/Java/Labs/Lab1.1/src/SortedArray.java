import java.util.ArrayList;
import java.util.Arrays;
import java.util.stream.IntStream;

interface SortArrayInterface{
    public void add(int element);
    public int get(int index);
    public void makeset();
}

public class SortedArray implements SortArrayInterface{
    private int[] arraydata;
    private int size;
    private int count;

    public SortedArray(int size){
        this.arraydata=new int[size+1];
        this.size = size;
        this.count=0;
    }//инициализация без начальных значений

    public SortedArray(int[] arr){
        this.arraydata=new int[arr.length+1];
        for(int i = 1;i<arr.length+1;i++){
            this.arraydata[i]=arr[i-1];
        }
        this.arraydata[0]=0;
        this.size=arr.length;
        this.count=arr.length;
        doSort(size-count,size);
    }//инициализация с начальными значениями

    public void add(int element){
        if (count>=size){
            throw new  ArrayIndexOutOfBoundsException("Массив переполнен");
        }
        else {

            arraydata[size-count] = element;
            doSort(size-count,size);
            count++;
        }
    }//Добавление элемента в массив

    public int get(int index){
        int result,buf;
        result = arraydata[index+1];
        return result;
    }//Достать index  элемент из массива

    public void makeset(ArrayList<Integer> a){
        Object[] arr = a.toArray();
        int index = 1;
        for(index=this.size-count;index<this.size;index++){
            if (this.arraydata[index]==this.arraydata[index+1]){
                this.count--;
                this.arraydata[index]=0;
                for(int i=index;i>=1;i--){
                    this.arraydata[i]=this.arraydata[i-1];
                }
            }
        }
    }//Удаление повторяющихся элементов

    public void quickSort() {
        int startIndex = 0;
        int endIndex = this.size;
        doSort(startIndex, endIndex);
    }

    private void doSort(int start, int end) {
        if (start >= end)
            return;
        int i = start, j = end;
        int cur = i - (i - j) / 2;
        while (i < j) {
            while (i < cur && (this.arraydata[i] <= this.arraydata[cur])) {
                i++;
            }
            while (j > cur && (this.arraydata[cur] <= this.arraydata[j])) {
                j--;
            }
            if (i < j) {
                int temp = this.arraydata[i];
                this.arraydata[i] = this.arraydata[j];
                this.arraydata[j] = temp;
                if (i == cur) {
                    cur = j;
                }
                else if (j == cur) {
                    cur = i;
                }
            }
        }
        doSort(start, cur);
        doSort(cur+1, end);
    }//Сортировка

    @Override
    public String toString(){
        String s = "[";
        for(int i=this.size-this.count+1;i<this.size;i++){
            s+=this.arraydata[i];
            s+=',';
        }
        s+=this.arraydata[this.size];
        s+="]";
        return s;
    }//Вывод всех добавленных элементов
}

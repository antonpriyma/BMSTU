import java.util.ArrayList;
import java.util.Comparator;

interface SequenceInterface<T>{
    void addElement(T element);
     void delete(int Number);
     void remove(T element);
     void sort();
     int characteristic();
}

public class Sequence<T extends Number>implements Comparable<Sequence>,SequenceInterface<T>{
    private ArrayList<T> data;

    public Sequence(ArrayList<T> arr){
        data=arr;
    }// Иницилизация с помощью arraylist

    public  Sequence(T[] arr) {
        for (int i = 0; i < arr.length; i++) {
            data.add(arr[i]);
        }
    }//Иницилизация с помощь массива

    public  Sequence(){
        data = new ArrayList<>();
    }//Инициализция без значени

    public void addElement(T element){
        data.add(element);
    }//Добавления элемента

    public void delete(int Number){
        data.remove(Number);
    }//Удаление эдемента по индексу

    public void remove(T element){
        while(data.contains(element)) {
            data.remove(element);
        }
    }//Удалеие всех element

    public void sort(){
        data.sort(new Comparator<T>() {
            @Override
            public int compare(T o1, T o2) {
                if (o1.doubleValue()>o2.doubleValue()) return 1;
                if (o1.doubleValue()<o2.doubleValue()) return -1;
                return 0;
            }
        });
    }//Сортировка последовательности

    public int characteristic(){
        ArrayList<T> help = new ArrayList<>();
        for (int i=0;i<data.size();i++){
            if (!help.contains(data.get(i))) {
                help.add(data.get(i));
            }
        }
        return help.size();
    }//Количество различных элементов в последовательности


    @Override
    public String toString(){
        String s = data.toString();
        return s;
    }

    public int compareTo(Sequence obj){
        int ch1=this.characteristic(),ch2=obj.characteristic();
        if (ch1-ch2>0) return 1;
        if (ch1-ch2<0) return -1;
        return 0;
    }
}

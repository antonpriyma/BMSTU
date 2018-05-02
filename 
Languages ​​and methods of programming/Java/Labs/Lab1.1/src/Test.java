public class Test {
    public static void main(String[] argc){
        SortedArray sample = new SortedArray(100);
        for (int i=4;i>0;i--)
            sample.add(i);
        sample.add(4);
        sample.add(-200);

    /** Проверка основных функций, сортировки, удаления повторений*/
        System.out.println(sample);
        sample.makeset();
        System.out.println(sample);
        sample.add(-100);
        System.out.println(sample);
        sample.add(500);
        System.out.println(sample);


    /** Проверка переполнения массива*/
        SortedArray sample1 = new SortedArray(1);
        sample1.add(5);
        System.out.println(sample1);
        try {
            sample1.add(6);
        }
        catch (ArrayIndexOutOfBoundsException e){
            System.out.println("Out of bounds error");
        }
        System.out.println(sample1);
    /** Проверка обращения за пределы массива*/
        try {
            sample1.get(100);
        } catch (ArrayIndexOutOfBoundsException e){
            System.out.println("Out of bounds error");
        }

    /** Проверка 2 конструктора */
        int[] numbers = {1,2,1,1,3,2,2,1};
        SortedArray sample3 = new SortedArray(numbers);
        System.out.println(sample3);
        sample3.makeset();
        System.out.println(sample3);
    }
}

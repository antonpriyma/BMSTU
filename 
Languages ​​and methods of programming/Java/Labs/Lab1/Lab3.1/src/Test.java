import java.util.Arrays;

public class Test {

    public static void print(Sequence[] q){
        for (int i=0;i<q.length;i++){
            System.out.println(q[i]);
        }
    }
    public static void main(String[] argc){
        /** Проверка ввода, создания элемента*/
        Sequence p = new Sequence();
        p.addElement(3);
        p.addElement(4);
        p.addElement(5);
        p.addElement(3);

        System.out.println("Sequence p: "+p);
        System.out.println("P characteristic: " + p.characteristic());
        p.sort();
        System.out.println("Sort sequence p: "+p);

        Sequence g = new Sequence();

        g.addElement(3);
        g.addElement(3);
        g.addElement(3);
        /** Проверка сортировки*/
        Sequence q[] = new Sequence[2];
        q[0]=p;
        q[1]=g;
        System.out.print("Array of sequences: ");
        Test.print(q);
        Arrays.sort(q);
        System.out.print("Array of sequences after sort: ");
        Test.print(q);

        g.remove(3);
        System.out.println("g sequence after remove: "+g);
        p.delete(2);
        System.out.println("p sequence after delete: "+p);

        /** Все исключения в arraylist*/
    }
}

import LinkedList.LinkedList;

public class Test {
    public static void main(String[] argc){
        LinkedList sample = new LinkedList<String>();


        sample.add("Hello");
        sample.add("World");


        for(Object a: sample)
            System.out.println(a);
    }
}

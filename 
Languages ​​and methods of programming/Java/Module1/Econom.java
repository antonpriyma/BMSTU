import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;
import java.util.regex.Pattern;

public class Econom{
    private final char Operation1 = '@';
    private final char Operation2 = '#';
    private final char Operation3 = '$';
    private final char Openbracket = '(';
    private static final char Closebracket = ')';
    private static String baseExpresion = new String();
    private Map<Character, StringBuilder> hashtable = new HashMap<Character, StringBuilder>();

    public Econom() {
        Scanner in = new Scanner(System.in);
        baseExpresion = in.nextLine();
    }//Ввод обрабатываемой строки

    public static int pars() {
        int count = 0;
        int len = baseExpresion.length();
        ArrayList<StringBuilder> buffer = new ArrayList<>();
        Character operation = new Character('n');

        for (; baseExpresion.length() > 1;) {

            /*
            switch (c){
                case Openbracket:
                    buffer.append(c);
                    break;
                case Closebracket:
                    if (hashtable.containsValue(buffer)==true){
                        break;
                    }
                    else{
                        hashtable.put(operation,buffer);
                        break;
                        count++;
                    }
                case
                */
            int index1,index = baseExpresion.indexOf(Econom.Closebracket);
            String buf = "";

            buf=baseExpresion.substring(index-4,index+1);
            baseExpresion = baseExpresion.replace(buf, " ");
            count++;

        }
        return count;
    }




    public static void main(String[] argc) {
        new Econom();
        System.out.println(pars());
    }
}


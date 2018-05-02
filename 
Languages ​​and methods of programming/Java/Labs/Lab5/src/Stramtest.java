import java.util.ArrayList;
import java.util.*;

public class Stramtest {


    public static ArrayList<Optional<Integer>> max(ArrayList<ArrayList<Integer>> c){
        ArrayList<Optional<Integer>> maxValues = new ArrayList<>();//Список max элементов
        c.stream().forEach((i)->maxValues.add(i.stream().max(Integer::compareTo)));//Добавление элемента
        return maxValues;
    }


}

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;


public class test {
    public static void main(String[] argc){
        ArrayList<Double> samplecoor = new ArrayList<>();
        samplecoor.add(1.0);
        samplecoor.add(2.0);
        samplecoor.add(3.0);
        Vectors sample = new Vectors(samplecoor);
        samplecoor=new ArrayList<>();
        samplecoor.add(2.0);
        samplecoor.add(3.0);
        samplecoor.add(1.0);
        sample.add(samplecoor);
        samplecoor=new ArrayList<>();
        samplecoor.add(10.0);
        samplecoor.add(3.0);
        samplecoor.add(1.0);
        sample.add(samplecoor);
        samplecoor=new ArrayList<>();
        samplecoor.add(6.0);
        samplecoor.add(9.0);
        samplecoor.add(3.0);
        sample.add(samplecoor);
        samplecoor=new ArrayList<>();
        samplecoor.add(6.0);
        samplecoor.add(9.0);
        samplecoor.add(3.0);
        samplecoor.add(3.0);//проверка, что такой угол найден не будет
        sample.add(samplecoor);
        samplecoor=new ArrayList<>();
        samplecoor.add(0.4);
        samplecoor.add(0.6);
        samplecoor.add(0.2);
        sample.add(samplecoor);
        samplecoor=new ArrayList<>();
        samplecoor.add(4.0);
        samplecoor.add(6.0);
        samplecoor.add(2.0);
        List<Vector> isLz = sample.findCollinear(samplecoor,0.1).collect(Collectors.toList());
        Map<Boolean, List<Vector>> result = isLz.stream().collect(Collectors.partitioningBy(v->v.length()<=1));
        System.out.println(result);
        System.out.println(sample.maxAngle().get());
    }
}

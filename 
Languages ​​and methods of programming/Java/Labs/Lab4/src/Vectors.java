import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.stream.Collectors;
import java.util.stream.Stream;


class Vector {
    private int dim;
    private List<Double> coor;

    public Vector(ArrayList<Double> coordinates){
        coor=coordinates;
        dim = coordinates.size();
    }

    public boolean isCollinear(Vector vec, double e) {
        if (vec.dim!=this.dim) return false;
        double attitudes = coor.get(0) / vec.coor.get(0);
        List<Double> buf=coor.stream().filter(num -> (Math.abs(num / vec.coor.get(coor.indexOf(num)) - attitudes) <= e)).collect(Collectors.toList());
        if (buf.size() == dim) return true;
        return false;
    }

    public Double length(){
        Double buf= new Double(coor.stream().reduce((s1,s2)->s1+s2*s2).orElse(0.0))-coor.get(0)+coor.get(0)*coor.get(0);
        return Math.sqrt(buf);
    }

    public Double angle(Vector vec){//вычисление угла
        if (this.dim!=vec.dim) return new Double(0);
        List<Double> scal =coor.stream().map((s1)->s1*vec.coor.get(coor.indexOf(s1))).collect(Collectors.toList());
        Double sum = scal.stream().reduce((s1,s2)->s1+s2).orElse(0.0);
        return Math.acos(sum.floatValue()/(this.length().floatValue()*vec.length().floatValue()));
    }

    @Override
    public String toString(){
        return coor.toString();
    }
}


public class Vectors {
    private int dim;
    private ArrayList<Vector> coor;

    public Vectors(ArrayList<Vector> arr, int size) {
        coor = arr;
        dim = size;
    }

    public Vectors(ArrayList<Double> arr) {
        coor = new ArrayList<>();
        coor.add(new Vector(arr));
        dim = 1;
    }

    public void add(ArrayList<Double> vec){
        dim++;
        coor.add(new Vector(vec));
    }

    public Stream<Vector> findCollinear(Vector vec, int E){
        return coor.stream().filter(value->value.isCollinear(vec,E));

    }

    public Stream<Vector> findCollinear(ArrayList<Double> arr, double E){
        Vector vec=new Vector(arr);
        return coor.stream().filter(value->value.isCollinear(vec,E));

    }

    public Optional<Double> maxAngle(){
        List<Double> anglelist = coor.stream()
                .filter(s->coor.indexOf(s)<coor.size()-1).map((s)->s.angle(coor.get(coor.indexOf(s)+1)))
                .collect(Collectors.toList());
        return anglelist.stream().max(Double::compareTo);
    }

    public Map<Boolean,List<Vector>> groupByLength(){
        Map<Boolean,List<Vector>> result = this.coor.stream().collect(Collectors.partitioningBy(v->v.length()<1));
        return result;
    }

}

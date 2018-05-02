import java.util.ArrayList;

public class Element<T> {
    private T value;
    private int rank;
    private Element<T> father =this;

    public Element(T value1){
        value=value1;
        rank=0;
    }

    public T x(){
        //System.out.println(this.value);
        return value;
    }

    public Element<T> finddather(){
        if (this.father==this) return this;

            return this.father = this.father.finddather();

    }

    public Boolean equivalent(Element<T> value1){
        return (this.finddather()==value1.finddather());
    }

    public void union(Element<T> b) {
        Element<T> help = this.finddather();
        Element<T> help1 = b.finddather();
        if (help != help1) {
            if (help.rank < help1.rank) {
                Element<T> buf = help;
                help = help1;
                help1 = buf;
            }
            help1.father = help;
            if (help.rank == help1.rank)
                ++help.rank;
        }
        }
    }





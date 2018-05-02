/*import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class vectorsTest{
    private ArrayList<Double> samplecoor;
    private vectors sample;
    @org.junit.jupiter.api.BeforeEach
    void setUp() {
        samplecoor = new ArrayList<>();
        samplecoor.add(1.0);
        samplecoor.add(2.0);
        samplecoor.add(3.0);
        sample = new vectors(samplecoor);
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
        samplecoor.add(4.0);
        samplecoor.add(6.0);
        samplecoor.add(2.0);
    }

    @Test
    void testAll(){
        List<vector> isLz = sample.findCollinear(samplecoor,0).collect(Collectors.toList());
        Map<Boolean, List<vector>> result = isLz.stream().collect(Collectors.partitioningBy(v->v.length()<=1));
        String stringValue = result.toString();
        assertEq
    }

    @org.junit.jupiter.api.AfterEach
    void tearDown() {
    }



}*/

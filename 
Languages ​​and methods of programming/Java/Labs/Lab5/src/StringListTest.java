import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import static junit.framework.TestCase.assertEquals;
import static junit.framework.TestCase.assertFalse;
import static junit.framework.TestCase.assertTrue;

public class StringListTest{
    private StringList sample;

    @Before
    public void init(){
        sample=new StringList();
        sample.add("addd");
        sample.add("dddk");
    }

    @After
    public void tearDown(){
        sample=null;
    }

    @Test
    public void Test1(){
        String result = sample.toString();
        assertEquals(result,"ddd dddk");
        assertEq
    }

    @Test
    public void Test2(){
        sample.clear();
        sample.add("addd");
        sample.add("ddd");
        sample.add("irprjf");
        sample.add("btryew");
        sample.add("oorpol");
        sample.add("bbgjtb");
        sample.add("jmyzdx");
        sample.add("xcbtvtdd");
        sample.add("ddaaaaaaaaaaaaaaaaaaaaaaaaaaa");
        String s = sample.toString();
        assertEquals(sample.toString(),"ddd x dd ddaaaaaaaaaaaaaaaaaaaaaaaaaaa");
    }

    @Test
    public void TestContains(){
        assertTrue(sample.contains("ddd"));
        assertFalse(sample.contains("l"));
    }

}

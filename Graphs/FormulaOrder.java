import javax.swing.text.html.HTML;
import java.util.*;
import java.util.function.IntPredicate;



public class FormulaOrder {

    class Position {
        private String text;
        private int index, line, col;

        public Position(String text) {
            this(text, 0, 1, 1);
        }

        private Position(String text, int index, int line, int col) {
            this.text = text;
            this.index = index;
            this.line = line;
            this.col = col;
        }

        public int getChar() {
            return index < text.length() ? text.codePointAt(index) : -1;
        }

        public boolean satisfies(IntPredicate p) {
            return p.test(getChar());
        }

        public Position skip() {
            int c = getChar();
            switch (c) {
                case -1:
                    return this;
                case '\n':
                    return new Position(text, index+1, line+1, 1);
                default:
                    return new Position(text, index + (c > 0xFFFF ? 2 : 1), line, col+1);
            }
        }

        public Position skipWhile(IntPredicate p) {
            Position pos = this;
            while (pos.satisfies(p)) pos = pos.skip();
            return pos;
        }

        public String toString() {
            return String.format("(%d, %d)", line, col);
        }
    }

    enum Tag{
        PLUS,MINUS,MULT,SUB,IDENT, NUMBER, LPAREN, RPAREN, END, COMMA,EQUAL;
        public String toString() {
            switch (this) {
                case IDENT: return "identifier";
                case NUMBER: return "number";
                //case STRING: return "string";
                case LPAREN: return "'('";
                case RPAREN: return "')'";
                case END: return "end of text";
            }
            throw new RuntimeException("unreachable code");
        }
    }


    static class Token {
        private  ArrayList<Tag> tags =new ArrayList<>();
        private  boolean isRight = false;
        private  String line;
        private  String newIdent=new String() ;
        private int count = 0;
        private  int count_equal = 0;
        private  int count_var = 0;
        private int count_comma=0;

        public Token(String text) throws Exception {
            this.line=text;
        }

        public void find(int k, HashMap<String, Integer> a) throws Exception {
            String s = line;
            for (int i=0;i<s.length();i++){
                switch (s.charAt(i)){
                    case '+':
                        if (count_equal == 0)
                            throw new Exception();
                        tags.add(Tag.PLUS);
                        break;
                    case '-':
                        if (count_equal == 0)
                            throw new Exception();
                        tags.add(Tag.MINUS);
                        break;
                    case '/':
                        if (count_equal == 0)
                            throw new Exception();
                        tags.add(Tag.SUB);
                        break;
                    case '*':
                        if (count_equal == 0)
                            throw new Exception();
                        tags.add(Tag.MULT);
                        break;
                    case ' ':
                        continue;
                    case '(':
                        tags.add(Tag.LPAREN);
                        continue;
                    case ')':
                        tags.add(Tag.RPAREN);
                        continue;
                    case '=':
                        isRight = true;
                        count_equal++;
                        if (count_equal > 1)
                            throw new Exception();
                        tags.add(Tag.EQUAL);
                        continue;
                    case ',':
                        tags.add(Tag.COMMA);
                        continue;
                    default:
                        if (Character.isDigit(s.charAt(i))){
                            tags.add(Tag.NUMBER);
                            while (i < s.length() && Character.isDigit(s.charAt(i))) {
                                i++;
                            }
                            i--;
                            continue;
                        }
                        if (Character.isAlphabetic(s.charAt(i))){
                            tags.add(Tag.IDENT);
                            count_var++;
                            newIdent=new String();
                            newIdent+=s.charAt(i++);
                            while (i<s.length() && (Character.isAlphabetic(s.charAt(i))|| Character.isDigit(s.charAt(i)))){
                                newIdent+=s.charAt(i++);
                            }
                            i--;
                            if (!isRight){
                                if (!a.containsKey(newIdent)){
                                    a.put(newIdent,k);
                                }else {
                                    throw new Exception();
                                }
                            }
                            continue;
                        }
                        throw new Exception();
                }
            }
        }

        void parse_left () throws Exception {
            if (tags.get(count) == Tag.IDENT) {
                count++;
                if (tags.get(count) == Tag.COMMA) {
                    count++;
                    count_comma++;
                    parse_left();
                }
                if (count < tags.size() && tags.get(count) == Tag.EQUAL) {
                    count++;
                    if (count_comma == 0)
                        parse_right();
                    else {
                        parse_right_comma();
                        if (count_comma != 0)
                            throw new Exception();
                    }
                }
            }
            else
                throw new Exception();
        }

        void parse_right_comma() throws Exception {
            if (tags.get(count) == Tag.IDENT || tags.get(count) == Tag.NUMBER) {
                count++;
                if (count < tags.size() && tags.get(count) == Tag.COMMA) {
                    count++;
                    count_comma--;
                    parse_right_comma();
                }
            }
        }

        void parse_right() throws Exception {
            parse_t();
            parse_e2();
        }

        void parse_t() throws  Exception {
            parse_f();
            parse_t2();
        }

        void parse_e2() throws Exception {
            if (count < tags.size() && (tags.get(count) == Tag.MINUS || tags.get(count) == Tag.PLUS)) {
                count++;
                parse_t();
                parse_e2();
            }
            if (count == tags.size()-1 && (tags.get(count) ==Tag.NUMBER || tags.get(count) == Tag.COMMA))
                throw new Exception();
        }

        void parse_t2() throws Exception{
            if (count < tags.size() && (tags.get(count) == Tag.MULT || tags.get(count) == Tag.SUB)) {
                count++;
                parse_f();
                parse_t2();
            }

        }

        void parse_f() throws Exception {
            if (tags.get(count) == Tag.MINUS || tags.get(count) == Tag.LPAREN || tags.get(count) == Tag.NUMBER || tags.get(count) == Tag.IDENT) {
                count++;
                if (tags.get(count-1) == Tag.MINUS)
                    parse_f();
                if (tags.get(count-1) == Tag.LPAREN) {
                    parse_right();
                    if (tags.get(count) == Tag.RPAREN)
                        count++;
                    else
                        throw new Exception();
                }
            }
            else
                throw new Exception();
        }

        void for_HM (int n, HashMap<String, Integer> a, Graph graph) throws Exception {
            isRight = false;
            String s = line;
            for (int i = 0; i < s.length(); i++) {
                if (s.charAt(i) == '=')
                    isRight = true;
                if (isRight) {
                    if (Character.isAlphabetic(s.charAt(i))) {
                        String var = s.charAt(i) + "";
                        i++;
                        while (i < s.length() && (Character.isAlphabetic(s.charAt(i)) || Character.isDigit(s.charAt(i)))) {
                            var += s.charAt(i);
                            i++;
                        }
                        i--;
                        if (!a.containsKey(var))
                            throw new Exception();
                        int x = a.get(var);
                        graph.getGraph().get(n).vert.add(graph.getGraph().get(x));
                        continue;
                    }
                }
            }
        }



    }

    static class Edge{
        Set<Edge> vert;
        int count;
        int color;
        public Edge(int count){
            this.vert=new HashSet<>();
            this.count=count;
            this.color=0;
        }


    }

    static class Graph {
        ArrayList<Edge> graph=new ArrayList<>();
        private boolean[]  used ;
        private ArrayList<Integer>  ans = new ArrayList<>();
        public ArrayList<Edge> getGraph() {
            return graph;
        }

        public void dfs1(int v) throws Exception {
            if (graph.get(v).color==1){
                throw new Exception();
            }
            graph.get(v).color=1;
            for (int i=0;i<graph.get(v).vert.size();i++){
                Edge help[] = new Edge[graph.size()-1];
                graph.get(v).vert.toArray(help);
                int to = help[i].count;
                dfs1(to);
            }
            graph.get(v).color=2;
        }

        public void dfs (int v) throws Exception {
            if (graph.get(v).color==1){
                throw new Exception();
            }
            graph.get(v).color=1;
            used[v] = true;
            for (int i=0; i<graph.get(v).vert.size(); ++i) {
                Edge help[] = new Edge[graph.size()-1];
                graph.get(v).vert.toArray(help);
                int to = help[i].count;
                if (!used[to]) {
                    dfs(to);
                }
                else if (graph.get(to).color==1){
                    throw new Exception();
                }
            }
            ans.add(v);
            graph.get(v).color=2;
        }

        void topological_sort() throws Exception {

            used=new boolean[graph.size()];
            for (int i=0; i<graph.size(); ++i)
                used[i] = false;
            ans.clear();
            for (int i=0; i<graph.size(); ++i)
                if (!used[i]) {
                /*
                    try {
                        dfs1(i);
                    } catch (Exception e) {
                        System.out.println("cycle");
                        System.exit(0);
                    }
                    */
                     try {
                        dfs(i);
                    }catch (Exception e){
                        System.out.println("cycle");
                        System.exit(0);
                    }
                }
        }
    }

    public static void main(String[] argc) throws Exception {
        Graph graph= new Graph();
        HashMap<String,Integer> h = new HashMap<>();
        ArrayList<Token> str = new ArrayList<>();
        Scanner in = new Scanner(System.in);
        while (in.hasNextLine()){
            str.add(new Token(in.nextLine()));
            graph.getGraph().add(new Edge(str.size()-1));
            try {
                str.get(str.size()-1).find(str.size()-1,h);
            } catch (Exception e) {
                System.out.println("syntax error");
                System.exit(0);
            }
            try {
                str.get(str.size()-1).parse_left();
            }catch (Exception e){
                System.out.println("syntax error");
                System.exit(0);
            }
        }

        for (int i = 0; i < str.size(); i++) {
            try {
                str.get(i).for_HM(i, h, graph);
            }catch (Exception e){
                System.out.println("syntax error");
                System.exit(0);
            }
        }

        graph.topological_sort();

        for (int i:graph.ans){
            System.out.println(str.get(i).line);
        }
    }
}





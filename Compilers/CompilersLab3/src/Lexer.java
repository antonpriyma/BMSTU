
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.*;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

enum TYPES {
    SINGLE,
    INT,
    FLOAT,
    MULTI
}

class Position {
    private int line, pos;

    public int getLine() {
        return line;
    }

    public int getPos() {
        return pos;
    }

    Position(int line, int pos) {
        this.line = line;
        this.pos = pos;
    }

    void next_line() {
        this.line += 1;
        this.pos = 1;
    }

    void next_pos(int pos) {
        this.pos += pos;
    }

    public String toString() {
        return "(" + this.line + "," + this.pos + ")";
    }
}

class Match {
    private String value;
    private TYPES type;
    private int line;
    private int pos;

    public Match(Position position, String value, TYPES type) {
        this.line = position.getLine();
        this.pos = position.getPos();
        this.value = value;
        this.type = type;
    }

    @Override
    public String toString() {
        return
                type + " " + "(" + line + "," + pos + ")" +
                        " " + value;

    }
}

class MatchesIterator implements Iterator<Match> {
    List<Match> matches;
    int position;

    public MatchesIterator(List<Match> matches) {
        this.matches = matches;
        position = 0;
    }

    @Override
    public boolean hasNext() {
        return position < matches.size();
    }

    @Override
    public Match next() {
        Match match = matches.get(position);
        position++;
        return match;
    }
}

public class Lexer {
    private static MatchesIterator test_match(String line, Position pos) {
        List<Match> matches = new ArrayList<>();
        String integerDigitTemplate = "\\d+";
        String floatDigitTemplate = "[-+]?\\d*((\\.[0-9]+([eE][-+]?\\d+)?)|([eE][-+]?\\d+))";
        String singleLineComment = "--.*";
        String multiLineComment = "\\{-(.|\\n)*?-\\}";

        String pattern = "(?<single>^" + singleLineComment + ")|(?<float>^" + floatDigitTemplate + ")|(?<int>^" + integerDigitTemplate + ")" + "|(?<multi>^" + multiLineComment + ")";

        Pattern p = Pattern.compile(pattern);
        Matcher m;

        while (!line.equals("")) {
            m = p.matcher(line);
            if (m.find()) {
                if (m.group("int") != null) {
                    String item = m.group("int");
                    Match match = new Match(pos, item, TYPES.INT);
                    matches.add(match);
                    pos.next_pos(item.length());
                    line = line.substring(line.indexOf(item) + item.length());

                } else if (m.group("float") != null) {
                    String item = m.group("float");
                    Match match = new Match(pos, item, TYPES.FLOAT);
                    matches.add(match);
                    pos.next_pos(item.length());
                    line = line.substring(line.indexOf(item) + item.length());

                } else if (m.group("single") != null) {
                    String item = m.group("single");
                    Match match = new Match(pos, item, TYPES.SINGLE);
                    matches.add(match);
                    pos.next_pos(item.length());
                    line = line.substring(line.indexOf(item) + item.length());

                } else if (m.group("multi") != null) {
                    String item = m.group("multi");
                    Match match = new Match(pos, item, TYPES.MULTI);
                    matches.add(match);
                    pos.next_pos(item.length());
                    line = line.substring(line.indexOf(item) + item.length());
                }
            } else {
                if (line.charAt(0) == '\n') {
                    line = line.substring(1);
                    pos.next_line();
                } else if (Character.isWhitespace(line.charAt(0))) {
                    while (Character.isWhitespace(line.charAt(0))) {
                        line = line.substring(1);
                        pos.next_pos(1);
                    }
                } else {
                    System.out.println("syntax error " + pos.toString());
                    while (!m.find() && !line.equals("")) {
                        line = line.substring(1);
                        pos.next_pos(1);
                        m = p.matcher(line);
                    }
                }
            }
        }
        return new MatchesIterator(matches);
    }

    public static void main(String[] args) {
        if (args.length != 1) {
            System.err.println("You should give 1 argument: javac -encoding utf8 <file>");
            System.exit(1);
        } else {
            String file_name = args[0];
            Position pos = new Position(1, 1);
            try {
                String text = Files.readString(Paths.get(file_name));
                MatchesIterator iterator = test_match(text, pos);
                while (iterator.hasNext()) {
                    Match match = iterator.next();
                    System.out.println(match);
                }
            } catch (IOException ex) {
                ex.printStackTrace();
            }
        }
    }

}
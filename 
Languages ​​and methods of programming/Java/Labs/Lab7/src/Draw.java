import javax.swing.*;
import java.awt.*;

public class Draw extends JPanel {
    @Override
    protected void paintComponent(Graphics g) {
        super.paintComponent(g);
        draw1(g);
    }

    private void draw1(Graphics g){
        g.drawOval(10,10,20,20);
    }
}
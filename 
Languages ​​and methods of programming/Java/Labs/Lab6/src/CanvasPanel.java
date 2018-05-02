import javax.swing.*;
import java.awt.*;

public class CanvasPanel extends JPanel{
    private int radius = 20;

    public void setRadius(int r){
        radius=r;
        repaint();
    }

    public void paintComponent(Graphics g){
        super.paintComponent(g);
        g.setColor(Color.RED);
        g.drawOval(100,100,radius,radius);
    }
}

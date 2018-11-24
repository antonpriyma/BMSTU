

import javax.swing.*;
import javax.swing.event.ChangeEvent;
import javax.swing.event.ChangeListener;


import java.awt.*;


public class Pictureform {
    private JPanel mainPanel;
    private JSpinner a;
    private JSpinner b;
    private JPanel Drawfield;
    //private JTextField Areafield;

    public Pictureform(){
        a.addChangeListener(new ChangeListener() {
            @Override
            public void stateChanged(ChangeEvent e) {
                int radius = (int)a.getValue();
                //canvasPanel.setRadius(radius);
                double area = Math.PI*radius*radius;

                //Areafield.setText(String.format("%.2f",area));
            }
        });
        a.setValue(20);
    }

    public static void main(String[] args) {
        JFrame frame = new JFrame("Round");
        //frame.setSize(200,200);
        frame.add(new Pictureform().mainPanel);

       // frame.setContentPane(new Pictureform().mainPanel);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.pack();
        frame.setVisible(true);
    }
}





import javax.swing.*;
import javax.swing.event.ChangeEvent;
import javax.swing.event.ChangeListener;

public class PictureForm {
    public static void main(String[] args) {
        JFrame frame = new JFrame("Эллипс");
        frame.setContentPane(new PictureForm().MainPanel);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.pack();
        frame.setVisible(true);

    }

    public PictureForm(){
        radiusSpinner.addChangeListener(new ChangeListener() {
            @Override
            public void stateChanged(ChangeEvent e) {
                int radius = (int)radiusSpinner.getValue();
                double area = Math.PI*radius*radius;
                areaField.setText(String.format("%.2f",area));
                canvasPanel.setRadius(radius);
            }
        });
//        canvasPanel.setRadius(20);
        radiusSpinner.setValue(20);
    }

    private JPanel MainPanel;
    private JTextField areaField;
    private JSpinner radiusSpinner;
    private CanvasPanel canvasPanel;
}

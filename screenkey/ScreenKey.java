import java.io.*;
import java.awt.*;
import java.awt.image.*;
import javax.imageio.*;
import javax.swing.*;

public class ScreenKey extends JFrame {
	Keyer k = new Keyer();
	Robot r = null;
	public ScreenKey() {
		k.setKey(Color.GREEN);
		k.setMaxDistance(160.0f);
		try {
			r = new Robot();
		} catch(Exception e){
			System.out.println("Robot failed: "+e.toString());
			e.printStackTrace();
			System.exit(1);
		}
		setDefaultCloseOperation(WindowConstants.EXIT_ON_CLOSE);
		setContentPane(new DrawPane());
	}
	class DrawPane extends JPanel {
		public void paintComponent(Graphics g) {
			BufferedImage result = null;
			try {
				BufferedImage base = r.createScreenCapture(new Rectangle(0,0, 640, 480));
				BufferedImage layer = r.createScreenCapture(new Rectangle(640,0, 640, 480));
				BufferedImage mask = k.generateMask(layer);
				result = k.key(layer, mask, base);
			} catch (Exception e) {
				System.out.println("Screencap: "+e.toString());
				e.printStackTrace();
			}
			g.drawImage(result, 0, 0, Color.BLACK, new ImageObserver() {
				public boolean imageUpdate(Image img, int infoflags, int x, int y, int width, int height) {
					return true;
				}
			});
			repaint();
		}
	}

	public static void main(String[] args) {
		JFrame f = new ScreenKey();
		f.setSize(200, 200);
		f.setVisible(true);
	}
}

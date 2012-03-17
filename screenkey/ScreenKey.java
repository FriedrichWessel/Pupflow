import java.io.*;
import java.awt.*;
import java.awt.image.*;
import javax.imageio.*;

public class ScreenKey {
	public static void main(String[] args) {
		try {
			Robot r = new Robot();
			BufferedImage bi = r.createScreenCapture(new Rectangle(0,0, 200, 200));
			File f = new File("test.png");
			ImageIO.write(bi, "png", f);
		} catch (Exception e) {
			System.out.println("Ficken: "+e.toString());
		}
	}
}

import java.io.*;
import java.awt.*;
import java.awt.image.*;
import javax.imageio.*;

public class ScreenKey {
	public static void main(String[] args) {
		Keyer k = new Keyer();
		k.setKey(Color.WHITE);
		k.setMaxDistance(40.0f);
		try {
			Robot r = new Robot();
			BufferedImage base = r.createScreenCapture(new Rectangle(0,0, 200, 200));
			BufferedImage layer = r.createScreenCapture(new Rectangle(200,200, 200, 200));
			BufferedImage mask = k.generateMask(layer);
			BufferedImage result = k.key(layer, mask, base);
			File f = new File("mask.png");
			ImageIO.write(mask, "png", f);
			f = new File("result.png");
			ImageIO.write(result, "png", f);
		} catch (Exception e) {
			System.out.println("Ficken: "+e.toString());
			e.printStackTrace();
		}
	}
}

import java.awt.*;
import java.awt.image.*;

public class Keyer {
	private Color key;
	private float maxdistance;
	private final int[] KEYED = new int[]{0,0,0}, UNKEYED = new int[]{255,255,255};

	public Keyer()	{
		key = Color.BLACK;
		maxdistance = 0.0f;
	}

	public Color getKey() {
		return key;
	}

	public void setKey(Color k) {
		key = k;
	}

	public float getMaxDistance() {
		return maxdistance;
	}

	public void setMaxDistance(float d) {
		maxdistance = d;
	}

	public boolean isKeyed(Color c) {
		int deltared = c.getRed() - key.getRed();
		int deltagreen = c.getGreen() - key.getGreen();
		int deltablue = c.getBlue() - key.getBlue();
		int distancesquared = deltared*deltared + deltagreen*deltagreen + deltablue*deltablue;
		return distancesquared <= maxdistance*maxdistance;
	}

	public BufferedImage generateMask(BufferedImage i) {
		BufferedImage mask = new BufferedImage(i.getWidth(), i.getHeight(), BufferedImage.TYPE_BYTE_BINARY);
		WritableRaster maskr = mask.getRaster();
		Raster r = i.getData();
		int[] buf = new int[3];
		for(int y = 0; y < i.getHeight(); y++) {
			for(int x = 0; x < i.getWidth(); x++) {
				r.getPixel(x, y, buf);
				Color c = new Color(buf[0], buf[1], buf[2]);
				if(isKeyed(c)) {
					maskr.setPixel(x, y, KEYED);
				} else {
					maskr.setPixel(x, y, UNKEYED);
				}
			}
		}
		return mask;
	}

	public BufferedImage key(BufferedImage layer, BufferedImage mask, BufferedImage base) {
		BufferedImage i = new BufferedImage(base.getWidth(), base.getHeight(), BufferedImage.TYPE_INT_RGB);
		WritableRaster r = i.getRaster();
		WritableRaster maskr = mask.getRaster();
		Raster layerr = layer.getData();
		Raster baser = base.getData();
		int[] buf = new int[3];

		for(int y = 0; y < i.getHeight(); y++) {
			for(int x = 0; x < i.getWidth(); x++) {
				maskr.getPixel(x, y, buf);
				if(buf[0] == 0) {
					baser.getPixel(x, y, buf);
					r.setPixel(x, y, buf);
				} else {
					layerr.getPixel(x, y, buf);
					r.setPixel(x, y, buf);
				}
			}
		}
		return i;
	}
}

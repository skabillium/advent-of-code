package common;

import java.io.InputStream;
import java.io.IOException;

public interface Solver {
    Object solve(InputStream in) throws IOException;
}

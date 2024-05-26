package common;

import java.io.InputStream;
import java.io.IOException;

public interface Solver {
    int solve(InputStream in) throws IOException;
}

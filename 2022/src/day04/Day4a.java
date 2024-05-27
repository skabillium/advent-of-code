package day04;

import java.io.*;

import common.Solver;

public class Day4a implements Solver {
    public int solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));
        var count = 0;

        while (true) {
            var line = reader.readLine();
            if (line == null) {
                break;
            }

            var pairs = line.split(",");
            var range1 = pairs[0].split("-");
            var start1 = Integer.parseInt(range1[0]);
            var end1 = Integer.parseInt(range1[1]);

            var range2 = pairs[1].split("-");
            var start2 = Integer.parseInt(range2[0]);
            var end2 = Integer.parseInt(range2[1]);

            if ((start1 >= start2 && end1 <= end2) || (start2 >= start1 && end2 <= end1)) {
                count++;
            }
        }
        return count;
    }
}
package day06;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;

import common.Solver;

public class Day6a implements Solver {
    public Object solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));
        var builder = new StringBuilder();

        // Read entire input stream into a string
        while (true) {
            var line = reader.readLine();
            if (line == null) {
                break;
            }
            builder.append(line);
        }

        var position = 0;
        var chars = builder.toString().toCharArray();
        for (var i = 0; i < chars.length - 13; i++) {
            var one = chars[i];
            var two = chars[i + 1];
            var three = chars[i + 2];
            var four = chars[i + 3];

            if (one != two && two != three && three != four && four != one && three != one && two != four) {
                position = i + 4; // Format the response correctly
                break;
            }
        }
        return position;
    }
}
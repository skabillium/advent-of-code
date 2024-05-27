package day06;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.HashSet;

import common.Solver;

public class Day6b implements Solver {
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
            var charset = new HashSet<Character>();
            for (var j = 0; j < 14; j++) {
                charset.add(chars[i + j]);
            }

            if (charset.size() == 14) {
                position = i + 14;
                break;
            }
        }
        return position;
    }
}
package day03;

import java.io.*;
import java.util.HashSet;

import common.Solver;

public class Day3a implements Solver {
    public int solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));

        var priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
        var sum = 0;
        while (true) {
            var line = reader.readLine();
            if (line == null) {
                break;
            }

            var first = line.substring(0, line.length() / 2);
            var second = line.substring(line.length() / 2, line.length());

            var firstChars = new HashSet<String>();
            for (var c : first.toCharArray()) {
                firstChars.add(String.valueOf(c));
            }

            var seen = new HashSet<String>();
            for (var c : second.toCharArray()) {
                var str = String.valueOf(c);
                if (firstChars.contains(str) && !seen.contains(str)) {
                    sum += priorities.indexOf(str) + 1;
                    seen.add(str);
                }
            }

        }
        return sum;
    }
}

package day03;

import java.io.*;
import java.util.HashSet;

import common.Solver;

public class Day3b implements Solver {
    public Integer solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));

        var priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
        var sum = 0;
        while (true) {
            var first = reader.readLine();
            var second = reader.readLine();
            var third = reader.readLine();
            if (first == null || second == null || third == null) {
                break;
            }

            var chars = new HashSet<Character>();
            for (var c : first.toCharArray()) {
                chars.add(c);
            }

            var secondChars = new HashSet<Character>();
            for (var c : second.toCharArray()) {
                secondChars.add(c);
            }

            var thirdChars = new HashSet<Character>();
            for (var c : third.toCharArray()) {
                thirdChars.add(c);
            }

            chars.retainAll(secondChars);
            chars.retainAll(thirdChars);

            for (var c : chars) {
                sum += priorities.indexOf(c) + 1;
            }
        }
        return sum;
    }
}

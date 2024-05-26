package day01;

import java.util.ArrayList;
import java.util.Collections;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.InputStream;

import common.Solver;

public class Day1b implements Solver {
    public int solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));
        var currentCalories = 0;
        var elfs = new ArrayList<Integer>();
        String line;

        while (true) {
            line = reader.readLine();
            if (line == null) {
                elfs.add(currentCalories);
                break;
            }

            if (line.equals("")) {
                elfs.add(currentCalories);
                currentCalories = 0;
                continue;
            }
            var calories = Integer.parseInt(line);
            currentCalories += calories;
        }

        Collections.sort(elfs);

        return elfs.get(elfs.size() - 1) + elfs.get(elfs.size() - 2) + elfs.get(elfs.size() - 3);

    }
}
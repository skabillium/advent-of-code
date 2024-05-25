package day01;

import java.util.ArrayList;
import java.util.Collections;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

import common.Solver;

public class Day1b implements Solver {
    public void solve() {
        var reader = new BufferedReader(new InputStreamReader(System.in));
        var currentCalories = 0;
        var elfs = new ArrayList<Integer>();
        String line;
        try {
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

            System.err.println(elfs.get(elfs.size() - 1) + elfs.get(elfs.size() - 2) + elfs.get(elfs.size() - 3));
        } catch (IOException e) {
            System.err.println(e);
            return;
        }
    }
}
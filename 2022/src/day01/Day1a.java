package day01;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.InputStream;

import common.Solver;

public class Day1a implements Solver {
    public int solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));
        var maxCalories = 0;
        var currentCalories = 0;
        String line;

        while ((line = reader.readLine()) != null) {
            if (line.equals("")) {
                if (currentCalories > maxCalories) {
                    maxCalories = currentCalories;
                }
                currentCalories = 0;
                continue;
            }
            var calories = Integer.parseInt(line);
            currentCalories += calories;
        }
        return maxCalories;
    }
}

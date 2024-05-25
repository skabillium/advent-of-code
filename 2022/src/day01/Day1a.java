package day01;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

import common.Solver;

public class Day1a implements Solver {
    public void solve() {
        var reader = new BufferedReader(new InputStreamReader(System.in));
        var maxCalories = 0;
        var currentCalories = 0;
        String line;
        try {
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
            System.out.println(maxCalories);
        } catch (IOException e) {
            System.err.println(e);
            return;
        }
    }
}

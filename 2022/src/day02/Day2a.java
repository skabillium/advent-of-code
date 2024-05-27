package day02;

import common.Solver;

import java.util.HashMap;
import java.io.InputStream;
import java.io.IOException;
import java.io.BufferedReader;
import java.io.InputStreamReader;

public class Day2a implements Solver {
    class Choice {
        final int value;
        final String weakTo;
        final String strongTo;
        final String equal;

        Choice(int value, String equal, String weakTo, String strongTo) {
            this.value = value;
            this.equal = equal;
            this.weakTo = weakTo;
            this.strongTo = strongTo;
        }
    }

    public Integer solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));
        var total = 0;

        while (true) {
            var line = reader.readLine();
            if (line == null || line.equals("")) {
                break;
            }

            // Rock A X
            // Paper B Y
            // Scissors C Z

            var choices = new HashMap<String, Choice>();
            choices.put("X", new Choice(1, "A", "B", "C")); // Rock
            choices.put("Y", new Choice(2, "B", "C", "A")); // Paper
            choices.put("Z", new Choice(3, "C", "A", "B")); // Scissors

            var round = line.split(" ");
            var elf = round[0];
            var my = round[1];

            var move = choices.get(my);
            var points = move.value;
            if (move.equal.equals(elf)) {
                points += 3;
            } else if (move.strongTo.equals(elf)) {
                points += 6;
            }

            total += points;
        }
        return total;
    }
}
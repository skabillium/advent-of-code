package day02;

import java.util.HashMap;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.BufferedReader;
import java.io.IOException;

import common.Solver;

public class Day2b implements Solver {
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

        var choices = new HashMap<String, Choice>();
        choices.put("A", new Choice(1, "A", "B", "C")); // Rock
        choices.put("B", new Choice(2, "B", "C", "A")); // Paper
        choices.put("C", new Choice(3, "C", "A", "B")); // Scissors

        while (true) {
            var line = reader.readLine();
            if (line == null || line.equals("")) {
                break;
            }

            var round = line.split(" ");
            var elf = round[0];
            var result = round[1];

            var elfMove = choices.get(elf);
            var points = switch (result) {
                case "X" -> choices.get(elfMove.strongTo).value;
                case "Y" -> elfMove.value + 3;
                case "Z" -> choices.get(elfMove.weakTo).value + 6;
                default -> 0;
            };

            total += points;
        }
        return total;
    }
}
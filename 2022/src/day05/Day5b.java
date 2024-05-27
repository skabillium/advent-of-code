package day05;

import java.io.*;
import java.util.ArrayList;
import java.util.Stack;

import common.Solver;

public class Day5b implements Solver {
    public String solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));

        // Build the stacks
        var lines = new ArrayList<String>();
        while (true) {
            var line = reader.readLine();
            if (line.equals("")) {
                break;
            }
            lines.add(line);
        }

        var stacks = new ArrayList<Stack<String>>();
        var numbers = lines.get(lines.size() - 1).toCharArray();
        for (var i = 0; i < numbers.length; i++) {
            var c = numbers[i];
            if (Character.isDigit(c)) {
                // Gather all elements for stack
                var stack = new Stack<String>();
                for (var j = lines.size() - 2; j >= 0; j--) {
                    var line = lines.get(j).toCharArray();
                    var ch = line[i];
                    if (Character.isAlphabetic(ch)) {
                        stack.add(String.valueOf(ch));
                    } else {
                        stacks.add(stack);
                        break;
                    }

                    if (j == 0) {
                        stacks.add(stack);
                    }
                }

            }
        }

        while (true) {
            var line = reader.readLine();
            if (line == null) {
                break;
            }

            var tokens = line.split(" ");
            var amount = Integer.parseInt(tokens[1]);
            var from = Integer.parseInt(tokens[3]) - 1;
            var to = Integer.parseInt(tokens[5]) - 1;

            var crates = new ArrayList<String>();
            for (var i = 0; i < amount; i++) {
                crates.addFirst(stacks.get(from).pop());

            }

            for (var cr : crates) {
                stacks.get(to).add(cr);
            }
        }

        var result = "";
        for (var stack : stacks) {
            result = result.concat(stack.pop());
        }

        return result;
    }
}
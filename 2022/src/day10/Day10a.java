package day10;

import java.io.*;

import common.Solver;

public class Day10a implements Solver {

    public Integer solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));

        var clock = 0;
        var wait = false;
        var register = 1;
        var nextClock = 20;
        var sum = 0;
        var add = 0;
        while (true) {
            clock++;

            if (clock == nextClock) {
                sum += clock * register;
                nextClock = clock + 40;
                if (clock == 220) {
                    break;
                }
            }

            if (wait) {
                register += add;
                wait = false;
                continue;
            }

            var line = reader.readLine();
            if (line == null) {
                break;
            }

            if (line.equals("noop")) {
                continue;
            }

            add = Integer.parseInt(line.split(" ")[1]);
            wait = true;
        }

        return sum;
    }
}
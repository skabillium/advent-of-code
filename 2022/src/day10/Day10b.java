package day10;

import java.io.*;

import common.Solver;

public class Day10b implements Solver {
    public Boolean solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));

        var clock = 0;
        var wait = false;
        var register = 1;
        var nextClock = 40;
        var add = 0;
        while (true) {
            clock++;

            // Draw pixel
            if ((clock - 1) % 40 >= register - 1 && (clock - 1) % 40 <= register + 1) {
                System.out.print("#");
            } else {
                System.out.print(".");
            }

            if (clock == nextClock) {
                // sum += clock * register;
                System.out.println();
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

        return true;
    }
}
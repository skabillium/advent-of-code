import day01.Day1a;
import day01.Day1b;
import day02.Day2a;
import day02.Day2b;
import day03.Day3a;
import day03.Day3b;
import day04.Day4a;
import day04.Day4b;
import common.Solver;

import java.io.FileInputStream;
import java.io.IOException;

public class Main {
    public static Solver[][] solvers = new Solver[25][2];

    public static void main(String[] args) throws IOException {
        registerSolvers();

        if (args.length < 2) {
            System.out.println("Not enough arguments need [day] [part]");
            return;
        }

        var day = Integer.parseInt(args[0]);
        var part = Integer.parseInt(args[1]);

        if (day < 1 || day > 25) {
            System.out.println("Day needs to be between 1 and 25");
            return;
        }

        if (part != 1 && part != 2) {
            System.out.println("Part needs to be either 1 or 2");
            return;
        }

        var solver = solvers[day - 1][part - 1];
        if (solver == null) {
            System.out.printf("Solution for day %d part %d has not been implemented yet \n", day, part);
            return;
        }

        var in = System.in;
        if (args.length == 3) {
            in = new FileInputStream(args[2]);
        }

        var solution = solver.solve(in);
        System.out.printf("Day %d, part %d: %d \n", day, part, solution);
    }

    public static void registerSolvers() {
        // Day 1
        solvers[0][0] = new Day1a();
        solvers[0][1] = new Day1b();
        // Day 2
        solvers[1][0] = new Day2a();
        solvers[1][1] = new Day2b();
        // Day 3
        solvers[2][0] = new Day3a();
        solvers[2][1] = new Day3b();
        // Day 4
        solvers[3][0] = new Day4a();
        solvers[3][1] = new Day4b();
    }
}

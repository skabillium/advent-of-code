import day01.Day1a;
import day01.Day1b;
import common.Solver;

public class Main {
    public static Solver[][] solvers = new Solver[25][2];

    public static void main(String[] args) {
        initSolvers();
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

        solver.solve();
    }

    public static void initSolvers() {
        solvers[0][0] = new Day1a();
        solvers[0][1] = new Day1b();
    }
}

package day08;

import java.io.*;
import java.util.ArrayList;

import common.Solver;

public class Day8a implements Solver {
    class Direction {
        int column;
        int row;

        Direction(int column, int row) {
            this.column = column;
            this.row = row;
        }
    }

    private int[][] readForest(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));
        var lines = new ArrayList<int[]>();
        String line;
        while ((line = reader.readLine()) != null) {
            var row = new int[line.length()];
            for (var i = 0; i < line.length(); i++) {
                row[i] = (int) line.charAt(i);
            }
            lines.add(row);
        }

        var forest = new int[lines.size()][];
        for (var i = 0; i < lines.size(); i++) {
            forest[i] = lines.get(i);
        }
        return forest;
    }

    private boolean isVisible(int[][] forest, int row, int column, Direction direction) {
        var tree = forest[row][column];
        while (true) {
            row += direction.row;
            column += direction.column;
            if (row < 0 || row >= forest.length || column < 0 || column >= forest[0].length) {
                return true;
            }
            if (forest[row][column] >= tree) {
                return false;
            }
        }
    }

    public Integer solve(InputStream in) throws IOException {
        var forest = readForest(in);

        var directions = new Direction[4];
        directions[0] = new Direction(-1, 0); // Left
        directions[1] = new Direction(1, 0); // Right
        directions[2] = new Direction(0, -1); // Up
        directions[3] = new Direction(0, 1); // Down

        var count = 2 * (forest.length - 1) + 2 * (forest[0].length - 1);
        for (var i = 1; i < forest.length - 1; i++) {
            for (var j = 1; j < forest[0].length - 1; j++) {
                var visible = false;
                for (var k = 0; k < directions.length; k++) {
                    if (isVisible(forest, i, j, directions[k])) {
                        visible = true;
                        break;
                    }
                }
                if (visible) {
                    count++;
                }
            }
        }
        return count;
    }
}
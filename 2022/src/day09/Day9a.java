package day09;

import java.io.*;
import java.util.HashSet;

import common.Solver;

public class Day9a implements Solver {
    class Position {
        int row;
        int col;

        Position() {
            this.row = 0;
            this.col = 0;
        }

        Position(int row, int col) {
            this.row = row;
            this.col = col;
        }

        Position distance(Position position) {
            return new Position(position.row - row, position.col - col);
        }

        boolean equals(Position position) {
            return row == position.row && col == position.col;
        }

        @Override
        public String toString() {
            return String.format("{row: %d, col: %d}", row, col);
        }
    }

    private final HashSet<String> visited = new HashSet<String>();

    public Integer solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));

        var head = new Position();
        var tail = new Position();
        addVisited(tail);

        String line;
        while ((line = reader.readLine()) != null) {
            var split = line.split(" ");
            var dir = split[0];
            var steps = Integer.parseInt(split[1]);

            var direction = switch (dir) {
                case "U" -> new Position(-1, 0);
                case "D" -> new Position(1, 0);
                case "R" -> new Position(0, 1);
                case "L" -> new Position(0, -1);
                default -> new Position();
            };

            for (var i = 0; i < steps; i++) {
                var previous = head;
                head = moveBy(head, direction);
                tail = moveTail(head, previous, tail);
                addVisited(tail);
            }
        }

        return visited.size();
    }

    private Position moveTail(Position head, Position previous, Position tail) {
        if (tail.equals(head)) {
            return tail;
        }

        var distance = tail.distance(head);
        if (Math.abs(distance.row) == 1 && Math.abs(distance.col) == 1) {
            return tail;
        }

        if (distance.row == 0) {
            if (Math.abs(distance.col) != 1) {
                tail.col += distance.col > 0 ? 1 : -1;
            }
            return tail;
        }

        if (distance.col == 0) {
            if (Math.abs(distance.row) != 1) {
                tail.row += distance.row > 0 ? 1 : -1;
            }
            return tail;
        }

        tail.col += distance.col > 0 ? 1 : -1;
        tail.row += distance.row > 0 ? 1 : -1;

        return tail;
    }

    private Position moveBy(Position position, Position direction) {
        return new Position(position.row + direction.row, position.col + direction.col);
    }

    private void addVisited(Position position) {
        visited.add(String.format("%d:%d", position.row, position.col));
    }
}
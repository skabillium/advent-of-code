package day11;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.function.IntUnaryOperator;

import common.Solver;

public class Day11a implements Solver {
    class Monkey {
        ArrayList<Integer> items;
        IntUnaryOperator operation;
        int divisibleBy;
        int ifTrue;
        int ifFalse;
        int inspections = 0;

        Monkey(ArrayList<Integer> items, IntUnaryOperator operation, int divisibleBy, int ifTrue, int ifFalse) {
            this.items = items;
            this.operation = operation;
            this.divisibleBy = divisibleBy;
            this.ifTrue = ifTrue;
            this.ifFalse = ifFalse;
        }

        int getNext(int n) {
            return n % divisibleBy == 0 ? ifTrue : ifFalse;
        }

        int getWorryLevel(int n) {
            return operation.applyAsInt(n);
        }

        void addItem(int item) {
            items.add(item);
        }

        int getInspections() {
            return inspections;
        }

        void print() {
            System.out.print("Items:");
            items.forEach(item -> System.out.printf(" %d", item));
            System.out.println();
            System.out.printf("Divisible by: %d \n", divisibleBy);
            System.out.printf("If true -> %d \n", ifTrue);
            System.out.printf("If false -> %d \n", ifFalse);
        }

    }

    public Integer solve(InputStream in) throws IOException {
        var monkeys = readMonkeys(in);
        var rounds = 20;

        for (var i = 0; i < rounds; i++) {
            for (var j = 0; j < monkeys.size(); j++) {
                var monkey = monkeys.get(j);
                var items = new ArrayList<Integer>();
                for (var item : monkey.items) {
                    monkey.inspections++;
                    var worryLevel = monkey.getWorryLevel(item);
                    var val = worryLevel / 3;
                    var next = monkey.getNext(val);
                    if (next == j) {
                        items.add(val);
                    }
                    monkeys.get(next).addItem(val);
                }
                monkey.items = items;
            }
        }

        monkeys.sort(Comparator.comparingInt(Monkey::getInspections));

        return monkeys.get(monkeys.size() - 1).getInspections() * monkeys.get(monkeys.size() - 2).getInspections();
    }

    private ArrayList<Monkey> readMonkeys(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));
        var monkeys = new ArrayList<Monkey>();

        String line;
        while ((line = reader.readLine()) != null) {
            line = reader.readLine();
            var starting = line.substring("  Starting items: ".length()).split(", ");
            var items = new ArrayList<Integer>();
            for (var i = 0; i < starting.length; i++) {
                items.add(Integer.parseInt(starting[i]));
            }

            line = reader.readLine();
            var expr = line.substring(" Operation: new = ".length()).split(" ");
            var rhs = expr[3];
            IntUnaryOperator operation = switch (expr[2]) {
                case "*" -> {
                    if (rhs.equals("old")) {
                        yield (int i) -> i * i;
                    }
                    yield (int i) -> i * Integer.parseInt(rhs);
                }
                case "+" -> {
                    if (rhs.equals("old")) {
                        yield (int i) -> i + i;
                    }
                    yield (int i) -> i + Integer.parseInt(rhs);
                }
                default -> null;
            };

            line = reader.readLine();
            var div = Integer.parseInt(line.substring("  Test: divisible by ".length()));

            line = reader.readLine();
            var ifTrue = Integer.parseInt(line.substring("    If true: throw to monkey ".length()));

            line = reader.readLine();
            var ifFalse = Integer.parseInt(line.substring("    If false: throw to monkey ".length()));

            monkeys.add(new Monkey(items, operation, div, ifTrue, ifFalse));
            reader.readLine();
        }
        return monkeys;
    }

}
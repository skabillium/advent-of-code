package day11;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.function.UnaryOperator;

import common.Solver;

public class Day11b implements Solver {
    class Monkey {
        ArrayList<BigInteger> items;
        UnaryOperator<BigInteger> operation;
        Integer divisibleBy;
        int ifTrue;
        int ifFalse;
        int inspections = 0;

        Monkey(ArrayList<BigInteger> items, UnaryOperator<BigInteger> operation, int divisibleBy, int ifTrue,
                int ifFalse) {
            this.items = items;
            this.operation = operation;
            this.divisibleBy = divisibleBy;
            this.ifTrue = ifTrue;
            this.ifFalse = ifFalse;
        }

        int getNext(BigInteger n) {
            return n.mod(BigInteger.valueOf(divisibleBy.intValue()))
                    .compareTo(BigInteger.valueOf(Integer.valueOf(0))) == 0 ? ifTrue : ifFalse;
        }

        BigInteger getWorryLevel(BigInteger n) {
            return operation.apply(n);
        }

        void addItem(BigInteger item) {
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
        var rounds = 10_000;

        for (var i = 0; i < rounds; i++) {
            for (var j = 0; j < monkeys.size(); j++) {
                var monkey = monkeys.get(j);
                var items = new ArrayList<BigInteger>();
                for (var item : monkey.items) {
                    monkey.inspections++;
                    var worryLevel = monkey.getWorryLevel(item);
                    var val = worryLevel;
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
            var items = new ArrayList<BigInteger>();
            for (var i = 0; i < starting.length; i++) {
                items.add(new BigInteger(starting[i]));
            }

            line = reader.readLine();
            var expr = line.substring(" Operation: new = ".length()).split(" ");
            var rhs = expr[3];
            UnaryOperator<BigInteger> operation = switch (expr[2]) {
                case "*" -> {
                    if (rhs.equals("old")) {
                        yield (BigInteger i) -> i.pow(2);
                    }
                    yield (BigInteger i) -> i.multiply(new BigInteger(rhs));
                }
                case "+" -> {
                    if (rhs.equals("old")) {
                        yield (BigInteger i) -> i.add(i);
                    }
                    yield (BigInteger i) -> i.add(new BigInteger(rhs));
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
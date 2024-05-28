package day07;

import java.io.*;
import java.util.HashMap;
import java.util.List;
import java.util.ArrayList;

import common.Solver;

public class Day7a implements Solver {
    class FsNode {
        String name;
    }

    class FsDir extends FsNode {
        FsDir parent;
        HashMap<String, FsDir> dirs;
        HashMap<String, FsFile> files;

        FsDir(String name, FsDir parent) {
            this.name = name;
            this.parent = parent;
            this.dirs = new HashMap<String, FsDir>();
            this.files = new HashMap<String, FsFile>();
        }

        FsDir getChildDir(String childName) {
            var child = dirs.get(childName);
            var found = child != null;
            if (!found) {
                addChildDir(childName);
            }
            return child;
        }

        void addChildDir(String name) {
            dirs.put(name, new FsDir(name, this));
        }

        void addFile(FsFile file) {
            files.put(file.name, file);
        }

        int size() {
            var sum = 0;
            for (var file : files.values()) {
                sum += file.size;
            }
            for (var dir : dirs.values()) {
                sum += dir.size();
            }
            return sum;
        }

        void sum(List<Integer> sizes) {
            var sz = size();
            if (sz < 100_000) {
                sizes.add(sz);
            }
            for (var dir : dirs.values()) {
                dir.sum(sizes);
            }
        }
    }

    class FsFile extends FsNode {
        int size;

        FsFile(String name, int size) {
            this.name = name;
            this.size = size;
        }
    }

    public Integer solve(InputStream in) throws IOException {
        var reader = new BufferedReader(new InputStreamReader(in));
        reader.readLine(); // Skip first line

        var root = new FsDir("/", null);
        var cwd = root;
        while (true) {
            var line = reader.readLine();
            if (line == null) {
                break;
            }

            if (line.startsWith("$ ls")) {
                continue;
            }

            if (line.startsWith("$ cd")) {
                var dir = line.split(" ")[2];
                if (dir.equals("..")) {
                    cwd = cwd.parent;
                    continue;
                }
                var child = cwd.getChildDir(dir);
                cwd = child;
                continue;
            }

            if (line.startsWith("dir")) {
                cwd.addChildDir(line.split(" ")[1]);
                continue;
            }

            // File
            var split = line.split(" ");
            cwd.addFile(new FsFile(split[1], Integer.parseInt(split[0])));
        }

        var sizes = new ArrayList<Integer>();
        root.sum(sizes);

        var sum = sizes.stream().mapToInt(Integer::intValue).sum();
        return sum;
    }
}
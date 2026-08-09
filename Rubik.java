// Rubik.java
import java.util.*;

public class Rubik {
    private char[][][] faces = new char[6][4][4];

    public Rubik() {
        char[] colors = {'W', 'Y', 'R', 'O', 'B', 'G'};
        for (int f=0; f<6; f++) {
            for (int i=0; i<4; i++) {
                for (int j=0; j<4; j++) {
                    faces[f][i][j] = colors[f];
                }
            }
        }
    }

    private void rotateFaceCW(int idx) {
        char[][] newFace = new char[4][4];
        for (int i=0; i<4; i++) {
            for (int j=0; j<4; j++) {
                newFace[i][j] = faces[idx][3-j][i];
            }
        }
        faces[idx] = newFace;
    }

    private void rotateFaceCCW(int idx) {
        char[][] newFace = new char[4][4];
        for (int i=0; i<4; i++) {
            for (int j=0; j<4; j++) {
                newFace[i][j] = faces[idx][j][3-i];
            }
        }
        faces[idx] = newFace;
    }

    public void applyMove(String move) {
        switch (move) {
            case "U": rotateFaceCW(0); break;
            case "U'": rotateFaceCCW(0); break;
            case "R": rotateFaceCW(2); break;
            case "R'": rotateFaceCCW(2); break;
            case "F": rotateFaceCW(4); break;
            case "F'": rotateFaceCCW(4); break;
            default: System.out.println("Неизвестный ход");
        }
    }

    public List<String> scramble(int moves) {
        String[] possible = {"U", "U'", "R", "R'", "F", "F'"};
        List<String> seq = new ArrayList<>();
        Random rand = new Random();
        for (int i=0; i<moves; i++) {
            String m = possible[rand.nextInt(possible.length)];
            seq.add(m);
            applyMove(m);
        }
        return seq;
    }

    public void display() {
        System.out.println("\n--- Кубик Рубика 4x4 ---");
        String[] names = {"U", "D", "R", "L", "F", "B"};
        for (int f=0; f<6; f++) {
            System.out.println(names[f] + ":");
            for (int i=0; i<4; i++) {
                for (int j=0; j<4; j++) {
                    System.out.print(faces[f][i][j] + " ");
                }
                System.out.println();
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        Rubik cube = new Rubik();
        System.out.println("Добро пожаловать в симулятор кубика Рубика 4x4!");
        System.out.println("Команды: U, U', R, R', F, F', scramble, reset, quit, help");
        while (true) {
            System.out.print("> ");
            String cmd = scanner.nextLine().trim();
            if (cmd.equals("quit")) break;
            else if (cmd.equals("help")) {
                System.out.println("U, U', R, R', F, F', D, D', L, L', B, B' - ходы");
                System.out.println("scramble - перемешать");
                System.out.println("reset - собрать");
                System.out.println("quit - выход");
            } else if (cmd.equals("reset")) {
                cube = new Rubik();
                System.out.println("Куб собран.");
            } else if (cmd.equals("scramble")) {
                List<String> seq = cube.scramble(50);
                System.out.println("Перемешивание: " + String.join(" ", seq));
            } else {
                cube.applyMove(cmd);
            }
            cube.display();
        }
        scanner.close();
    }
}

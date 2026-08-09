// rubik.cpp
#include <iostream>
#include <string>
#include <vector>
#include <cstdlib>
#include <ctime>
#include <algorithm>

using namespace std;

class Rubik4x4 {
public:
    char faces[6][4][4];

    Rubik4x4() {
        char colors[6] = {'W', 'Y', 'R', 'O', 'B', 'G'};
        for (int f=0; f<6; f++) {
            for (int i=0; i<4; i++) {
                for (int j=0; j<4; j++) {
                    faces[f][i][j] = colors[f];
                }
            }
        }
    }

    void rotateFaceCW(int idx) {
        char newFace[4][4];
        for (int i=0; i<4; i++) {
            for (int j=0; j<4; j++) {
                newFace[i][j] = faces[idx][3-j][i];
            }
        }
        for (int i=0; i<4; i++)
            for (int j=0; j<4; j++)
                faces[idx][i][j] = newFace[i][j];
    }

    void rotateFaceCCW(int idx) {
        char newFace[4][4];
        for (int i=0; i<4; i++) {
            for (int j=0; j<4; j++) {
                newFace[i][j] = faces[idx][j][3-i];
            }
        }
        for (int i=0; i<4; i++)
            for (int j=0; j<4; j++)
                faces[idx][i][j] = newFace[i][j];
    }

    void applyMove(const string& move) {
        if (move == "U") rotateFaceCW(0);
        else if (move == "U'") rotateFaceCCW(0);
        else if (move == "R") rotateFaceCW(2);
        else if (move == "R'") rotateFaceCCW(2);
        else if (move == "F") rotateFaceCW(4);
        else if (move == "F'") rotateFaceCCW(4);
        else cout << "Неизвестный ход" << endl;
    }

    vector<string> scramble(int moves) {
        vector<string> possible = {"U", "U'", "R", "R'", "F", "F'"};
        vector<string> seq;
        for (int i=0; i<moves; i++) {
            string m = possible[rand() % possible.size()];
            seq.push_back(m);
            applyMove(m);
        }
        return seq;
    }

    void display() {
        cout << "\n--- Кубик Рубика 4x4 ---" << endl;
        string names[6] = {"U", "D", "R", "L", "F", "B"};
        for (int f=0; f<6; f++) {
            cout << names[f] << ":" << endl;
            for (int i=0; i<4; i++) {
                for (int j=0; j<4; j++) {
                    cout << faces[f][i][j] << " ";
                }
                cout << endl;
            }
            cout << endl;
        }
    }
};

int main() {
    srand(time(0));
    Rubik4x4 cube;
    string cmd;
    cout << "Добро пожаловать в симулятор кубика Рубика 4x4!" << endl;
    cout << "Команды: U, U', R, R', F, F', scramble, reset, quit, help" << endl;
    while (true) {
        cout << "> ";
        getline(cin, cmd);
        if (cmd == "quit") break;
        else if (cmd == "help") {
            cout << "U, U', R, R', F, F', D, D', L, L', B, B' - ходы" << endl;
            cout << "scramble - перемешать" << endl;
            cout << "reset - собрать" << endl;
            cout << "quit - выход" << endl;
        } else if (cmd == "reset") {
            cube = Rubik4x4();
            cout << "Куб собран." << endl;
        } else if (cmd == "scramble") {
            auto seq = cube.scramble(50);
            cout << "Перемешивание: ";
            for (const auto& m : seq) cout << m << " ";
            cout << endl;
        } else {
            cube.applyMove(cmd);
        }
        cube.display();
    }
    return 0;
}

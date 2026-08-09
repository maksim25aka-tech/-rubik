// rubik.js
const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

class Rubik4x4 {
    constructor() {
        const colors = ['W', 'Y', 'R', 'O', 'B', 'G'];
        this.faces = [];
        for (let f=0; f<6; f++) {
            this.faces[f] = [];
            for (let i=0; i<4; i++) {
                this.faces[f][i] = [];
                for (let j=0; j<4; j++) {
                    this.faces[f][i][j] = colors[f];
                }
            }
        }
    }

    rotateFaceCW(idx) {
        const face = this.faces[idx];
        const newFace = Array.from({length:4}, () => Array(4));
        for (let i=0; i<4; i++) {
            for (let j=0; j<4; j++) {
                newFace[i][j] = face[3-j][i];
            }
        }
        this.faces[idx] = newFace;
    }

    rotateFaceCCW(idx) {
        const face = this.faces[idx];
        const newFace = Array.from({length:4}, () => Array(4));
        for (let i=0; i<4; i++) {
            for (let j=0; j<4; j++) {
                newFace[i][j] = face[j][3-i];
            }
        }
        this.faces[idx] = newFace;
    }

    applyMove(move) {
        switch (move) {
            case 'U': this.rotateFaceCW(0); break;
            case "U'": this.rotateFaceCCW(0); break;
            case 'R': this.rotateFaceCW(2); break;
            case "R'": this.rotateFaceCCW(2); break;
            case 'F': this.rotateFaceCW(4); break;
            case "F'": this.rotateFaceCCW(4); break;
            default: console.log('Неизвестный ход');
        }
    }

    scramble(moves=50) {
        const possible = ['U', "U'", 'R', "R'", 'F', "F'"];
        const seq = [];
        for (let i=0; i<moves; i++) {
            const m = possible[Math.floor(Math.random() * possible.length)];
            seq.push(m);
            this.applyMove(m);
        }
        return seq;
    }

    display() {
        console.log('\n--- Кубик Рубика 4x4 ---');
        const names = ['U', 'D', 'R', 'L', 'F', 'B'];
        for (let f=0; f<6; f++) {
            console.log(`${names[f]}:`);
            for (let i=0; i<4; i++) {
                console.log(this.faces[f][i].join(' '));
            }
            console.log();
        }
    }
}

const cube = new Rubik4x4();
console.log('Добро пожаловать в симулятор кубика Рубика 4x4!');
console.log('Команды: U, U\', R, R\', F, F\', scramble, reset, quit, help');

function prompt() {
    rl.question('> ', (cmd) => {
        cmd = cmd.trim();
        if (cmd === 'quit') {
            rl.close();
            return;
        } else if (cmd === 'help') {
            console.log("U, U', R, R', F, F', D, D', L, L', B, B' - ходы");
            console.log('scramble - перемешать');
            console.log('reset - собрать');
            console.log('quit - выход');
        } else if (cmd === 'reset') {
            Object.assign(cube, new Rubik4x4());
            console.log('Куб собран.');
        } else if (cmd === 'scramble') {
            const seq = cube.scramble(50);
            console.log('Перемешивание:', seq.join(' '));
        } else {
            cube.applyMove(cmd);
        }
        cube.display();
        prompt();
    });
}

prompt();

// rubik.rs
use std::io::{self, Write};
use rand::Rng;

type Face = [[char; 4]; 4];

struct Rubik4x4 {
    faces: [Face; 6],
}

impl Rubik4x4 {
    fn new() -> Self {
        let colors = ['W', 'Y', 'R', 'O', 'B', 'G'];
        let mut faces: [Face; 6] = [[[' '; 4]; 4]; 6];
        for f in 0..6 {
            for i in 0..4 {
                for j in 0..4 {
                    faces[f][i][j] = colors[f];
                }
            }
        }
        Rubik4x4 { faces }
    }

    fn rotate_face_cw(&mut self, face_idx: usize) {
        let mut new_face: Face = [[' '; 4]; 4];
        let face = self.faces[face_idx];
        for i in 0..4 {
            for j in 0..4 {
                new_face[i][j] = face[3-j][i];
            }
        }
        self.faces[face_idx] = new_face;
    }

    fn rotate_face_ccw(&mut self, face_idx: usize) {
        let mut new_face: Face = [[' '; 4]; 4];
        let face = self.faces[face_idx];
        for i in 0..4 {
            for j in 0..4 {
                new_face[i][j] = face[j][3-i];
            }
        }
        self.faces[face_idx] = new_face;
    }

    fn apply_move(&mut self, move_str: &str) {
        match move_str {
            "U" => self.rotate_face_cw(0),
            "U'" => self.rotate_face_ccw(0),
            "R" => self.rotate_face_cw(2),
            "R'" => self.rotate_face_ccw(2),
            "F" => self.rotate_face_cw(4),
            "F'" => self.rotate_face_ccw(4),
            _ => println!("Неизвестный ход"),
        }
    }

    fn scramble(&mut self, moves: usize) -> Vec<String> {
        let possible = ["U", "U'", "R", "R'", "F", "F'"];
        let mut rng = rand::thread_rng();
        let mut seq = Vec::new();
        for _ in 0..moves {
            let m = possible[rng.gen_range(0..possible.len())].to_string();
            seq.push(m.clone());
            self.apply_move(&m);
        }
        seq
    }

    fn display(&self) {
        println!("\n--- Кубик Рубика 4x4 ---");
        let names = ["U", "D", "R", "L", "F", "B"];
        for f in 0..6 {
            println!("{}:", names[f]);
            for i in 0..4 {
                for j in 0..4 {
                    print!("{} ", self.faces[f][i][j]);
                }
                println!();
            }
            println!();
        }
    }
}

fn main() {
    let mut cube = Rubik4x4::new();
    println!("Добро пожаловать в симулятор кубика Рубика 4x4!");
    println!("Команды: U, U', R, R', F, F', scramble, reset, quit, help");
    loop {
        print!("> ");
        io::stdout().flush().unwrap();
        let mut input = String::new();
        io::stdin().read_line(&mut input).unwrap();
        let cmd = input.trim();
        match cmd {
            "quit" => break,
            "help" => {
                println!("U, U', R, R', F, F', D, D', L, L', B, B' - ходы");
                println!("scramble - перемешать");
                println!("reset - собрать");
                println!("quit - выход");
            }
            "reset" => {
                cube = Rubik4x4::new();
                println!("Куб собран.");
            }
            "scramble" => {
                let seq = cube.scramble(50);
                println!("Перемешивание: {}", seq.join(" "));
            }
            _ => {
                cube.apply_move(cmd);
            }
        }
        cube.display();
    }
}

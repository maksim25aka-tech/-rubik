// rubik.cs
using System;
using System.Collections.Generic;

class Rubik4x4
{
    private char[,,] faces = new char[6,4,4];

    public Rubik4x4()
    {
        char[] colors = {'W', 'Y', 'R', 'O', 'B', 'G'};
        for (int f=0; f<6; f++)
            for (int i=0; i<4; i++)
                for (int j=0; j<4; j++)
                    faces[f,i,j] = colors[f];
    }

    private void RotateFaceCW(int idx)
    {
        char[,] newFace = new char[4,4];
        for (int i=0; i<4; i++)
            for (int j=0; j<4; j++)
                newFace[i,j] = faces[idx,3-j,i];
        for (int i=0; i<4; i++)
            for (int j=0; j<4; j++)
                faces[idx,i,j] = newFace[i,j];
    }

    private void RotateFaceCCW(int idx)
    {
        char[,] newFace = new char[4,4];
        for (int i=0; i<4; i++)
            for (int j=0; j<4; j++)
                newFace[i,j] = faces[idx,j,3-i];
        for (int i=0; i<4; i++)
            for (int j=0; j<4; j++)
                faces[idx,i,j] = newFace[i,j];
    }

    public void ApplyMove(string move)
    {
        switch (move)
        {
            case "U": RotateFaceCW(0); break;
            case "U'": RotateFaceCCW(0); break;
            case "R": RotateFaceCW(2); break;
            case "R'": RotateFaceCCW(2); break;
            case "F": RotateFaceCW(4); break;
            case "F'": RotateFaceCCW(4); break;
            default: Console.WriteLine("Неизвестный ход"); break;
        }
    }

    public List<string> Scramble(int moves)
    {
        string[] possible = {"U", "U'", "R", "R'", "F", "F'"};
        var seq = new List<string>();
        var rand = new Random();
        for (int i=0; i<moves; i++)
        {
            string m = possible[rand.Next(possible.Length)];
            seq.Add(m);
            ApplyMove(m);
        }
        return seq;
    }

    public void Display()
    {
        Console.WriteLine("\n--- Кубик Рубика 4x4 ---");
        string[] names = {"U", "D", "R", "L", "F", "B"};
        for (int f=0; f<6; f++)
        {
            Console.WriteLine(names[f] + ":");
            for (int i=0; i<4; i++)
            {
                for (int j=0; j<4; j++)
                    Console.Write(faces[f,i,j] + " ");
                Console.WriteLine();
            }
            Console.WriteLine();
        }
    }

    static void Main()
    {
        var cube = new Rubik4x4();
        Console.WriteLine("Добро пожаловать в симулятор кубика Рубика 4x4!");
        Console.WriteLine("Команды: U, U', R, R', F, F', scramble, reset, quit, help");
        while (true)
        {
            Console.Write("> ");
            string cmd = Console.ReadLine().Trim();
            if (cmd == "quit") break;
            else if (cmd == "help")
            {
                Console.WriteLine("U, U', R, R', F, F', D, D', L, L', B, B' - ходы");
                Console.WriteLine("scramble - перемешать");
                Console.WriteLine("reset - собрать");
                Console.WriteLine("quit - выход");
            }
            else if (cmd == "reset")
            {
                cube = new Rubik4x4();
                Console.WriteLine("Куб собран.");
            }
            else if (cmd == "scramble")
            {
                var seq = cube.Scramble(50);
                Console.WriteLine("Перемешивание: " + string.Join(" ", seq));
            }
            else
            {
                cube.ApplyMove(cmd);
            }
            cube.Display();
        }
    }
}

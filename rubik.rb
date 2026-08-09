# rubik.rb
class Rubik4x4
  def initialize
    colors = ['W', 'Y', 'R', 'O', 'B', 'G']
    @faces = Array.new(6) { Array.new(4) { Array.new(4, ' ') } }
    6.times do |f|
      4.times do |i|
        4.times do |j|
          @faces[f][i][j] = colors[f]
        end
      end
    end
  end

  def rotate_face_cw(idx)
    face = @faces[idx]
    new_face = Array.new(4) { Array.new(4) }
    4.times do |i|
      4.times do |j|
        new_face[i][j] = face[3-j][i]
      end
    end
    @faces[idx] = new_face
  end

  def rotate_face_ccw(idx)
    face = @faces[idx]
    new_face = Array.new(4) { Array.new(4) }
    4.times do |i|
      4.times do |j|
        new_face[i][j] = face[j][3-i]
      end
    end
    @faces[idx] = new_face
  end

  def apply_move(move)
    case move
    when 'U' then rotate_face_cw(0)
    when "U'" then rotate_face_ccw(0)
    when 'R' then rotate_face_cw(2)
    when "R'" then rotate_face_ccw(2)
    when 'F' then rotate_face_cw(4)
    when "F'" then rotate_face_ccw(4)
    else puts "Неизвестный ход"
    end
  end

  def scramble(moves=50)
    possible = ['U', "U'", 'R', "R'", 'F', "F'"]
    seq = []
    moves.times do
      m = possible.sample
      seq << m
      apply_move(m)
    end
    seq
  end

  def display
    puts "\n--- Кубик Рубика 4x4 ---"
    names = ['U', 'D', 'R', 'L', 'F', 'B']
    6.times do |f|
      puts "#{names[f]}:"
      4.times do |i|
        4.times do |j|
          print "#{@faces[f][i][j]} "
        end
        puts
      end
      puts
    end
  end
end

cube = Rubik4x4.new
puts "Добро пожаловать в симулятор кубика Рубика 4x4!"
puts "Команды: U, U', R, R', F, F', scramble, reset, quit, help"
loop do
  print "> "
  cmd = gets.chomp.strip
  case cmd
  when 'quit' then break
  when 'help'
    puts "U, U', R, R', F, F', D, D', L, L', B, B' - ходы"
    puts "scramble - перемешать"
    puts "reset - собрать"
    puts "quit - выход"
  when 'reset'
    cube = Rubik4x4.new
    puts "Куб собран."
  when 'scramble'
    seq = cube.scramble(50)
    puts "Перемешивание: #{seq.join(' ')}"
  else
    cube.apply_move(cmd)
  end
  cube.display
end

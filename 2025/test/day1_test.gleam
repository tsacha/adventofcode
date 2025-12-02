import day1
import gleeunit/should

pub fn day1_example_part1_test() {
  let input =
    "L68
L30
R48
L5
R60
L55
L1
L99
R14
L82"

  should.equal(day1.solve_part1(input), 3)
}

pub fn day1_example_part2_test() {
  let input =
    "L68
L30
R48
L5
R60
L55
L1
L99
R14
L82"

  should.equal(day1.solve_part2(input), 6)
}

pub fn day1_rotation_r_test() {
  should.equal(day1.dial(11, 8, 0), #(19, 0))
}

pub fn day1_rotation_l_test() {
  should.equal(day1.dial(19, -19, 0), #(0, 1))
}

pub fn day1_rotation_l_subzero_test() {
  should.equal(day1.dial(5, -10, 0), #(95, 1))
}

pub fn day1_rotation_l_two_turns_test() {
  should.equal(day1.dial(5, -110, 0), #(95, 2))
}

pub fn day1_off_by_one_test() {
  let input =
    "L51
R1
"

  should.equal(day1.solve_part2(input), 2)
}

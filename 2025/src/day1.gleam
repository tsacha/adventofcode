import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import utils

pub fn dial(state: Int, clicks: Int, zero_acc: Int) -> #(Int, Int) {
  case clicks {
    0 -> #(state, zero_acc)
    n if n < 0 -> {
      let new_state = state - 1 |> int.modulo(100) |> result.unwrap(0)
      let new_zero_acc = case new_state {
        0 -> zero_acc + 1
        _ -> zero_acc
      }
      dial(new_state, clicks + 1, new_zero_acc)
    }
    _ -> {
      let new_state = state + 1 |> int.modulo(100) |> result.unwrap(0)
      let new_zero_acc = case new_state {
        0 -> zero_acc + 1
        _ -> zero_acc
      }
      dial(new_state, clicks - 1, new_zero_acc)
    }
  }
}

pub fn do_solve(rotations: List(String), any_zero any_zero: Bool) -> Int {
  let #(_final_state, final_acc) =
    list.fold(rotations, #(50, 0), fn(state, rotation) {
      let #(current_state, current_acc) = state

      let clicks =
        rotation
        |> string.drop_start(up_to: 1)
        |> int.parse
        |> result.unwrap(0)

      let #(new_state, new_acc) = case string.starts_with(rotation, "R") {
        True -> dial(current_state, clicks, current_acc)
        False -> dial(current_state, clicks * -1, current_acc)
      }
      case any_zero, new_state {
        False, 0 -> #(new_state, current_acc + 1)
        False, _ -> #(new_state, current_acc)
        True, _ -> #(new_state, new_acc)
      }
    })

  final_acc
}

pub fn solve_part1(input: String) -> Int {
  input
  |> string.split(on: "\n")
  |> do_solve(any_zero: False)
}

pub fn solve_part2(input: String) -> Int {
  input
  |> string.split(on: "\n")
  |> do_solve(any_zero: True)
}

pub fn main() -> Nil {
  utils.get_puzzle_input(2025, 1)
  |> result.unwrap("")
  |> solve_part1
  |> int.to_string
  |> io.println

  utils.get_puzzle_input(2025, 1)
  |> result.unwrap("")
  |> solve_part2
  |> int.to_string
  |> io.println
}

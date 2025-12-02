import dot_env as dot
import dot_env/env
import gleam/http/request
import gleam/http/response
import gleam/httpc
import gleam/int
import gleam/result
import simplifile

pub fn get_puzzle_input(year: Int, day: Int) -> Result(String, httpc.HttpError) {
  let cache_file =
    "cache/" <> int.to_string(year) <> "-" <> int.to_string(day) <> ".txt"
  case simplifile.read(cache_file) {
    Ok(content) -> Ok(content)
    Error(_) -> {
      // Prepare a HTTP request record
      let assert Ok(base_req) =
        request.to(
          "https://adventofcode.com/"
          <> int.to_string(year)
          <> "/day/"
          <> int.to_string(day)
          <> "/input",
        )

      dot.new()
      |> dot.set_path(".env")
      |> dot.set_debug(False)
      |> dot.load
      let session = case env.get_string("AOC_SESSION") {
        Ok(value) -> value
        Error(_) -> echo "something went wrong"
      }

      let req =
        request.prepend_header(base_req, "cookie", "session=" <> session)

      // Send the HTTP request to the server
      use resp <- result.try(httpc.send(req))

      // We get a response record back
      assert resp.status == 200

      let content_type = response.get_header(resp, "content-type")
      assert content_type == Ok("text/plain")

      let _ = simplifile.write(cache_file, resp.body)
      Ok(resp.body)
    }
  }
}

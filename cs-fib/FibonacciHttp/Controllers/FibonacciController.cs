using Microsoft.AspNetCore.Mvc;

namespace FibonacciHttp.Controllers
{
  [ApiController]
  [Route("fib")]
  public class FibonacciController : ControllerBase
  {
    [HttpGet("v")]
    public string GetServiceName()
    {
      return Fibonacci.Service.ServiceName();
    }

    [HttpGet("n/{num}")]
    public IActionResult GetNumber(int num)
    {
      if (num < 0)
      {
        return NotFound("");
      }

      var result = Fibonacci.Service.Fibonacci(num).ToString() + "\n";
      return Ok(result);
    }

    [HttpGet("s/{num?}")]
    public IActionResult GetSequence(int num = 10)
    {
      if (num <= 0)
      {
        num = 10;
      }

      var sequence = Fibonacci.Service.Sequence().Take(num);
      var result = string.Join("\n", sequence) + "\n";
      return Ok(result);
    }
  }
}

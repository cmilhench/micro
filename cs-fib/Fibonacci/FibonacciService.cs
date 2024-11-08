using System.Reflection;
using System.IO;

namespace Fibonacci
{
  public static class Service
  {
    public static string Name = "Fibonacci Service";
    public static string Revision
    {
      get
      {
        var assembly = Assembly.GetExecutingAssembly();
        var resourceName = "Fibonacci..version";

        using (Stream stream = assembly.GetManifestResourceStream(resourceName)!)
        using (StreamReader reader = new StreamReader(stream))
        {
          return reader.ReadToEnd();
        }
      }
    }

    public static string ServiceName()
    {
      return $"{Name} {Revision}";
    }

    public static int Fibonacci(int n)
    {
      if (n <= 1)
      {
        return n;
      }
      return Fibonacci(n - 1) + Fibonacci(n - 2);
    }

    public static IEnumerable<int> Sequence()
    {
      int x = 0, y = 1;
      while (true)
      {
        int temp = x + y;
        x = y;
        y = temp;
        yield return x;
      }
    }
  }
}

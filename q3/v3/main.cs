using System;
using System.Collections.Generic;
using System.IO;
using System.Text;

class Program
{
    static void Run(Stream input, Stream output)
    {
        var br = new StreamReader(input, Encoding.ASCII);
        var bw = new StreamWriter(output);
        bw.AutoFlush = true;

        int t = int.Parse(br.ReadLine());        
        while (t > 0)
        {
            t--;

            int n = int.Parse(br.ReadLine());
            var even = new Dictionary<int, int>(n);
            var odd = new Dictionary<int, int>(n);
            var all = new Dictionary<int, int>(n); 
            long count = 0;
            
            while (n > 0) {
                n--;
                
                var he = new HashCode();
                var ho = new HashCode();
                bool isOdd = false;
                bool hoExists = false;
                
                while (true) {
                    byte c = (byte)br.Read();
                    if (c == '\n') 
                        break;
                    
                    if (isOdd) {
                        isOdd = false;
                        ho.Add(c);
                        hoExists = true;
                    } else {
                        isOdd = true;
                        he.Add(c);
                    }
                }

                int v;
                int e = he.ToHashCode();
                if (even.TryGetValue(e, out v)) {
                    even[e] = v + 1;
                    count += v;
                } else {
                    even.Add(e, 1);
                }

                if (!hoExists)
                    continue;

                int o = ho.ToHashCode();
                if (odd.TryGetValue(o, out v)) {
                    odd[o] = v + 1;
                    count += v;
                } else {
                    odd.Add(o, 1);
                }

                int s = e ^ o;
                if (all.TryGetValue(s, out v)) {
                    all[s] = v + 1;
                    count -= v;
                } else {
                    all.Add(s, 1);
                }
            }

            bw.WriteLine(count);
        }
    }

    static void Main(string[] args)
    {
        // DateTime start = DateTime.Now;
        Run(Console.OpenStandardInput(), Console.OpenStandardOutput());
        // Console.WriteLine($"{(DateTime.Now - start).TotalMilliseconds}");
    }
}

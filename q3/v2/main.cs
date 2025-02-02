using System;
using System.Collections.Generic;
using System.IO;
using System.Text;

class Program
{
    static long Solve(string[] ss)
    {
        var even = new Dictionary<string, int>(ss.Length);
        var odd = new Dictionary<string, int>(ss.Length);
        var all = new Dictionary<string, int>(ss.Length); 
        long count = 0;

        for (int i = 0; i < ss.Length; i++)
        {
            string s = ss[i];
            string e = ExtractEvenChars(s);
            string o = ExtractOddChars(s);

            if (!even.ContainsKey(e)) {
                even.Add(e, 1);
            } else {
                count += even[e];
                even[e]++;
            }

            if (string.IsNullOrEmpty(o))
                continue;

            if (!odd.ContainsKey(o)) {
                odd.Add(o, 1);
            } else {
                count += odd[o];
                odd[o]++;
            }
            
            if (!all.ContainsKey(s)) {
                all.Add(s, 1);
            } else {
                count -= all[s];
                all[s]++;
            }
        }

        return count;
    }

    static string ExtractEvenChars(string s)
    {
        var sb = new StringBuilder(s.Length / 2 + 1);
        for (int i = 0; i < s.Length; i += 2)
        {
            sb.Append(s[i]);
        }
        return sb.ToString();
    }

    static string ExtractOddChars(string s)
    {
        var sb = new StringBuilder(s.Length / 2 + 1);
        for (int i = 1; i < s.Length; i += 2)
        {
            sb.Append(s[i]);
        }
        return sb.ToString();
    }

    static void Run(Stream input, Stream output)
    {
        var br = new StreamReader(input, bufferSize: 4096);
        var bw = new StreamWriter(output, bufferSize: 4096);
        bw.AutoFlush = true;

        int t = int.Parse(br.ReadLine());

        for (int i = 1; i <= t; i++)
        {
            int n = int.Parse(br.ReadLine());

            string[] ss = new string[n];
            for (int j = 0; j < n; j++)
            {
                string s = br.ReadLine();
                ss[j] = s.Trim();
            }

            long ans = Solve(ss);
            bw.WriteLine(ans);
        }
    }

    static void Main(string[] args)
    {
        Run(Console.OpenStandardInput(), Console.OpenStandardOutput());
    }
}

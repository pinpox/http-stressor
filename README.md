# http-stressor
A tool to run stress tests on HTTP services

**WARNING:** Use this tool only at your own risk, agaist your own servers or
those you have permission to test. It may result in a Denial-of-Service!

NO WARRENTY whatsoever.

# Usage

```
usage: main [<flags>] <url>

Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
  -v, --verbose            Verbose mode, shows requests
      --worker-requests=0  number of requests per worker, use 0 for infinite
      --workers=1000       number of workers

Args:
  <url>  URL to test

```

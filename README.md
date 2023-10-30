# Network Programming with Go

## About

This is a collection of code examples ("tests") from the book <i>Network Programming with Go</i> by Adam Woodbeck.

## Run

- Run all tests from the repository directory: `go test -timeout 300s -race -bench=. ./...`.
- Run specific test, for example, `listen_test.go`: `go test -v -timeout 300s -race -bench=. ./ch03/listen_test.go`

## File Structure

- `ch03` Reliable TCP Data Streams
- `ch04` Sending TCP Data
- `ch05` Unreliable UDP Communication

## Notes on General Network Service Metrics

The following notes are quoted from a different book, <i>Devops for the Desparate</i> by Bradley Smith.

<i>Golden Signals</i>:
- Latency: the time it takes for a service to process a request.
- Traffic: how many requests an application is receiving.
- Errors: the number of errors an application is reporting.
- Saturation: how full a service is. For example, measure CPU usage to determine how much headroom is left on the system before the application or host becomes slow or unresponsive.

<i>RED</i>:
- Rate: the number of requests per second a service is receiving.
- Error: the number of failed requests per second that the service encounters.
- Duration: the amount of time it takes to serve a request, or how long it takes to return the data requested from your service to the client.

<i>USE</i>:
- Utilization: the average time the resource is busy doing work.
- Saturation: the extra work the system could not get to.
- Errors: the number of errors a system is having.

## Notes on Useful Linux Commands for Troubleshooting

The following notes are quoted from a different book, <i>Devops for the Desparate</i> by Bradley Smith.

<i>High load average</i>:
- `uptime`.
- `top`.
- For more information about particular processes: `vmstat`, `strace`, `lsof`.

<i>High memory usage</i>:
- `free -hm` \<flags: human-readable, mebibyte\>.
- `vmstat 1 5` \<parameter: delay in seconds\> \<parameter: count\>.
- `ps -efly --sort=rss | head` \<flags: show all processes in long format, sort by resident set size (amount of non-swappable physical memory a process uses)\> \<pipe: `head` command (displays first ten lines by default)\>.

<i>High iowait</i>:
- `iostat -xz 1 20` \<flags: show active devices with extended stat format\> \<parameter: delay in seconds\> \<parameter: count\>.
- `iotop -oPab` \<flags: show processes performing I/O with accumulative stats in batch mode\>.

<i>Hotname resolution failure</i>:
- `/etc/resolv.conf`.
- `resolvectl dns`.
- `dig @<upstream dns ip> <hostname>`.

<i>Out of disk space</i>:
- `df -h` \<flags: human-readable\>
- `find / -type f -size +100M -exec du -ah {} + | sort -hr | head` \<flags: type: file, size: more than 100MB, execute: file size on disk in human-readable format\> \<pipe: `sort` command (sorts descending) with human-readable flag\> \<pipe: `head` command (displays first ten lines by default)\>
- `lsof /var/log/<filename>.log` (note: the example from the book is that a large log file caused the out-of-disk-space problem, hence why the `lsof` command is being used to get info on an open log file and the following `logrotate` command is recommended).
- `logrotate` and it's config file `/etc/logrotate.d/`.

<i>Connection refused/timeout</i>:
- `curl <url>`.
- `ss -l -n -p | grep <port>` \<flags: shows listening sockets, do not resolve service names, shows the process using the socket\> \<pipe: `grep` command to filter for specific port\> (note: "ss" stands for "socket statistics").
- `tcpdump -ni any tcp port <port number>` \<flags: do not resolge host or port names, specifies the network interface (`any` means listen on all interfaces), specifies type of packets (tcp) and which port number\>.

<i>Searching logs</i>:
- `systemd` and `journal` services.
  - `journalctl -r` \<flags: reverse order (so newest is on top)\>.
  - `journalctl -r --since "2 hours ago"`.
  - `journalctl -r -u ssh` \<flags: filter on service name ("u" stands for "unit")\>.
  - `journalctl -r -u ssh -p err` \<flags: specify priority level\>.
  - `journalctl -r -u ssh -g "session opened"` \<flags: match on regular expression (similar to piping to `grep`)\>.
- Parsing logs:
  - `grep <string or pattern to match> <filename or stream>`.
    - `grep "10.0.2.33" /var/log/syslog`.
    - `grep -B 5 "user NOT in sudoers" /var/log/auth.log` \<flags: show \<number\> of lines before matched line\>.
  - `awk <string or pattern to match> <filename or stream>` similar to `grep` but can filter on columns.
    - `awk '{print $1}' /var/log/nginx/access.log` \<parameter: print first column\>.
    - `awk '($9 ~ /500/)' /var/log/nginx/access.log` meaning: search for all HTTP 500 response codes (usually in the ninth column) in the Nginx access log file. Inside the parantheses, the tilde (~) is a field number that tells `awk` to apply the search pattern only to a specific column.
    - `awk '($9 ~ /404/) {if (/POST/) print}' /var/log/nginx/access.log` similar to the one above: match the value at column $9 to the number 404, then pass an `if` block that states, "if the line from the column $9 match contains the word POST anywhere in it, pring that whole log line."
    - Another similar example, this time using the pipe symbol as a logical OR operator: `awk '($9 ~ /401|403/)' /var/log/nginx/access.log`.
- Common log files on linux systems:
  - `/var/log/syslog`
  - `/var/log/auth.log`
  - `/var/log/kern.log`
  - `/var/log/dmesg`



<i>Miscellaneous</i>:
- `strace -s 128 -p 19419` \<flags: message output size (128 bytes), specifies a process ID or PID\>, traces system calls and signals (note: this command can be very verbose and may cause performance issues).
  - `strace -f -c` \<flags: forked: follow any new processes created, summary flag produces on-going overview of what system calls the process is using\>.
  - `strace -p 28485 -e openat -o mytracefile.txt` \<flags: specifies a process ID or PID, specifies a particular system call, output to \<filename\>\>.
  - Similar to `strace`: `ltrace` reports dynamic library calls, `dtrace` can trace kernel-level issues. 
- `nmap` testing network connections.
- `iptables` and it's helpful wrapper `ufw` for allowing, blocking, and filtering connections.
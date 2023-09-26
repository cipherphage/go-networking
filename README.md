# Network Programming with Go

## About

This is a collection of code examples from the book <i>Network Programming with Go</i> by Adam Woodbeck.

## Notes on Service Metrics

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
- `uptime`
- `top`
- For more information about particular processes: `vmstat`, `strace`, `lsof`

<i>High memory usage</i>:
- `free -hm` \<flags: human-readable, mebibyte\>
- `vmstat 1 5` \<parameter: delay in seconds\> \<parameter: count\>
- `ps -efly --sort=rss | head` \<flags: show all processes in long format, sort by resident set size (amount of non-swappable physical memory a process uses)\> \<pipe: `head` command (displays first ten lines by default)\>

<i>High iowait</i>:
- `iostat -xz 1 20` \<flags: show active devices with extended stat format\> \<parameter: delay in seconds\> \<parameter: count\>
- `iotop -oPab` \<flags: show processes performing I/O with accumulative stats in batch mode\>

<i>Hotname resolution failure</i>:
- `/etc/resolv.conf`
- `resolvectl dns`
- `dig @<upstream dns ip> <hostname>`

<i>Out of disk space</i>:
- `df -h` \<flags: human-readable\>
- `find / -type f -size +100M -exec du -ah {} + | sort -hr | head` \<flags: type: file, size: more than 100MB, execute: file size on disk in human-readable format\> \<pipe: `sort` command (sorts descending) with human-readable flag\> \<pipe: `head` command (displays first ten lines by default)\>
- `lsof /var/log/<filename>.log` (note: the example from the book is that a large log file caused the out-of-disk-space problem, hence why the `lsof` command is being used to get info on an open log file and the following `logrotate` command is used).
- `logrotate` and it's config file `/etc/logrotate.d/`

<i>Connection refused/timeout</i>:
- `curl <url>`
- `ss -l -n -p | grep <port>` \<flags: shows listening sockets, do not resolve service names, shows the process using the socket\> \<pipe: `grep` command to filter for specific port\> (note: "ss" stands for "socket statistics").
- `tcpdump -ni any tcp port <port number>` \<flags: do not resolge host or port names, specifies the network interface (`any` means listen on all interfaces), specifies type of packets (tcp) and which port number\>

- `strace`


<i>Testing network connections and allowing, blocking, filtering them</i>:
- `nmap`
- `iptables` and it's helpful wrapper `ufw`
# CloudFlare Whitelisted IP Scanner

A utility to help you scan and export cloudflare IPs to find the ones that are clean and not blocked within your internet provider.

## How does scan work?

1. Fetch the list of Cloudflare IPv4 Networks from "https://cloudflare.com/ips-v4".
2. For each network in cloudflare networks:

   1. Convert network to list of IPs
   2. For each ip in network IPs:

      1. Send an HTTPS request to the IP with "www.cloudflare.com" as SNI

         - Succeeded: IP is Clean!
         - Failed: IP is Dirty or Offline.

We store the result of scans within boltDb on filesystem (`scan.db`).

## How does export work?

The export command relies on the scan result from previous scans.
It returns a list of clean IPs, each IP takes one line in the output.

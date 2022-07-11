# Black Hat Go: Notes and code samples

## Disclaimer

Don't use this to port scan arbitrary hosts (you can freely use `scanme.nmap.org` to play). Your ISP WILL block you. You have been warned. Also, check the license for this code. TL;DR If you fuck up, it's your problem.

## Structure

This repo follows a typical Go command library structure. You 'll find the different commands under `cmd/`.

## Building and running

Run `make` and you 're set. You 'll get a bunch of executables under `bin/`.

## List of Commands

* `rude_scanner`: A simple connect port scanner, nothing elaborate. Nice parallelism pattern though.
* `echo_server`: A generic TCP server echoing back whatever you send it. Use telnet with it and be amazed.
* `tcp_proxy`: Proxies access to any site (currently `example.com`) to `:80`. Fully functional, you could use it to bypass a firewall by running it in an intermediate, "approved" server. 
* `command_proxy`: Pretty basic TCP server: Fires up a shell and gives remote access to anyone. Don't fuck around in production with this.
* `shodan`: A simple command to search for a host in [shodan.io](https://shodan.io) and return results.
* `metasploit`: A simple command to demonstrate how to interact with Metasploit's Meterpreter, assuming you 've managed to compromise a machine somewhere.

## Prerequisites

* To run the `shodan` command, you need to have the `SHODAN_API_KEY` var in your environment. 
  This implies you have an account with Shodan. You don't need to pay for executing a few example
  scans, but anything more elaborate and you need to cough up the cash (not cheap).

* To run the `metasploit` command, you need to:

    1. have the [Metasploit community edition](https://docs.metasploit.com/docs/using-metasploit/getting-started/nightly-installers.html) installed. 
    2. Follow the guide in the page to go through initial setup. 
    3. kickoff an rpc server in metasploit as follows:

       ```
       $ msfconsole
       [...]
       msf6> load msgrpc Pass=s3cr3t ServerHost=10.0.1.6
       ```
    4. Create two environment variables to hold your beer:
       
       ```
       $ export MSFHOST=10.0.1.6:55552
       $ export MSFPASS=s3cr3t
       ```
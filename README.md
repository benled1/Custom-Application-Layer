# What is this?
This project is for the purpose of learning the different application layer protocols.
Some of the protocols I will be implementing are:
- HTTP 
- FTP 
- SSH

Some more may be added in the future, but for now these protocols are the goal.

# Specs
- the user of this project should be able to launch either a client or a server via the CLI tool
- the CLI should provide additional options that then specify the protocol as well as args for that specific protocol

for example: launch a ssh client pointed at localhost:22
```
custom-application-layer -client -ssh 127.0.0.1 22
```
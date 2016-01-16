# trapforwarder

## Quick start

Add the following lines to /etc/snmp/snmptrapd.conf:

    disableAuthorization yes
    traphandle default /path/to/trapforwarder log

Start snmpd in a terminal:

    $ snmptrapd -f -Le -Dsnmptrapd

Send a trap in another terminal:

    $ snmptrap -v 1  -c public localhost 1.3.6.1.4.1.5471.2  $(hostname) 6 20 "" SNMPv2-MIB::sysLocation.0 s "Just here"

Check syslog to see the result:

    $ journalctl -t trapforwarder

## Recipients

### Log
Send the raw trap to syslog

### Sensu
Send the trap to the sensu client. Work in progress ...

### Icinga
Not implmented yet.

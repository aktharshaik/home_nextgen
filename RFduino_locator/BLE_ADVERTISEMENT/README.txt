
Beacon technology —> No need to connect. Data sent is accessed during scan(by other device) itself.

/*
BLE --> 1.Advertise mode 2. Scan mode 3. Master device 4. Slave device

In order to establish a connection, one device has to be in advertising mode
(and allow for a connection) and the other device in Initiator mode. 
It is similar to the scanner mode but with the intention to establish a connection.
The initiator scans for a desirable device-advertising packet and consequently 
sends a connection request. Once a connection is established, the initiator 
assumes the role of master device and the advertiser becomes a Slave device. 
Slave devices may have only one connection at a time, while master devices may 
have multiple connections with different slave devices simultaneously.

Also, the active scanner may request up to 31 bytes of additional information 
from the advertiser if the advertising mode allows such an operation. 
It means that a sizable portion of data can be received from the advertising 
device even without establishing a connection. 
*/


HOW TO:

1. Changes advertisement data every 10 seconds. Check out for manufactured data using NRF android app for testing.
/*
Package stbl provides facilities for handling String Tables, as defined in
Valve's networking and dem file format.

String tables are replicated data containers with indexed entries that contain
a text string and optional binary user data (4 kB maximum). String tables are
created on the server and updates are replicated instantly and reliable to all
clients.

https://developer.valvesoftware.com/wiki/Networking_Events_%26_Messages#String_Tables
*/
package stbl

# domain-name-index
Domain name index responses whether domain name in it or not with consideration of wildcard

## Installation
```bash
go get github.com/AntonPates/domain-name-index
```

## Validation and normalization of domain name
 No validation and normalization of domain name is done in this module. It is assumed that domain name is already validated and normalized.



## Wildcard domain name
A "wildcard domain name" is defined by having its initial (i.e.,
   leftmost or least significant) label be, in binary format:

      0000 0001 0010 1010 (binary) = 0x01 0x2a (hexadecimal)

   The first octet is the normal label type and length for a 1-octet-
   long label, and the second octet is the ASCII representation [RFC20]
   for the '*' character.

   https://www.rfc-editor.org/rfc/rfc4592#section-2.1.1

## TODO
 - Add more tests



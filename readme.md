# Work in progress

**TODO**
* HTML interface
* document build steps

# swfc
 
*Simple Web File Commander*

## About
A simple web file commander implemented in the Go language. 
It compiles into one single binary which bundles all HTML, javascript etc. 
The protocol for the API is JSON RPC 2.0 through HTTP posts.
You can define virtual file system roots to make certain parts of the file system accessible.  

## Dependencies
The only dependency required to build the application is the go framework https://golang.org/

## Configuration 
Before you can utilize the file commander, you have to define your (virtual) file system roots.
 

## Protocol

### JSON RPC 2.0
For complete specification of the JSON RPC protocol, please visit http://www.jsonrpc.org/specification

### List roots (df)
Lists available (virtual) file system roots by name  

Example of command:
```
{ 
  "jsonrpc":"2.0",
  "method": "df",
  "params": [],
  "id": "3"
}
```
Example of successful response:
```
{
    "jsonrpc":"2.0",
    "result":["incoming","temp"],
    "id":"3"
}
```
### Copy (cp)
Copy a file to an other file
    
Example of command:
```
{ 
  "jsonrpc":"2.0",
  "method": "cp",
  "params": ["private/animals/cat.jpg", "public/animals/cat.jpg"],
  "id": "3"
}
```
Example of successful response:
```
{
    "jsonrpc":"2.0",
    "result":null,
    "id":"3"
}
```
Example of unsuccessful response:
``` 
{
    "jsonrpc":"2.0",
    "id":"3",
    "error":
    {
        "code":-32603,
        "message":"open /temp/animals/cat.jpg: The system cannot find the file specified."
    }
}
```
### Delete (rm)
Deletes a file  

Example of command:
```
{ 
  "jsonrpc":"2.0",
  "method": "rm",
  "params": ["incoming/public/dog.jpg"],
  "id": "3"
}
```
### Move (mv)
Moves/renames a file or directory  
Note. A file cannot be moved between two phycically different drives  

Example of command:
```
{ 
  "jsonrpc":"2.0",
  "method": "mv",
  "params": ["incoming/public/dog.jpg", "public/animals/dog.jpg"],
  "id": "3"
}
```
Example of successful response:
```
{
    "jsonrpc":"2.0",
    "result":null,
    "id":"3"
}
```
### List files (ls)
Lists all files and directories for a certain path  

Example of command:
```
{ 
  "jsonrpc":"2.0",
  "method": "ls",
  "params": ["incoming/public"],
  "id": "3"
}
```
Example of successful response:
Note: Type d: Directory, Type f: File 
```
{
    "jsonrpc":"2.0",
    "result":[{"Type":"d","Name":"dir1"},{"Type":"f","Name":"file1.txt"},{"Type":"f","Name":"file2.txt"}]
}
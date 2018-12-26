<p align="center">
    <img src="https://user-images.githubusercontent.com/5860071/50458636-35c99780-096d-11e9-955a-cec7b82432c3.png" width="250px" border="0" />
    <br/>
    <a href="https://github.com/vrachieru/ws/releases/latest">
        <img src="https://img.shields.io/badge/version-1.0.0-brightgreen.svg?style=flat-square" alt="Version">
    </a>
    <a href="https://travis-ci.org/vrachieru/ws">
        <img src="https://img.shields.io/travis/vrachieru/ws.svg?style=flat-square" alt="Version">
    </a>
    <br/>
    WebSocket command line client
</p>

A simplistic tool for sending and receiving websocket messages from the command line.  

## Install

```
$ go get github.com/vrachieru/ws
```

## Usage

```
$ ws ws://echo.websocket.org
Trying ws://echo.websocket.org...
Connected to ws://echo.websocket.org.
Exit with CTRL+C.

> echo
< echo
> 
```

## License

MIT
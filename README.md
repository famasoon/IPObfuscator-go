# IPObfuscator-go
IPObfuscator-go is to convert the IP address to obfuscated forms.
This is the Go language implementation of [OsandaMalith/IPObfuscator](https://github.com/OsandaMalith/IPObfuscator)

```
go run .
Enter IP Address: 127.0.0.1
2130706433
0x7f.0x00.0x00.0x01
0177.00.00.01
0x000000007f.0x0000000000.0x0000000000.0x0000000001
00000000177.00000000000.00000000000.00000000001

0x7f.0x00.0x00.1
0x7f.0x00.0.1
0x7f.0.0.1

0177.00.00.1
0177.00.0.1
0177.0.0.1

0x7f.0x00.1
0177.00.1
0x7f.00.1
0x7f.1
0177.1
0x7f.0x0.00.01
0x7f.00.00.01
```

<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="http://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png" /></a><br /><span xmlns:dct="http://purl.org/dc/terms/" property="dct:title">IP Obfuscator</span> by <a xmlns:cc="http://creativecommons.org/ns#" href="http://osandamalith.wordpress.com" property="cc:attributionName" rel="cc:attributionURL">Osanda Malith Jayathissa</a> is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/">Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License</a>.<br />Based on a work at <a xmlns:dct="http://purl.org/dc/terms/" href="http://osandamalith.wordpress.com" rel="dct:source">https://github.com/OsandaMalith/IPObfuscator/</a>.
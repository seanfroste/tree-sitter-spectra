## SAVE

This statement specifies the file name for saving the calculation result.

```
[format] SAVE FILE = <char> [ TYPE=<char> TSTEP=<char> OPTION=<char> ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| FILE | char | file name for saving the calculation result |
| TYPE | char | type of specified file ( TEXT , DATA or RESTART) DATA specifies spectra_plot file (\*.dat) and RESTART specifies restart file (default:RESTART) |
| TSTEP | progression | the time step number to create the output data file, whose name is defined as the input filename except ".in" + "\_t" + TSTEP number +".dat" in more than 3 digits cases. In the other cases, TSTEP number is written as 3 digits (see notice). |
| OPTION | char | designation of saving file in repeating calculation by DEFINE statement( FINAL or ALL) FINAL or ALL denote the final data or all data to be saved, respectively. (default: ALL) |

```
[ ex. ] SAVE FILE=MOS-Tr-1.save OPTION=FINAL
        SAVE FILE=ccd-1.txt TYPE=TEXT
            DEFINE NAME=Vsub VALUE=10 VSTEP=1 NSTEP=10
        SAVE FILE=es-vsub-{@Vsub-10}.dat TYPE=DATA
        SAVE FILE=es-vsub-{%3.1f:@Vsub-10}.dat TYPE=DATA
        SAVE FILE=trans-1.dat TYPE=DATA TSTEP=(0, 5, 10, 15, 20)
```

> Notice:

1. In the case of "TYPE=DATA", the file name can include equations with user parameters in {}. The number format of the equation in {} can be designated by using C language format definition before ":" in {}. The default format is "%g".\
   By the last SAVE statement and DEFINE statement in the above example, 11 data output files, from "es-vsub-0.0 dat" to "es-vsub-5.0 dat", are created.
2. The above example with TSTEP=(0, 5, 10, 15, 20) creates the 5 output files whose file names are trans-1_t000 dat, trans-1_t005 dat, trans-1_t010 dat, trans-1_t015 dat, and trans-1_t020 dat.

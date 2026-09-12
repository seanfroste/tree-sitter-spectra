## DEFINE

This statement denotes the name and the value of a user defined parameter. The parameter has both a numeric value and a character string.

```
[format] DEFINE NAME = <char> VALUE=<real>, <series>, or <func>
    [ VSTEP=<real> NSTEP=<real> FILE=<char>
    CHARACTER=<char> or <char array>
    OPTIMIZE=<char> or <func> TARGET=<real> ERROR=<real> ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| NAME | char | parameter name (upper-case and lower-case are distinguished) "@" or "\$" is necessary before a parameter, when the parameter is quoted in the other statements. "@ + name" returns its numeric value and "\$ + name" returns its character string. \$INPUTFILE has been already defined as the input file name without the suffix of ".in" |
| VALUE | real, series, or function | numeric value of a parameter (default:0) The value changes as numeric series defined in the parenthesis. If "#" + the other parameter name + numeric series is denoted, the value changes as numeric series combined with the parameter (see example). |
| CHARACTER | char or char array | character string or its array of a parameter (default: NULL) The value changes as character string array defined in the parenthesis. If "#" + the other parameter name + character string array is denoted, the value changes as character string array combined with the parameter (see example). |
| VSTEP | real | increment value of the defined parameter in a loop (default: 0) |
| NSTEP | real | repeat time of a loop calculation (default: 0) |
| FILE | char | file name for define statement in which user defined parameter is described |
| OPTIMIZE | char or func | extracted parameter name or their function for optimization (see appendix L) |
| TARGET | real | target value of the extracted parameter |
| ERROR | real | acceptable difference between the target and the extracted parameter value |

```
[ex.]   DEFINE NAME=Width VALUE=5.5 CHAR=Gate width

        DEFINE NAME=Length VALUE=4.8 VSTEP=0.2 NSTEP=5

        DOPE xmin=0 xmax={@Width-0.5} ymin={@Length/4} ymax=@Length

        EXTRACT PARA=VE ENAME=GATE PARA=VMAX FILE=test-{@Width}-1.txt
            (file name is "test-5.5-1.txt" in the above case)

        EXTRACT PARA=VE ENAME=GATE PARA=VMAX FILE=test-{\$Width}-1.txt
            (file name is "test-Gate_width-1.txt" in the above case)

        DEFINE NAME=Size VALUE=(0.3 0.5 0.8 1.2 2.0)
            (loop calculation is executed changing Size value as 0.3, 0.5, 0.8, 1.2, and 2.0 in the above case)

        DEFINE NAME=Size VALUE=(0.3 0.5 0.8 1.2 2.0)

        DEFINE NAME=Width VALUE=#Size(0.1 1.0 0.4 1.8 3.0)
            In the above case, the combination of (Size, Width) changes as (0.3, 0.1), (0.5, 1.0), (0.8, 0.4), (1.2, 1.8), and (2.0, 3.0).

        DEFINE NAME=Size VALUE=(0.3 0.5 0.8 1.2)

        DEFINE NAME=File_name CHAR=#Size(abc.txt def.txt ghi.txt jkl.txt)
            In the above case, the combination of (Size, File_name) changes as (0.3, abc. txt), (0.5, def. txt), (0.8, ghi. txt), and (1.2, jkl. txt)

        DEFINE NAME=Vsub VALUE=20 VSTEP=3 NSTEP=8 OPTIMIZE=#Pwell_Barrier
            TARGET=0. 3 ERROR=0. 001

        EXTRACT FILE=optimize-1.vm PARA=@Vsub
            PARA=BARRIER AXIS=Z XMIN=3.5 ZMIN=0.3 ZMAX=5 NAME=Pwell_Barrier In the above case, Vsub is automatically adjusted to make the extracted parameter BARRIER named as Pwell_Barrier along Z-axis in the region of XMIN=3.5, ZMIN=0.3.

        DEFINE NAME=extfile CHARACTER={\$INPUTFILE}.ext
            When the input file name is filename.in, \$extfile is defined as filename.ext in the above case.
```

> Notice: The maximum number of DEFINE statement which include VSTEP or NSTEP parameter is 30.

DEFINE NAME=Value1 VALUE=1 VSTEP=1 NSTEP=2
DEFINE NAME=Value2 VALUE=2 VSTEP=2 NSTEP=3
DEFINE NAME=Value3 VALUE=3 VSTEP=3 NSTEP=4

The above DEFINE statements mean the following calculation loops.

```
DO 3000 I1=0, 2
  DO 2000 I2=0, 3
    DO 1000 I3=0, 4
      SPECTRA calculation
1000    Value3=Value3+3
2000  Value2=Value2+2
3000 Value1=Value1+1
```

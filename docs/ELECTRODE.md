## ELECTRODE

This statement defines electrode voltage and its coordinate. The shape of the electrode is box or triangle pillar.

```
[format] ELECTRODE XMIN = <real> YMIN = <real> ZMIN = <real> XMAX = <real>
    YMAX= <real> ZMAX = <real> VOLTAGE = <real> [ PATTERN <char>
    NAME = <char> MASK = <func> ANDMASK = <func> ANDNEGA = <func>
    XPLUS=<real> YPLUS = <real> XMULTIPLE = <real>
    YMULTIPLE = <real> WORK = <real> SUPREM4 = <char>
    RESISTANCE = <real> LOCOS=<char> TRIM=<func> NEGATIVE=<char>
    EXPAND = <real> &[] ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| XMIN | real | electrode edge minimum X-coordinate (unit: $\mu$ m) |
| YMIN | real | electrode edge minimum Y-coordinate (unit: $\mu$ m) |
| ZMIN | real | electrode edge minimum Z-coordinate (unit: $\mu$ m) |
| XMAX | real | electrode edge maximum X-coordinate (unit: $\mu$ m) |
| YMAX | real | electrode edge maximum Y-coordinate (unit: $\mu$ m) |
| ZMAX | real | electrode edge maximum Z-coordinate (unit: $\mu$ m) |
| VOLTAGE | real or func | electrode voltage(unit:Volts) In transient analysis, voltage source can be designated as a time dependent fuction (see Time dependent functions in Model). If FILE(file_name) is denoted as the function, pairs of time (sec) and voltage (Volts) are input from the file and a piece wise linear voltage source is assumed. |
| PATTERN | char | the direction from the hypotenuse center to the right-angle in the triangle pillar case X+Y+, X+Y-, X-Y+, X-Y-, Y+Z+, Y+Z-, Y-Z+, Y-Z-, X+Z+, X+Z-X-Z+, or X-Z- (if this parameter is not defined, box is assumed) |
| NAME | char | electrode name (default: "E" + ELECTRODE statement input order) |
| MASK | func | mask shape definition POLYGON(X1, Y1 ... Xn, Yn), FILE(file\_name, cell\_name, layer, data\_type) or MFILE(file\_name, layer\_name) can be denoted as the function. X1, Y1 ... Xn, Yn is coordinate progression. If the progression is not closed, it is automatically closed. GDSII format data or CADENCE CAD data can be used by FILE function, which denotes file name, cell name, layer number, and data type. Data type can be omitted. GDSII and CADENCE data are automatically distinguished. MENTOR CAD data can be used by MFILE function, which denotes file name, and layer name. In CADENCE and MENTOR cases the data file must be ASCII data converted from CAD data. Thickness of electrode is defined by ZMIN and ZMAX in all cases. Plural mask parameters can be designated in an ELECTRODE statement, where the mask data by OR logical operation between the designated MASK data and the mask data created before the MASK parameter is created. |
| ANDMASK | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the designated ANDMASK data and the mask data created before the ANDMASK parameter is created. |
| ANDNEGA | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the negative data of designated ANDNEGA data and the mask data created before the ANDNEGA parameter is created. |
| XPLUS | real | shift value of MASK X-coordinate (unit: $\mu$ m) (default: 0) |
| YPLUS | real | shift value of MASK Y-coordinate (unit: $\mu$ m) (default: 0) |
| XMULTIPLE | real | multiplication factor for MASK X-coordinate (default: 1) |
| YMULTIPLE | real | multiplication factor for MASK Y-coordinate (default: 1) |
| WORK | real | work function difference between electrode and silicon(unit:Volts) this parameter is valid in oxide, nitride, or vacuum (default:-0.55) |
| RESISTANCE | char | resistance between the electrode and the voltage source (unit:0hm) (default:0) |
| SUPREM4 | char | "Y"or"N". If "Y" is specified, electrode region is assumed to the box region defined by XMIN, YMIN, ZMIN, XMAX, YMAX, ZMAX, and PATTERN except SUPREM4 calculation region (default:"N") |
| LOCOS | char | "Y"or"N". If "Y" is specified, electrode region is assumed out of LOCOS field oxide which is defined in REGION statement. (default: N) |
| TRIM | func | trimming or periodical arrangement of mask pattern (RECTANGLE() or PERIODIC()) RECTANGLE(x0 y0 x1 y1) executes trimming with rectangle region of $x0 \le x \le x1$ and $y0 \le y \le y1$ . PERIODIC(x0 y0 x1 y1) executes periodic arrangement with boundary of x=x0, x=x1, y=y0, and y=y1. |
| NEGATIVE | char | Negative mask pattern designation ("Y" or "N") (default: "N") |
| EXPAND | real | Expansion (positive value) or shrinkage (negative value) for MASK polygons (unit:μm) (default:0) |
| & [ ] | operator | The order of MASK operation can be designated by puting MASK=, ANDMASK=, or ANDNEGA= into brackets [], where AND operation is performed in the case that "%" exists just before [ and OR operation is perfomed in the other cases |

```
[ ex. ] ELECTRODE XMI=2.0 YMI=3 ZMI=0 XMA=5.5 YMA=7.8 ZMA=0.1 V=5 PAT=X+Y+ NAME=Drain
        ELECTRODE XMI=1 YMI=2 ZMI=-0.5 XMA=5.5 YMA=7.8 ZMA=-0.1 NAME=Gate1
            V=PWL(002E-954E-956E-90) SUPREM4=Y
        ELECTRODE XMI=8.0 YMI=3 ZMI=0 XMA=8 YMA=3 ZMA=0.1 V=0
            NAME=CS RESIST=1E8
        ELECTRODE ZMIN=-0.2 ZMAX=-0.1 MASK=POLY(0, 4 3, 4 3, 7 8, 7 8, 9 0, 9) V=0
        ELECTRODE ZMIN=-0.3 ZMAX=-0.1 MASK=FILE(maskfile.txt, gate, 8) V=5
        ELECTRODE ZMIN=-0.3 ZMAX=-0.1 XPLUS=3 YPLUS=4
            MASK=FILE(maskfile.txt.drain.6.2) V=3.3
        ELECTRODE XMI=1 XMA=2 YMI=2 YMA=3 ZMI=-0.3 ZMA=-0.1 V=FILE(source.txt)
        ELECTRODE XMI=8.0 YMI=3 ZMI=0 XMA=8 YMA=3 ZMA=0.1 LOCOS=Y V=0
        ELECTRODE ZMIN=-0.3 ZMAX=-0.1 MASK=FILE(maskfile.txt,rg,9) V=0
            TRIM=PERIODIC(0.2 0.3 5.1 7.2)
        ELECTRODE ZMIN=-0.4 ZMAX=-0.1 MASK=FILE(maskfile.txt,cs,7) V=0
            NEGA=Y
        ELECTRODE ZMIN=-0.4 ZMAX=-0.1 V=3.3
            MASK=FILE (maskfile.gds, pixel, 4)
            MASK=FILE (maskfile.gds, pixel, 5)
            ANDMASK=FILE (maskfile.gds, pixel, 6)
            ANDNEGA=FILE (maskfile.gdst, pixel, 7)
        ELECTRODE ZMIN=-0.2 ZMAX=-0.05 V=0
            MASK=FILE (maskfile.gds, gate, 8) EXPAND=-0.1
        ELECTRODE ZMIN=-0.3 ZMAX=-0.1 V=2
            MASK=FILE (maskfile.txt, pixel, 4)
            &[MASK=FILE(maskfile.txt,pixel,5)
            ANDMASK=FILE (maskfile.txt, pixel, 6)
            ANDNEGA=FILE (maskfile.txt, pixel, 7)]
```

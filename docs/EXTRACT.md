## EXTRACT

This statement specifies extracted parameter name and output file name. The parameters are extracted to the specified file in every bias step or every time step, if BIAS statement is used or time parameter is specified in convergence statement. At most 100 parameters can be extracted to one file.

```
[format] EXTRACT PARAMETER=<char> FILE=<char> [ ENAME=<char> XMIN=<real>
    YMIN=<real> ZMIN=<real> XMAX=<real> YMAX=<real>
    ZMAX=<real> AXIS=<char> PATTERN=<char> OPTION=<char>
    MASK=<func> ANDMASK=<func> ANDNEGA=<func> NAME=<char>
    XPLUS=<real> YPLUS=<real> VALUE=<real> PEAK=<char>
    RADIUS=<real> TRIM=<func> STEP=<char> MATERIAL=<char>
    XMULTIPLE=<real> YMULTIPLE=<real> WIDTH=<real>
    COORDINATE=<func> DELIMITER=<char> ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| PARAMETER | char | extracted parameter name. "ELECTRON", "HOLE", "VMIN", "VMAX", "NFERMI", "PFERMI", "NCURRENT", "PCURRENT", "NSADDLE", "PSADDLE", "EMAX, "EMIN", "BARRIER", "TIME", "GII", "GRCURRENT", "VELECTRODE", "QELECTRODE", "TTRANSIT", "NGRCURRENT", "PGRCURRENT", "CNMAX", "CPMAX", "XVMIN", "YVMIN", "ZVMIN", "XVMAX", "YVMAX", "ZVMAX", "XEMIN", "YEMIN", "ZEMIN", "XEMAX", "YEMAX", "ZEMAX", "XCNMAX", "YCNMAX", "ZCNMAX", "XCPMAX", "YCNPMAX", "ZCPMAX", "XNSADDLE", "YNSADDLE", "ZNSADDLE", "XPSADDLE", "YPSADDLE", "ZPSADDLE", "LAMBDA", "THETA", "PHI", "NMAX", "XNMAX", "YNMAX", "ZNMAX", "NMIN", "XNMIN", "YNMIN", "ZNMIN", "PMAX", "XPMAX", "YPMAX", "ZPMAX", "PMIN", "XPMIN", "YPMIN", "ZPMIN", "ZPMIN", "BARRIER", "NBARRIER", "PBARRIER", "NBMAX", "XNBMAX", "YNBMAX", "ZNBMAX", "NBMIN", "XNBMIN", "YNBMIN", "ZNBMIN", "PBMAX", "XPBMAX", "YPBMAX", "ZPBMAX", "PBMIN", "XPBMIN", "YPBMIN", "ZPBMIN", "NFMAX", "XNFMAX", "YNFMAX", "ZNFMAX", "PFMAX", "XPFMAX", "YPFMAX", "ZPFMAX", "NFMIN", "XNFMIN", "YNFMIN", "ZNFMIN", "PFMIN", "XPFMIN", "YPFMIN", "ZPFMIN", "DONOR", "ACCEPTOR", "NMAX", "PMAX", "EVOLUME", "VALUE" or user defined parameter beginning from "@" or "\$\$" can be specified @ + user parameter name: the numerical value of the user parameter \$\$ + user parameter name: the character string of the user parameter |
| FILE | char | output file name(not more than 72 characters) File name can include equations with user parameters in {}. The number format of the equation in {} can be designated by using C language format before ":" in {}. The default format is "%g". |
| ENAME | char | electrode name(substrate electrode name is "SUB") |
| XMIN | real | minimum X-coordinate of the specified region (unit: $\mu$ m) (default:0) |
| YMIN | real | minimum Y-coordinate of the specified region (unit: $\mu$ m) (default:0) |
| ZMIN | real | minimum Z-coordinate of the specified region (unit: $\mu$ m) (default:0) |
| XMAX | real | maximum X-coordinate of the specified region (unit: $\mu$ m) (default: XMAX of calculation region) |
| YMAX | real | maximum Y-coordinate of the specified region (unit: $\mu$ m) (default:YMAX of calculation region) |
| ZMAX | real | maximum Z-coordinate of the specified region (unit: $\mu$ m) (default:ZMAX of calculation region) |
| AXIS | char | Axis for NSADDLE, PSADDLE, XNSADDLE, YNSADDLE, ZNSADDLE, XPSADDLE, YPSADDLE, ZPSADDLE, EMIN, or BARRIER calculation ("X", "Y", "Z", "-X", "-Y", "-Z", "X+Y+", "X-Y+", "X+Y-", "X-Y-", "X+Z+", "X-Z+", "X+Z-", "X-Z-", "Y+Z+", "Y-Z+", "Y+Z-", "Y-Z-", or "N", default: N) |
| PATTERN or AXIS | char | the direction from the hypotenuse center to the right-angle in the triangle pillar case for ELECTRON, HOLE, VMAX, and VMIN X+Y+, X+Y-, X-Y+, X-Y-, Y+Z+, Y+Z-, Y-Z+, Y-Z-, X+Z+, X+Z-X-Z+, or X-Z- (if this parameter is not defined, box is assumed) |
| OPTION | char | designation of output style for FLOAT=N or P denoted in BIAS statement and decimal places for extracted parameter values (FINAL, PRECISE, NEWBARRIER or NONE) (default: NONE) In OPTION=FINAL case, only final converged value (zero current) is output to the extract file. In OPTION=PRECISE case, the decimal places for extracted parameters are increased from 6 (default) to 12. In OPTION=NEWBARRIER case, NBARRIER and PBARRIER are calculated by the maximum and minimum potential along the specified axis even in the case without peak or valley, where the maximum potential must exist before the minimum one along the axis for NBARRIER and the minimum potential must exist before the maximum one along the axis for PBARRIER |
| MASK | func | mask shape definition for ELECTRON, HOLE, VMAX, and VMIN POLYGON(X1, Y1 ... Xn, Yn), FILE(file_name, cell_name, layer, data_type) or MFILE(file_name, layer_name) can be denoted as the function. X1, Y1 ... Xn, Yn is coordinate progression. If the progression is not closed, it is automatically closed. GDSII format data or CADENCE CAD data can be used by FILE function, which denotes file name, cell name, layer number, and data type. Data type can be omitted. GDSII and CADENCE data are automatically distinguished. MENTOR CAD data can be used by MFILE function, which denotes file name, and layer name. In CADENCE and MENTOR cases the data file must be ASCII data converted from CAD data. Plural mask parameters can be designated for an extracted parameter, where the mask data by OR logical operation between the designated MASK data and the mask data created before the MASK parameter is created. |
| ANDMASK | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the designated ANDMASK data and the mask data created before the ANDMASK parameter is created. |
| ANDNEGA | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the negative data of designated ANDNEGA data and the mask data created before the ANDNEGA parameter is created. |
| TRIM | func | trimming or periodical arrangement of mask pattern (RECTANGLE() or PERIODIC()) RECTANGLE(x0 y0 x1 y1) executes trimming with rectangle region of $x0 \le x \le x1$ and $y0 \le y \le y1$. PERIODIC(x0 y0 x1 y1) executes periodic arrangement with boundary of x=x0, x=x1, y=y0, and y=y1. |
| NAME | char | MASK name for MASK plot (default: "X" + EXTRACT statement input order + "-" + PARAMETER input order) |
| XPLUS | real | shift value of MASK X-coordinate (unit: $\mu$ m) (default: 0) |
| YPLUS | real | shift value of MASK Y-coordinate (unit: $\mu$ m) (default: 0) |
| XMULTIPLE | real | multiplication factor for MASK X-coordinate (default: 1) |
| YMULTIPLE | real | multiplication factor for MASK Y-coordinate (default: 1) |
| VALUE | real | the absolute value of electric field in the case of PARAMAETER=EVOLUME (unit: V/cm) and the numerical value in the case of PARAMTER=VALUE, where the numerical operation including functions and user parameters can be designated in {} (see notice (5)) (default: 0) |
| PEAK | char | elimination of edge point in the case of PARAMETER=NSADDLE, PSADDLE, XNSADDLE, YNSADDLE, ZNSADDLE, XPSADDLE, YPSADDLE, or ZPSADDLE (Y or N) (default: N) |
| MATERIAL | char | Material name of the extract region for EMAX, XEMAX, YEMAX, or ZEMAX ("SILICON", "OXIDE", "NITRIDE", "INSULATOR", "METAL", "ALUMINUM", "POLY", "VACUUM", "AIR" or "ALL") (default: SILICON) |
| RADIUS | real | radius of the cylindrical region along diagonal axis specified by AXIS=X+Y+Z+, X+Y+Z-, X+Y-Z+, X+Y-Z-, X-Y+Z+, X-Y+Z-, X-Y-Z+, or X-Y-Z- (defult: $0.5 \mu m$ ) |
| STEP | real | extract point number of the cylindrical region along diagonal axis specified by AXIS=X+Y+Z+, X+Y+Z-, X+Y-Z+, X+Y-Z-, X-Y+Z+, X-Y+Z-, X-Y-Z+, or X-Y-Z- (default: 20) |
| COORDINATE | func | By denoting PWL(A1, B1,..., An, Bn) as a function, piece-wise linear axis is defined. (A, B denotes X, Y or X, Z or Y, Z: the coordinate on the plane which is defined by AXIS parameter) |
| WIDTH | real | width along the direction to be perpendicular to the piece-wise linear axis specified by COORDINATE = PWL() (unit: $\mu$ m) (default: 1) |
| DELIMITER | char | delimiter among parameters (SPACE, COMMA, TAB, or NONE) (default: NONE) |

> Notice:

1. If "ELECTRON", "HOLE", "VMIN", "VMAX", "NFERMI", "PFERMI",\
   "GRCURRENT", "NGRCURRENT", "PGRCURRENT",\
   "BARRIER", "EMIN", "EMAX", "CNMAX", "CPMAX", "XVMIN", "YVMIN",\
   "ZVMIN", "XVMAX", "YVMAX", "ZVMAX", "XEMIN", "YEMIN", "ZEMIN", "XEMAX",\
   "YEMAX", "ZEMAX", "XCNMAX", "YCNMAX", "ZCNMAX", "XCPMAX", "YCPMAX",\
   "ZCPMAX", "NSADDLE", "PSADDLE", "XNSADDLE", "YNSADDLE", "ZNSADDLE",\
   "XPSADDLE", "YPSADDLE", "ZPSADDLE", "NMAX", "XNMAX", "YNMAX", "ZNMAX",\
   "NMIN", "XNMIN", "YNMIN", "ZNMIN", "PMAX", "XPMAX", "YPMAX", "ZPMAX",\
   "PMIN", "XPMIN", "YPMIN", "ZPMIN" or "TTRANSIT" is specified,\
   XMIN, YMIN, ZMIN, XMAX, YMAX, and ZMAX must be defined.
2. If NCURRENT, PCURRENT, VELECTRODE, or QELECTRODE is specified, ENAME must be defined.
3. If NSADDLE, PSADDLE, XNSADDLE, YNSADDLE, ZNSADDLE, XPSADDLE, YPSADDLE, ZPSADDLE, BARRIER or EMIN is specified, AXIS must be defined.
4. If MASK=FILE(...) or MASK=MFILE(...) is specified, the designate layer or data-type should consist of only one polygon or one rectangle data.
5. The extracted parameters can be used as the parameter values of Xmin, Xmax, Ymin, Ymax, Zmin, and Zmax by designating their NAME parameters following "#" in braces {}, where the parameter values must consist of not more than 80 characters in a single row.

```
[ ex. ] EXTRACT PARA=VE ENAME=DRAIN FILE=mostr.cur
            PARA=NCUR ENAME=DRAIN
            PARA=ELEC XMIN=5 YMIN=2 ZMIN=0 XMAX=8 YMAX=6 ZMAX=1
            PARA=NSAD XMIN=5 YMIN=2 ZMIN=0 XMAX=8 YMAX=6 ZMAX=1 AXIS=X
            PARA=EMIN XMIN=3 YMIN=0 ZMIN=0 XMAX=6 YMAX=8 ZMAX=2 AXIS=-X
            PARA=@user\_defined\_parameter
            PARA=\$\$user\_character\_string
            PARA=ELEC MASK=FILE(mask.txt, PD1, 2) ZMIN=0.2 ZMAX=2 NAME=pd1

        EXTRACT PARA=VE ENAME=DRAIN PARA=NC ENAME=DRAIN FILE=mostr-{%3.1f:@Vg}.nc
            PARA=EVOL XMIN=1 YMIN=2 ZMIN=0 XMAX=8 YMAX=6 ZMAX=3 VALUE=2E5
            PARA=NSADDLE AXIS=-X XMIN=1 XMAX=7 ZMAX=2 PEAK=Y
            PARA=PSADDLE AXIS=N XMIN=5 XMAX=8 ZMAX=0.3 ZMAX=3
            PARA=PBARRIER AXIS=X-Y-Z- XMIN=2 XMAX=6 YMIN=3 YMAX=5
                ZMAX=0.1 ZMAX=1 RADIUS=0.3 STEP=40
            PARA=EMAX XMIN=5 YMIN=2 ZMIN=0 XMAX=8 YMAX=6 ZMAX=3 MATE=OX
            PARA=NGR XMIN=1.0 YMIN=1.5 ZMIN=0.2 XMAX=1.8 YMAX=2.2 ZMAX=1

        EXTRACT FILE={\$INPUTFILE}.ext
            PARA=VMAX ZMAX=2
            MASK=FILE(ext-mask.txt,cis,1)
            MASK=FILE (ext-mask. txt, cis, 2)
            ANDMASK=FILE(ext-mask.txt, cis, 3)
            ANDNEGA=FILE(ext-mask.txt,cis,4)
        EXTRACT FILE={\$INPUTFILE}.ele OPTION=PRECISE DELIMIT=TAB
            PARA=@VFD PARA=ELEC ZMAX=2
            PARA=VMAX XMIN=0.5 XMAX=1.0 ZMIN=0.1 ZMAX=0.8 NAME=XVMAX 1
            PARA=NSADDLE AXIS=-X XMIN=0 XMAX={#XVMAX 1+0.1}
            PARA=VALUE VALUE={#Electron-1 \* 1E-3}
```

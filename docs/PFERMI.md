## PFERMI

This statement specifies the hole quasi-Fermi potential and the coordinate of box or triangle pillar region.

```
[format] PFERMI XMIN= <real> YMIN=<real> ZMIN=<real> XMAX=<real>
    YMAX=<real> ZMAX=<real> VOLTAGE=<real> [ PATTERN=<char>
    AXIS=<char> DV=<real> NAME=<char> MASK=<func> ANDMASK=<func>
    ANDNEGA=<func> TRIM=<func> XPLUS=<real> YPLUS=<real>
    XMULTIPLE= <real> YMULTIPLE=<real> EXPAND=<real> &[ ] ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| XMIN | real | specified region X-coordinate minimum(unit: $\mu$ m) |
| YMIN | real | specified region Y-coordinate minimum(unit: $\mu$ m) |
| ZMIN | real | specified region Z-coordinate minimum(unit: $\mu$ m) |
| XMAX | real | specified region X-coordinate maximum(unit: $\mu$ m) |
| YMAX | real | specified region Y-coordinate maximum(unit: $\mu$ m) |
| ZMAX | real | specified region Z-coordinate maximum(unit: $\mu$ m) |
| VOLTAGE | real | hole quasi-Fermi potential of the box region (unit:Volts) |
| AXIS | char | axis for the linear difference of quasi-Fermi potential (X, Y, or Z ) $$ |
| DV | real | linear difference of quasi-Fermi potential along The specified axis (unit:Volts) (default: 0) |
| PATTERN | char | the direction from the hypotenuse center to the right-angle in the triangle pillar case X+Y+, X+Y-, X-Y+, X-Y-, Y+Z+, Y+Z-, Y-Z+, Y-Z-, X+Z+, X+Z-X-Z+, or X-Z- (if this parameter is not defined, box is assumed) |
| NAME | char | hole quasi-Fermi potential region name (default: "P" + PFERMI statement input order) |
| MASK | func | mask shape definition POLYGON(X1, Y1 ... Xn, Yn), FILE(file_name, cell_name, layer, data_type) or MFILE(file_name, layer_name) can be denoted as the function. X1. Y1 ... Xn. Yn is coordinate progression. If the progression is not closed, it is automatically closed. GDSII format data or CADENCE CAD data can be used by FILE function, which denotes file name, cell name, layer number, and data type. Data type can be omitted. GDSII and CADENCE data are automatically distinguished. MENTOR CAD data can be used by MFILE function, which denotes file name, and layer name. In CADENCE and MENTOR cases the data file must be ASCII data converted from CAD Plural MASK parameters can be designated, where the mask data by OR logical operation between the designated MASK data and the mask data created before the MASK paramter is created. |
| ANDMASK | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the designated ANDMASK data and the mask data created before the ANDMASK parameter is created. |
| ANDNEGA | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the negative data of designated ANDNEGA data and the mask data created before the ANDNEGA parameter is created. |
| XPLUS | real | X coordinate shift value (unit: $\mu$ m) (default: 0) |
| YPLUS | real | Y coordinate shift value (unit: $\mu$ m) (default: 0) |
| XMULTIPLE | real | multiplication factor for MASK X-coordinate (default: 1) |
| YMULTIPLE | real | multiplication factor for MASK Y-coordinate (default: 1) |
| TRIM | func | trimming or periodical arrangement of mask pattern RECTANGLE() or PERIODIC() RECTANGLE(x0 y0 x1 y1) executes trimming with rectangle region of $x0 \le x \le x1$ and $y0 \le y \le y1$. PERIODIC(x0 y0 x1 y1) executes periodic arrangement with boundary of x=x0, x=x1, y=y0, and y=y1. |
| EXPAND | real | Expansion (positive value) or shrinkage (negative value) for MASK polygons (unit: μm) (default:0) |
| &[] | operator | The order of MASK operation can be designated by putting MASK=..., ANDMASK=..., or ANDNEGA=... into brackets [], where AND operation is performed in the case that "%" exists just before [ and OR operation is perfomed in the other cases |

> Notice:

PFERMI of the unspecified region is set to the minimum value of the PFERMI.

```
[ ex. ] PFERMI XMIN=2.0 YMIN=3 ZMIN=0 XMAX=5.5 YMAX=7.8 ZMAX=0.1 V=5 PAT=X-Y- NAME=DRAIN
        PFERMI MASK=FILE (mask.txt, PDP, 3) name=pd3 V=5
        PFERMI NAME=pdp ZMIN=0 ZMAX=10 V=3.3
            MASK=FILE (mask. txt, cell-1, 5)
            MASK=FILE (mask. txt, cell-1, 6)
            ANDMASK=FILE (mask. txt, cell-1, 7)
            ANDNEGA=FILE (mask. txt, cell-1, 5)
        PFERMI NAME=PD XMIN=3 XMAX=7 YMIN=2 YMAX=8 ZMIN=0 ZMAX=10 V=3.3
            AXIS=Z DV=10
        PFERMI ZMIN=0 ZMAX=2 V=-3
            MASK=FILE (mask. gds, cell-1, 9) EXPAND=0.2
        PFERMI NAME=pdp ZMIN=0 ZMAX=3 V=0
            MASK=FILE (mask. txt, cell-3, 2)
            &[ MASK=FILE(mask.txt,cell-3,3)
            ANDMASK=FILE (mask. txt, cell-3, 4)
            ANDNEGA=FILE (mask. txt, cell-3, 5) ]
```

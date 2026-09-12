## INTERFACE

This statement specifies the box region coordinate where the interface exists, the charge concentration, and the carrier generation recombination velocity of the interface.

```
[format] INTERFACE [ MATERIAL=<char> QF=<real> SN=<real> SP=<real>
    XMIN=<real> YMIN=<real> ZMIN=<real> XMAX=<real> YMAX=<real> ZMAX=<real>
    TRIM=<func> NEGATIVE=<char> CTRAP=<real> DTRAP=<real> TRAPMOB=<real>
    XPLUS=<real> YPLUS=<real> NAME=<char> MASK=<func> ANDMASK=<func>
    ANDNEGA=<func> XMULTIPLE=<real> YMULTIPLE=<real> EXPAND=<real> &[ ] ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| MATERIAL | . char | material name which formes the interface "SILICON", "OXIDE", "NITRIDE", "METAL", "ALUMINUM", "INSULATOR", "POLY", "VACUUM", "M1", "M2", "M3", "M4", "M5", "M6", "M7", or "M8" (default: SILICON and OXIDE) |
| QF | real | charge concentration of the interface (unit:cm ⁻² ) (default:0) |
| SN | real | electron generation recombination velocity (unit:cm/s) (default:0) |
| SP | real | hole generation recombination velocity (unit:cm/s) (default:0) |
| XMIN | real | box region minimum X-coordinate (unit: $\mu$ m) (default:0) |
| YMIN | real | box region minimum Y-coordinate (unit: $\mu$ m) (default:0) |
| ZMIN | real | box region minimum Z-coordinate (unit: $\mu$ m) (default:0) |
| XMAX | real | box region maximum X-coordinate (unit: $\mu$ m) (default:0) |
| YMAX | real | box region maximum Y-coordinate (unit: $\mu$ m) (default:0) |
| ZMAX | real | box region maximum Z-coordinate (unit: $\mu$ m) (default:0) |
| TRIM | func | trimming or periodical arrangement of mask pattern (RECTANGLE() or PERIODIC()) RECTANGLE(x0 v0 x1 v1) executes trimming with rectangle region of $x0 \le x \le x1$ and $y0 \le y \le y1$ . PERIODIC(x0 y0 x1 y1) executes periodic arrangement with boundary of x=x0, x=x1, v=v0, and y=y1. |
| CTRAP | real | trap density of Si-SiO2 interface (unit:cm⁻²) (default: 1E11) |
| DTRAP | real | thickness of Si-SiO2 interface trap region (unit: $\mu$ m) (default: 1E-4) |
| TRAPMOB | real | mobility along Z-axsis of Si-SiO2 interface trap region (unit: cm²/(Vsec)) (default: 1E-5) |
| NAME | char | MASK name for MASK plot (default: "I" + INTERFACE statement input order) |
| MASK | func | mask shape definition POLYGON(X1, Y1 ... Xn, Yn), FILE(file_name, cell_name, layer, data_type) or MFILE(file_name, layer name) can be denoted as the function. X1, Y1 ... Xn, Yn is coordinate progression. If the progression is not closed, it is automatically closed. GDSII format data or CADENCE CAD data can be used by FILE function, which denotes file name, cell name, layer number, and data type. Data type can be omitted. GDSII and CADENCE data are automatically distinguished. MENTOR CAD data can be used by MFILE function, which denotes file name, and layer name. In CADENCE and MENTOR cases the data file must be ASCII data converted from CAD data. Plural MASK parameters can be designated, where the mask data by OR logical operation between the designated MASK data and the mask data created before the MASK paramter is created. |
| ANDMASK | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the designated ANDMASK data and the mask data created before the ANDMASK parameter is created. |
| ANDNEGA | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the negative data of designated ANDNEGA data and the mask data created before the ANDNEGA parameter is created. |
| XPLUS | real | shift value of MASK X-coordinate (unit: $\mu$ m) (default: 0) |
| YPLUS | real | shift value of MASK Y-coordinate (unit: $\mu$ m) (default: 0) |
| XMULTIPLE | real | multiplication factor for MASK X-coordinate (default: 1) |
| YMULTIPLE | real | multiplication factor for MASK Y-coordinate (default: 1) |
| EXPAND | real | Expansion (positive value) or shrinkage (negative value) for MASK polygons (unit: $\mu$ m) (default:0) |
| & [ ] | operator | The order of MASK operation can be designated by puting MASK=, ANDMASK=, or ANDNEGA= into brackets [], where AND operation is performed in the case that "%" exists just before [ and OR operation is perfomed in the other cases |

```
[ ex. ] INTERFACE MATE=NI MATE=SI XMI=2 YMI=3 ZMI=1 XMA=8 YMA=8 ZMA=2
            QF=5E10 SN=50 SP=50
        INTERFACE MATE=OXIDE MATE=SI MASK=FILE(maskfile.txt,gate,8,3)
            ZMIN=-1 ZMAX=1 TRAPMOB=1E-6 CTRAP=1E12 DTRAP=1E-5
        INTERFACE MATE=SI ZMIN=-1 ZMAX=1 QF=8E10
            MASK=FILE(maskfile.txt, cis, 1, 3)
            ANDMASK=FILE(maskfile.txt,cis,2,3)
            ANDNEGA=FILE (maskfile.txt, cis, 2, 4)
            MASK=FILE(maskfile.txt, cis, 3, 3)
        INTERFACE MATE=M1 MATE=M2 ZMIN=-1 ZMAX=-0.5 QF=2E11
            MASK=FILE (maskfile.gds, cell-1, 1, 5) EXPAND=0.2
        INTERFACE MATE=M1 MATE=M2 ZMIN=-1 ZMAX=-0.5 QF=-1E11
            MASK=FILE (maskfile.txt, pixel, 4)
            &[ MASK=FILE(maskfile.txt,pixel,5)
            ANDMASK=FILE (maskfile.txt, pixel, 6)
            ANDNEGA=FILE (maskfile.txt, pixel, 7) ]
```

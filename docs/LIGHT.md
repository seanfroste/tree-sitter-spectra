## LIGHT

This statement specifies wave length, intensity, angle, illuminated region, etc. of incident light.

```
[format] LIGHT [ LAMBDA= <real> POWER=<real> or <func> THETA=<real>
    PHI=<real> XMIN=<real> YMIN=<real> XMAX=<real>
    YMAX=<real> YMAX=<real> ASI or ABSORPTION=<real> NSI=<char> or <real>
    NAME=<char> XREPEAT=<char> YREPEAT=<char> TYPE=<char>
    FILE=<char> SKIP=<real> NSTEP=<real> LMAX=<real>
    REVERSE=<char> MIRROR=<char> XPLUS=<real> YPLUS=<real>
    ZPLUS=<real> BACK=<char> MASK=<func> ANDMASK=<func>
    ANDNEGA=<func> TRIM=<func> XMULTIPLE=<real> YMULTIPLE=<real>
    EXPAND=<real> &[ ] ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| LAMBDA | real | wave length (unit:nm) (default:550) |
| PoWER | real or func | intensity(unit:W/cm²) (default:1) PWL, EXP, PULSE, or SIN can be denoted as a time dependent function (see Time dependent function in Model) (unit for function: sec and W/cm²) |
| THETA | real | incident angle against Z-axis (Z-axis direction is 0 degree) (unit: degree) 0 ≤ THETA ≤ 90 (default:0) |
| PHI | real | incident angle against X-axis on XY-plane (X-axis direction is 0 degree) (unit: degree) (default: 0) |
| XMIN | real | minimum X-coordinate of illuminated region or light generation current region for FILE designation (unit: $\mu$ m) (default:0) |
| YMIN | real | minimum Y-coordinate of illuminated region or input region for FILE designation (unit: $\mu$ m) (default:0) |
| ZMIN | real | minimum Z-coordinate of input region for FILE designation (unit: $\mu$ m) (default:0) |
| XMAX | real | maximum X-coordinate of illuminated region or input region for FILE designation (unit: $\mu$ m) (default: XMAX of calculation region) |
| YMAX | real | maximum Y-coordinate of illuminated region or input region for FILE designation (unit: $\mu$ m) (default: YMAX of calculation region) |
| ZMAX | real | maximum Z-coordinate of input region for FILE designation (unit: $\mu$ m) (default: ZMAX of calculation region) |
| ASIor ABSORPTΙΟΝ | char or real | absorption coefficient of Si ( TOCCATA, SPECTRA or value) (default: SPECTRA, see Table 1) In TOCCATA case, ASI defined in TOCCATA. |
| NSI | char or real | refractive index of Si ( TOCCATA, SPECTRA or value) (default: SPECTRA, see Table 2) In TOCCATA case, NSI defined in TOCCATA. |
| XREPEAT | char | repeat the illuminated region for X-direction ("Y" or "N") (default:"N") |
| YREPEAT | char | repeat the illuminated region for X-direction ("Y" or "N") (default:"N") |
| TYPE | char | type of TOCCATA, if TOCCATA result is used. ("WAVE", "RAY", or "FDTD") (default: auto-detect) |
| FILE | char | TOCCATA output file (\*.rtd) or 3-D light intensity text file name |
| SKIP | real | skip number for reading TOCCATA output file, if it has some wave length or incident angle data. (default:0) |
| NSTEP | real | repeat time of calculation, if wave length is a changing parameter. If FILE parameter is specified, NLIGHT parameter of TOCCATA should not be less than this parameter. In the other case, (LMAX-LAMBDA)/NSTEP is the wave length difference among each calculation step. (default:0) |
| LMAX | real | maximum wave length in the case of wave length changing calculation If FILE parameter is specified, this parameter is ignored. (unit:nm )(default:0) |
| XPLUS | real | X-coordinate shift value of TOCCATA and 3-D light intensity data (unit: $\mu$ m) (default: 0) |
| YPLUS | real | Y-coordinate shift value of TOCCATA and 3-D light intensity data (unit: $\mu$ m) (default: 0) |
| ZPLUS | real | Z-coordinate shift value of TOCCATA and 3-D light intensity data (unit: $\mu$ m) (default: 0) |
| XMULTIPLE | real | multiplication factor for MASK X-coordinate (default: 1) |
| YMULTIPLE | real | multiplication factor for MASK Y-coordinate (default: 1) |
| REVERSE | char | reverse light intensity distribution along the specified axis is assumed in the case of TYPE=WAVE ("X" or "Y") (default:none) |
| NAME | char | Name of the light statement for drawing the illuminated region (default: L + the number of the LIGHT statement) |
| BACK | char | back side illumination ("Y" or "N") (default: "N") |
| MASK | func | mask shape definition POLYGON(X1, Y1 Xn, Yn), FILE(file_name, cell_name, layer, data_type) or MFILE(file_name, layer_name) can be denoted as the function. X1, Y1 Xn, Yn is coordinate progression. If the progression is not closed, it is automatically closed. GDSII format data or CADENCE CAD data can be used by FILE function, which denotes file name, cell name, layer number, and data type. Data type can be omitted. GDSII and CADENCE data are automatically distinguished. MENTOR CAD data can be used by MFILE function, which denotes file name, and layer name. In CADENCE and MENTOR cases the data file must be ASCII data converted from CAD data. Plural MASK parameters can be designated, where the mask data by OR logical operation between the designated MASK data and the mask data created before the MASK paramter is created. |
| ANDMASK | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the designated ANDMASK data and the mask data created before the ANDMASK parameter is created. |
| ANDNEGA | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the negative data of designated ANDNEGA data and the mask data created before the ANDNEGA parameter is created. |
| TRIM | func | trimming or periodical arrangement of mask pattern (RECTANGLE() or PERIODIC()) RECTANGLE</a>(x0 y0 x1 y1) executes trimming with rectangle region of $x0 \le x \le x1$ and $y0 \le y \le y1$. PERIODIC(x0 y0 x1 y1) executes periodic arrangement with boundary of x=x0, x=x1, y=y0, and y=y1. |
| EXPAND | real | Expansion (positive value) or shrinkage (negative value) for MASK polygons (unit: $\mu$ m) (default:0) |
| &[ ] | operator | The order of MASK operation can be designated by puting MASK=..., ANDMASK=..., or ANDNEGA=... into brackets [], where AND operation is performed in the case that "%" exists just before [ and OR operation is performed in the other cases |

```
[ ex. ] LIGHT LA=400 P0=1E-3 XMI=4 YMI=2 XMAX=8 YMAX=9 THETA=10 PHI=90
        LIGHT FILE=toccata-1.rtd TYPE=WAVE POWER=1E-3 SKIP=3
        LIGHT FILE=toccata-2.rtd TYPE=WAVE REVERSE=X POWER=1E-2 NSTEP=15
        LIGHT LA=500 POWER=PWL(0, 0 1E-9, 0 2E-9, 1E-3 8E-9, 3E-3) XMI=4 YMI=2 XMAX=8 YMAX=9 BACK=Y
        LIGHT FILE=toccata-4.rtd ASI=T NSI=T XMI=1 YMI=2 XMAX=8 YMAX=6 ZMAX=3
        LIGHT LAMBDA=550 POW=1E-3 MASK=FILE (maskfile.txt,cell-1,5,1)
        LIGHT LAMBDA=550 POW=1E-3
            MASK=FILE (maskfile.txt, cell-1, 5, 1)
            MASK=FILE (maskfile.txt, cell-1, 6, 1)
            ANDMASK=FILE (maskfile.txt,cell-1,7,1)
            ANDNEGA=FILE (maskfile.txt,cell-1,8,1)
        LIGHT FILE=light power.txt Zmax=10 Power=PWL(0, 0 1E-9, 0 1E-9, 3E-3)
        LIGHT POW=1E-3 MASK=FILE (maskfile.gds, cell-1, 9, 1) EXPAND=0.2
        LIGHT LAMBDA=550 POW=1E-3
            MASK=FILE (maskfile.txt, pixel, 4)
            &[ MASK=FILE (maskfile.txt, pixel, 5)
            ANDMASK=FILE (maskfile.txt, pixel, 6)
            ANDNEGA=FILE (maskfile.txt, pixel, 7) ]
```

> Notice:

1. If FILE is specified, power and skip parameter can be defined, where the power value is multiplied to intensity.
2. In the case of 3-D light intensity text file, the file format must be as follows, which is same as 3-D impurity file format in DOPE statement. (one row consists of less than 1024 characters and at least one space is necessary between numbers)

   - [Format number 0]\
     1st row: 0 Nx Ny\
     (O is format number and Nx, Ny are mesh number along X and Y axis, respectively)\
     2nd∼Nx+1th row: X-coordinate (unit: \mu m, Nx)\
     Nx+2th~Nx+Ny+1th row: Y-coordinate (unit: \mu m, Ny)\
     Nx+Ny+2th row: X-coordinate Y-coordinate N1\
     (N1 is mesh number along depth axis of 1st mesh point on XY plane)\
     Nx+Ny+3th $\sim$ Nx+Ny+N1+2th row: Z-coordinate intensity (unit: W/cm²)\
     Mth row: X-coordinate Y-coordinate Nn\
     (Nn is mesh number along depth axis of nth mesh point on XY plane)\
     M+1th $\sim$ M+Nnth row: Z-coordinate intensity (unit: W/cm²)\
     where\
     n = Nx \* Ny\
     M = Nx + Ny + 1 + $\sum_{k=1}^n (N_k + 1)$
   - [Format number 1]\
      1st row: 1 Nx Ny Nz\
      (1 is format number and Nx, Ny, and Nz are total grid numbers along X, Y, and Z axes, respectively)
     2nd $\sim$ Nx+1th row: X-coordinate (unit: $\mu$ m, Nx)\
      Nx+2th $\sim$ Nx+Ny+1th row: Y-coordinate (unit: $\mu$ m, Ny)\
      Nx+Ny+2th $\sim$ Nx+Ny+Nz+1th row: Z-coordinate (unit: $\mu$ m, Nz)\
      Nx+Ny+Nz+2th $\sim$ Nx+Ny+Nz+Nx\*Ny\*Nz+1th row: Ip (light intensity, unit: W/cm^2)\
      .\
      .\
      .\
      The order of Ip is shown below by C language program.

     ```c
     for (k=0; k< Nz; k++) {
        for (j=0; j< Ny; j++) {
            for (i=0; i< N_X; i++) {
                printf("%e\n", E[i], E[i], E[i], E[k]);
            }
        }
     }
     ```

     where X[i], Y[j], and Z[k] denote i-th X, j-th Y, and k-th Z coordinates, respectively.

## REGION

This statement specifies the coordinate, the material, and carrier generation-recombination velocity of the box or triangle pillar region.

```
[format] REGION MATERIAL=<char> [ XMIN=<real> XMAX=<real> YMIN=<real>
    YMAX=<real> ZMIN=<real> or <series> ZMAX=<real> or <series>
    PATTERN=<char> XORIGIN=<real> YORIGIN=<real> AXIS=<char>
    FILE=<char> GR=<real> or <func> MASK=<func> ANDMASK=<func> ANDNEGA=<func> XPLUS=<real> YPLUS=<real> BEAK=<real> or <series>
    ZBEAK=<real> EDGE=<char> TN0=<real> TP0=<real>
    TRIM=<func> RN=<func> RP=<func> MULTIPLY=<real> or <func>
    XMULTIPLE=<real> YMULTIPLE=<real> ZMULTIPLE=<real>
    FXMULTIPLE=<real> FYMULTIPLE=<real> XSHIFT=<real>
    YSHIFT=<real> ZSHIFT=<real> CURRENT=<real> NCURRENT=<real>
    PCURRENT=<real> EMOB=<real> HMOB=<real> SHAPE=<char>
    NAME=<char> EXPAND=<real> &[ ] ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| MATERIAL | char | material name. "SILICON", "OXIDE", "NITRIDE", "INSULATOR", "POLY", "VACUUM", "LOCOS", "SUPREM4", "SSUPREM4", "TIF", "DFISE2", "ISE2", "M1", "M2", "M3", "M4", "M5", "M6", "M7", or "M8" (default: SILICON when Z > 0, OXIDE when Z < 0). See Appendix J for MATERIAL=LOCOS. |
| XMIN | real | specified region minimum X-coordinate(unit: $\mu$ m) |
| YMIN | real | specified region minimum Y-coordinate(unit: $\mu$ m) |
| ZMIN | real or series | specified region minimum Z-coordinate(unit: $\mu$ m) (The maximum number of series element is 10.) |
| XMAX | real | specified region maximum X-coordinate(unit: $\mu$ m) |
| YMAX | real | specified region maximum Y-coordinate(unit: $\mu$ m) |
| ZMAX | real or series | specified region maximum Z-coordinate(unit: $\mu$ m) (The maximum number of series element is 10.) |
| PATTERN | char | the direction from the hypotenuse center to the right-angle in the triangle case (X+Y+, X+Y-, X-Y+, X-Y-, Y+Z+, Y+Z-, Y-Z+, Y-, Z-, X+Z+, X+Z-, X-Z+, or X-Z-) (if this parameter is not defined, the region shape is regarded as box) |
| XORIGIN | real | X-coordinate of SUPREM4 point to correspond to SPECTRA point (XMIN, YMIN, ZMIN) (unit: $\mu$ m) |
| YORIGIN | real | Y-coordinate of SUPREM4 point to correspond to SPECTRA point (XMIN, YMIN, ZMIN) (unit: $\mu$ m) |
| AXIS | char | the normal vector direction of SUPREM4 calculation plane (+X, +Y, -X, or -Y) |
| MASK | func | mask shape definition POLYGON(X1, Y1 ... Xn, Yn), FILE(file_name, cell_name, layer, data_type) or MFILE(file_name, layer_name) can be denoted as the function. X1, Y1 ... Xn, Yn is coordinate progression. If the progression is not closed, it is automatically closed. GDSII format data or CADENCE CAD data can be used by FILE function, which denotes file name, cell name, layer number, and data type. Data type can be omitted. GDSII and CADENCE data are automatically distinguished. MENTOR CAD data can be used by MFILE function, which denotes file name, and layer name. In CADENCE and MENTOR cases the data file must be ASCII data converted from CAD data. Plural MASK parameters can be designated, where the mask data by OR logical operation between the designated MASK data and the mask data created before the MASK parameter is created. |
| ANDMASK | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the designated ANDMASK data and the mask data created before the ANDMASK parameter is created. |
| ANDNEGA | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the negative data of designated ANDNEGA data and the mask data created before the ANDNEGA parameter is created. |
| XPLUS | real | shift value of MASK X-coordinate (unit: $\mu$ m) (default: 0) |
| YPLUS | real | shift value of MASK Y-coordinate (unit: $\mu$ m) (default: 0) |
| FILE | char | SUPREM4 calculation result file name |
| TN0 | real | life time of electron(unit:sec)(default:3.95E-4) |
| TP0 | real | life time of hole(unit:sec)(default:3.52E-5) |
| RN | func | Generation-recombination rate of electron in SRH model. FILE(file_name) should be denoted, where file_name indicates the name of the file including the generation-recombination rate (unit: $s^{-1}$ ) and coordinate data. |
| RP | func | Generation-recombination rate of hole in SRH model. FILE(file_name) should be denoted, where file_name indicates the name of the file including the generation-recombination rate (unit: $\rm s^{-1}$ ) and coordinate data. |
| GR | real or func | carrier generation-recombination velocity (unit: A/cm³) (default:0) In the case of GR=FILE(file_name), generation-recombination velocity is input from the file denoted by file_name. |
| CURRENT | real | Total injection current (unit: A) Electron and hole currents are equally injected in the brick region defined by Xmin, Xmax, Ymin, Ymax, Zmin, and Zmax. (default:0) |
| NCURRENT | real | Total injected electron current (unit: A) Electron current is injected in the brick region defined by Xmin, Xmax, Ymin, Ymax, Zmin, and Zmax. (default:0) |
| PCURRENT | real | Total injected hole current (unit: A) Hole current is injected in the brick region defined by Xmin, Xmax, Ymin, Ymax, Zmin, and Zmax. (default:0) |
| BEAK | real or series | bird's beak length in the case of MATERIAL=LOCOS or SHAPE=LOCOS (unit: $\mu$ m) (default:0) (See Appendix J) (The maximum number of series element is 10.) |
| ZBEAK | real | Z-coordinate of bird's beak edge in the case of MATERIAL=LOCOS or SHAPE=LOCOS (unit: $\mu$ m) (default:0) (See Appendix J) |
| EDGE | char | MASK edge in the case of MATERIAL=LOCOS or SHAPE=LOCOS (See Appendix J) (ACTIVE or FIELD) (default: ACTIVE) |
| TRIM | func | trimming or periodical arrangement of mask pattern (RECTANGLE() or PERIODIC()) |
| MULTIPLY | real or func | multiplied value to the carrier generation- recombination rate and current value denoted by GR, CURRENT, NCURRENT, or PCURRENT PWL, EXP, PULSE, or SIN can be denoted as a time dependent function (see Time dependent function in Model) (default: 1) |
| XMULTIPLE | real | multiplication factor for MASK X-coordinate (default: 1) |
| YMULTIPLE | real | multiplication factor for MASK Y-coordinate (default: 1) |
| ZMULTIPLE | real | multiplication factor for Z-coordinate of the carrier generation-recombination rate file (default: 1) |
| FXMULTIPLE | real | multiplication factor for X-coordinate of the carrier generation-recombination rate file (default: 1) |
| FYMULTIPLE | real | multiplication factor for Y-coordinate of the carrier generation-recombination rate file (default: 1) |
| XSHIFT | real | shift value of X-coordinate for carrier generation-recombination rate file (unit: $\mu m)$ (default: 0) |
| YSHIFT | real | shift value of Y-coordinate for the carrier generation-recombination rate file (unit: $\mu m)$ (default: 0) |
| ZSHIFT | real | shift value of Z-coordinate for the the carrier generation-recombination rate file (unit: $\mu m)$ (default: 0) |
| EMOB | real | Electron mobility under low concentration, low electric field, and room temperature (unit: cm²·s⁻¹·V⁻¹) (default: 1430). |
| HMOB | real | Hole mobility under low concentration, low electric field, and room temperature (unit: cm²·s⁻¹·V⁻¹) (default: 460). |
| SHAPE | char | designation of shape of the material If "LOCOS" is denoted, LOCOS shape is created. (see Appendix J) |
| NAME | char | region name (default: "R" + REGION statement input order) |
| EXPAND | real | Expansion (positive value) or shrinkage (negative value) for MASK polygons (unit:μm) (default:0) |
| & [ ] | operator | The order of MASK operation can be designated by puting MASK=, ANDMASK=, or ANDNEGA= into brackets [], where AND operation is performed in the case that "&" exists just before [ and OR |

```
[ ex. ] REGION MATERIAL=NI XMI=2 YMI=3 ZMI=1 XMA=8 YMA=8 ZMA=2
        REGION MATERIAL=SUPREM4 FILE='mostr.str' XMI=1 YMI=2 ZMI=0
            XMA=3 YMA=5 ZMA=5 AXIS=-X XO=0 YO=-0.1
        REGION MATERIAL=LOCOS BEAK=1.0 ZMIN=-0.6 ZMAX=0.7
            MASK=POLYGON (1, 1 3, 1 4, 2 4, 5 2, 5 1, 4) NAME=LOCOS-1
        REGION MATERIAL=OXIDE ZMI=0 ZMA=0.4 MASK=FILE (maskfile.txt, trench, 9)
        REGION MATERIAL=LOCOS ZMI=-0.5 ZMA=0.4 BEAK=0.9
            MASK=FILE(maskfile.txt,locos, 3, 1) XPLUS=2 YPLUS=3
        REGION XMI=2 YMI=3 ZMI=1 XMA=8 YMA=8 ZMA=2 TNO=1E-5 TPO=1E-6
        REGION MATERIAL=OXIDE ZMI=0 ZMA=0.4 MASK=FILE(maskfile.txt,oxide,8)
            TRIM=RECT (0. 1 0. 2 7. 1 9. 2)
        REGION XMI=0 YMI=0 ZMI=0 XMA=6 YMA=7 ZMA=8 RN=FILE(srhfile-1.txt)
        REGION XMI=1 YMI=0 ZMI=0 XMA=8 YMA=8 ZMA=8 GR=FILE(grfile-1.txt)
            MULTI=1E-2 XPLUS=0.5 YPLUS=1.2 ZPLUS=0.2
        REGION XMI=1 YMI=2 ZMI=0 XMA=5 YMA=3 ZMA=2 EMOB=800
        REGION MATERIAL=LOCOS MASK=POLYGON(1, 1 3, 1 4, 2 4, 5 2, 5 1, 4)
            BEAK= (0. 15 0. 29 0. 4 0. 5) ZMIN= (-0. 05 -0. 08 -0. 15 -0. 2)
            ZMAX=(0.05 0.12 0.15 0.2) EDGE=FIELD NAME=LOCOS-2
        REGION NC=1E-6 XMI=1.5 YMI=2.1 ZMI=0 XMA=1.8 YMA=2.7 ZMA=1.2
        REGION XMI=1 YMI=0 ZMI=0 XMA=8 YMA=8 ZMA=8 GR=FILE(grfile-1.txt)
            MULTI=1E-3 FXMUL=0.5 FYMUL=0.5 XSHIFT=2 YSHIFT=3 ZSHIFT=0.1
        REGION MATERIAL=LOCOS ZMI=-0.2 ZMA=0.2 BEAK=0
            MASK=FILE (maskfile.txt, cell-2, 3, 1)
            MASK=FILE (maskfile.txt, cell-2, 4, 1)
            ANDMASK=FILE (maskfile.txt, cell-2, 5, 1)
            ANDNEGA=FILE (maskfile.txt, cell-2, 6, 1)
        REGION MATERIAL=INSULATOR ZMIN=3.0 ZMA=4.0 ZBEAK=3.5 BEAK=0.5
            SHAPE=LOCOS MASK=FILE (maskfile.txt,cell-2,3,1)
        REGION GR=FILE(grfile-1.txt) XMI=3 YMI=0 ZMI=0 XMA=8 YMA=9 ZMA=3
            MULTI=PWL (0, 0 1E-7, 0 5E-7, 1 8E-7, 1 1E-6, 0)
        REGION MATERIAL=M1 ZMI=0 ZMA=0.4 MASK=FILE(maskfile.gds,cell-1,9)
            EXPAND=0.2
        REGION MATERIAL=M1 ZMI=-0.3 ZMA=-0.1
            MASK=FILE (maskfile.txt, cell-3, 3, 1)
            &[ MASK=FILE(maskfile.txt,cell-3,4,1)
            ANDMASK=FILE (maskfile.txt, cell-3, 5, 1)
            ANDNEGA=FILE (maskfile.txt, cell-3, 6, 1) ]
```

> Notice:

1. SILVACO SSUPREM4 structure file can be input by denoting MATERIAL=SSUPREM4
2. ISE 2-D simulation data can be input by denoting MATERIAL=ISE2. Then, the grid information file name must be denoted as FILE parameter.
3. Synopsis 2-D structure file can be input by denoting MATERIAL=TIF.
4. Synopsis Sentaurus 2-D process simulation data file of DFISE format can be input by denoting MATERIAL= DFISE2. Then, the grid information file name must be denoted as FILE parameter.
5. In the case of MATERIAL=LOCOS or SHAPE=LOCOS, inner part of the designated region is active region and outer part is field region (see Appendix J).
6. In the case of MATERIAL-LOCOS or SHAPE-LOCOS in plural REGION statements for different materials with different depths, the depth ranges designated by Zmin and Zmax must not be overlapped.
7. The data file format is shown below in the case of GR=FILE(file_namae), RN=FILE(file_name), or RP=FILE(file_name). It is same as 3-D impurity file format in DOPE statement.
    (one row consists of less than 1024 characters and at least one space is necessary between numbers)
    - [Format number 0]
      $1^{\text{st}}$ row: 0 Nx Ny
      (0 is format number and Nx, Ny are mesh number along X and Y axis, respectively)
      2nd $\sim$ Nx+1th row: X-coordinate (unit: $\mu$ m, Nx)
      Nx+2th $\sim$ Nx+Ny+1th row: Y-coordinate (unit: $\mu$ m, Ny)
      Nx+Ny+2th row: X-coordinate Y-coordinate N1
      (N1 is mesh number along depth axis of $1^{\rm st}$ mesh point on XY plane)
      Nx+Ny+3th $\sim$ Nx+Ny+N1+2th row: Z-coordinate generation-recombination velocity (A/cm $^3$ ) or SRH generation-recombination rate (s $^{-1}$ )
      .
      .
      .
      Mth row: X-coordinate Y-coordinate Nn
      (Nn is mesh number along depth axis of nth mesh point on XY plane)
      M+1th $\sim$ M+Nnth row: Z-coordinate generation-recombination velocity (A/cm $^3$ ) or SRH generation-recombination rate (s $^{-1}$ )
      where n = Nx \* Ny, $M = Nx+Ny+1+\sum_{k=1}^{n} (Nk+1)$
    - [Format number 1]
      1st row: 1 Nx Ny Nz
      (1 is format number and Nx, Ny, and Nz are total grid numbers along X, Y, and Z axes, respectively)
      2nd∼Nx+1th row: X-coordinate (unit: \mu m, Nx)
      Nx+2th $\sim$ Nx+Ny+1th row: Y-coordinate (unit: \mu m, Ny)
      Nx+Ny+2th $\sim$ Nx+Ny+Nz+1th row: Z-coordinate (unit: \mu m, Nz)
      Nx+Ny+Nz+2th $\sim$ Nx+Ny+Nz+Nx*Ny*Nz+1th row: GR (generation-recombination rate, unit: A/cm³) or SRH (generation-recombination rate of SRH model, unit: s^{-1})
      The order of GR is shown bellow by C language program.

        ```c
        for (k=0; k< Nz; k++) {
            for (j=0; j< Ny; j++) {
                for (i=0; i< Nx; i++) {
                    printf(" %e \n", GR(X[i], Y[j], Z[k]));
                }
            }
        }
        ```

    where X[i], Y[j], and Z[k] denote i-th X, j-th Y, and k-th Z coordinates, respectively.

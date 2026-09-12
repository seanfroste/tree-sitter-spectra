## DOPE

This statement defines the impurity profile. The region to introduced the impurity profile is defined by the doping mask coordinate. The profile is defined by box, Gaussian, SUPREM2, SUPREM3, SUPREM4 or pairs of depth and concentration. The impurity profile lateral spreading is performed by the error function. The mask shape is defined by rectangles, right-angled triangles and polygons.

```
[format] DOPE [TYPE=<char> or ELEMENT=<char>] PROFILE=<char>
    XMIN=<real> YMIN=<real> XMAX=<real> YMAX=<real>
    [FILE=<char> ZMIN=<real> ZMAX=<real> CMAX=<real> CBACK=<real>
    XJ=<real> RP=<real> LD=<real> THETA=<real> PHI=<real>
    MULTIPLY=<real> PATTERN=<char> STEP=<real> or <char> GFILE=<char>
    XORIGIN=<real> YORIGIN =<real> AXIS=<char> NAME=<char> MASK=<func>
    ANDMASK=<func> ANDNEGA=<func> CRATE =<real> DETECT=<char>
    APPEND=<char> TRIM=<func> XPLUS= <real> YPLUS=<real> ZPLUS=<real>
    XMULTIPLE= <real> YMULTIPLE=<real> ZMULTIPLE=<real> XSHIFT= <real>
    YSHIFT=<real> ZSHIFT=<real> FXMULTIPLE= <real> FYMULTIPLE=<real>
    NEGATIVE=<char> XREPEAT= <real> or <char> YREPEAT= <real> or <char>
    LOCOS= <real> MATERIAL= <real> SKIP=<real> EXPAND=<real> & [ ] ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| TYPE | char | impurity type "P" or "N" |
| ELEMENT | char | impurity atom "B", "P", "As", or "Ge" |
| PROFILE | char | impurity profile type "GAUSSIAN", "SUPREM2", "SUPREM3", "SUPREM4", "SSUPREM4", "ISE2", "TIF", "DFISE2", "BOX", "FREE", or "3D" |
| FILE | char | file name of SUPREM2, SUPREM3, SUPREM4, ISE2 or TIF |
| GFILE | char | ISE2 or DFISE2 grid information file name |
| XMIN | real | mask edge minimum X-coordinate (unit: $\mu$ m) |
| YMIN | real | mask edge minimum Y-coordinate (unit: $\mu$ m) |
| XMAX | real | mask edge maximum X-coordinate (unit: $\mu$ m) |
| YMAX | real | mask edge maximum Y-coordinate (unit: $\mu$ m) |
| ZMIN | real | profile defined region minimum Z-coordinate in box or SUPREM4 case(unit: $\mu\mathrm{m})$ |
| ZMAX | real | profile defined region maximum Z-coordinate in box or SUPREM4 case(unit: $\mu\mathrm{m})$ |
| CMAX | real | maximum concentration in GAUSSIAN or box case (unit: ${\rm cm}^{-3}$ ) |
| CBACK | real | substrate concentration in GAUSSIAN case (unit: $cm^{-3}$ ) |
| XJ | real | junction depth in GAUSSIAN case(unit: $\mu$ m) |
| RP | real | peak depth in GAUSSIAN case(unit: $\mu$ m) |
| LD | real | lateral spreading ratio against depth direction (default:1) |
| THETA | real | ion implantation angle against Z-axis in the case Of PROFILE=GAUSSIAN or FREE (Z-axis direction is 0 degree) (unit:degree) $0 \le \text{THETA} \le 90$ (default:0) |
| PHI | real | ion implantation angle against X-axis on XY-plane in the case of PROFILE=GAUSSIAN or FREE (X-axis direction is 0 degree) (unit:degree) (default:0) |
| MULTIPLY | real | multiple constant of impurity concentration in the case of GAUSSIAN, SUPREM2, SUPREM3, or FREE (default:1) |
| PATTERN | char | the direction from the hypotenuse center to the right-angle in the triangle case (X+Y+, X+Y-, X-Y+, or X-Y-) (if this parameter is not defined, the mask shape is regarded as a rectangle) |
| STEP | real or char | division number of a oblique side of the polygon (an oblique side of a polygon is transformed to steps) STEP=AUTO automatically decides the STEP value avoiding twisted polygon ("AUTO" or number) (default:10) |
| XORIGIN | real | X-coordinate of SUPREM4 point to correspond to SPECTRA point (XMIN, YMIN, ZMIN) (unit: $\mu$ m) |
| YORIGIN | real | Y-coordinate of SUPREM4 point to correspond to SPECTRA point (XMIN, YMIN, ZMIN) (unit: $\mu$ m) |
| AXIS | char | the normal vector direction of SUPREM4 calculation plane in PROFILE=SUPREM4 case or the impurity doping plane in the case of PROFILE=GAUSSIAN or FREE (X, Y, Z, -X, -Y, -Z, X+Y+, X+Y-, X-Y+, X-Y-, IN, or OUT) X, Y, Z, -X, -Y, -Z: normal axis of the doping plane X+Y+, X+Y-, X-Y+, X-Y-: diagonal direction normal to the doping plane IN, OUT: inside or outside of the designated polygon |
| NAME | char | the doping mask name (default: "D" + DOPE statement input order ) |
| MASK | func | mask shape definition POLYGON(X1, Y1 ... Xn, Yn), FILE(file_name, cell_name, layer, data_type) or MFILE(file_name, layer_name) can be denoted as the function. X1, Y1 ... Xn, Yn is coordinate progression. If the progression is not closed, it is automatically closed. GDSII format data or CADENCE CAD data can be used by FILE function, which denotes file name, cell name, layer number, and data type. Data type can be omitted. GDSII and CADENCE data are automatically distinguished. MENTOR CAD data can be used by MFILE function, which denotes file name, and layer name. In CADENCE and MENTOR cases the data file must be ASCII data converted from CAD data. Plural mask parameters can be designated in a DOPE statement, where the mask data by OR logical operation between the designated MASK data and the mask data created before the MASK paramter is created. |
| ANDMASK | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the designated ANDMASK data and the mask data created before the ANDMASK parameter is created. |
| ANDNEGA | func | Same functions as MASK parameter can be designated. The mask data by AND logical operation between the negative data of designated ANDNEGA data and the mask data created before the ANDNEGA parameter is created. |
| XPLUS | real | shift value of MASK X-coordinate (unit: $\mu$ m) (default: 0) |
| YPLUS | real | shift value of MASK Y-coordinate (unit: $\mu$ m) (default: 0) |
| ZPLUS | real | shift value of Z-coordinate of one dimesional impurity profile (unit: $\mu$ m) (default: 0) |
| XMULTIPLE | real | multiplication factor of MASK X-coordinate (default: 1) |
| YMULTIPLE | real | multiplication factor of MASK Y-coordinate (default: 1) |
| ZMULTIPLE | real | multiplication factor of Z-coordinate of 1-D impurity profile (default: 1) |
| XSHIFT | real | shift value of X-coordinate of 3D impurity file (unit: $\mu$ m) (default: XPLUS) |
| YSHIFT | real | shift value of Y-coordinate of 3D impurity file (unit: $\mu$ m) (default: YPLUS) |
| ZSHIFT | real | shift value of Z-coordinate of 3D impurity file (unit: $\mu$ m) (default: ZPLUS) |
| FXMULTIPLE | real | multiplication factor of X-coordinate of 3D impurity file (default: 1) |
| FYMULTIPLE | real | multiplication factor of Y-coordinate of 3D impurity file (default: 1) |
| CRATE | real | ratio between the concentration to estimate the diffusion length and the maximum concentration in SUPREM2. SUPREM3 or FREE case. When this parameter is not defined, CBACK is used to estimate the diffusion length. |
| DETECT | char | the position to estimate the diffusion length against the peak concentration position in SUPREM2, SUPREM3 or FREE case. ("SHALLOW" or "DEEP") (default: "DEEP") |
| APPEND | char | append impurity profile ("Y" or "N") (valid only in the case of PROFILE=SUPREM4) (default: "N") |
| TRIM | func | trimming or periodical arrangement of mask pattern (RECTANGLE() or PERIODIC()) RECTANGLE(x0 y0 x1 y1) executes trimming with rectangle region of $x0 \le x \le x1$ and $y0 \le y \le y1$ . PERIODIC(x0 y0 x1 y1) executes periodic arrangement with boundary of x=x0, x=x1, y=y0, and y=y1. |
| NEGATIVE | char | Negative mask pattern designation ("Y" or "N") (default: "N") |
| XREPEAT | real or char | repeat number of structure along X-axis for 3-D lateral extension from 1-D impurity profile. In the case of even number, the half number of the structures are repeated for negative and positive directions, respectively. In the case of negative odd number, negative direction includes one more structure than positive direction. In the case of positive odd number, positive direction includes one more structure than negative direction. When "AUTO" is designated, the repeat number is automatically decided. (see appendix B) ("AUTO" or number) (default: "A") |
| YREPEAT | real or char | repeat number of structure along Y-axis for 3-D lateral extension from 1-D impurity profile. In the case of even number, the half number of the structures are repeated for negative and positive directions, respectively. In the case of negative odd number, negative direction includes one more structure than positive direction. In the case of positive odd number, positive direction includes one more structure than negative direction. When "AUTO" is designated, the repeat number is automatically decided. (see appendix B) ("AUTO" or number) (default: "A") |
| LOCOS | char | GAUSSIAN or FREE impurity concentration depending on the distance from interfaces between silicon and another material is added around the bird's beak of LOCOS structure designated by REGION mask name as same as this parameter value in the range of Zmin $\langle$ Z $\langle$ Zmax of the REGION statement. |
| MATERIAL | char | material name for an interface with silicon. GAUSSIAN or FREE impurity concentration depending on the distance from the above interface in a pillar region designated by Zmin, Zmax and MASK or Xmin, Xmax, Ymin, and Ymax is added to whole silicon region in $Z \geq 0$ . |
| SKIP | real | skip number of pair of depth and concentration per input pair in the case of PROF=FREE, which reduces the number of input pairs to 1/(SKIP+1) (default: 0) |
| EXPAND | real | Expansion (positive value) or shrinkage (negative value) for MASK polygons (unit: μm) (default:0) |
| &[ ] | operator | The order of MASK operation can be designated by puting MASK=..., ANDMASK=..., or ANDNEGA=... into brackets [], where AND operation is performed in the case that "&" exists just before [ and OR operation is performed in the other cases |

> Notice:

1. TYPE, CMAX, CBACK, XJ, and RP are necessary in the case of PROFILE=GAUSSIAN
2. TYPE, CMAX, ZMIN and ZMAX are necessary in the case of PROFILE=box
3. ELEMENT and FILE are necessary in the case of PROFILE=SUPREM2 or SUPREM3
4. FILE, XORIGIN, YORIGIN, AXIS, ZMIN and ZMAX are necessary in the case of PROFILE=SUPREM4, where the file format must be Medici structure file format
5. TYPE (ELEMENT for only Germanium case) and FILE are necessary in the case of PROFILE=FREE
6. PROFILE, XMIN, YMIN, XMAX and YMAX are necessary in every case
7. In the case of PROFILE=FREE, if the neighboring two points with same depth exist, the depth is regarded as Si-SiO2 interface (Z=0). Unless the above points, the first depth is regarded as Si-SiO2 interface (Z=0).
8. more than 1000th pair of depth and concentration is ignored in the case of PROFILE=FREE or SUPREM3
9. at least one pair of depth and concentration less than CBACK is necessary in the case of PROFILE=FREE, SUPREM2 or SUPREM3
10. SILVACO SSUPREM4 structure file can be input by denoting PROFILE=SSUPREM4
11. ISE 2-D simulation data can be input by denoting PROFILE=ISE2. Then "Net Active" and "Total" doping concentration are necessary and the grid information file name must be denoted as GFILE parameter value.
12. Synopsis 2-D process simulation data TIF file can be input by denoting PROFILE=TIF. Then "Net" and "Total" doping concentrations are necessary.
13. Synopsis Sentaurus 2-D process simulation data file of DFISE format can be input by denoting PROFILE= DFISE2. Then "BActive", "PActive", and "AsActive" doping concentrations must be included in the file and the grid information file name must be denoted as GFILE parameter value.
14. If FILE() or MFILE() is denoted as the MASK parameter value, any overlapping region of polygons must not be included in every layer of the designated file, because the impurity is doubly doped in the overlapping region.
15. In the case of ELEMENT=Ge, PROFILE must be BOX or FREE. Then, the Ge compound rate x of $\mathrm{Si}_{1-x}\mathrm{Ge}_x$ must be denoted as CMAX for BOX profile or the concentration in FREE file.
16. In the case of PROFILE=3D, the format of impurity file must be as follows. (one row consists of less than 1024 characters and at least one space is necessary between numbers)
    - [Format number 0]
    $1^{\text{st}}$ row: 0 Nx Ny

        `(0 is format number and Nx, Ny are total grid numbers along X and Y axis, respectively, which must not exceed 9900)`

        Nx+Ny+2th row: X-coordinate Y-coordinate N1

        `(N1 is mesh number along depth axis of $1^{\rm st}$ mesh point on XY plane)`

        Nx+Ny+3th $\sim$ Nx+Ny+N1+2th row: Z-coordinate net-concentration(Cn-Cp) total-concentration(Cn+Cp)

        `(Cn, Cp denotes donor and acceptor concentration, respectively)`

        Mth row: X-coordinate Y-coordinate Nn

        `(Nn is mesh number along depth axis of nth mesh point on XY plane)`

        M+1th∼M+Nnth row: Z-coordinate net-concentration(Cn-Cp) total-concentration(Cn+Cp)

        where:

        ```
        n = N_X * N_Y
        M = N_X + N_Y + 1 + \sum_{k=1}^{n} (N_k + 1)
        ```

    - [Format number 1]
    1st row: 1 Nx Ny Nz

        `(1 is format number and Nx, Ny, and Nz are total grid numbers along X, Y, and Z axis, respectively, which must not exceed 9900)`

        2nd∼Nx+1th row: X-coordinate (unit: \mu m, Nx)

        Nx+2th~Nx+Ny+1th row: Y-coordinate (unit: \mu m, Ny)

        Nx+Ny+2th \sim Nx+Ny+Nz+1th row: Z-coordinate (unit: \mu m, Nz)

        Nx+Ny+Nz+2th\sim

        Nx+Ny+Nz+Nx*Ny*Nz+1th row: net-concentration(Cnet≡Cn-Cp) total-concentration (Ctotal≡Cn+Cp)

        `(Cn, Cp denotes donor and acceptor concentration, respectively)`

        .

        .

        .

        The order of Cnet and Ctotal is shown below by C language program.

        ```c
        for (k=0; k< Nz; k++) {
            for (j=0; j<Ny; j++) {
                for (i=0; i< Nx; i++) {
                    printf ("%e %e\n", Cnet(X[i], Y[j], Z[k]), Ctotal(X[i], Y[j], Z[k]));
                }
            }
        }
        ```

        where X[i], Y[j], and Z[k] denote i-th X, j-th Y, and k-th Z coordinates, respectively.

```
[ex.]   DOPE ELEM=P PROF=SUPREM3 FILE=storage.sprm XMI=5 YMI=5 XMA=10 YMA=8
        DOPE TYPE=N PROF=GAUS XMI=2 YMI=3 XMA=5.5 YMA=9.8
            CM=1. 2E17 CB=1. 5E15 X T=0. 5 RP=0. 1 LD=0. 9 NAME=CS

        DOPE TYPE=N PROF=BOX XMI=0 YMI=0 XMA=5 YMA=7 ZMI=0 ZMA=8 CM=2E14 CB=2E15
        DOPE PROF=SUPREM4 FILE=cell.sp4 XMI=0 YMI=0 ZMI=0 XMA=8 YMA=9 ZMA=10 XO=1 YO=0.2 AXIS=+Y APPEND=Y
        DOPE ELEM=B PROF=SUPREM3 FILE=pwell.sprm XMI=0 YMI=0 XMA=10 YMA=8 MUL=1.5 CRATE=0.01 DETECT=S NAME=Pwell
        DOPE TYPE=N PROF=GAUSS CM=1.2E17 CB=2E15 XJ=0.8 RP=0.1 MASK=POLY(1, 1 5, 1 7, 3 7, 8 4, 8 4, 5 1, 5) NAME=PD
        DOPE ELEM=P PROF=SUPREM3 FILE=pd.spr MASK=FILE(maskfile.txt,pd-cell,7)
        DOPE ELEM=As PROF=SUPREM3 FILE=vccd.spr XPLUS=2 YPLUS=3 MASK=FILE(maskfile.txt,vccd-cell,5,1)
        DOPE ELEM=Ge PROF=FREE FILE=SiGe-1.txt XMI=0 YMI=0 XMA=10 YMA=10
        DOPE ELEM=B PROF=SUPREM3 FILE=sd. spr XPLUS=2 YPLUS=3 ZPLUS=4 MASK=MFILE(mask.txt,sd-layer)
        DOPE TYPE=N PROF=FREE FILE=pd. fre MASK=FILE(maskfile.txt, npd, 9) TRIM=PERIODIC(0.2 0.3 5.1 7.2)
        DOPE PROF=ISE2 FILE=ise-2d. dat GFILE=ise-2d. grd X0=-1 Y0=-2 XMI=0 YMI=0 ZMI=0 XMA=8 YMA=9 ZMA=10
        DOPE PROF=3D FILE=prof3d.imp XMI=0 YMI=0 ZMI=0 XMA=8 YMA=9 ZMA=10 APPEND=Y XPLUS=2 YPLUS=3 ZPLUS=4
        DOPE ELEM=B PROF=SUPREM3 FILE=cs.spr
            MASK=FILE(maskfile.txt,vccd-cell,5,1) NEGA=Y STEP=AUTO
        DOPE TYPE=N PROF=FREE FILE=pd.txt XMIN=4 YMIN=4 XMAX=6 YMAX=6 THETA=10 PHI=30
        DOPE TYPE=P PROF=FREE FILE=cs.txt XMIN=2 XMAX=3 ZMIN=0.5 ZMAX=1 YMAX=2.5 AXIS=-Y
        DOPE TYPE=N PROF=FREE FILE=sd.txt ZMAX=4.5 AXIS=-Z MASK=FILE(maskfile.txt, backside-cell, 3, 1)
        DOPE TYPE=P PROF=FREE FILE=p+. txt ZMIN=3.5 ZMAX=4.0 AXIS=X-Y+XMIN=2.5 YMIN=1.5 XMAX=3.5 YMAX=2.5
        DOPE TYPE=P PROF=FREE FILE=cs.txt ZMIN=3.5 ZMAX=4.0 AXIS=IN MASK=FILE(maskfile.txt, side-wall, 4, 1)
        DOPE TYPE=P PROF=FREE FILE=cs.txt
            MASK=FILE (maskfile.txt, pixel, 1, 1)
            MASK=FILE (maskfile.txt, pixel, 2, 1)
            ANDMASK=FILE (maskfile.txt, pixel, 3, 1)
            ANDNEGA=FILE (maskfile.txt, pixel, 4, 1)
        DOPE TYPE=P PROF=FREE FILE=cs.txt
            MASK=FILE (maskfile.txt, pixel, 1, 1) SKIP=2
        DOPE TYPE=P PROF=FREE FILE=dope-1.txt LOCOS=LOCOS-1
        DOPE TYPE=P PROF=FREE FILE=dope-3.txt MATERIAL=NITRIDE MASK=FILE(maskfile.txt,cell-3,5,1) ZMIN=0.1 ZMAX=4
        DOPE PROF=3D FILE=prof3d.imp XMI=0 YMI=0 ZMI=0 XMA=16 YMA=9 ZMA=10 XSHIFT=8 YSHIFT=0 ZSHIFT=1 FXMUL=-1
        DOPE TYPE=P PROF=FREE FILE=ppd.txt
            MASK=FILE(maskfile.gds,pixel,6,1) EXPAND=-0.1
        DOPE TYPE=N PROF=FREE FILE=npd.txt
            MASK=FILE (maskfile.txt, pixel, 5)
            &[MASK=FILE(maskfile.txt,pixel,6)
            ANDMASK=FILE (maskfile.txt, pixel, 7)
            ANDNEGA=FILE(maskfile.txt,pixel,8)]
```

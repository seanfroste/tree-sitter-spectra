## GRID

This statement specifies the grid number and the grid space width in the calculation region. The grid space width is given by geometric progression.

```
[format] GRID [ RATIO or RATE= <real> SPACE=<real> XMIN=<real> XMAX=<real>
    YMIN=<real> YMAX=<real> ZMIN=<real> ZMAX=<real> TMIN=<real>
    TMAX=<real> XPLUS=<real> YPLUS=<real> ZPLUS=<real>
    TPLUS=<real> ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| RATIO or RATE | real | the ratio of the first and the last grid space (default:1) |
| SPACE | real | the number of grid space (default:1) |
| XMIN | real | X-axis first grid coordinate(unit: $\mu$ m) (default: the last XMAX of the GRID statement, if XMAX was defined in the former GRID statement, or 0 in the other case.) |
| XMAX | real | X-axis last grid coordinate(unit: $\mu$ m) (default:0) |
| YNIN | real | Y-axis first grid coordinate(unit: $\mu$ m) (default: the last YMAX of the GRID statement, if YMAX was defined in the former GRID statement, or 0 in the other case.) |
| YMAX | real | Y-axis last grid coordinate(unit: $\mu$ m) (default:0) |
| ZMIN | real | Z-axis first grid coordinate(unit: $\mu$ m) (default: the last ZMAX of the GRID statement, if ZMAX was defined in the former GRID statement, or 0 in the other case.) |
| ZMAX | real | Z-axis last grid coordinate(unit: $\mu$ m) (default:0) |
| TMIN | real | Time axis first grid coordinate(unit: sec) (default: the last TMAX of the GRID statement, if TMAX was defined in the former GRID statement, or 0 in the other case.) |
| TMAX | real | Time axis last grid coordinate (unit: sec) (default:0) |
| XPLUS | real | shift value of XMIN and XMAX (unit: $\mu$ m) (default: 0) |
| YPLUS | real | shift value of YMIN and YMAX (unit: $\mu$ m) (default: 0) |
| ZPLUS | real | shift value of ZMIN and ZMAX (unit: $\mu$ m) (default: 0) |
| TPLUS | real | shift value of TMIN and TMAX (unit:sec) (default: 0) |

[notice] One of XMAX, YMAX and ZMAX must be specified in this statement.
 If XGSUM, YGSUM, ZGSUM, and TGSUM are the summation of X, Y, Z, and time axis grid space number respectively, the following equations must be satisfied:

XGSUM $\leq$ 9900, YGSUM $\leq$ 9900, ZGSUM $\leq$ 9900, TGSUM $\leq$ 9900 (XGSUM+3)\*(YGSUM+3)\*(ZGSUM+2) < GMAX  (GMAX depends on the system and the program version.)

If TMIN or TMAX is denoted in a GRID statement, STEP and TRATE of the CONVERGENCE statement are ignored.

```
[ ex. ] GRID ZMIN=-0.1 ZMAX=0 STEP=10
        GRID RATIO=20 ZMAX=20 STEP=30
```

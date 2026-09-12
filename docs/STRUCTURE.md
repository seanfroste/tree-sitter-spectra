## STRUCTURE

This statement specifies the calculation region and boundary conditions.

```
[format] STRUCTURE XBOUNDARY=<char> YBOUNDARY=<char> XMAX=<real>
    YMAX=<real> ZMAX=<real> ZMIN=<real> [ CELL=<char> ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| XBOUNDARY | char | X-direction boundary condition: PERIODIC, REFLECTION, or MIRROR. REFLECTION and MIRROR have the same meaning. |
| YBOUNDARY | char | Y-direction boundary condition: PERIODIC, REFLECTION, or MIRROR. REFLECTION and MIRROR have the same meaning. |
| XMAX | real | Maximum X-coordinate of calculation region (minimum: 0, unit: μm). |
| YMAX | real | Maximum Y-coordinate of calculation region (minimum: 0, unit: μm). |
| ZMAX | real | Maximum Z-coordinate of calculation region (unit: μm). |
| ZMIN | real | Minimum Z-coordinate of calculation region, which must be negative (unit: μm). |
| CELL | char | Cell structure: BOX or TETRAHEDRON (default: BOX). See Appendix O. |

```
[ex.] STRUCT XB=P YB=R XMAX=10 YMAX=10 ZMAX=20 ZMIN=-0.08 CELL=T
```

> Notice: ZMIN must be negative.

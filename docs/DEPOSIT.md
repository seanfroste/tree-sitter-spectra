## DEPOSIT

This statement deposit charge in the box region.

```
[format] DEPOSIT TYPE = <char> CHARGE = <char>
    XMIN = <real> YMIN = <real> ZMIN = <real>
    XMAX = <real> YMAX = <real> ZMAX = <real>
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| TYPE | char | charge type, "P" or "N" ( P:Hole, N:Electron ) |
| CHARGE | real or char | charge concentration (unit:electrons/cm³) or "RESTART" (charge distribution of the file defined by RESTART statement) |
| XMIN | real | box region minimum X-coordinate (unit: $\mu$ m) |
| YMIN | real | box region minimum Y-coordinate (unit: $\mu$ m) |
| ZMIN | real | box region minimum Z-coordinate (unit: $\mu$ m) |
| XMAX | real | box region maximum X-coordinate (unit: $\mu$ m) |
| YMAX | real | box region maximum Y-coordinate (unit: $\mu$ m) |
| ZMAX | real | box region maximum Z-coordinate (unit: $\mu$ m) |

```
[ex.]   DEPO TYPE=N CHARGE=1E15 XMIN=5.5 YMIN=6.0 ZMIN=0.1
            XMAX=6.5 YMAX=8.5 ZMAX=0.3
```

> Notice: In the zero carrier model case, the charge distribution defined by this statement in the defined region is not changed.

## IF

This statement controles the execution of the statements before ENDIF according to truth or falsehood of the equation in the square brackets just after "IF".

```
[format] IF [ <func> ] ELSE ENDIF
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| ELSE | - | The statements after the square brackets to ELSE are executed in truth case, and the statements after ELSE to ENDIF are executed in falsehood case. ELSE can be omitted. |
| ENDIF | - | If ELSE exists, the statements after ELSE to ENDIF are executed in falsehood case. If ELSE is omitted, the statements after the square brackets to ENDIF are executed in truth case. |

```
[ ex. ] IF [@Cmax <= 5E16]
            ELECTRODE Xmin=0 Xmax=5 Ymin=0 Ymax=5 Zmin=-0.05 Zmax=-0.03 V=0
        ELSE
            ELECTRODE Ymin=0 Ymax=5 Vmin=0 Ymax=5 Zmin=-0.05 Zmax=-0.03 V=5
        ENDIF
```

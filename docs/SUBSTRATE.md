## SUBSTRATE

This statement specifies impurity concentration, impurity type, and electrode voltage of substrate.

```
[format] SUBSTRATE TYPE=<char> CONCENTRATION=<real>
    VOLTAGE= <real> or <func> or <char>
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| TYPE | char | impurity type of the substrate (P or N) |
| CONCENTRATION | real | impurity concentration of the substrate (unit:cm-3) |
| VOLTAGE | real or func or char | substrate electrode voltage (unit:Volts) voltage source can be specified as a time dependent function in transient analysis (see Time dependent functions in Model) If FLOAT is specified, the z-component of electric field at the bottom of the calculation region is assumed to zero. |

```
[ ex. ] SUBS TYPE=P C=1E15 V=10
        SUBS TYPE=N C=1E14 V=FLOAT
        SUBS TYPE=P C=2E15 V=PWL(0, 0 1E-7, -2 5E-7, -2 1E-6, 0)
```

> Notice:

If VOLTAGE=FLOAT is specified, an oxide layer with thickness of $2\times10^{-9}$ m and one grid of z-axis are added to the bottom of the substrate.

## RESTART

This statement specifies the file name to load the calculation result and continue the calculation.

```
[format] RESTART FILE=<char> [ NEW_FERMI=<char> ADD_IMPURITY=<char>
    OLD_ELECTRODE= <char> ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| FILE | char | saved file name for restart |
| NEW_FERMI | char | new definition of NFERMI and PFERMI "Y" or "N" (default: N) |
| ADD_IMPURITY | char | designation of addition for impurity concentration "Y" or "N" (default: N) "Y" denotes the addition of the impurity concentration of DOPE statement to that of the restart file |
| OLD_ELECTRODE | char | designation of utilization of the electrode name in the restart file for changing the elctrode voltage without denoting coordinates "Y" or "N" (default: N) |

```
[ ex. ] RESTART FILE=MOS-Tr-1. save
        RESTART FILE=CCD-1.save NEW_FERMI=Y
        RESTART FILE=CIS-1.save ADD_IMP=Y
        RESTART FILE=MOS-1. save OLD_ELE=Y ELECTRODE NAME=TG V=3.3
```

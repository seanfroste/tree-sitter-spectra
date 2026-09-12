## CONVERGENCE

This statement defines the convergence condition.

```
[format] CONVERGENCE
    [ VERROR = <real> FERROR = <real> CARRIER = <real>
    PHASE1 = <real> PHASE2 = <real> PHASE3 = <real>
    TIME = <real> STEP = <real> TRATE = <real> SKIP = <char>
    PSOR = <real> PCYCLE = <real> NEWTON = <real> FSOR = <real>
    FCYCLE = <real> ABORT = <real> METHOD = <char>
    POISSON = <char> RANGECHECK = <char> FPRINT = <char> ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| VERROR | real | potential convergence value(maximum residual of Gummel method) (unit:Volts) (default:1E-6) |
| FERROR | real | quasi-Fermi potential convergence value(maximum residual of Gummel method) (unit:Volts) (default:1E-6) |
| CARRIER | real | carrier model: 0,1,-1 or 2 ( 1: electron one carrier model, -1: hole one carrier model) (default:0) |
| PHASE1 | real | PHASE1 maximum iteration number (default:5) |
| PHASE2 | real | PHASE2 maximum iteration number (default:50) |
| PHASE3 | real | PHASE3 maximum iteration number (default:30) |
| TIME | real | maximum time in transient analysis (unit:sec) (default:0) |
| STEP | real | the number of time steps in transient analysis (default:1) |
| TRATE | real | ratio of the first time step and the last time step in transient analysis (the time step is described as geometric progression)(default:1) |
| SKIP | char | skip the potential or current calculation ("Y", "N" or "ALL") In "ALL" case, impurity distribution calculation is also skipped. (default: N) |
| PS0R | real | SOR minimum iteration number for potential calculation (default:10) |
| CYCLE | real | ICCG inner loop for potential calculation (default:10) |
| NEWTON | real | ICCG outer loop for potential calculation (default:100) |
| FSOR | real | SOR maximum iteration number for current calculation (default:100) |
| FCYCLE | real | ICCG maximum iteration number for current calculation (default:1000) |
| ABORT | real | ICCG maximum outer loop for current calculation When the current calculation is not converged at the maximum iteration number (=FCYCLE), ICCG calculation is restarted after changing initial condition or aborted at the maximum outer loop which is denoted by "ABORT" value. (ABORT=0 means that ICCG calculation is never aborted) (default:1) |
| METHOD | char | solution method ( Gummel or Potetial-Dominant- Gummel) (default:Gummel) |
| POISSON | char | Poisson equation solving calculation at the beginning of PHASE 1 at every time step in transient analysis. (Everytime or None) (default: None) |
| RANGECHECK | char | POLYGON and Z-coordinate check for REGION, EXTRACT, NFERMI, PFERMI, LIGHT, and INTERFACE statement ("Y" or "N") (default: N) |
| FPRINT | char | output of the coordinate information resulting FN<0 or FP<0 to execution window and \*.out file ("Y" or "N") (default: N) |

> Notice: If TMIN or TMAX is denoted in a GRID statement, STEP and TRATE of the CONVERGENCE statement are ignored.

```
[ex.]   CONV VE=1E-5 FE=1E-5 CA=2 TIME=1E-9 STEP=3
        TRATE=5 NE=1000 FC=800 ABORT=3 METHOD=P POISSON=E RANGEC=Y FPRINT=Y
```

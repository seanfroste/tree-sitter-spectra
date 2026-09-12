## BIAS

This statement defines electrode name, potential step, repeat time, and quasi-Fermi region name which is changed with the potential.

```
[format] BIAS ENAME=<char> VSTEP=<real> NSTEP=<real> ERROR=<real>
    [ NFNAME=<char> PFNAME=<char> FLOAT=<char> BMODEL=<char> ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| ENAME | char | electrode name(electrode name of substrate is "SUB") |
| VSTEP | real | potential variation step(unit:Volts ) (default:0) |
| NSTEP | real | repeat time of calculation |
| ERROR | real | current convergence judgment value when FLOAT parameter is defined(unit:A) |
| NFNAME | char | region name of electron quasi-Fermi potential |
| PFNAME | char | region name of hole quasi-Fermi potential |
| FLOAT | char | floating electrode definition. "N" denotes electron current and "P" denotes hole current which are automatically converged to zero. (default: not define) |
| BMODEL | char | The model of convergence to zero electrode current in the case of FLOAT=N or P. ( Linear or Exponential) (default:Exponential) |

```
[ex.]   BIAS ENAME=DRAIN VSTEP=0.5 NSTEP=10 NFNAME=N2

        BIAS ENAME=PD VSTEP=0.5 NSTEP=5 NFNAME=PD1 FLOAT=N ERROR=1E-12

        BIAS ENAME=FD VSTEP=0.2 NSTEP=7 NFNAME=NFD FLOAT=N ERROR=1E-12 BMOD=L
```

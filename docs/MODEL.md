## MODEL

This statement specifies the calculation model.

```
[format] MODEL [ MOBILITY=<char> BGNARROW=<char> SRH=<char>
    AUGER=<char> TRAP=<char> TUNNELING=<char> FNCURRENT=<char> ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| MOBILITY | char | Mobility model: Y, N, or Planar (default: N). Y uses the mobility model, N uses constant EMOB and HMOB, and Planar assumes the Si-SiO2 interface exists only at Z=0. |
| BGNARROW | char | Band-gap narrowing model: Y or N (default: N). |
| SRH | char | Shockley-Read-Hall model: Y or N (default: N). |
| TUNNELING | char | Tunneling effect consideration in the SRH model: Y or N (default: N). |
| AUGER | char | Auger ionization model: Y or N (default: N). |
| IMPACT | char | Impact ionization model: Y, I, or N (default: N). Y calculates by averaged electric field and current on grids; I calculates by electric field and current at middle points among grids; N neglects impact ionization. |
| TRAP | char | Trap model at Si-SiO2 interface: Y, P, or N (default: N). Y means electron trap, P means hole trap, N means no trap. |
| FNCURRENT | char | Fowler-Nordheim oxide tunneling current model (default: N). Y considers electron and hole currents; E considers only electron current; H considers only hole current; N does not use FN current model. |

```
[ex.] MODEL MOBILITY=Y BGN=Y SRH=Y AUGER=Y IMPACT=Y FNC=Y
```

> Notice: If MOBILITY=Y is specified, the distance between every mesh point and every interface is calculated. In MOBILITY=P, calculation time is shorter than MOBILITY=Y because only the Si-SiO2 interface at Z=0 is assumed for mobility calculation.

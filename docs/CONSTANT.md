## CONSTANT

This statement specifies physical constant.

```
[format] CONSTANT
    [ TEMPERATURE=<real> EOX=<real> ESI=<real> ENI=<real>
    EIN=<real> EME=<real> EAL=<real> EPS=<real>
    EM1=<real> EM2=<real> EM3=<real> EM4=<real>
    EM5=<real> EM6=<real> EM7=<real> EM8=<real>
    EMOB=<real> HMOB=<real> NI=<real> CN=<real> CP=<real>
    TNO=<real> or <progression> TPO=<real> or <progression>
    VBG=<real> CBG=<real> CSRH=<real>
    AN=<real> AP=<real> BN=<real> BP=<real>
    XREF=<real> ESREF=<real> HSREF=<real>
    SREFN=<real> SREFP=<real> ALPHAN =<real>
    ALPHAP=<real> BETAN=<real> BETAP=<real>
    FACTN=<real> FACTP=<real> CREFN=<real> CREFP=<real>
    VCN=<real> VCP=<real> VSN=<real> VSP=<real>
    EMMIN=<real> HMMIN=<real> NSI=<real> or <progression>
    ASI=<real> or <progression> ASG=<real> or <progression>
    AGE=<real> or <progression> EGGE= <progression>
    QE=<real> or <progression> CTRAP=<real> DTRAP=<real>
    TRAPMOB=<real> TMASS=<real> DRANGE=<real> ET=<real>
    FNMN=<real> FNMP=<real> FNBN=<real> FNBP=<real> ]
```

| Parameter | Type | Meaning |
| --- | --- | --- |
| TEMPERATURE | real | device temperature(unit:K) (default:300) |
| EOX | real | relative permittivity of oxide (default:3.9) |
| ESI | real | relative permittivity of silicon (default:11.7) |
| ENI | real | relative permittivity of nitride (default:7.5) |
| EIN | real | relative permittivity of insulator (default:3.9) |
| EME | real | relative permittivity of metal (default:1E5) |
| EAL | real | relative permitivity of Aluminum (default:1E5) |
| EPS | real | relative permittivity of poly-silicon (default:1E5) |
| EM1 - EM8 | real | relative permittivity of M1 - M8 (default:1.0) |
| EMOB | real | electron mobility in low impurity concentration, low electric field and room temperature (unit:cm ² s ⁻¹ V ⁻¹ ) (default:1430) |
| HMOB | real | hole mobility in low impurity concentration, low electric field and room temperature (unit: $cm^2 \cdot s^{-1} \cdot V^{-1}$ ) (default:460) |
| NI | real | intrinsic carrier concentration in low impurity concentration (unit:cm ⁻³ ) (default: see "physical model") |
| CN | real | Auger coefficient of electron(unit:cm ⁶ s ⁻¹ ) (default: see "physical model") |
| CP | real | Auger coefficient of hole(unit:cm ⁶ s ⁻¹ ) (default: see "physical model") |
| TNO | real or Progression | life time of electron(unit:sec)(default:3.95E-4) pairs of concentration (unit:cm ⁻³ ) and lifetime (unit: sec) should be described in parenthesis (maximum: 100 pairs) |
| TP0 | real or Progression | life time of hole(unit:sec)(default:3.52E-5) pairs of concentration (unit:cm ⁻³ ) and lifetime (unit: sec) should be described in parenthesis (maximum: 100 pairs) |
| CSRH | real | SRH model parameter dependent on impurity concentration(unit:cm ⁻³ ) (default:7.1E15) |
| AN | real | linear part of electron impact ionization coefficient(unit:cm⁻¹) (default: see "physical model") |
| AP | real | linear part of hole impact ionization coefficient (unit:cm⁻¹) (default: see "physical model") |
| BN | real | exponential part of electron impact ionization coefficient(unit:cm ⁻¹ ) (default: see "physical model") |
| BP | real | exponential part of hole impact ionization coefficient (unit:cm ⁻¹ ) (default: see "physical model") |
| VBG | real | energy coefficient of band-gap narrowing effect (unit: eV) (default: 9.0E-3) |
| CBG | real | impurity concentration coefficient of band-gap narrowing effect(unit: cm ⁻³ ) (default:1.0E17) |
| ALPHAN | real | temperature coefficient of electron mobility in low impurity concentration and low electric field (default: -2.0) |
| ALPHAP | real | temperature coefficient of hole mobility in low impurity concentration and low electric field (default: -2.18) |
| EMMIN | real | minimum electron mobility in Caughey and Thomas high impurity concentration model (unit: Vcm/s)(default: see "physical model") |
| HMMIN | real | minimum hole mobility in Caughey and Thomas high impurity concentration model (unit: Vcm/s)(default: see "physical model") |
| CREFN | real | impurity concentration coefficient of electron mobility in Caughey and Thomas high impurity concentration model (unit:cm⁻³) (default: see "physical model") |
| CREFP | real | impurity concentration coefficient of hole mobility in Caughey and Thomas high impurity concentration model (unit:cm⁻³) (default: see "physical model") |
| BETAN | real | exponential part of electron mobility coefficient in Caughey and Thomas high impurity concentration model (default: see "physical model") |
| BETAP | real | exponential part of hole mobility coefficient in Caughey and Thomas high impurity concentration model (default: see "physical model") |
| XREF | real | specific depth of Seavey interface mobility model (unit: nm) |
| ESREF | real | interface electron mobility of Seavey interface mobility model in low electric field (unit: cms ⁻¹ V ⁻¹ ) (default: see "physical model") |
| HSREF | real | interface hole mobility of Seavey interface mobility model in low electric field (unit: cms ⁻¹ V ⁻¹ ) (default: see "physical model") |
| SREFN | real | electron mobility coefficient dependent on electric field of Seavey interface mobility model (unit: Volts/cm) (default:7.0E4) |
| SREFP | real | hole mobility coefficient dependent on electric field of Seavey interface mobility model (unit: Volts/cm) (default:2.7E4) |
| VCN | real | electron mobility longitudinal acoustic phonon velocity of Thornber high electric field model (unit: cm/s) (default:4.9E6) |
| VCP | real | hole mobility longitudinal acoustic phonon velocity of Thornber high electric field model (unit: cm/s) (default:2.928E6) |
| VSN | real | electron saturation velocity of Thornber high electric field model(unit: cm/s) (default: see "physical model") |
| VSP | real | hole saturation velocity of Thornber high electric field model(unit: cm/s) (default: see "physical model") |
| FACTN | real | electron mobility fitting parameter of Thornber high electric field model (default:8.8) |
| FACTP | real | hole mobility fitting parameter of Thornber high electric field model (default:1.6) |
| NSI | real or progression | refractive index of Silicon (default: see "physical model") pairs of wavelength (unit:nm) and refractive index should be described in parenthesis (maximum: 800 pairs) |
| ASI | real or progression | absorption coefficient of Silicon (default: see "physical model") pairs of wavelength (unit:nm ) and absorption coefficient (unit: $\mu$ m ⁻¹ ) should be described in parenthesis (maximum: 800 pairs) |
| ASG | real or progression | absorption coefficient of $Si_{0.5}Ge_{0.5}$ (default: see "physical model") pairs of wavelength (unit:nm ) and absorption coefficient (unit: $\mu$ m $^{-1}$ ) should be described in parenthesis (maximum: 800 pairs) |
| AGE | real or progression | absorption coefficient of Germanium (default: see "physical model") pairs of wavelength (unit:nm ) and absorption coefficient (unit: $\mu$ m $^{-1}$ ) should be described in parenthesis (maximum: 800 pairs) |
| EGGE | progression | Band gap of Silicon Germanium (default: see "physical model") pairs of Germanium concentration (x of Si _1-x Ge _x ) band gap (unit: eV) should be described in parenthesis (maximum: 100 pairs) |
| QE | real or progression | Quantum efficiency of Silicon (default: see "physical model") pairs of wavelength (unit:nm) and quantum efficiency should be described in parenthesis (maximum: 800 pairs) |
| CTRAP | real | trap density of Si-SiO2 interface (unit: cm ⁻² ) (default: 1E11) |
| DTRAP | real | trap depth of Si-SiO2 interface (unit: $\mu\mathrm{m})$ (default: 1E-4) |
| TRAPMOB | real | trapped carrier mobility along depth direction (unit: cm²/(Vsec)) (default: 1E-5) |
| TMASS | real | ratio between effective carrier math and math of electron in SRH tunneling model (default: 0.25) |
| DRANGE | real | diffusion considration coefficient for lateral extention from 1-D impurity profile in DOPE statemet (default: 2.0) (see appendix B) |
| ET | real | Shift value of trap energy level from mid gap energy level (positive for conduction band side) (unit: eV) (default: 0) |
| FNMN | real | FN current model effective mass ratio for electrons (default: 0.26) |
| FNMP | real | FN current model effective mass ratio for holes (default: 0.26) |
| FNBN | real | FN current model oxide barrier height for electrons (unit: eV) (default: 3.1) |
| FNBP | real | FN current model oxide barrier height for holes (unit: eV) (default: 3.8) |

> Notice: If TEMPERATURE is denoted, any parameters denoted before TEMPERATURE are ignored.

```
[ex.]   CONST TEMP=330 EMOB=1200 HMOB=500 EM1=3.9
        CONST ASI=(400, 1.1, 500, 0.8, 600, 0.7, 700, 0.8)
```

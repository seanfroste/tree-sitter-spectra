# Input file format

The input file format is defined by the following rules.

1. An input file consists of statement names, parameter names, equal marks, and parameter values.

   ```
   [ex.] substrate type=p conc=2e15 v=0 "substrate" is a statement name. "type", "conc", and "v" are parameter names. "p", "2e15", and "0" are parameter values.
   ```

2. A row consists of 511 columns. Characters after 511th column are neglected.
3. Upper-case and lower-case are not distinguished except a title, file names, electrode names, quasi-Fermi region names, and text parameters enclosed by single or double quotation marks.
4. A statement may consist of plural rows.

   ```
   [ex.] Dope type=p profile=Gaussian Xmin=0 Ymin=0 Xmax=10 Ymax=5.2 Cmax=2e15 Cb=1.0e14 Xj=3.5 Rp=0
   ```

5. A space or comma is necessary between a parameter value and a next parameter name. [ex.] xmin=0 ymin=3.0,xmax=7.5
6. Spaces among a parameter name, an equal mark, and a parameter name are neglected.

   ```
   [ex.] xmin = $0$ is equivalent to xmin= $0$
   ```

7. A text parameter including a space should be enclosed by single or double quotation marks.

   ```
   [ex.] plot potential dim=3 axis=3 zmin=0 comment="CCD Potential Analysis"
   ```

8. A row whose first column character is "#", "\$", or "\*" is regarded as a comment.
9. The parameter which has "@" at its head is regarded as the user defining parameter which is specified in DEFINE statement.
10. The parameter which has "#" at its head is regarded as the extracted parameter specified in EXTRACT statement or user defining parameter defined in DEFINE statement.
11. The parameter which has "{" at the head and "}" at the tail is regarded as the user defining function, which consists of numbers, user defining parameters, "+", "-", "\*", "\", "\", and the numerical functions shown in Appendix M.

    ```
    [ex.] time=\{1e-10*2^{(@ntime-1)}\}, theta=\{atan2(@y,@x)*180/PI\}
    ```

12. The parameter which has "(" at the head and ")" at the tail is regarded as the numeric series.

    ```
    [ex.] DEFINE NAME=Size VALUE=(0.3 0.5 0.8 1.2 2.0)
    ```

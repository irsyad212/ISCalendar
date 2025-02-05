# IS Calendar
This is a calendar that is so simple (mathematically).
More on that in the "docs.pdf" file.

## Calculations

The suffix 'G' is for Gregorian and 'I' for IS
Results for IS Calendar is $I_y$, $I_m$ and $I_d$
Results for Gregorian Calendar is $G_y$, $G_m$ and $G_d$

### Converting Gregorian to IS 

$G_y = Gmonth()$ \
$G_m = Gmonth()$ \
$G_d = Gday()$ \
<bs>
<bs>
$j = julian(G_y, G_m, G_d) - 2451544.5$ \
$I_y =$ $\lfloor j \div 360 \rfloor$ \
$a = (j \div 360) - I_y$ \
$I_m =$ $\lfloor a \times 12 \rfloor$ \
$b = (a \times 12) - I_m$ \
$I_d =$ $\lfloor b \times 30 \rfloor$ \


### Converting IS to Gregorian

$I_y = Imonth()$ \
$I_m = Imonth()$ \
$I_d = Iday()$ \
<bs>
<bs>
$a = I_d \div 30$ \
$b = I_m + a$ \
$c = b \div 12$ \
$d = I_y + c$ \
$e = (d \times 360) + 2451544.5$ \
$G_y$, $G_m$, $G_d$ $=gregorian(e)$ \




import time, math

def greg2julian(yr, mn, dy):
	dy = dy - 0.5
	a = (14 - mn) // 12
	y = yr + 4800 - a
	m = mn + 12 * a - 3
	return (dy + ((153 * m + 2) // 5) + 365 * y + y // 4 - y // 100 + y // 400 - 32045) 

def greg2IS():
	y = int(input("Year: "))
	m = int(input("Month: "))
	d = int(input("Day: "))

	a = greg2julian(y,m,d) - 2451544.5

	b = a / 360

	c = math.trunc(b) #year

	d = b - c

	e = round(d*12, 5)

	f = math.trunc(e) #month 0-11

	g = e - f

	h = round(g*30, 5)

	i = round(h, 2) 

	j = math.trunc(i) #day 0-29

	print(f"Date in IS is: {c}/{f}/{j}")

def juli2greg(julian):
	a = 1

	b = 1

	j = julian + 0.5
	
	i = math.floor(j)

	f = j - i

	if(i > 2299160):
		a = math.floor((i - 1867216.25)/36524.25)
		b = i + a - (a // 4) + 1
	else :
		b = i

	c = b + 1524

	d = math.floor( (c-122.1) / 365.25)

	e = math.floor(365.25 * d)

	g = math.floor( (c - e) / 30.6001 )

	day = c - e + f - math.floor(30.6001 * g)

	month = 1

	if(g < 13.5):
		month = g - 1
	else :
		month = g - 13

	year = 1

	if(month > 2.5):
		year = d - 4716
	else :
		year = d - 4715

	return year, month, day  

def IS2greg():
	y = int(input("Year: "))
	m = int(input("Month: "))
	d = int(input("Day: "))

	a = d / 30

	b = a + m

	c = b / 12

	d = c + y

	e = d * 360
	
	f = e + 2451544.5

	Y, M, D = juli2greg(f)

	D = math.trunc(D)

	print(f"Date in Greg is: {Y}/{M}/{D}")

def main():
	n = str(input("Type \"1\" to convert greg to IS, Type \"2\" to convert IS to greg: "))

	if(n == '1'):
		greg2IS()
	else:
		IS2greg()

if(__name__ == '__main__'):
	main()

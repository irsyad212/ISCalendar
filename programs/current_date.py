import time, math

def greg2julian(Y,M,D):
	a = (14 - M) // 12
	y = Y + 4800 - a
	m = M + 12 * a - 3
	return (D + ((153 * m + 2) // 5) + 365 * y + y // 4 - y // 100 + y // 400 - 32045) - 0.5

def main():
	y = time.localtime().tm_year
	m = time.localtime().tm_mon
	d = time.localtime().tm_mday

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

if(__name__ == '__main__'):
	main()

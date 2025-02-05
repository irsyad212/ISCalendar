import math

def main():
	y = int(input("Enter year (IS): "))
	m = int(input("Enter month (IS): "))
	d = int(input("Enter day (IS): "))

	days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"]

	a = d / 30

	b = m + a

	c = b / 12

	e = y + c

	f = (e*360)-2

	g = f / 7

	h = g - math.trunc(g)

	i = h*7

	j = int(round(i, 5))

	print(days[j])

main()

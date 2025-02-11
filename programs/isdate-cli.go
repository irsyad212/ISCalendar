package main

import (
	"fmt"
	"log"
	"os"
	"context"
	"strconv"
	"time"

	"github.com/urfave/cli/v3"
)

func greg2julian(Y int, M int, D int) float64 {
	return float64((1461*(int64(Y)+4800+(int64(M)-14)/12))/4 + (367*(int64(M)-2-12*((int64(M)-14)/12)))/12 - (3*((int64(Y)+4900+(int64(M)-14)/12)/100))/4 + int64(D) - 32075) - 0.5
}

func round(n float64, d int) float64 {
	out := float64(0.3)
	f := fmt.Sprintf("%%.%df", d)
	s := fmt.Sprintf(f, n)
	if s, err := strconv.ParseFloat(s, 64); err == nil {
		out = s
	}
	return out
}

func greg2IS(Y int, M int, D int) (int, int, int) {
	A := float64(greg2julian(Y, M, D)) - float64(2451544.5)

	a := int(A)

	b := float64(a) / float64(360)

	c := int(b) //year

	d := float64(b) - float64(c)

	e := round(float64(d) * float64(12), 5)

	f := int(e) //month 0-11

	g := float64(e) - float64(f)

	h := round(float64(g) * float64(30), 5)

	i := round(float64(h), 2)

	j := int(i) //day 0-29

	return c, f, j
}

func juli2greg(julian float64) (int, int, int) {
	a := float64(1)

	b := float64(1)

	j := float64(julian) + float64(0.5)
	
	i := float64(int(j))

	f := float64(j) - float64(i)

	if(i > float64(2299160)) {
		a = float64(int(( float64(i) - float64(1867216.25)) / float64(36524.25) ))

		b = float64(i) + float64(a) - float64((int(a / 4))) + float64(1)
	} else {
		b = float64(i)
	}

	c := float64(b) + float64(1524)

	d := float64( int( ( ( float64(c)-float64(122.1)) / float64(365.25) ) ) )

	e := float64( int(float64(365.25) * float64(d)))

	g := float64( int(( ( float64(c) - float64(e) ) / float64(30.6001) )))

	day := float64(c) - float64(e) + float64(f) - float64( int( (30.6001 * g) ) )

	month := float64(1)

	if(float64(g) < float64(13.5)) {
		month = float64(g) - float64(1)
	} else {
		month = float64(g) - float64(13)

	}

	year := float64(1)

	if(float64(month) > float64(2.5)) {
		year = float64(d) - float64(4716)
	} else {
		year = float64(d) - float64(4715)
	}

	Y := int(year)

	M := int(month)

	D := int(day)

	return Y, M, D
}

func IS2greg(Y int, M int, D int) (int, int, int) {
	a := float64(D) / float64(30)

	b := float64(a) + float64(M)

	c := float64(b) / float64(12)

	d := float64(c) + float64(Y)

	e := float64(d) * float64(360)
	
	f := float64(e) + float64(2451544.5)

	return juli2greg(f)

}

func main() {
	cmd := &cli.Command{
		Name:  "isdate",
		Usage: "Tool for the IS Calendar",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "convert",
				Aliases: []string{"c"},
				Value: "gti, itg",
				Usage: "converter. gti (gregorian to IS), itg (IS to gregorian); date format to enter is YYYY-MM-DD",
			},
			&cli.StringFlag{
				Name: "format",
				Aliases: []string{"f"},
				Value: "default",
				Usage: "currently default, formats the date; only supports IS format only, not gregorian",
			},			
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			now := time.Now()
			date := "0000-00-00"

			if cmd.NArg() > 0 {

				date = cmd.Args().Get(0)

			}

			 if cmd.String("convert") == "gti" {

				t, err := time.Parse("2006-01-02", date);
				if err != nil {
					log.Fatal(err)
				}
				year, month, day := t.Date()
				a, b, c := greg2IS(year, int(month), day)
				fmt.Printf("%d/%d/%d\n", a, b, c)

			} else if cmd.String("convert") == "itg" {

				d := 1
				e := 1
				f := 1
				_, err := fmt.Sscanf(date, "%d/%d/%d", &d, &e, &f) 
				if err != nil { 
					log.Fatal(err)
				} 
				a, b, c := IS2greg(d, e, f)
				fmt.Printf("%04d-%02d-%02d\n", a, b, c)

			} else if cmd.String("format") == "default" {
				
				if date == "0000-00-00" {
					a, b, c := greg2IS(now.Year(), int(now.Month()), now.Day())
					fmt.Printf("%d/%d/%d\n", a, b, c)
				} else {
					d := 1
					e := 1
					f := 1
					_, err := fmt.Sscanf(date, "%d/%d/%d", &d, &e, &f) 
					if err != nil { 
						log.Fatal(err)
					} 
					fmt.Printf("%d-%d-%d\n", d, e, f)
				}
			} else {

				a, b, c := greg2IS(now.Year(), int(now.Month()), now.Day())
				fmt.Printf("%d/%d/%d\n", a, b, c)

			}

			return nil
		},

	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

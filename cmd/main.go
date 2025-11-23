package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lkoman/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "redovalnicaApp",
		Usage: "Upravljanje ocen študentov preko CLI aplikacije.",
		Flags: []cli.Flag{

			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Najmanjše število ocen potrebnih za pozitivno oceno.",
				Value: 3,
			},

			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Najmanjša dovoljena ocena.",
				Value: 5,
			},

			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Največja dovoljena ocena.",
				Value: 10,
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {

			// Preberi vrednosti iz CLI
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")

			fmt.Println("CLI parametri:")
			fmt.Println("Minimalno število ocen:", stOcen)
			fmt.Println("Min ocena:", minOcena)
			fmt.Println("Max ocena:", maxOcena)
			fmt.Println()

			var studenti map[string]redovalnica.Student
			studenti = make(map[string]redovalnica.Student)

			var s1, s2 redovalnica.Student
			s1.Ime = "Laraana"
			s1.Priimek = "Koman"
			s1.Ocene = []int{9, 10, 7}

			s2 = redovalnica.Student{
				Ime:     "Janez",
				Priimek: "Novak",
				Ocene:   []int{5, 5, 5},
			}

			s3 := redovalnica.Student{
				Ime:     "John",
				Priimek: "Wick",
				Ocene:   []int{10, 10, 10},
			}

			studenti["63210152"] = s1
			studenti["63210043"] = s2
			studenti["63210067"] = s3

			redovalnica.IzpisVsehOcen(studenti)

			studenti = redovalnica.DodajOceno(studenti, "63210152", 9)

			redovalnica.IzpisVsehOcen(studenti)

			fmt.Println("Povprecje ocen za studenta", studenti["63210152"].Ime, ":", redovalnica.IzpisiKoncniUspeh(studenti, "63210152"))

			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

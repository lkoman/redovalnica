*Modul redovalnica je preprosta Go knjižnica za delo s podatki o ocenah študentov.*

# Opis
**Struktura Student:** hrani podatke o študentih:
- Ime string
- Priimek string
- Ocene []int

**IzpisVsehOcen(map[string]Student):** izpiše vse študente in njihove ocene.

**DodajOceno(map[string]Student, string, int) map[string]Student:** doda oceno študentu z dano vpisno številko.

**IzpisiKoncniUspeh(map[string]Student, string) float64:** vrne končno povprečje ocen študenta (uporablja interno privatno funkcijo).

**povprecje(...) float64:** privatna funkcija, ki izračuna povprečje ocen.

# Struktura projketa
```
redovalnica/
│   go.mod
│   README.txt
│
├── cmd/
│   └── main.go
│
└── redovalnica/
    └── redovalnica.go
```

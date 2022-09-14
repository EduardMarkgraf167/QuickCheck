# QuickCheck
## Vergleich mit vorgegebenen QuickCheck-Implementierungen
Die QuickCheck-Implementierung in C++ verwendet Templates, um die generische QuickCheck-Funktion zu realisieren. Um die Logik für die Zuordnung der verschiedenen Arbitrary-Funktionen zur generischen QuickCheck-Funktion zu konstruieren wird die überladene Arbitrary-Funktion genutzt. Hierbei gibt es Implementierungen der Arbitrary-Funktion für jeden benötigten Typen.

Die QuickCheck-Implementierung in Haskell verwendet eine generische Arbitrary-Klasse, welche eine Methode fordert, die eine Arbitrary-Methode mit dem dem Typenparameter entsprechenden Rückgabetypen enthält. Hierfür werden verschiedene Instanzen, mit den entsprechenden Implementierungen der Arbitrary-Methode genutzt. Um die Funktionalität zum Wählen eines zufälligen Elements und zur Generierung einer Sammlung mit Zufallszahlen werden Hilfsfunktionen verwendet. Zuletzt wird eine generische QuickCheck-Funktion entworfen. Hierbei werden nur Typenparameter angenommen, welche die Constraints "Show" und "Arbitrary" erfüllen. Das heißt, es muss Implementierungen der Klasse Arbitrary, sowie der Klasse Show geben, die mit dem Typen umgehen können müssen. So kann sichergestellt werden, dass die Arbitrary-Methode und die Konsolenausgabe erfolgreich durchgeführt weden können. Die Zuordnung wird so durch Überladung vollzogen. 

Um die generische QuickCheck-Funktion in Go zu implementieren, wird das Sprachfeature Generics in Go benutzt. An dieser Stelle tritt ein Problem auf, das dem Design der Sprache Go geschuldet ist. Go begrenzt Funktionsüberladungen auf den Receiver. Somit ist die Zuordnung der jeweiligen Arbitrary-Implementierung anders zu lösen. Hierfür wird der Ansatz aus der QuickCheck-Implementierung in C++ verwendet. So operieren die Arbitrary-Funktionen auf Proxy Receivern. Hier findet sich eine Art von AdHoc-Polymorphie. Das Arbitrary-Interface wird durch Duck-Typing von jeder Struktur implementiert, die eine Generate-Mehtode besitzt. Das Arbitrary-Interface ist ein generisches Interface, wobei der definierte Typen-Parameter den Rückgabetypen bestimmt. Folglich implementieren die einzelnen Arbitrary-Typen das Arbitrary-Interface eines bestimmten Typen, welcher durch den Rückgabetyp der zu implementierenden Generate-Methode bestimmt wird. Um nun die Zuordnung einer Arbitrary-Funktion wie in der C++-Implementierung umzusetzen, wird der QuickCheck-Funktion eine Arbitrary-Struktur übergeben. Hierbei muss die Arbitrary-Struktur das Arbitrary-Interface implementieren, das denselben Typenparameter besitzt, wie der Typenparameter der QuickCheck-Funktion. Somit ist sichergestellt, dass keine inkompatiblen Generator-Strukturen übergeben werden.

## QuickCheck Implementierung in der Standard-Bibliothek von Go
Im Standardpaket "testing/quick" der Go-Standardbibliothek ist eine QuickCheck-Implementierung enthalten. 

Hier gibt es folgende öffentliche Funktionen:

### func Check(f any, config *Config) error
Die Funktion f ist eine Funktion, die einen boolschen Wert als Rückgabetyp enthält. Die Funktion f wird mit verschiedenen zufälligen Eingabewerten ausgeführt, bis entweder das Limit für die Versuche erreicht wird oder f false zurückgibt.
#### Beispiel
```golang
func TestOddMultipleOfThree(t *testing.T) {
  f := func(x int) bool {
    y := OddMultipleOfThree(x)
    return y%2 == 1 && y%3 == 0
  }
  if err := quick.Check(f, nil); err != nil {
    t.Error(err)
  }
}
```

Im obigen Beispiel wird zuerst eine Funktion f definiert, die die zu testende Funktion "OddMultipleOfThree" aufruft und anschließend überprüft, ob der Rückgabewert ungerade und durch 3 teilbar ist. Diese wird mithilfe von quick.Check() überprüft und anschließend ein Fehler geworfen, falls die Überprüfung einen Eingabewert ergeben hat, der dazu führt, dass die zu prüfende Funktion "OddMultipleOfThree" die in f definierten Bedingungen nicht erfüllt.

```golang
func OddMultipleOfThree(x int) int {
  start := x * 3
  cnt := 0
  for {
  	ret := start + 3*cnt
  	if ret%2 == 1 {
  		return ret
  	}
  	cnt += 1
  }
}
```
Im obigen Quellcode-Ausschnitt ist eine beispielhafte Implementierung von OddMultipleOfThree zu sehen. Sie nimmt einen ganzzahligen Wert als Parameter, multipliziert sie mit 3 und berechnet die weitere Dreierreihe bis eine ungerade Summe entsteht. Auf den ersten Blick scheint die beispielhafte Implementierung korrekt zu sein. Auf den zweiten Blick lässt sich leicht erkennen, dass der Rückgabewert der Funktion bei einem Aufruf mit 0 als Parameter die aufgestellte Bedingung nicht erfüllt. Durch den Einsatz von QuickCheck wird jedoch selbst wenn der Nullfall abgedeckt wurde kein erfolgreicher Durchlauf erzielt. Der Fehler könnte in etwa lauten: "failed on input 4601851300195147788". Hier lässt sich ein Fehler erkennen, der sonst wahrscheinlich nicht aufgefallen wäre. Wird ein hoher Integerwert mit 3 multiplizieren, so übersteigt er den maximalen Wert, den eine Integer-Variable speichern kann. Daher wird die Funktion angepasst:

```golang
func OddMultipleOfThree(x int) int {
  if x == 0 || x > math.MaxInt / 3 {
  	// Error handling
  }
  
  start := x * 3
  cnt := 0
  for {
  	ret := start + 3*cnt
  	if ret%2 == 1 {
  		return ret
  	}
  	cnt += 1
  }
}
```
Nun wird QuickCheck erneut eingesetzt und meldet immernoch einen Fehler. Dieser tritt für hohe negative Werte für x auf, wie z.B. -4576219449590091309. Also wird die Funktion erneut angepasst:

```golang
func OddMultipleOfThree(x int) int {
  if x <= 0 && x > math.MaxInt / 3 {
  	// Error handling
  }
  
  start := x * 3
  cnt := 0
  for {
  	ret := start + 3*cnt
  	if ret%2 == 1 {
  		return ret
  	}
  	cnt += 1
  }
}
```

Ein paar letzte Durchläufe mit QuickCheck zeigen keine weiteren Fehler an. Somit ist die Funktion vollständig.

Eine weitere Möglichkeit die Sonderfälle abzufangen wäre es, die QuickCheck Property anzupassen. So könnte auch definiert werden, dass auch negative Zahlen angenommen werden. Hierfür müsste das Property angepasst werden und wie folgt lauten: y%2 != 0 && y%3 == 0. Ebenfalls wäre es möglich, die Sonderfälle in der Property-Definition abzufangen und zu überspringen. So kann das Property flexibel angepasst werden.

#### Verwendung
### func CheckEqual(f, g any, config *Config) error
Vergleicht zwei Funktionen f und g. Zufällige Eingabewerte werden generiert bis die Funktionen unterschiedliche Ergebnisse liefern. Das Ergebnis ist ein CheckEqualError

### func Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool)
Es wird ein zufälliges Element des übergebenen Typs zurückgegeben. Voraussetzung ist, dass t dem Generator-Interface genügt.

### Errors
Die "testing/quick" Implementierung definiert eine die Fehlertypen CheckEqualError und CheckError. Diese werden von den oben aufgelisteten Funktionen zurückgegeben, falls die Überprüfung fehlschlägt. Zudem gibt es einen SetupError, welcher dann geworfen wird, wenn eine fehlerhafte Bedienung der Funktionen stattfindet.

### Config
Eine Config-Struktur kann verwendet werden, um die Funktionen Check() und CheckEqual() zu konfigurieren. Hierbei können vier Optionen gesetzt werden.

#### MaxCount
MaxCount bestimmt, wieviele Iterationen in der Überprüfung durchlaufen werden, bis sie als erfolgreich gilt. Wird MaxCount auf 0 gesetzt, so wird stattdessen auf MaxCountScale zurückgegriffen.

#### MaxCountScale
MaxCountScale ist eine nicht-negative float-Zahl, die ein Vielfaches der Standardeinstellung für die maximalen Iterationen angibt. Dieses ist standardmäßig auf 100 gesetzt, kann aber mithilfe der Flag -quickchecks gesetzt werden.

#### Rand
Rand gibt einen Pool zufälliger Zahlen an. Ist Rand nil, so wird ein Pool pseudorandomisierter Zahlen generiert.

#### Values
Values spezifiziert eine Funktion zur Generierung einer Slice mit zufälligen Werten, welche zur zu testenden Funktion passen.

### Generator
Das Generator-Interface setzt folgende Funktion voraus:

```golang
Generate(rand *rand.Rand, size int) reflect.Value
```

Hierbei gibt Generate einen zufälligen Wert ebendieses Typs zurück, der das Generator-Interface implementiert.
Der Parameter "size" wird als Hinweis auf die zu generierenden Werte genutzt. 

#### Implementierung
```golang
func Check(f any, config *Config) error {
	if config == nil {
		config = &defaultConfig
	}

	fVal, fType, ok := functionAndType(f)
	if !ok {
		return SetupError("argument is not a function")
	}

	if fType.NumOut() != 1 {
		return SetupError("function does not return one value")
	}
	if fType.Out(0).Kind() != reflect.Bool {
		return SetupError("function does not return a bool")
	}

	arguments := make([]reflect.Value, fType.NumIn())
	rand := config.getRand()
	maxCount := config.getMaxCount()

	for i := 0; i < maxCount; i++ {
		err := arbitraryValues(arguments, fType, config, rand)
		if err != nil {
			return err
		}

		if !fVal.Call(arguments)[0].Bool() {
			return &CheckError{i + 1, toInterfaces(arguments)}
		}
	}

	return nil
}
```
Zu Beginn wird geprüft, ob eine Konfiguration übergeben wurde. Falls nicht wird auf die Standardkonfiguration zurückgegriffen. Anschließend wird die übergebene Funktion f überprüft und mithilfe von Reflection durch die Funktion "functionAndType" der Typ der Funktione ermittelt. Die Überprüfung wird in der Variable "ok" gespeichert und stellt sicher, dass der übergebene Pointer auf eine Funktion zeigt. Des Weiteren wird geprüft, dass die übergebene Funktion nur einen Rückgabewert hat und dieser den Typ Boolean hat.
```golang
func functionAndType(f any) (v reflect.Value, t reflect.Type, ok bool) {
	v = reflect.ValueOf(f)
	ok = v.Kind() == reflect.Func
	if !ok {
		return
	}
	t = v.Type()
	return
}
```
Wenn mit den Eingabeparametern alles in Ordnung ist wird die Funktion initialisiert. Hierbei wird eine Slice erstellt, die die Argumente der zu prüfenden Funktion darstellt. Der Zufallsgenerator wird aus der Konfiguration entnommen, ebenso wie der maxCount-Wert.

In der Überprüfungsphase werden die Argumente mithilfe der Funktion "arbitraryValues" zufällig generiert und die übergebene Funktion damit geprüft. Gibt diese False zurück, so bricht die Funktion ab und wirft einen CheckError.

Die Funktion "arbitraryValues" basiert auf der Funktion "Value", die wiederum auf der Funktion "sizedValue" basiert:
```golang
// sizedValue returns an arbitrary value of the given type. The size
// hint is used for shrinking as a function of indirection level so
// that recursive data structures will terminate.
func sizedValue(t reflect.Type, rand *rand.Rand, size int) (value reflect.Value, ok bool) {
	if m, ok := reflect.Zero(t).Interface().(Generator); ok {
		return m.Generate(rand, size), true
	}

	v := reflect.New(t).Elem()
	switch concrete := t; concrete.Kind() {
	case reflect.Bool:
		v.SetBool(rand.Int()&1 == 0)
	case reflect.Float32:
		v.SetFloat(float64(randFloat32(rand)))
	case reflect.Float64:
		v.SetFloat(randFloat64(rand))
	case reflect.Complex64:
		v.SetComplex(complex(float64(randFloat32(rand)), float64(randFloat32(rand))))
	case reflect.Complex128:
		v.SetComplex(complex(randFloat64(rand), randFloat64(rand)))
	case reflect.Int16:
		v.SetInt(randInt64(rand))
	case reflect.Int32:
		v.SetInt(randInt64(rand))
	case reflect.Int64:
		v.SetInt(randInt64(rand))
	case reflect.Int8:
		v.SetInt(randInt64(rand))
	case reflect.Int:
		v.SetInt(randInt64(rand))
	case reflect.Uint16:
		v.SetUint(uint64(randInt64(rand)))
	case reflect.Uint32:
		v.SetUint(uint64(randInt64(rand)))
	case reflect.Uint64:
		v.SetUint(uint64(randInt64(rand)))
	case reflect.Uint8:
		v.SetUint(uint64(randInt64(rand)))
	case reflect.Uint:
		v.SetUint(uint64(randInt64(rand)))
	case reflect.Uintptr:
		v.SetUint(uint64(randInt64(rand)))
	case reflect.Map:
		numElems := rand.Intn(size)
		v.Set(reflect.MakeMap(concrete))
		for i := 0; i < numElems; i++ {
			key, ok1 := sizedValue(concrete.Key(), rand, size)
			value, ok2 := sizedValue(concrete.Elem(), rand, size)
			if !ok1 || !ok2 {
				return reflect.Value{}, false
			}
			v.SetMapIndex(key, value)
		}
	case reflect.Pointer:
		if rand.Intn(size) == 0 {
			v.Set(reflect.Zero(concrete)) // Generate nil pointer.
		} else {
			elem, ok := sizedValue(concrete.Elem(), rand, size)
			if !ok {
				return reflect.Value{}, false
			}
			v.Set(reflect.New(concrete.Elem()))
			v.Elem().Set(elem)
		}
	case reflect.Slice:
		numElems := rand.Intn(size)
		sizeLeft := size - numElems
		v.Set(reflect.MakeSlice(concrete, numElems, numElems))
		for i := 0; i < numElems; i++ {
			elem, ok := sizedValue(concrete.Elem(), rand, sizeLeft)
			if !ok {
				return reflect.Value{}, false
			}
			v.Index(i).Set(elem)
		}
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			elem, ok := sizedValue(concrete.Elem(), rand, size)
			if !ok {
				return reflect.Value{}, false
			}
			v.Index(i).Set(elem)
		}
	case reflect.String:
		numChars := rand.Intn(complexSize)
		codePoints := make([]rune, numChars)
		for i := 0; i < numChars; i++ {
			codePoints[i] = rune(rand.Intn(0x10ffff))
		}
		v.SetString(string(codePoints))
	case reflect.Struct:
		n := v.NumField()
		// Divide sizeLeft evenly among the struct fields.
		sizeLeft := size
		if n > sizeLeft {
			sizeLeft = 1
		} else if n > 0 {
			sizeLeft /= n
		}
		for i := 0; i < n; i++ {
			elem, ok := sizedValue(concrete.Field(i).Type, rand, sizeLeft)
			if !ok {
				return reflect.Value{}, false
			}
			v.Field(i).Set(elem)
		}
	default:
		return reflect.Value{}, false
	}

	return v, true
}
```
``golang
func Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool) {
	return sizedValue(t, rand, complexSize)
}
```
Hierbei lässt sich ein TypeSwitch erkennen, welcher basierend auf dem zu generierenden Typ eine entsprechende Arbitrary-Funktion auswählt, um einen zufälligen Wert des gesuchten Typs zu generieren. Interessant ist hier, dass auch komplexe Typen wie Structs, Arrays, Slices und Strings generiert werden können. Am interessantesten ist die implementierung für Strukturen, da sie beliebig aufgebaut sein und verschiedene Typen beinhalten können. In Zeile 166 wird dieser Fall abgedeckt. 
Zu Beginn wird die Menge der Felder der jeweiligen Struktur ermittelt und in n gespeichert. Anschließend wird der Wert für sizeLeft berechnet, welcher dazu genutzt wird, sizeValue rekursiv aufzurufen. Auf ähnliche Weise funktionieren auch die übrigen Komplex-Typen.

### Fazit
Bei der Implementierung von QuickCheck in Go wird der Generierungsprozess standardisiert und macht sich den Umstand zu Nutze, dass jeder übergebene Typ letztlich aus primitiven Typen aufgebaut ist. Hier ist der Nutzer auf die Implementierung angewiesen und muss sich an diese Anpassen.

In der hier entwickelten Version basiert QuickCheck darauf, dass Benutzer ihre eigenen Typen entwickeln können und die Funktionalität so nach Belieben erweitern können. Dies geht für einen Benutzer zwar mit einem gewissen Mehraufwand einher, bietet jedoch mehr Flexibilität. Die Standard-Implementierung erspart das eigene Implementieren des Nutzers, reduziert damit jedoch die Erweiterbarkeit und Flexibilität. 

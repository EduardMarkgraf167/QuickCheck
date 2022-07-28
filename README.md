# QuickCheck
## Vergleich mit vorgegebenen QuickCheck-Implementierungen
Die QuickCheck-Implementierung in C++ verwendet Templates, um die generische QuickCheck-Funktion zu realisieren. Um die Logik für die Zuordnung der verschiedenen Arbitrary-Funktionen zur generischen QuickCheck-Funktion zu konstruieren wird die überladene Arbitrary-Funktion genutzt. Hierbei gibt es Implementierungen der Arbitrary-Funktion für jeden benötigten Typen.

Die QuickCheck-Implementierung in Haskell verwendet eine generische Arbitrary-Klasse, welche eine Methode fordert, die eine Arbitrary-Methode mit dem dem Typenparameter entsprechenden Rückgabetypen enthält. Hierfür werden verschiedene Instanzen, mit den entsprechenden Implementierungen der Arbitrary-Methode genutzt. Um die Funktionalität zum Wählen eines zufälligen Elements und zur Generierung einer Sammlung mit Zufallszahlen werden Hilfsfunktionen verwendet. Zuletzt wird eine generische QuickCheck-Funktion entworfen. Hierbei werden nur Typenparameter angenommen, welche die Constraints "Show" und "Arbitrary" erfüllen. Das heißt, es muss Implementierungen der Klasse Arbitrary, sowie der Klasse Show geben, die mit dem Typen umgehen können müssen. So kann sichergestellt werden, dass die Arbitrary-Methode und die Konsolenausgabe erfolgreich durchgeführt weden können. Die Zuordnung wird so durch Überladung vollzogen. 

Um die generische QuickCheck-Funktion in Go zu implementieren, wird das Sprachfeature Generics in Go benutzt. An dieser Stelle tritt ein Problem auf, das dem Design der Sprache Go geschuldet ist. Go begrenzt Funktionsüberladungen auf den Receiver. Somit ist die Zuordnung der jeweiligen Arbitrary-Implementierung anders zu lösen. Hierfür wird der Ansatz aus der QuickCheck-Implementierung in C++ verwendet. So operieren die Arbitrary-Funktionen auf Proxy Receivern. Hier findet sich eine Art von AdHoc-Polymorphie. Das Arbitrary-Interface wird durch Duck-Typing von jeder Struktur implementiert, die eine Generate-Mehtode besitzt. Das Arbitrary-Interface ist ein generisches Interface, wobei der definierte Typen-Parameter den Rückgabetypen bestimmt. Folglich implementieren die einzelnen Arbitrary-Typen das Arbitrary-Interface eines bestimmten Typen, welcher durch den Rückgabetyp der zu implementierenden Generate-Methode bestimmt wird. Um nun die Zuordnung einer Arbitrary-Funktion wie in der C++-Implementierung umzusetzen, wird der QuickCheck-Funktion eine Arbitrary-Struktur übergeben. Hierbei muss die Arbitrary-Struktur das Arbitrary-Interface implementieren, das denselben Typenparameter besitzt, wie der Typenparameter der QuickCheck-Funktion. Somit ist sichergestellt, dass keine inkompatiblen Generator-Strukturen übergeben werden.

## QuickCheck Implementierungen in Go
Im Standardpaket "testing/quick" der Go-Standardbibliothek ist eine QuickCheck-Implementierung enthalten. 

Hier gibt es folgende öffentlich Funktionen:

### func Check(f any, config *Config) error
Sucht nach einem Eingabewert für f (Funktion die bool zurückgibt), sodass f false zurückgibt. Wird ein solcher Wert gefunden, so wird ein CheckError zurückgegeben

### func CheckEqual(f, g any, config *Config) error
Vergleicht zwei Funktionen f und g. Zufällige Eingabewerte werden generiert bis die Funktionen unterschiedliche Ergebnisse liefern. Ergebnis ist ein CheckEqualError

### func Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool)
Es wird ein zufälliges Element des übergebenen Typs zurückgegeben. Voraussetzung ist, dass t dem Generator-Interface genügt.

Bei der Implementierung werden keine generischen Typen verwendet. Stattdessen wird der Typ als Parameter übergeben. Dieser wird durch Reflection zur Laufzeit ermittelt und durch ein übergreifendes Switch-Statement zu einer Arbitrary-Funktion zugeordnet.

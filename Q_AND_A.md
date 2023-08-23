## Wie ist die Laufzeit Ihres Programms?

Die Hauptmerging-Funktion hat eine Laufzeit von O(n), wobei n die Anzahl der Intervalle ist. (Weitere Informationen finden Sie im Code selbst)

## Wie kann die Robustheit sichergestellt werden, vor allem auch mit Hinblick auf sehr große Eingaben?

Ich habe die Eingabe durch reguläre Ausdrücke validiert und spezifische Grenzprüfungen durchgeführt. Für sehr große Eingaben werden durch die Grenzwerte (maximal 20 Paare, Zahlen nicht größer als 10.000) Überlauf und extreme Eingabewerte vermieden. Umfangreiche Tests helfen, um die Funktionalität in verschiedenen Szenarien zu gewährleisten.

## Wie verhält sich der Speicherverbrauch ihres Programms?

Der Speicherverbrauch ist hauptsächlich O(n), proportional zur Größe der Eingabeintervalle. Hilfsstrukturen nehmen zusätzlichen Speicher in Anspruch, sind aber im Gesamtkontext vernachlässigbar. (Weitere Informationen finden Sie im Code selbst)
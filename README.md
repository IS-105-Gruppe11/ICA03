                                                      ICA03 - Gruppe 11
                                                      
                                                      
                                                      
                                                      Oppgave 1A
Formål: - bli kjent med ASCII tabellen og kunne manipulere disse i et programmeringsmiljø 
bruk filen ascii.go i REP_ICA03 som utgangspunkt, og skriv ut en tabell av alle tegn i “string literal” ascii (deklarert i filen ascii.go) 

Resultatet i oppgave 1a ble at vi skriver ut en tabell som skriver alle tegn i “string literal” ascii. Nedenfor vises det en “screenshot” av ascii tabellen. Dette løste vi ved å opprette en funksjon i ascii.go og main.go, og videre en FOR løkke i ascii.go som bygger opp på den, og tilslutt en print funksjon som printer resultatet. 
//link til repository: https://github.com/IS-105-Gruppe11/ICA03/tree/master/oppgave1(a)


På virtuell server ble resultatet (noe av det):

// Screenshot til oppgave 1A: http://imgur.com/a/o3Q55












                                                      Oppgave 1b
lag en funksjon greetingASCII() i samme filen ascii.go, som skriver ut "Hello :-)"

Funksjonen nedenfor skal kunne generere en utskrift fra en sekvens av bytes, altså det fant vi ved å finne den “heksdesimale” representasjonen av tegnene, som var som følgende; "\x22\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29\x22", videre lager vi en FOR løkke som går over stengene, og tilslutt en printF funksjon som viser oss resultatet.,
Link til github repository: https://github.com/IS-105-Gruppe11/ICA03/tree/master/Oppgave1(b)

Virtuell maskin:

Resultatet på lokal og virtuell maskin ble likt.


Screenshot til oppgave 1b: http://imgur.com/a/qFDNg




                                                      Oppgave 1c
Implementer en test for funksjonen greetingASCII() i egen fil ascii_test.go, som tester om returverdier (av type string) inneholder kun ASCII-tegn. 

Lagde en test fil ascii_test.go, med to funskjoner; “TestGreetingASCII” og “isASCII”. TestGreetingASCII er en testfunksjon som tester om GreetingASCII funkjsonen kun inneholder 7 bit Ascii tegn. Funksjonen har en if test som referer til isASCII funksjonen. Den har en løkke som tester hver karakter “c” i stringen den får inn som parameter. Den igjen har en if test som returnerer false om karakteren er et desimaltall større enn 127 (Innenfor 7 bit Ascii). Dette vil føre til at testen feiler.
Link: https://github.com/IS-105-Gruppe11/ICA03/tree/master/Oppgave1(c)/test
Bilde: 

Screenshot til oppgave 1c: http://imgur.com/a/bssLZ













                                                      Oppgave 2a
Formål: - Bli kjent med ISO/IEC 8859 serier for 8-bits koding av typografiske symboler. - Illustrere forskjell på ASCII og utvidet ASCII kode gjennom golang rammeverk for behandling av tekststrenger (på engelsk brukes det nesten alltid begrepet “strings” for tekststrenger). 

Funksjonen itererer (går i en løkke over) over tegn med byte-verdier fra 0x80 til 0xFF, dvs. det utvidede ASCII settet; funksjonen skal implementeres i filen iso.go, som dere finner i mappen iso i REP_ICA03

Denne oppgaven var mye lik oppgave 1a, men i dette tilfelle var det begrensing på ascii verdiene fra 0x80 til 0xFF. Dette implementerte vi via en funksjon i iso.go -> videre en for løkke som går over ascii verdiene, og tilslutt en main funksjon for å kjøre den. Som nevnt var oppgaven mye lik Oppgave 1A
https://github.com/IS-105-Gruppe11/ICA03/tree/master/Oppgave%20%202A

Vi kan se at karakterene har en verdi på 8 bits.
Server:

Screenshot til oppgave 2a: http://imgur.com/a/Qfnk4










                                                      Oppgave2b
lag en funksjon greetingExtendedASCII() i samme filen iso.go, som skriver ut "Salut, ça va °-) €50"

Gjør det samme som i oppgave 1B, men finner hexadesimal-tallet til bokstavene i “Salut, ça va °-) €50"
//link til github: https://github.com/IS-105-Gruppe11/ICA03/tree/master/Oppgave2b/iso



windows:

mac:

server:
 // screenshot: http://imgur.com/a/ud3TR

Vi kan se at det er noen forskjeller i resultat på de ulike plattformene. Ingen av plattformene viser eurotegnet. 

                                                     
                                                     
                                                     
                                                     Oppgave2c
Implementer en test for funksjonen greetingExtendedASCII() i egen fil iso_test.go, som tester om returverdier (av type string) inneholder kun tegn fra en Extended ASCII. 

Lik som i oppgave 1C. Endrer på isASCII funksjonen slik at den sjekket om karakteren “c” er mellom 127 og 255, som er området for ExtendendAscii. 
Github repository: https://github.com/IS-105-Gruppe11/ICA03/tree/master/Oppgave2c



Screenshot: http://imgur.com/a/oBVW3





                                                      Oppgave 3a
Forklar de store forskjellene i resultater av utskrift i eksemplene ovenfor (%s, %q, %+q og %c). Hva må man endre i bytesekvensen for at fmt.Printf("%s", ) returnerer "½?=? ⌘" 

Svar:
Grunnen til den store forskjellen i utskrift i de forskjellige eksemplene er fordi de har ulik formaterings grunnlag. For eksempel %s printer ut kodene som string, mens % X skriver ut hver sekvens av byte med mellomrom. Så selv om tallene som blir brukt er like, blir de printet ut på ulikt grunnlag.
Slik må byte sekvensen se ut for å returnere "½?=? ⌘": \xc2\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98
Bilde:

screenshot: http://imgur.com/a/HMvAr




                                                      Oppgave 3c
Skriv en funksjon PrintTreasureUTF8(treasure_string string) []byte {} som returnerer den teksten som skjuler seg bak koden kodet

Denne oppgaven ble løst ved å skrive en PrintTreasureUTF8 funksjon, som retunerer den teksten som skjuler seg bak koden. Kodesettet som skjulte seg bak koden fant vi ut ved å importere “bytes” inn i  “treasure filen” og videre en func replace, som det er skrevet i ICA-03 filen, ikke minst en func i main.go
https://github.com/IS-105-Gruppe11/ICA03/tree/master/Oppgave3C



Screenshot:  http://imgur.com/a/Zv4Tw

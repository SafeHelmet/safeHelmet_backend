# TODO
- [ ] CRUD database (sui pacchetti MQTT scompattati)
- [ ] Gestione MQTT

## Requisiti del Progetto Digital Twin

### Generali
- [ ] Ogni oggetto (casco, cantiere) deve avere un Digital Twin.
- [ ] Associare ogni operatore a un casco.
- [ ] Creare tabelle nel database per memorizzare i dati dei caschi e dei cantieri.

### Caschi
- [ ] Creare una tabella per memorizzare i dati di ogni casco.
  - [ ] ID Casco
  - [ ] ID Operatore
  - [ ] Dati dei sensori (es. inquinamento, temperatura, umidità)
  - [ ] Timestamp delle letture
- [ ] Aggiungere letture dei sensori al database.
- [ ] Visualizzare i dati di ogni casco.

### Cantieri
- [ ] Creare una tabella per memorizzare i dati del cantiere.
  - [ ] ID Cantiere
  - [ ] Eventi importanti (es. media inquinamento, incidenti)
  - [ ] Timestamp degli eventi
- [ ] Calcolare la media dei valori di inquinamento del monossido di carbonio di tutti i caschi nel cantiere.
- [ ] Salvare i valori medi loggati ogni mezz'ora.
- [ ] Controllare l'esposizione degli operatori a valori di inquinamento.

### Gestione Utenti
- [ ] Creare nuovi utenti.
- [ ] Gestire assegnazioni utenti ai cantieri.
- [ ] Visualizzare le assegnazioni e i dati degli utenti.

## Funzioni CRUD
- [ ] Aggiungere letture sensori al database.
- [ ] Creare nuovi utenti.
- [ ] Gestire assegnazioni utenti ai cantieri.
- [ ] Visualizzare dati di caschi e cantieri.

## MQTT
- [ ] Gestire la comunicazione MQTT per ricevere i dati dai sensori.
- [ ] Scompattare i pacchetti MQTT e salvare i dati nel database.

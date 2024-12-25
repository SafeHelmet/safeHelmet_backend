# API Documentation

## **Lavoratori e Cantieri**

### Lista dei lavoratori in un cantiere
- **Endpoint:** `GET /worksites/:worksiteId/workers`
- **Descrizione:** Restituisce l'elenco dei lavoratori attualmente assegnati al cantiere specificato.
- **Parametri:**
  - `worksiteId` (path): ID del cantiere.

---

### Dettagli lavoratore
- **Endpoint:** `GET /workers/:workerId`
- **Descrizione:** Restituisce i dettagli di un lavoratore, incluse specializzazioni e assegnazioni attuali.
- **Parametri:**
  - `workerId` (path): ID del lavoratore.

---

## **Letture**

### Letture specifiche per operaio, cantiere e giorno
- **Endpoint:** `GET /workers/:workerId/worksites/:worksiteId/readings`
- **Descrizione:** Ottieni le letture di un casco associato a un lavoratore per un determinato cantiere e giorno.
- **Parametri:**
  - `workerId` (path): ID del lavoratore.
  - `worksiteId` (path): ID del cantiere.
  - `date` (query): Data specifica in formato `YYYY-MM-DD`.

---

### Letture di un cantiere in una finestra temporale
- **Endpoint:** `GET /worksites/:worksiteId/readings`
- **Descrizione:** Restituisce tutte le letture di un cantiere in un intervallo temporale.
- **Parametri:**
  - `worksiteId` (path): ID del cantiere.
  - `from` (query): Inizio finestra temporale (`YYYY-MM-DD HH:mm:ss`).
  - `to` (query): Fine finestra temporale (`YYYY-MM-DD HH:mm:ss`).

---

### Letture anomale di un cantiere
- **Endpoint:** `GET /worksites/:worksiteId/readings/anomalous`
- **Descrizione:** Restituisce tutte le letture anomale associate a un cantiere.
- **Parametri:**
  - `worksiteId` (path): ID del cantiere.

---

### Letture anomale di un lavoratore
- **Endpoint:** `GET /workers/:workerId/readings/anomalous`
- **Descrizione:** Restituisce tutte le letture anomale associate a un lavoratore.
- **Parametri:**
  - `workerId` (path): ID del lavoratore.

---

### Ore di funzionamento del casco in una finestra temporale
- **Endpoint:** `GET /helmets/:helmetId/operating-hours`
- **Descrizione:** Calcola le ore di funzionamento di un casco in un intervallo temporale.
- **Parametri:**
  - `helmetId` (path): ID del casco.
  - `from` (query): Inizio finestra temporale (`YYYY-MM-DD HH:mm:ss`).
  - `to` (query): Fine finestra temporale (`YYYY-MM-DD HH:mm:ss`).

---

### Tutte le letture di un casco in una finestra temporale
- **Endpoint:** `GET /helmets/:helmetId/readings`
- **Descrizione:** Restituisce tutte le letture di un casco in un intervallo temporale.
- **Parametri:**
  - `helmetId` (path): ID del casco.
  - `from` (query): Inizio finestra temporale (`YYYY-MM-DD HH:mm:ss`).
  - `to` (query): Fine finestra temporale (`YYYY-MM-DD HH:mm:ss`).

---

## **Gestione e Configurazione**

### Lista dei cantieri
- **Endpoint:** `GET /worksites`
- **Descrizione:** Restituisce l'elenco di tutti i cantieri attivi.

---

### Assegna lavoratore a cantiere
- **Endpoint:** `POST /worksites/:worksiteId/workers`
- **Descrizione:** Assegna un lavoratore a un cantiere con un casco specifico.
- **Parametri:**
  - `worksiteId` (path): ID del cantiere.
- **Body:**
  ```json
  {
    "workerId": 123,
    "helmetId": 456
  }

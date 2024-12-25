# API Documentation

## **Workers and Worksites**

### List of Workers in a Worksite
- **Endpoint:** `GET /worksites/:worksiteId/workers`
- **Description:** Returns the list of workers currently assigned to the specified worksite.
- **Parameters:**
  - `worksiteId` (path): ID of the worksite.

---

### Worker Details
- **Endpoint:** `GET /workers/:workerId`
- **Description:** Returns details about a worker, including specializations and current assignments.
- **Parameters:**
  - `workerId` (path): ID of the worker.

---

## **Readings**

### Specific Readings for Worker, Worksite, and Date
- **Endpoint:** `GET /workers/:workerId/worksites/:worksiteId/readings`
- **Description:** Retrieves helmet readings associated with a worker for a specific worksite and date.
- **Parameters:**
  - `workerId` (path): ID of the worker.
  - `worksiteId` (path): ID of the worksite.
  - `date` (query): Specific date in `YYYY-MM-DD` format.

---

### Worksite Readings in a Time Window
- **Endpoint:** `GET /worksites/:worksiteId/readings`
- **Description:** Returns all readings for a worksite within a specified time window.
- **Parameters:**
  - `worksiteId` (path): ID of the worksite.
  - `from` (query): Start of the time window (`YYYY-MM-DD HH:mm:ss`).
  - `to` (query): End of the time window (`YYYY-MM-DD HH:mm:ss`).

---

### Anomalous Readings for a Worksite
- **Endpoint:** `GET /worksites/:worksiteId/readings/anomalous`
- **Description:** Returns all anomalous readings associated with a worksite.
- **Parameters:**
  - `worksiteId` (path): ID of the worksite.

---

### Anomalous Readings for a Worker
- **Endpoint:** `GET /workers/:workerId/readings/anomalous`
- **Description:** Returns all anomalous readings associated with a worker.
- **Parameters:**
  - `workerId` (path): ID of the worker.

---

### Helmet Operating Hours in a Time Window
- **Endpoint:** `GET /helmets/:helmetId/operating-hours`
- **Description:** Calculates the operating hours of a helmet within a specified time window.
- **Parameters:**
  - `helmetId` (path): ID of the helmet.
  - `from` (query): Start of the time window (`YYYY-MM-DD HH:mm:ss`).
  - `to` (query): End of the time window (`YYYY-MM-DD HH:mm:ss`).

---

### All Helmet Readings in a Time Window
- **Endpoint:** `GET /helmets/:helmetId/readings`
- **Description:** Returns all readings for a helmet within a specified time window.
- **Parameters:**
  - `helmetId` (path): ID of the helmet.
  - `from` (query): Start of the time window (`YYYY-MM-DD HH:mm:ss`).
  - `to` (query): End of the time window (`YYYY-MM-DD HH:mm:ss`).

---

## **Management and Configuration**

### List of Worksites
- **Endpoint:** `GET /worksites`
- **Description:** Returns a list of all active worksites.

---

### Assign Worker to Worksite
- **Endpoint:** `POST /worksites/:worksiteId/workers`
- **Description:** Assigns a worker to a worksite with a specific helmet.
- **Parameters:**
  - `worksiteId` (path): ID of the worksite.
- **Body:**
  ```json
  {
    "workerId": 123,
    "helmetId": 456
  }

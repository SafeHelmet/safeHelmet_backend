# SafeHelmet Backend API Documentation

## Table of Contents
- [Attendances](#attendances)
  - [Get All Attendances](#get-all-attendances)
  - [Get Attendance Details](#get-attendance-details)
  - [Get Last Attendance Details](#get-last-attendance-details)
- [Bosses](#bosses)
  - [Get All Bosses](#get-all-bosses)
  - [Get Boss Details](#get-boss-details)
  - [Create Boss](#create-boss)
  - [Update Boss](#update-boss)
  - [Delete Boss](#delete-boss)
- [Helmets](#helmets)
  - [Get All Helmets](#get-all-helmets)
  - [Get Helmet Details](#get-helmet-details)
  - [Get Helmet ID](#get-helmet-id)
  - [Get Helmet Attendance](#get-helmet-attendance)
  - [Get Helmet Readings](#get-helmet-readings)
  - [Create Helmet](#create-helmet)
  - [Update Helmet](#update-helmet)
  - [Delete Helmet](#delete-helmet)
- [Readings](#readings)
  - [Get All Readings](#get-all-readings)
  - [Get Reading Details](#get-reading-details)
  - [Get Reading Worker](#get-reading-worker)
  - [Get Reading Worksite](#get-reading-worksite)
  - [Create Reading](#create-reading)
  - [Update Reading](#update-reading)
  - [Delete Reading](#delete-reading)
- [Weather](#weather)
  - [Get All Weather Readings](#get-all-weather-readings)
  - [Get All Worksite Weather](#get-all-worksite-weather)
  - [Get Last Worksite Weather](#get-last-worksite-weather)
- [Workers](#workers)
  - [Get All Workers](#get-all-workers)
  - [Get Worker Details](#get-worker-details)
  - [Get Worker Worksite](#get-worker-worksite)
  - [Get Worker Readings](#get-worker-readings)
  - [Get Worker Attendance](#get-worker-attendance)
  - [Get Last Worker Attendance](#get-last-worker-attendance)
  - [Create Worker](#create-worker)
  - [Create Worker Attendance](#create-worker-attendance)
  - [Update Worker](#update-worker)
  - [Update Worker Attendance](#update-worker-attendance)
  - [Delete Worker](#delete-worker)
- [Worksites](#worksites)
  - [Get All Worksites](#get-all-worksites)
  - [Get Worksite Details](#get-worksite-details)
  - [Get Workers in Worksite](#get-workers-in-worksite)
  - [Get Worksite Readings](#get-worksite-readings)
  - [Get Worksite Attendance](#get-worksite-attendance)
  - [Create Worksite](#create-worksite)
  - [Assign Worker to Worksite](#assign-worker-to-worksite)
  - [Update Worksite](#update-worksite)
  - [Delete Worksite](#delete-worksite)

## Attendances

### Get All Attendances
- **URL**: `/api/v1/attendances`
- **Method**: `GET`
- **Description**: Retrieves a list of all attendances.

Example:
```
/api/v1/attendances
```

### Get Attendance Details
- **URL**: `/api/v1/attendances/:attendance_id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific attendance by ID.

Example:
```
/api/v1/attendances/1
```

### Get Last Attendance Details
- **URL**: `/api/v1/attendances/attendance-details`
- **Method**: `GET`
- **Description**: Retrieves the last attendance details for a specific worker, worksite, and helmet.
- **Body**:
  ```json
  {
    "worker_id": 1,
    "worksite_id": 1,
    "helmet_id": 1
  }
  ```

Example:
```
/api/v1/attendances/attendance-details
```

## Bosses

### Get All Bosses
- **URL**: `/api/v1/bosses`
- **Method**: `GET`
- **Description**: Retrieves a list of all bosses.

Example:
```
/api/v1/bosses
```

### Get Boss Details
- **URL**: `/api/v1/bosses/:boss-id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific boss by ID.

Example:
```
/api/v1/bosses/1
```

### Create Boss
- **URL**: `/api/v1/bosses`
- **Method**: `POST`
- **Description**: Creates a new boss.
- **Body**:
  ```json
  {
    "name": "BOSSJohn",
    "surname": "Doe",
    "email": "bossjohn.doe@example.com",
    "password": "password",
    "phone": "1234567800",
    "active": true,
    "fiscal_code": "BOSJHN123456"
  }
  ```

Example:
```
/api/v1/bosses
```

### Update Boss
- **URL**: `/api/v1/bosses/:boss-id`
- **Method**: `PUT`
- **Description**: Updates a specific boss by ID.
- **Body**:
  ```json
  {
    "name": "Updated BOSSJohn",
    "surname": "Updated Doe",
    "email": "updated.bossjohn.doe@example.com",
    "phone": "0987654321",
    "active": false,
    "fiscal_code": "UPDBOSJHN123456"
  }
  ```

Example:
```
/api/v1/bosses/1
```

### Delete Boss
- **URL**: `/api/v1/bosses/:boss-id`
- **Method**: `DELETE`
- **Description**: Deletes a specific boss by ID.

Example:
```
/api/v1/bosses/1
```

## Helmets

### Get All Helmets
- **URL**: `/api/v1/helmets`
- **Method**: `GET`
- **Description**: Retrieves a list of all helmets. Supports sorting in ascending or descending order on fields.
- **Query Parameters**:
  - `sortBy` (optional): The field to sort by (e.g., `model`, `manufacture_date`).
  - `order` (optional): The order of sorting (`asc` for ascending, `desc` for descending). Default is `asc`.

Example:
```
/api/v1/helmets?sortBy=model&order=desc
```

### Get Helmet Details
- **URL**: `/api/v1/helmets/:helmet-id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific helmet by ID.

Example:
```
/api/v1/helmets/1
```

### Get Helmet ID
- **URL**: `/api/v1/helmets/mac-address/:mac-address`
- **Method**: `GET`
- **Description**: Retrieves the ID of a helmet by its MAC address.

Example:
```
/api/v1/helmets/mac-address/mac0
```

### Get Helmet Attendance
- **URL**: `/api/v1/helmets/:helmet-id/attendance`
- **Method**: `GET`
- **Description**: Retrieves a list of attendances for a specific helmet.

Example:
```
/api/v1/helmets/1/attendance
```

### Get Helmet Readings
- **URL**: `/api/v1/helmets/:helmet-id/readings`
- **Method**: `GET`
- **Description**: Retrieves a list of readings for a specific helmet.

Example:
```
/api/v1/helmets/1/readings
```

### Get Helmet Readings
- **URL**: `/api/v1/helmets/helmet-categories`
- **Method**: `GET`
- **Description**: Retrieves a list of all helmets categories.

Example:
```
/api/v1/helmets/helmet-categories
```

### Get Helmet Readings
- **URL**: `/api/v1/helmets/helmet-categories/:category-id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific helmet category by ID.

Example:
```
/api/v1/helmets/1/readings
```

### Create Helmet
- **URL**: `/api/v1/helmets`
- **Method**: `POST`
- **Description**: Creates a new helmet.
- **Body**:
  ```json
  {
    "category_id": 1,
    "mac_address": "mac0"
  }
  ```

Example:
```
/api/v1/helmets
```

### Update Helmet
- **URL**: `/api/v1/helmets/:helmet-id`
- **Method**: `PUT`
- **Description**: Updates a specific helmet by ID.
- **Body**:
  ```json
  {
    "category_id": 1,
    "mac_address": "updated_mac0"
  }
  ```

Example:
```
/api/v1/helmets/1
```

### Delete Helmet
- **URL**: `/api/v1/helmets/:helmet-id`
- **Method**: `DELETE`
- **Description**: Deletes a specific helmet by ID.

Example:
```
/api/v1/helmets/1
```

## Readings

### Get All Readings
- **URL**: `/api/v1/readings`
- **Method**: `GET`
- **Description**: Retrieves a list of all readings. Supports sorting in ascending or descending order on fields.
- **Query Parameters**:
  - `sortBy` (optional): The field to sort by (e.g., `timestamp`, `value`).
  - `order` (optional): The order of sorting (`asc` for ascending, `desc` for descending). Default is `asc`.

Example:
```
/api/v1/readings?sortBy=timestamp&order=desc
```

### Get Reading Details
- **URL**: `/api/v1/readings/:reading-id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific reading by ID.

Example:
```
/api/v1/readings/1
```

### Get Reading Worker
- **URL**: `/api/v1/readings/:reading-id/worker`
- **Method**: `GET`
- **Description**: Retrieves the worker associated with a specific reading.

Example:
```
/api/v1/readings/1/worker
```

### Get Reading Worksite
- **URL**: `/api/v1/readings/:reading-id/worksite`
- **Method**: `GET`
- **Description**: Retrieves the worksite associated with a specific reading.

Example:
```
/api/v1/readings/1/worksite
```

### Create Reading
- **URL**: `/api/v1/readings`
- **Method**: `POST`
- **Description**: Creates a new reading.
- **Body**:
  ```json
  {
    "attendance_id": 1,
    "read_at": "2023-10-01T00:00:00Z",
    "temperature": 10.5,
    "humidity": 45.0,
    "brightness": 20,
    "methane": false,
    "carbon_monoxide": true,
    "smoke_detection": false,
    "uses_welding_protection": false,
    "uses_gas_protection": true,
    "avg_x": 1.0,
    "avg_y": 2.0,
    "avg_z": 3.0,
    "avg_g": 4.0,
    "std_x": 0.1,
    "std_y": 0.2,
    "std_z": 0.3,
    "std_g": 0.4,
    "max_g": 4.5,
    "incorrect_posture": 0.0,
    "anomaly": false
  }
  ```

Example:
```
/api/v1/readings
```

### Update Reading
- **URL**: `/api/v1/readings/:reading-id`
- **Method**: `PUT`
- **Description**: Updates a specific reading by ID.
- **Body**:
  ```json
  {
    "attendance_id": 1,
    "read_at": "2023-10-01T00:00:00Z",
    "temperature": 10.5,
    "humidity": 45.0,
    "brightness": 20,
    "methane": false,
    "carbon_monoxide": true,
    "smoke_detection": false,
    "uses_welding_protection": false,
    "uses_gas_protection": true,
    "avg_x": 1.0,
    "avg_y": 2.0,
    "avg_z": 3.0,
    "avg_g": 4.0,
    "std_x": 0.1,
    "std_y": 0.2,
    "std_z": 0.3,
    "std_g": 0.4,
    "max_g": 4.5,
    "incorrect_posture": 0.0,
    "anomaly": false
  }
  ```

Example:
```
/api/v1/readings/1
```

### Delete Reading
- **URL**: `/api/v1/readings/:reading-id`
- **Method**: `DELETE`
- **Description**: Deletes a specific reading by ID.

Example:
```
/api/v1/readings/1
```

## Weather

### Get All Weather Readings
- **URL**: `/api/v1/weather`
- **Method**: `GET`
- **Description**: Retrieves a list of all weather readings.

Example:
```
/api/v1/weather
```

### Get All Worksite Weather
- **URL**: `/api/v1/weather/:worksite-id`
- **Method**: `GET`
- **Description**: Retrieves a list of all weather readings for a specific worksite.

Example:
```
/api/v1/weather/1
```

### Get Last Worksite Weather
- **URL**: `/api/v1/weather/:worksite-id/last`
- **Method**: `GET`
- **Description**: Retrieves the last weather reading for a specific worksite.

Example:
```
/api/v1/weather/1/last
```

## Workers

### Get All Workers
- **URL**: `/api/v1/workers`
- **Method**: `GET`
- **Description**: Retrieves a list of all workers. Supports sorting in ascending or descending order on fields.
- **Query Parameters**:
  - `sortBy` (optional): The field to sort by (e.g., `name`, `surname`).
  - `order` (optional): The order of sorting (`asc` for ascending, `desc` for descending). Default is `asc`.

Example:
```
/api/v1/workers?sortBy=name&order=desc
```

### Get Worker Details
- **URL**: `/api/v1/workers/:worker-id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific worker by ID.

Example:
```
/api/v1/workers/1
```

### Get Worker Worksite
- **URL**: `/api/v1/workers/:worker-id/worksite`
- **Method**: `GET`
- **Description**: Retrieves the worksite associated with a specific worker.

Example:
```
/api/v1/workers/1/worksite
```

### Get Worker Readings
- **URL**: `/api/v1/workers/:worker-id/readings`
- **Method**: `GET`
- **Description**: Retrieves a list of readings for a specific worker.

Example:
```
/api/v1/workers/1/readings
```

### Get Worker Attendance
- **URL**: `/api/v1/workers/:worker-id/attendance`
- **Method**: `GET`
- **Description**: Retrieves a list of attendances for a specific worker.

Example:
```
/api/v1/workers/1/attendance
```

### Get Last Worker Attendance
- **URL**: `/api/v1/workers/:worker-id/attendance/last`
- **Method**: `GET`
- **Description**: Retrieves the last attendance details for a specific worker.

Example:
```
/api/v1/workers/1/attendance/last
```

### Create Worker
- **URL**: `/api/v1/workers`
- **Method**: `POST`
- **Description**: Creates a new worker.
- **Body**:
  ```json
  {
    "name": "John",
    "surname": "Doe",
    "email": "john.doe@example.com",
    "password": "password",
    "phone": "1234567890",
    "active": true,
    "fiscal_code": "JHNDOE123456"
  }
  ```

Example:
```
/api/v1/workers
```

### Create Worker Attendance
- **URL**: `/api/v1/workers/attendance`
- **Method**: `POST`
- **Description**: Creates a new worker attendance.
- **Body**:
  ```json
  {
    "worker_id": 1,
    "worksite_id": 1,
    "helmet_id": 1,
    "start_at": "2023-10-01T00:00:00Z"
  }
  ```

Example:
```
/api/v1/workers/attendance
```

### Update Worker
- **URL**: `/api/v1/workers/:worker-id`
- **Method**: `PUT`
- **Description**: Updates a specific worker by ID.
- **Body**:
  ```json
  {
    "name": "Updated Worker Name",
    "surname": "Updated Worker Surname",
    "email": "updated.worker@example.com",
    "phone": "0987654321",
    "active": false,
    "fiscal_code": "UPDJHNDOE123456"
  }
  ```

Example:
```
/api/v1/workers/1
```

### Update Worker Attendance
- **URL**: `/api/v1/workers/attendance`
- **Method**: `PUT`
- **Description**: Updates a specific worker attendance.
- **Body**:
  ```json
  {
    "worker_id": 1,
    "worksite_id": 1,
    "helmet_id": 1,
    "end_at": "2023-10-01T12:00:00Z"
  }
  ```

Example:
```
/api/v1/workers/attendance
```

### Delete Worker
- **URL**: `/api/v1/workers/:worker-id`
- **Method**: `DELETE`
- **Description**: Deletes a specific worker by ID.

Example:
```
/api/v1/workers/1
```

## Worksites

### Get All Worksites
- **URL**: `/api/v1/worksites`
- **Method**: `GET`
- **Description**: Retrieves a list of all worksites. Supports sorting in ascending or descending order on fields.
- **Query Parameters**:
  - `sortBy` (optional): The field to sort by (e.g., `name`, `start_date_of_work`).
  - `order` (optional): The order of sorting (`asc` for ascending, `desc` for descending). Default is `asc`.

Example:
```
/api/v1/worksites?sortBy=name&order=desc
```

### Get Worksite Details
- **URL**: `/api/v1/worksites/:worksite-id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific worksite by ID.

Example:
```
/api/v1/worksites/1
```

### Get Workers in Worksite
- **URL**: `/api/v1/worksites/:worksite-id/workers`
- **Method**: `GET`
- **Description**: Retrieves a list of workers in a specific worksite. Supports sorting in ascending or descending order on fields.
- **Query Parameters**:
  - `sortBy` (optional): The field to sort by (e.g., `name`, `surname`).
  - `order` (optional): The order of sorting (`asc` for ascending, `desc` for descending). Default is `asc`.

Example:
```
/api/v1/worksites/:worksite-id/workers?sortBy=name&order=desc
```

### Get Worksite Readings
- **URL**: `/api/v1/worksites/:worksite-id/readings`
- **Method**: `GET`
- **Description**: Retrieves a list of readings for a specific worksite.

Example:
```
/api/v1/worksites/1/readings
```

### Get Worksite Attendance
- **URL**: `/api/v1/worksites/:worksite-id/attendance`
- **Method**: `GET`
- **Description**: Retrieves a list of attendances for a specific worksite.

Example:
```
/api/v1/worksites/1/attendance
```

### Create Worksite
- **URL**: `/api/v1/worksites`
- **Method**: `POST`
- **Description**: Creates a new worksite.
- **Body**:
  ```json
  {
    "name": "Worksite 1",
    "latitude": 45.4642,
    "longitude": 9.1900,
    "city": "Milan",
    "zip_code": "20100",
    "state": "Italy",
    "address": "Via Roma, 1",
    "start_at": "2023-10-01T00:00:00Z"
  }
  ```

Example:
```
/api/v1/worksites
```

### Assign Worker to Worksite
- **URL**: `/api/v1/worksites/assign-worker`
- **Method**: `POST`
- **Description**: Assigns a worker to a worksite.
- **Body**:
  ```json
  {
    "worksite_id": 1,
    "worker_id": 1,
    "assigned_by": 1
  }
  ```

Example:
```
/api/v1/worksites/assign-worker
```

### Update Worksite
- **URL**: `/api/v1/worksites/:worksite-id`
- **Method**: `PUT`
- **Description**: Updates a specific worksite by ID.
- **Body**:
  ```json
  {
    "name": "Updated Worksite Name",
    "latitude": 45.0,
    "longitude": 9.0,
    "start_date_of_work": "2023-10-01T00:00:00Z",
    "end_date_of_work": "2023-12-31T00:00:00Z"
  }
  ```

Example:
```
/api/v1/worksites/1
```

### Delete Worksite
- **URL**: `/api/v1/worksites/:worksite-id`
- **Method**: `DELETE`
- **Description**: Deletes a specific worksite by ID.

Example:
```
/api/v1/worksites/1
```

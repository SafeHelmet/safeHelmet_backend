# SafeHelmet Backend API Documentation

## Table of Contents
- [Worksites](#worksites)
  - [Get All Worksites](#get-all-worksites)
  - [Get Worksite Details](#get-worksite-details)
  - [Get Workers in Worksite](#get-workers-in-worksite)
  - [Get Worksite Readings](#get-worksite-readings)
  - [Create Worksite](#create-worksite)
  - [Assign Worker to Worksite](#assign-worker-to-worksite)
  - [Update Worksite](#update-worksite)
- [Workers](#workers)
  - [Get All Workers](#get-all-workers)
  - [Get Worker Details](#get-worker-details)
  - [Get All Bosses](#get-all-bosses)
  - [Update Worker](#update-worker)
- [Helmets](#helmets)
  - [Get All Helmets](#get-all-helmets)
- [Readings](#readings)
  - [Get All Readings](#get-all-readings)
  - [Get Reading Details](#get-reading-details)

## Worksites

### Get All Worksites
- **URL**: `/api/v1/worksites`
- **Method**: `GET`
- **Description**: Retrieves a list of all worksites.

### Get Worksite Details
- **URL**: `/api/v1/worksites/:worksite-id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific worksite by ID.

### Get Workers in Worksite
- **URL**: `/api/v1/worksites/:worksite-id/workers`
- **Method**: `GET`
- **Description**: Retrieves a list of workers in a specific worksite.

### Get Worksite Readings
- **URL**: `/api/v1/worksites/:worksite-id/readings`
- **Method**: `GET`
- **Description**: Retrieves a list of readings for a specific worksite.

### Create Worksite
- **URL**: `/api/v1/worksites`
- **Method**: `POST`
- **Description**: Creates a new worksite.
- **Body**:
  ```json
  {
    "name": "Worksite Name",
    "latitude": 45.0,
    "longitude": 9.0,
    "start_date_of_work": "2023-10-01T00:00:00Z",
    "end_date_of_work": "2023-12-31T00:00:00Z"
  }

### Assign Worker to Worksite
- **URL**: `api/v1/worksites/assign-worker`
- **Method**: `POST`
- **Description**: Assigns a worker to a worksite.
- **Body**:
  ```json
  {
    "worksite_id": 1,
    "worker_id": 1,
    "assigned_by": 1
  }

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

## Workers

### Get All Workers
- **URL**: `/api/v1/workers`
- **Method**: `GET`
- **Description**: Retrieves a list of all workers.

### Get Worker Details
- **URL**: `/api/v1/workers/:worker-id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific worker by ID.

### Get All Bosses
- **URL**: `/api/v1/workers/bosses`
- **Method**: `GET`
- **Description**: Retrieves a list of all bosses.

### Update Worker
- **URL**: `/api/v1/workers/:worker-id`
- **Method**: `PUT`
- **Description**: Updates a specific worker by ID.
- **Body**:
  ```json
  {
    "name": "Updated Worker Name",
    "surname": "Updated Worker Surname"
  }

## Helmets

### Get All Helmets
- **URL**: `/api/v1/helmets`
- **Method**: `GET`
- **Description**: Retrieves a list of all helmets.

## Readings

### Get All Readings
- **URL**: `/api/v1/readings`
- **Method**: `GET`
- **Description**: Retrieves a list of all readings.

### Get Reading Details
- **URL**: `/api/v1/readings/:reading-id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific reading by ID.

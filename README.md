# SafeHelmet Backend

This repository contains the code for the backend of the SafeHelmet project. The backend is responsible for processing the data sent from the user's smartphone, comparing values from various sensors with dynamic or static thresholds, and sending alerts to the user in case of dangerous situations.

The dynamic thresholds are:
- Weather temperature
- Weather humidity

Values read from the sensors are then compared with the weather conditions from the OpenWeather APIs. The weather for each worksite is fetched by a coroutine every hour for each worksite location from 8 am to 8 pm.

Thresholds regarding the user's posture and movement are static and are set by the worksite manager or modified from the frontend.

The backend also provides APIs for managing polling from workers' smartphones and notifying them in case an anomalous reading is detected.

Additionally, the backend provides APIs to the frontend for all CRUD operations and managing various worksites, workers, and work shifts.

[Frontend APIs documentation](https://github.com/SafeHelmet/safeHelmet_backend/blob/100fe108be156cfcd649adb5d58751167c701836/API_DOCUMENTATION.md#assign-worker-to-worksite)
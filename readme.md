# Airline Voucher Seat Assignment App

## Project Overview

This project is a **voucher seat assignment application** for an airline campaign. It includes two main parts:
- **Backend** (built with **Go**): Handles API requests for seat assignment.
- **Frontend** (built with **React**): Provides a user interface for the airline crew to generate seat vouchers.

## Table of Contents
- [Requirements](#requirements)
- [Installation](#installation)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [Running the Application](#running-the-application)
  - [Backend](#backend)
  - [Frontend](#frontend)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [License](#license)

## Requirements

Before setting up the project, ensure that the following are installed:

1. **Go** (v1.18 or later) for the backend.
2. **Node.js** and **npm** (v16 or later) for the frontend.
3. **Git** (to clone the repository).

## Installation

### Backend Setup (Go)

1. **Clone the repository** (or download the project files):

   ```bash
   git clone https://github.com/gunturwsn/airline-voucher-seat-assignment-app.git
   cd your-project-folder
   ```


2. **Install Go dependencies:**

    Navigate to the backend directory:
    ```
    cd backend
    ```

    Then, install the Go dependencies by running:
    ```
    go mod tidy
    ```

3. **Run the Backend:**

    Start the backend server by running:
    ```
    go run main.go
    ```
    The backend will be available at http://localhost:8000.

### Frontend Setup (React)

1. **Navigate to the frontend folder:**
    ```
    cd frontend
    ```

2. **install frontend dependencies:**

    Run the following command to install npm dependencies for React:
    ```
    npm install
    ```

3. **Start the Frontend:**

    To start the React app, use:
    ```
    npm start
    ```
    The frontend will be available at http://localhost:3000.



## Running the Application

### Backend

To run the backend manually, follow these steps:

1. **Navigate to the backend directory and run:**

    ```
    go run main.go
    ```
    The backend will be accessible at http://localhost:8000

### Frontend

To run the frontend manually, follow these steps:

1. **Navigate to the frontend directory and run:**

    ```
    npm start
    ```
    The frontend will be accessible at http://localhost:3000.


## API Endpoints

### POST /api/check

  - **Description**: Checks if the seat vouchers ar already generated for a specific flight and date.
  - **Request Body**:
    ```
    {
      "flightNumber: "GA102",
      ""date: "2025-07-12"
    }
    ```
  - **Response**:
    ```
    {
      "exists": true
    }
    ```

### POST /api/generate

  - **Description**: Generate 3 random seat assignments for a flight.
  - **Request Body**: 
    ```
    {
      "name": "Sarah",
      "id": "98123",
      "flightNumber": "ID102",
      "date": "2025-07-12",
      "aircraft": "Airbus 320"
    }
    ```
  - **Response**:
    ```
    {
      "success": true,
      "seats": ["3B, "7C", "14D"]
    }
    ```


## Project Structure

The project is divided into two main parts: **frontend** (React) and **backend** (Go), Here's a summary of the structure:

```
/root
  ├── /backend                # Go backend code
  │   ├── main.go             # Main entry point for the Go backend
  │   ├── /models             # Backend data models
  │   ├── /handlers           # API handlers for seat assignment
  │   ├── /routes             # Define routes and map them to handlers
  │   ├── /database           # Database connection and models (SQLite or other)
  │   ├── /interfaces         # Interfaces for service layer and dependency injection
  │   └── /services           # Business logic for seat assignment and other operations
  ├── /frontend               # React frontend code
  │   ├── /public             # Public assets for React app
  │   ├── /src                # Source code (components, views)
  │   └── package.json        # npm dependencies and scripts

```

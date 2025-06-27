# Car Viewer

A modern web application for browsing, comparing, and managing information about car models, their specifications, manufacturers, and more.  
Built with a Go backend and a responsive, user-friendly frontend.

---

## Features

- **Browse Cars:** View a list of car models with key specs and images.
- **Advanced Filtering:** Filter by manufacturer, category, and search by model name.
- **Car Details:** Click any car to view detailed specs and an image in a pop-up.
- **Favorites:** Like cars to add them to your favorites list. Remove them anytime.
- **Compare Cars:** Select multiple cars and compare their specs side-by-side.
- **Personalized Recommendations:** See recommended cars based on your interests.
- **Data Visualization:** Interactive bar and pie charts for horsepower and category distribution.
- **Modern UI:** Clean, Material-UI-inspired design with a sea green and white theme.

---

## Screenshots

*(Add screenshots here if desired)*

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (1.18+ recommended)
- [Node.js](https://nodejs.org/) (optional, only if you want to use npm for frontend tooling)
- Modern web browser

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/car-viewer.git
   cd car-viewer

2. **Prepare the database and images:**

    Ensure api/data.json contains your car data (see format below).
    Place car images in api/img/.

3. **Run the Go server:**

    go run main.go
     **The server will start on http://localhost:8080.

4. **Open the app:**

    Visit http://localhost:8080/mainpage.html in your browser.

**API Endpoints**

    /api/cars — Get all cars
    /api/cars/details?id=... — Get details for a specific car
    /api/manufacturers — Get list of manufacturers
    /api/categories — Get list of categories
    /api/favorites — Get user's favorite cars
    /api/like?id=... — Add a car to favorites
    /api/unlike?id=... — Remove a car from favorites

**Customization**
    Add more cars: Edit data.json<vscode_annotation details='%5B%7B%22title%22%3A%22hardcoded-credentials%22%2C%22description%22%3A%22Embedding%20credentials%20in%20source%20code%20risks%20unauthorized%20access%22%7D%5D'> and</vscode_annotation> add images to img.
    Change theme: Edit styles.css.
    Extend backend: Add new endpoints in handlers.go.



**Credits**
    Chart.js for data visualization
    Material-UI for design inspiration
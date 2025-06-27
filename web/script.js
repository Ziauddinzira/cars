let allCars = []
let allManufacturers = []
let allCategories = []

async function fetchCars() {
  const res = await fetch('/api/cars')
  const data = await res.json()
  allCars = data
  renderCars(data)
}

async function fetchManufacturers() {
  const res = await fetch('/api/manufacturers')
  allManufacturers = await res.json()
  const select = document.getElementById("manufacturerFilter")
  select.innerHTML = `<option value="">All Manufacturers</option>` + allManufacturers.map(m => `<option value="${m}">${m}</option>`).join("")
}

async function fetchCategories() {
  const res = await fetch('/api/categories')
  allCategories = await res.json()
  const select = document.getElementById("categoryFilter")
  select.innerHTML = `<option value="">All Categories</option>` + allCategories.map(c => `<option value="${c}">${c}</option>`).join("")
}

function renderCars(cars) {
  const tbody = document.querySelector("#carsTable tbody")
  tbody.innerHTML = ""

  cars.forEach(car => {
    const row = document.createElement("tr")
    row.innerHTML = `
      <td>${car.model}</td>
      <td>${car.year}</td>
      <td><input type="checkbox" class="compare-checkbox" value="${car.id}"></td>
      <td>
        <button onclick="showDetails('${car.id}')">Details</button>
        <button onclick="likeCar('${car.id}')">❤️ Like</button>
      </td>
    `
    tbody.appendChild(row)
  })
}

function filterCars() {
  const term = document.getElementById("search").value.toLowerCase()
  const manufacturer = document.getElementById("manufacturerFilter").value
  const category = document.getElementById("categoryFilter").value
  const filtered = allCars.filter(c =>
    c.model.toLowerCase().includes(term) &&
    (manufacturer === "" || c.manufacturer === manufacturer) &&
    (category === "" || c.category === category)
  )
  renderCars(filtered)
}

function compareSelectedCars() {
  const checked = document.querySelectorAll('.compare-checkbox:checked')
  if (checked.length < 2) {
    alert("Select at least two cars to compare.")
    return
  }

  const selectedIds = Array.from(checked).map(input => input.value)
  const selectedCars = allCars.filter(car => selectedIds.includes(car.id))

  const compareBox = document.getElementById("compareBox")
  compareBox.classList.remove("hidden")

  let html = `<h3>Comparison Table</h3><table><thead><tr><th>Spec</th>`
  selectedCars.forEach(car => {
    html += `<th>${car.model}</th>`
  })
  html += `</tr></thead><tbody>`

  const specs = ["year", "engine", "horsepower"]
  specs.forEach(spec => {
    html += `<tr><td>${spec}</td>`
    selectedCars.forEach(car => {
      html += `<td>${car[spec]}</td>`
    })
    html += `</tr>`
  })

  html += `</tbody></table>`
  compareBox.innerHTML = html
}

function trackView(carId) {
  let viewed = JSON.parse(localStorage.getItem("viewedCars") || "[]")
  if (!viewed.includes(carId)) {
    viewed.push(carId)
    localStorage.setItem("viewedCars", JSON.stringify(viewed))
  }
}

function showDetails(id) {
  fetch(`/api/cars/details?id=${id}`)
    .then(res => {
      if (!res.ok) throw new Error("Details not found")
      return res.json()
    })
    .then(car => {
      const box = document.getElementById("detailsBox")
      box.classList.remove("hidden")
      box.innerHTML = `
        <h3>${car.model} (${car.year})</h3>
        <img src="${car.image}" alt="${car.model}" style="max-width:300px;display:block;margin-bottom:1em;">
        <p><strong>Engine:</strong> ${car.engine}</p>
        <p><strong>Horsepower:</strong> ${car.horsepower}</p>
        ${showRecommendations(car)}
      `
      trackView(car.id)
    })
    .catch(() => alert("Details not found"))
}

function showRecommendations(currentCar) {
  // Recommend cars from the same manufacturer, not the current car
  const recommended = allCars.filter(car =>
    car.manufacturer === currentCar.manufacturer && car.id !== currentCar.id
  ).slice(0, 3); // Show up to 3 recommendations

  if (recommended.length === 0) return "";

  return `
    <h4>Recommended for you</h4>
    <ul>
      ${recommended.map(car => `<li>${car.model} (${car.year})</li>`).join("")}
    </ul>
  `;
}

function likeCar(carId) {
  fetch(`/api/like?id=${carId}`, { method: "POST" })
    .then(res => {
      if (res.ok) {
        loadFavorites()
      } else {
        alert("Failed to like car")
      }
    })
}

function loadFavorites() {
  fetch("/api/favorites")
    .then(res => res.json())
    .then(data => {
      const box = document.getElementById("favoritesBox")
      if (!data || data.length === 0) {
        box.classList.add("hidden")
        box.innerHTML = ""        // <-- Add this line
        return
      }
      box.classList.remove("hidden")
      box.innerHTML = "<ul>" + data.map(c =>
        `<li>${c.model} (${c.year}) <button onclick="removeFavorite('${c.id}')">Remove</button></li>`
      ).join("") + "</ul>"
    })
}

function removeFavorite(carId) {
  fetch(`/api/unlike?id=${carId}`, { method: "POST" })
    .then(res => {
      if (res.ok) {
        loadFavorites()
      } else {
        alert("Failed to remove favorite")
      }
    })
}



// Initial load
fetchCars().then(() => {
  fetchManufacturers()
  fetchCategories()
  loadFavorites()
})

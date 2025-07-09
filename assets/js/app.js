import "phoenix_html"
import {Socket} from "phoenix"
import {LiveSocket} from "phoenix_live_view"
import topbar from "../vendor/topbar"

let Hooks = {}
Hooks.AddressAutocomplete = {
  mounted() {
    console.log("AddressAutocomplete hook mounted on:", this.el);
    this.debouncedSearch = this.debounce(this.searchAddress, 500);
    this.el.addEventListener("input", (e) => {
      console.log("Input event fired. Query:", e.target.value);
      const query = e.target.value;
      if (query.length > 2) {
        this.debouncedSearch(query);
      } else {
        this.clearSuggestions();
      }
    });
  },

  debounce(func, delay) {
    let timeout;
    return function(...args) {
      const context = this;
      clearTimeout(timeout);
      timeout = setTimeout(() => func.apply(context, args), delay);
    };
  },

  searchAddress(query) {
    const url = `https://nominatim.openstreetmap.org/search?format=json&addressdetails=1&q=${query}`;
    fetch(url, {
      headers: {
        "User-Agent": "PulapApp/1.0 (your_email@example.com)" // Replace with your app name and email
      }
    })
      .then(response => response.json())
      .then(data => this.renderSuggestions(data))
      .catch(error => console.error("Error fetching address suggestions:", error));
  },

  renderSuggestions(suggestions) {
    const suggestionsDiv = document.getElementById("address_suggestions");
    suggestionsDiv.innerHTML = "";

    if (suggestions.length === 0) {
      suggestionsDiv.classList.add("hidden");
      return;
    }

    suggestionsDiv.classList.remove("hidden");

    suggestions.forEach(suggestion => {
      const div = document.createElement("div");
      div.classList.add("p-2", "cursor-pointer", "hover:bg-gray-100");
      div.textContent = suggestion.display_name;
      div.addEventListener("click", () => {
        console.log("Selected suggestion:", suggestion);
        this.el.value = suggestion.display_name;
        this.pushEvent("address_selected", { lat: suggestion.lat, lng: suggestion.lon, address: suggestion });
        setTimeout(() => this.clearSuggestions(), 100);
      });
      suggestionsDiv.appendChild(div);
    });
  },

  clearSuggestions() {
    const suggestionsDiv = document.getElementById("address_suggestions");
    suggestionsDiv.innerHTML = "";
    suggestionsDiv.classList.add("hidden");
  }
};

let csrfToken = document.querySelector("meta[name='csrf-token']").getAttribute("content")
let liveSocket = new LiveSocket("/live", Socket, {
  longPollFallbackMs: 2500,
  params: {_csrf_token: csrfToken},
  hooks: Hooks
})

topbar.config({barColors: {0: "#29d"}, shadowColor: "rgba(0, 0, 0, .3)"})
window.addEventListener("phx:page-loading-start", _info => topbar.show(300))
window.addEventListener("phx:page-loading-stop", _info => topbar.hide())

liveSocket.connect()

window.liveSocket = liveSocket
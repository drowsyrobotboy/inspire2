document.addEventListener("DOMContentLoaded", function () {
    const tabs = document.querySelectorAll(".tab");
    const hackerNewsList = document.getElementById("hackernews-list");
    const lobstersList = document.getElementById("lobsters-list");

    // Fetch Hacker News
    fetch("/api/hackernews")
        .then(response => {
            if (!response.ok) {
                throw new Error("Failed to fetch Hacker News");
            }
            return response.json();
        })
        .then(data => {
            hackerNewsList.innerHTML = ""; // Clear existing content
            data.forEach(item => {
                const link = document.createElement("a");
                link.href = item.url;
                link.textContent = item.title;
                link.target = "_blank"; // Open in a new tab
                const listItem = document.createElement("li");
                listItem.appendChild(link);
                hackerNewsList.appendChild(listItem);
            });
        })
        .catch(error => {
            console.error("Error fetching Hacker News:", error);
            hackerNewsList.innerHTML = "<li>Failed to load Hacker News. Please try again later.</li>";
        });

    // Fetch Lobsters
    fetch("/api/lobsters")
        .then(response => {
            if (!response.ok) {
                throw new Error("Failed to fetch Lobsters");
            }
            return response.json();
        })
        .then(data => {
            lobstersList.innerHTML = ""; // Clear existing content
            data.forEach(item => {
                const link = document.createElement("a");
                link.href = item.url;
                link.textContent = item.title;
                link.target = "_blank"; // Open in a new tab
                const listItem = document.createElement("li");
                listItem.appendChild(link);
                lobstersList.appendChild(listItem);
            });
        })
        .catch(error => {
            console.error("Error fetching Lobsters:", error);
            lobstersList.innerHTML = "<li>Failed to load Lobsters. Please try again later.</li>";
        });

    // Tab switching logic
    tabs.forEach(tab => {
        tab.addEventListener("click", function () {
            // Remove active class from all tabs and lists
            tabs.forEach(t => t.classList.remove("active"));
            document.querySelectorAll(".news-list").forEach(list => list.classList.remove("active"));

            // Add active class to the clicked tab and corresponding list
            this.classList.add("active");
            const source = this.getAttribute("data-source");
            document.getElementById(`${source}-list`).classList.add("active");
        });
    });
});
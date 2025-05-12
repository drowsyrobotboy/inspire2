document.addEventListener("DOMContentLoaded", function () {
    fetch("/news")
        .then(response => {
            if (!response.ok) {
                throw new Error("Failed to fetch news");
            }
            return response.json();
        })
        .then(data => {
            const newsList = document.getElementById("news-list");
            newsList.innerHTML = ""; // Clear existing content
            data.forEach(item => {
                const link = document.createElement("a");
                link.href = item.url;
                link.textContent = item.title;
                link.target = "_blank"; // Open in a new tab
                const listItem = document.createElement("li");
                listItem.appendChild(link);
                newsList.appendChild(listItem);
            });
        })
        .catch(error => {
            console.error("Error fetching news:", error);
            const errorMessage = document.createElement("p");
            errorMessage.textContent = "Failed to load news. Please try again later.";
            document.body.appendChild(errorMessage);
        });
});
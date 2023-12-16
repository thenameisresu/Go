// static/timesheet.js

document.addEventListener("DOMContentLoaded", function () {
    // Call backend API to get user details and determine user type
    fetch("/api/userDetails")
        .then(response => response.json())
        .then(data => {
            const downloadButton = document.getElementById("downloadButton");

            if (data.userType === "admin") {
                // Only show download option for admin users
                downloadButton.style.display = "block";
                downloadButton.addEventListener("click", downloadTimesheet);
            } else {
                downloadButton.style.display = "none";
            }
        })
        .catch(error => console.error("Error fetching user details:", error));
});

function downloadTimesheet() {
    // Call backend API to initiate the download
    fetch("/api/downloadTimesheet")
        .then(response => response.blob())
        .then(blob => {
            // Create a link element and trigger a click to download the blob
            const link = document.createElement("a");
            link.href = URL.createObjectURL(blob);
            link.download = "timesheet.xlsx"; // Set desired filename
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
        })
        .catch(error => console.error("Error downloading timesheet:", error));
}

// static/register.js

function submitRegister() {
    const userType = document.getElementById("userType").value;
    const email = document.getElementById("email").value;
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    // Call backend API to register a new user
    fetch("/api/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ userType, email, username, password }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            alert("Registration successful. Please login.");
            window.location.href = "/login";
        } else {
            alert("Registration failed. Please try again.");
        }
    })
    .catch(error => console.error("Error during registration:", error));
}

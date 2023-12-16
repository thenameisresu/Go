// static/login.js

function submitLogin() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    // Call backend API to validate login credentials
    fetch("/api/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.token) {
            // Store the token (you may want to use local storage or a cookie)
            alert("Login successful! Token: " + data.token);
        } else {
            alert("Invalid username or password");
        }
    })
    .catch(error => console.error("Error during login:", error));
}

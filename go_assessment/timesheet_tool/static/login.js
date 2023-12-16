// static/login.js

function submitLogin() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    // Call backend API to validate login credentials
    fetch("/api/timesheet/login", {
        method: "POST", 
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            // Redirect to the timesheet page or another authenticated page
            window.location.href = "/timesheet";
        } else {
            alert("Invalid username or password");
        }
    })
    .catch(error => console.error("Error during login:", error));
}

function cancelLogin() {
    // Assuming you have input elements with IDs 'username' and 'password'
    document.getElementById("username").value = "";
    document.getElementById("password").value = "";

    // Optionally, you can add more logic or UI changes for cancellation
    alert("Login canceled");
}

function forgotPassword() {
    // Show the forgot password modal
    const modal = document.getElementById("forgotPasswordModal");
    modal.style.display = "block";
}

function closeForgotPasswordModal() {
    // Close the forgot password modal
    const modal = document.getElementById("forgotPasswordModal");
    modal.style.display = "none";

    // Optionally, clear the entered email
    document.getElementById("forgotEmail").value = "";
}

function submitForgotPassword() {
    // Get the entered email
    const email = document.getElementById("forgotEmail").value;

    // Call backend API to check if the email exists in the user table
    fetch("/api/forgotPassword", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ email }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            alert("Password reset email sent successfully!");
            closeForgotPasswordModal();
        } else {
            alert("Email not found. Please enter a valid email.");
        }
    })
    .catch(error => console.error("Error during forgot password:", error));
}
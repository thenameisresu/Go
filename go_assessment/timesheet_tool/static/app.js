console.log("Start of app.js");
document.addEventListener("DOMContentLoaded", function () {
    const projectDropdown = document.getElementById("projectName");
    const subprojectDropdown = document.getElementById("subProjectName");

    // Add default options
    const defaultProjectOption = document.createElement("option");
    defaultProjectOption.value = ""; 
    defaultProjectOption.textContent = "--Select your project--"; 
    projectDropdown.appendChild(defaultProjectOption);

    const defaultSubprojectOption = document.createElement("option");
    defaultSubprojectOption.value = ""; 
    defaultSubprojectOption.textContent = "--Select your subproject--"; 
    subprojectDropdown.appendChild(defaultSubprojectOption);

    // Fetch projects and populate the projects dropdown
    fetch("/api/timesheet/projects")  
        .then(response => response.json())
        .then(data => {
            data.projects.forEach(project => {
                const option = document.createElement("option");
                option.value = project.project_id;
                option.textContent = `${project.project_id} | ${project.project_name}`;
                projectDropdown.appendChild(option);
            });
        })
        .catch(error => console.error("Error fetching projects:", error))
        .finally(() => {
            console.log("projectDropdown : ", projectDropdown); 
        });

    // Event listener for project dropdown change
    projectDropdown.addEventListener("change", function () {
        const selectedProjectID = projectDropdown.value;

        // Fetch subprojects based on the selected project
        fetch(`/api/timesheet/subprojects?projectID=${selectedProjectID}`)
            .then(response => response.json())
            .then(data => {
                subprojectDropdown.innerHTML = ""; // Clear existing options
                subprojectDropdown.appendChild(defaultSubprojectOption); // Add default option

                data.subprojects.forEach(subproject => {
                    const option1 = document.createElement("option");
                    option1.value = subproject.sub_project_id;
                    option1.textContent = `${subproject.sub_project_id} | ${subproject.sub_project_name}`;
                    subprojectDropdown.appendChild(option1);
                });
            })
            .catch(error => console.error("Error fetching subprojects:", error))
            .finally(() => {
                console.log("subprojectDropdown : ", subprojectDropdown); 
            });
    });
});

// fetchSubProjects

function fetchSubProjects() {
    const selectedProjectID = document.getElementById("projectName").value;
    console.log("Selected Project ID:", selectedProjectID); 

    fetch(`/api/timesheet/subprojects?projectID=${selectedProjectID}`)
        .then(response => response.json())
        .then(data => {
            const subprojectDropdown = document.getElementById("subProjectName");
            subprojectDropdown.innerHTML = ""; 

            data.subprojects.forEach(subproject => {
                const option = document.createElement("option");
                option.value = subproject.SubProjectID;
                option.textContent = subproject.SubProjectName;
                subprojectDropdown.appendChild(option);
            });
        })
        .catch(error => console.error("Error fetching subprojects:", error));
}

//submitTimesheet
function submitTimesheet() {
    const timesheetData = {
        project_id: document.getElementById("projectName").value,
        sub_project_id: document.getElementById("subProjectName").value,
        jira_snow_id: document.getElementById("jiraSnowID").value,
        task_description: document.getElementById("taskDescription").value,
        hours_spent: parseInt(document.getElementById("hoursSpent").value),
        comments: document.getElementById("comments").value
    };
    console.info("timesheetData : ",timesheetData)
    fetch("/api/timesheet", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(timesheetData)
    })
        .then(response => response.json())
        .then(data => {
            console.log("Timesheet submitted successfully:", data);
            alert("Timesheet submitted successfully!");
        })
        .catch(error => console.error("Error submitting timesheet:", error));
}

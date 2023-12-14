// main.go
package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"libraryManagementSystem/db"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	db.InitDB()
}

func main() {
	defer db.CloseDB()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}

/*type Component struct {
	CompID      int
	CompName    string
	Description string
}*/

func indexHandler(w http.ResponseWriter, r *http.Request) {
	components, err := db.GetComponents()
	if err != nil {
		http.Error(w, "Error retrieving components", http.StatusInternalServerError)
		return
	}

	err = tpl.ExecuteTemplate(w, "index.html", components)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Display the edit form
		componentIDStr := r.URL.Query().Get("compID")
		if componentIDStr == "" {
			http.Error(w, "Component ID is required", http.StatusBadRequest)
			return
		}

		componentID, err := strconv.Atoi(componentIDStr)
		if err != nil {
			http.Error(w, "Invalid Component ID", http.StatusBadRequest)
			return
		}

		component, err := db.GetComponentByID(componentID)
		if err != nil {
			http.Error(w, "Error retrieving component", http.StatusInternalServerError)
			return
		}

		err = tpl.ExecuteTemplate(w, "edit.html", component)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	} else if r.Method == http.MethodPost {
		// Handle form submission for updating component
		componentIDStr := r.FormValue("CompID")
		compName := r.FormValue("CompName")
		description := r.FormValue("Description")

		componentID, err := strconv.Atoi(componentIDStr)
		if err != nil {
			http.Error(w, "Invalid Component ID", http.StatusBadRequest)
			return
		}

		updatedComponent := db.Component{
			CompID:      componentID,
			CompName:    compName,
			Description: description,
		}

		err = db.UpdateComponent(updatedComponent)
		if err != nil {
			http.Error(w, "Error updating component", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Display the create form
		err := tpl.ExecuteTemplate(w, "create.html", nil)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	} else if r.Method == http.MethodPost {
		// Handle form submission for creating component
		compName := r.FormValue("CompName")
		description := r.FormValue("Description")

		newComponent := db.Component{
			CompName:    compName,
			Description: description,
		}

		_, err := db.CreateComponent(newComponent)
		if err != nil {
			http.Error(w, "Error creating component", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Display the delete confirmation form
		componentIDStr := r.URL.Query().Get("compID")
		if componentIDStr == "" {
			http.Error(w, "Component ID is required", http.StatusBadRequest)
			return
		}

		componentID, err := strconv.Atoi(componentIDStr)
		if err != nil {
			http.Error(w, "Invalid Component ID", http.StatusBadRequest)
			return
		}

		component, err := db.GetComponentByID(componentID)
		if err != nil {
			http.Error(w, "Error retrieving component", http.StatusInternalServerError)
			return
		}

		err = tpl.ExecuteTemplate(w, "delete.html", component)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	} else if r.Method == http.MethodPost {
		// Handle form submission for deleting component
		componentIDStr := r.FormValue("CompID")

		componentID, err := strconv.Atoi(componentIDStr)
		if err != nil {
			http.Error(w, "Invalid Component ID", http.StatusBadRequest)
			return
		}

		err = db.DeleteComponent(componentID)
		if err != nil {
			http.Error(w, "Error deleting component", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

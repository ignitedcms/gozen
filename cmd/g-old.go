// Generate Insert function
	builder.WriteString(fmt.Sprintf("// Insert inserts a new %s into the database\n", structName))
	builder.WriteString(fmt.Sprintf("func Create("))
	firstField := true
	for _, field := range fields {
		if field.Name != "id" && field.Name != "created_at" && field.Name != "updated_at" {
			if !firstField {
				builder.WriteString(", ")
			} else {
				firstField = false
			}
			builder.WriteString(fmt.Sprintf("%s %s", strings.ToLower(field.Name), field.Type))
		}
	}
	builder.WriteString(") (int64, error) {\n")
	builder.WriteString("\tquery := \"INSERT INTO " + table + "(")
	var insertFields []string
	for _, field := range fields {
		if field.Name != "id" && field.Name != "created_at" && field.Name != "updated_at" {
			insertFields = append(insertFields, field.Name)
		}
	}
	builder.WriteString(strings.Join(insertFields, ", "))
	builder.WriteString(", created_at, updated_at) OUTPUT INSERTED.id VALUES(")
	for i := range insertFields {
		builder.WriteString(fmt.Sprintf("@p%d", i+1))
		if i < len(insertFields)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(", @p" + fmt.Sprintf("%d", len(insertFields)+1))
	builder.WriteString(", @p" + fmt.Sprintf("%d", len(insertFields)+2) + ")\"\n")

	builder.WriteString("\tvar lastInsertID int64\n")
	builder.WriteString("\terr := db.DB.QueryRow(query")
	for _, field := range fields {
		if field.Name != "id" && field.Name != "created_at" && field.Name != "updated_at" {
			builder.WriteString(", sql.Named(\"" + strings.ToLower(field.Name) + "\", " + strings.ToLower(field.Name) + ")")
		}
	}
	builder.WriteString(", sql.Named(\"created_at\", time.Now()), sql.Named(\"updated_at\", time.Now())).Scan(&lastInsertID)\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\treturn 0, err\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\treturn lastInsertID, nil\n")
	builder.WriteString("}\n\n")

